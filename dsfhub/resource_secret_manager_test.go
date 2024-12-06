package dsfhub

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDSFSecretManager_Hashicorp(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		assetId      = testOnPremServerHostName + ":HASHICORP::8200"
		resourceName = "example-hashicorp"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfSecretManagerResourceType, resourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDSFSecretManagerConfig_Hashicorp(resourceName, gatewayId, assetId, "ec2", "vault-role-for-ec2"),
			},
			// validate import
			{
				ResourceName:      resourceTypeAndName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccSecretManagerId(state *terraform.State) (string, error) {
	log.Printf("[INFO] Running test testAccSecretManagerId \n")
	for _, rs := range state.RootModule().Resources {
		if rs.Type != dsfSecretManagerResourceType {
			continue
		}
		return fmt.Sprintf("%s", rs.Primary.ID), nil
	}
	return "", fmt.Errorf("error finding DSF secretManagerId")
}

func testCheckSecretManagerExists(secretManagerId string) resource.TestCheckFunc {
	log.Printf("[INFO] Running test testCheckDSFDataSourceExists \n")
	return func(state *terraform.State) error {
		res, ok := state.RootModule().Resources[secretManagerId]
		if !ok {
			return fmt.Errorf("DSF Data Source resource not found by secretManagerId: %s", secretManagerId)
		}
		serverType, ok := res.Primary.Attributes["server_type"]
		if !ok || serverType == "" {
			return fmt.Errorf("DSF Data Source Server Type does not exist for secretManagerId %s", secretManagerId)
		}
		client := testAccProvider.Meta().(*Client)
		_, err := client.ReadSecretManager(res.Primary.ID)
		if err != nil {
			return fmt.Errorf("DSF Data Source Server Type: %s (secretManagerId: %s) does not exist", serverType, secretManagerId)
		}
		return nil
	}
}

func testAccSecretManagerDestroy(state *terraform.State) error {
	log.Printf("[INFO] Running test testAccDSFDataSourceDestroy \n")
	client := testAccProvider.Meta().(*Client)
	for _, res := range state.RootModule().Resources {
		if res.Type != "secret_manager" {
			continue
		}
		secretManagerId := res.Primary.ID
		readDSFDataSourceResponse, err := client.ReadSecretManager(secretManagerId)
		if readDSFDataSourceResponse.Errors == nil {
			return fmt.Errorf("DSF Data Source %s should have received an error in the response", secretManagerId)
		}
		if err == nil {
			return fmt.Errorf("DSF Data Source %s still exists for gatewayId: %s", secretManagerId, testGatewayId)
		}
	}
	return nil
}
