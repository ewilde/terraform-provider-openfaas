package openfaas

import (
	"context"
	"fmt"
	"strings"

	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceOpenFaaSFunction() *schema.Resource {
	return &schema.Resource{
		Create: resourceOpenFaaSFunctionCreate,
		Read:   resourceOpenFaaSFunctionRead,
		Update: resourceOpenFaaSFunctionUpdate,
		Delete: resourceOpenFaaSFunctionDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"image": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"secrets": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},
			"labels": &schema.Schema{
				Type:             schema.TypeMap,
				Optional:         true,
				DiffSuppressFunc: labelsDiffFunc,
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
							Optional: true,
						},
						"cpu": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
							Optional: true,
						},
						"cpu": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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
	deploySpec := expandDeploymentSpec(d, name)
	config := meta.(Config)

	statusCode := config.Client.DeployFunction(context.Background(), deploySpec)
	if statusCode >= 300 {
		return fmt.Errorf("error deploying function %s status code %d", name, statusCode)
	}

	d.SetId(name)
	return nil
}

func resourceOpenFaaSFunctionRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Id()
	config := meta.(Config)

	function, err := config.Client.GetFunctionInfo(context.Background(), name, "")

	if err != nil {
		if isFunctionNotFound(err) {
			d.SetId("")
			return nil
		}

		return err
	}

	return flattenOpenFaaSFunctionResource(d, function)
}

func resourceOpenFaaSFunctionUpdate(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	deploySpec := expandDeploymentSpec(d, name)
	config := meta.(Config)
	statusCode := config.Client.DeployFunction(context.Background(), deploySpec)
	if statusCode >= 300 {
		return fmt.Errorf("error deploying function %s status code %d", name, statusCode)
	}

	return nil
}

func resourceOpenFaaSFunctionDelete(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	config := meta.(Config)

	err := config.Client.DeleteFunction(context.Background(), name, "")
	return err
}

func isFunctionNotFound(err error) bool {
	return strings.Contains(err.Error(), "404") ||
		strings.Contains(err.Error(), "No such function")
}

var whiteListLabels = map[string]string{
	"labels.com.openfaas.function": "",
	"labels.function":              "",
	"labels.uid":                   "",
}

const extraProviderLabelsCount = 3

func labelsDiffFunc(k, old, new string, d *schema.ResourceData) bool {
	if _, ok := whiteListLabels[k]; ok {
		return true
	}

	// TODO: call proxy.Versions, when it's merged and only do this is the provider is faas-swarm
	o, err := strconv.Atoi(old)
	if err != nil {
		return old == new
	}

	n, err := strconv.Atoi(new)
	if err != nil {
		return old == new
	}
	if o > 0 {
		o = o - extraProviderLabelsCount
	}

	return o == n
}
