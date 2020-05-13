package openfaas

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/openfaas/faas-cli/proxy"
	"github.com/openfaas/faas-cli/stack"
	"github.com/openfaas/faas-provider/types"
)

func expandDeploymentSpec(d *schema.ResourceData, name string) *proxy.DeployFunctionSpec {
	deploySpec := &proxy.DeployFunctionSpec{
		FunctionName: name,
		Image:        d.Get("image").(string),
		Update:       true,
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
		deploySpec.Annotations = expandStringMap(v.(map[string]interface{}))
	}

	request, ok := expandFunctionResourceRequest(d)
	if ok {
		deploySpec.FunctionResourceRequest = request
	}

	return deploySpec
}

func expandFunctionResourceRequest(d *schema.ResourceData) (proxy.FunctionResourceRequest, bool) {
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

func expandStringList(configured []interface{}) []string {
	list := make([]string, 0, len(configured))
	for _, v := range configured {
		val, ok := v.(string)
		if ok && val != "" {
			list = append(list, v.(string))
		}
	}
	return list
}

func expandStringMap(m map[string]interface{}) map[string]string {
	list := make(map[string]string, len(m))
	for i, v := range m {
		list[i] = v.(string)
	}
	return list
}

func flattenOpenFaaSFunctionResource(d *schema.ResourceData, function types.FunctionStatus) error {
	d.Set("name", function.Name)
	d.Set("image", function.Image)
	d.Set("f_process", function.EnvProcess)
	d.Set("labels", pointersMapToStringList(function.Labels))
	d.Set("annotations", pointersMapToStringList(function.Annotations))

	return nil
}

func pointersMapToStringList(pointers *map[string]string) map[string]string {
	if pointers != nil {
		return *pointers
	}

	return nil
}
