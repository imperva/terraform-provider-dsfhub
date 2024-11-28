package dsfhub

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDSFCloudAccount_Aws(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		assetId      = "arn:aws:iam::123456789012"
		resourceName = "aws-cloud-account"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfCloudAccountResourceType, resourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCloudAccountDestroy,
		Steps: []resource.TestStep{
			{Config: testAccDSFCloudAccountConfig_Aws(resourceName, gatewayId, assetId, "default")},
			{Config: testAccDSFCloudAccountConfig_Aws(resourceName, gatewayId, assetId, "iam_role")},
			{Config: testAccDSFCloudAccountConfig_Aws(resourceName, gatewayId, assetId, "key")},
			{Config: testAccDSFCloudAccountConfig_Aws(resourceName, gatewayId, assetId, "profile")},
			// validate import
			{
				ResourceName:      resourceTypeAndName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDSFCloudAccount_Azure(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		assetId      = "/subscriptions/11111111-2222-3333-4444-123456789012/asset"
		resourceName = "azure-cloud-account"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfCloudAccountResourceType, resourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCloudAccountDestroy,
		Steps: []resource.TestStep{
			{Config: testAccDSFCloudAccountConfig_Azure(resourceName, gatewayId, assetId, "client_secret")},
			{Config: testAccDSFCloudAccountConfig_Azure(resourceName, gatewayId, assetId, "auth_file")},
			{Config: testAccDSFCloudAccountConfig_Azure(resourceName, gatewayId, assetId, "managed_identity")},
			// validate import
			{
				ResourceName:      resourceTypeAndName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDSFCloudAccount_Gcp(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		assetId      = "my_service_account@project-name.iam.gserviceaccount.com:project-name"
		resourceName = "gcp-cloud-account"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfCloudAccountResourceType, resourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCloudAccountDestroy,
		Steps: []resource.TestStep{
			{Config: testAccDSFCloudAccountConfig_Gcp(resourceName, gatewayId, assetId, "default")},
			{Config: testAccDSFCloudAccountConfig_Gcp(resourceName, gatewayId, assetId, "service_account")},
			// validate import
			{
				ResourceName:      resourceTypeAndName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCloudAccountId(state *terraform.State) (string, error) {
	log.Printf("[INFO] Running test testAccCloudAccountId \n")
	for _, rs := range state.RootModule().Resources {
		if rs.Type != dsfCloudAccountResourceType {
			continue
		}
		return fmt.Sprintf("%s", rs.Primary.ID), nil
	}
	return "", fmt.Errorf("error finding DSF dataSourceId")
}

func testCheckCloudAccountExists(dataSourceId string) resource.TestCheckFunc {
	log.Printf("[INFO] Running test testCheckCloudAccountExists \n")
	return func(state *terraform.State) error {
		res, ok := state.RootModule().Resources[dataSourceId]
		if !ok {
			return fmt.Errorf("DSF Cloud Account Source resource not found by dataSourceId: %s", dataSourceId)
		}
		serverType, ok := res.Primary.Attributes["server_type"]
		if !ok || serverType == "" {
			return fmt.Errorf("DSF Cloud Account Server Type does not exist for dataSourceId %s", dataSourceId)
		}
		client := testAccProvider.Meta().(*Client)
		_, err := client.ReadCloudAccount(res.Primary.ID)
		if err != nil {
			return fmt.Errorf("DSF Cloud Account Server Type: %s (dataSourceId: %s) does not exist", serverType, dataSourceId)
		}
		return nil
	}
}

func testAccCloudAccountDestroy(state *terraform.State) error {
	log.Printf("[INFO] Running test testAccCloudAccountDestroy \n")
	client := testAccProvider.Meta().(*Client)
	for _, res := range state.RootModule().Resources {
		if res.Type != "dsfhub_data_source" {
			continue
		}
		cloudAccountId := res.Primary.ID
		readCloudAccountResponse, err := client.ReadCloudAccount(cloudAccountId)
		if readCloudAccountResponse.Errors == nil {
			return fmt.Errorf("DSF Cloud Account %s should have received an error in the response", cloudAccountId)
		}
		if err == nil {
			return fmt.Errorf("DSF Cloud Account %s still exists for gatewayId: %s", cloudAccountId, testGatewayId)
		}
	}
	return nil
}
