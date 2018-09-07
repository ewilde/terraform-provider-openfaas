package openfaas

import (
	"github.com/hashicorp/terraform/helper/resource"
	"fmt"
	"github.com/hashicorp/terraform/helper/acctest"
	"testing"
)

func TestAccOpenFaaSFunction_importBasic(t *testing.T) {
	resourceName := "openfaas_function.function_test"
	name := fmt.Sprintf("testaccopenfaasfunction-basic-%s", acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckOpenFaaSFunctionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccOpenFaaSFunctionConfig_basic(name),
			},

			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

