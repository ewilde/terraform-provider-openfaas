package openfaas

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceOpenFaaSFunction() *schema.Resource {
	return &schema.Resource{
		Create: resourceOpenFaaSFunctionCreate,
		Read:   resourceOpenFaaSFunctionRead,
		Update: resourceOpenFaaSFunctionUpdate,
		Delete: resourceOpenFaaSFunctionDelete,

		Schema: map[string]*schema.Schema{
			"service": &schema.Schema{
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
			"env_process": &schema.Schema{
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
				Type: schema.TypeSet,
				Optional: true,
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
				Type: schema.TypeSet,
				Optional: true,
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
	return nil
}

func resourceOpenFaaSFunctionRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceOpenFaaSFunctionUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceOpenFaaSFunctionDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}