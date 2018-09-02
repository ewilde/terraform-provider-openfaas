package openfaas

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/viveksyngh/faas-cli/proxy"
	"github.com/openfaas/faas-cli/stack"
)

func resourceOpenFaaSFunction() *schema.Resource {
	return &schema.Resource{
		Create: resourceOpenFaaSFunctionCreate,
		Read:   resourceOpenFaaSFunctionRead,
		Update: resourceOpenFaaSFunctionUpdate,
		Delete: resourceOpenFaaSFunctionDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"image": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Computed: true,
			},
			"network": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"f_process": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"env_vars": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"registry_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"constraints": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"secrets": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"labels": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"annotations": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
			"limits": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"memory": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"cpu": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"requests": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"memory": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"cpu": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"read_only_root_file_system": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}

func resourceOpenFaaSFunctionCreate(d *schema.ResourceData, meta interface{}) error {

	name := d.Get("name").(string)
	deploySpec := &proxy.DeployFunctionSpec{
		FunctionName: name,
		Image:        d.Get("image").(string),
	}

	if v, ok := d.GetOk("network"); ok {
		deploySpec.Network = v.(string)
	}

	if v, ok := d.GetOk("f_process"); ok {
		deploySpec.FProcess = v.(string)
	}

	if v, ok := d.GetOk("env_vars"); ok {
		deploySpec.EnvVars = expandStringMap(v.(map[string]interface{}))
	}

	if v, ok := d.GetOk("registry_auth"); ok {
		deploySpec.RegistryAuth = v.(string)
	}

	if v, ok := d.GetOk("constraints"); ok {
		deploySpec.Constraints = expandStringList(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("secrets"); ok {
		deploySpec.Secrets = expandStringList(v.(*schema.Set).List())
	}

	if v, ok := d.GetOk("labels"); ok {
		deploySpec.Labels = expandStringMap(v.(map[string]interface{}))
	}

	if v, ok := d.GetOk("annotations"); ok {
		deploySpec.Labels = expandStringMap(v.(map[string]interface{}))
	}

	request, ok := buildFunctionResourceRequest(d)
	if ok {
		deploySpec.FunctionResourceRequest = request
	}

	statusCode := proxy.DeployFunction(deploySpec)
	if statusCode >= 300 {
		return fmt.Errorf("error deploying function %s status code %d", name, statusCode)
	}

	d.SetId(name)
	return nil
}

func resourceOpenFaaSFunctionRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	config := meta.(Config)
	function, err := proxy.GetFunctionInfo(config.GatewayURI, name, config.TLSInsecure)

	if err != nil {
		if strings.Contains(err.Error(), "404") {
			d.SetId("")
			return nil
		}

		return err
	}

	d.Set("image", function.Image)
	d.Set("f_process", function.EnvProcess)
	d.Set("labels", pointersMapToStringList(function.Labels))
	return nil
}

func resourceOpenFaaSFunctionUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceOpenFaaSFunctionDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func buildFunctionResourceRequest(d *schema.ResourceData) (proxy.FunctionResourceRequest, bool) {
	rLimits, okLimits := d.GetOk("limits")
	rRequests, okRequests := d.GetOk("requests")

	if !okLimits && !okRequests {
		return *new(proxy.FunctionResourceRequest), false
	}

	var limits *stack.FunctionResources
	var requests *stack.FunctionResources
	if okLimits && len(rLimits.(*schema.Set).List()) > 0 {
		data := rLimits.(*schema.Set).List()[0].(map[string]interface{})
		limits = &stack.FunctionResources{
			Memory: data["memory"].(string),
			CPU:    data["cpu"].(string),
		}
	}

	if okRequests && len(rRequests.(*schema.Set).List()) > 0 {
		data := rRequests.(*schema.Set).List()[0].(map[string]interface{})
		requests = &stack.FunctionResources{
			Memory: data["memory"].(string),
			CPU:    data["cpu"].(string),
		}
	}

	return *&proxy.FunctionResourceRequest{
		Limits:   limits,
		Requests: requests,
	}, true
}
