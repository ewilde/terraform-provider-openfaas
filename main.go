package main

import (
	"github.com/ewilde/terraform-provider-openfaas/openfaas"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: openfaas.Provider})
}
