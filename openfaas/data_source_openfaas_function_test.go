package openfaas

import (
	"fmt"
	"testing"

	"strconv"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourceOpenFaaSFunction_basic(t *testing.T) {
	name := fmt.Sprintf("testaccopenfaasfunction-basic-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceOpenFaaSFunctionConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.openfaas_function.function_test", "name", name),
					resource.TestCheckResourceAttr("data.openfaas_function.function_test", "image", "functions/alpine:latest"),
					resource.TestCheckResourceAttr("data.openfaas_function.function_test", "f_process", "sha512sum"),
					resource.TestCheckResourceAttr("data.openfaas_function.function_test", "labels.%", strconv.FormatInt(2+extraProviderLabelsCount, 10)),
					resource.TestCheckResourceAttr("data.openfaas_function.function_test", "labels.Name", "TestAccOpenFaaSFunction_basic"),
					resource.TestCheckResourceAttr("data.openfaas_function.function_test", "labels.Environment", "Test"),
					resource.TestCheckResourceAttr("data.openfaas_function.function_test", "annotations.%", "1"),
					resource.TestCheckResourceAttr("data.openfaas_function.function_test", "annotations.CreatedDate", "Mon Sep  3 07:15:55 BST 2018"),
				),
			},
		},
	})
}

func testAccDataSourceOpenFaaSFunctionConfigBasic(functionName string) string {
	return fmt.Sprintf(`resource "openfaas_function" "function_test" {
  name      = "%s"
  image     = "functions/alpine:latest"
  f_process = "sha512sum"
  labels = {
    Name        = "TestAccOpenFaaSFunction_basic"
    Environment = "Test"
  }

  annotations = {
    CreatedDate = "Mon Sep  3 07:15:55 BST 2018"
  }
}

data "openfaas_function" "function_test" {
  name = openfaas_function.function_test.name
}`, functionName)
}
