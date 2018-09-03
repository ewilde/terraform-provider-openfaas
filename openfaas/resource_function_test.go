package openfaas

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/openfaas/faas/gateway/requests"
	"testing"
	"github.com/hashicorp/terraform/terraform"
	"github.com/viveksyngh/faas-cli/proxy"
	"errors"
)

func TestAccOpenFaaSFunction_basic(t *testing.T) {
	var conf requests.Function
	functionName := fmt.Sprintf("testaccopenfaasfunction-basic-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

	resource.Test(t, resource.TestCase{
		PreCheck:      func() { testAccPreCheck(t) },
		IDRefreshName: "openfaas_function.function_test",
		Providers:     testAccProviders,
		CheckDestroy:  testAccCheckOpenFaaSFunctionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccOpenFaaSFunctionConfig_basic(functionName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckOpenFaaSFunctionExists("openfaas_function.function_test", &conf),
					resource.TestCheckResourceAttr("openfaas_function.function_test", "name", functionName),
					resource.TestCheckResourceAttr("openfaas_function.function_test", "image", "functions/alpine:latest"),
					resource.TestCheckResourceAttr("openfaas_function.function_test", "f_process", "sha512sum"),
					resource.TestCheckResourceAttr("openfaas_function.function_test", "labels.%", "2"),
					resource.TestCheckResourceAttr("openfaas_function.function_test", "labels.Name", "TestAccOpenFaaSFunction_basic"),
					resource.TestCheckResourceAttr("openfaas_function.function_test", "labels.Environment", "Test"),
					resource.TestCheckResourceAttr("openfaas_function.function_test", "annotations.%", "1"),
					resource.TestCheckResourceAttr("openfaas_function.function_test", "annotations.CreatedDate", "Mon Sep  3 07:15:55 BST 2018"),
				),
			},
		},
	})
}

func testAccCheckOpenFaaSFunctionDestroy(s *terraform.State) error {
	config := testAccProvider.Meta().(Config)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "openfaas_function" {
			continue
		}

		_, err := proxy.GetFunctionInfo(config.GatewayURI, rs.Primary.ID, config.TLSInsecure)

		if err == nil {
			return fmt.Errorf("function %q still exists", rs.Primary.ID)
		}

		// Verify the error
		if isFunctionNotFound(err) {
			return nil
		} else {
			return fmt.Errorf("unexpected error checking function destroyed: %s", err)
		}
	}

	return nil
}

func testAccCheckOpenFaaSFunctionExists(n string, res *requests.Function) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("no function ID is set")
		}

		config := testAccProvider.Meta().(Config)

		function, err := proxy.GetFunctionInfo(config.GatewayURI, rs.Primary.ID, config.TLSInsecure)


		if err != nil {
			return err
		}

		*res = function
		return nil
	}
}

func testAccOpenFaaSFunctionConfig_basic(functionName string) string {
	return fmt.Sprintf(`resource "openfaas_function" "function_test" {
  name            = "%s"
  image           = "functions/alpine:latest"
  f_process       = "sha512sum"
  labels {
    Name        = "TestAccOpenFaaSFunction_basic"
    Environment = "Test"
  }

  annotations {
    CreatedDate = "Mon Sep  3 07:15:55 BST 2018"
  }
}`, functionName)
}
