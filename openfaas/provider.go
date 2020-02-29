package openfaas

import (
	"crypto/tls"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/openfaas/faas-cli/config"
	"github.com/openfaas/faas-cli/proxy"
)

var (
	defaultTimeout = 60 * time.Second
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	log.Printf("[DEBUG] returning provider schema")
	// The actual provider
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"uri": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "http://localhost:8080",
				Description: "OpenFaaS gateway uri",
			},
			"tls_insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "if true, skip tls verification (not recommended)",
			},
			"user_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "OpenFaaS gateway username",
			},

			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Default:     "",
				Description: "OpenFaaS gateway password",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"openfaas_function": dataSourceOpenFaaSFunction(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"openfaas_function": resourceOpenFaaSFunction(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	log.Printf("[DEBUG] configuring provider")

	gatewayURI := d.Get("uri").(string)
	auth := newCLIAuth("", gatewayURI)
	insecure := d.Get("tls_insecure").(bool)
	transport := GetDefaultCLITransport(insecure, &defaultTimeout)
	client := proxy.NewClient(auth, gatewayURI, transport, &defaultTimeout)

	providerConfig := Config{
		Client: client,
	}

	return providerConfig, nil
}

func newCLIAuth(token string, gateway string) proxy.ClientAuth {
	authConfig, _ := config.LookupAuthConfig(gateway)

	var (
		username    string
		password    string
		bearerToken string
	)

	if authConfig.Auth == config.BasicAuthType {
		username, password, _ = config.DecodeAuth(authConfig.Token)

		return &BasicAuth{
			username: username,
			password: password,
		}

	}

	// User specified token gets priority
	if len(token) > 0 {
		bearerToken = token
	} else {
		bearerToken = authConfig.Token
	}

	return &BearerToken{
		token: bearerToken,
	}
}

type BasicAuth struct {
	username string
	password string
}

func (auth *BasicAuth) Set(req *http.Request) error {
	if auth.username == "" {
		return nil
	}

	req.SetBasicAuth(auth.username, auth.password)
	return nil
}

type BearerToken struct {
	token string
}

func (c *BearerToken) Set(req *http.Request) error {
	req.Header.Set("Authorization", "Bearer "+c.token)
	return nil
}

func GetDefaultCLITransport(tlsInsecure bool, timeout *time.Duration) *http.Transport {
	if timeout != nil || tlsInsecure {
		tr := &http.Transport{
			Proxy:             http.ProxyFromEnvironment,
			DisableKeepAlives: false,
		}

		if timeout != nil {
			tr.DialContext = (&net.Dialer{
				Timeout: *timeout,
			}).DialContext

			tr.IdleConnTimeout = 120 * time.Millisecond
			tr.ExpectContinueTimeout = 1500 * time.Millisecond
		}

		if tlsInsecure {
			tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: tlsInsecure}
		}
		tr.DisableKeepAlives = false

		return tr
	}
	return nil
}
