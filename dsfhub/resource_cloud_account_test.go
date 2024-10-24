package dsfhub

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDSFCloudAccount_Aws(t *testing.T) {
	gatewayId := os.Getenv("GATEWAY_ID")
	if gatewayId == "" {
		t.Skip("GATEWAY_ID environment variable must be set")
	}

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
			// {Config: testAccDSFCloudAccountConfig_Aws(resourceName, gatewayId, assetId, "key")}, //TODO: fix "key" failing refresh
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

// Configs
func testAccDSFCloudAccountConfig_Aws(resourceName string, gatewayId string, assetId string, authMechanism string) string {
	var assetConnectionBlock string

	if authMechanism == "key" {
		assetConnectionBlock = fmt.Sprintf(`
			asset_connection {
				access_id = "my-access-id"
				auth_mechanism = "` + authMechanism + `"
				reason = "default"
				region = "us-east-1"
				secret_key = "my-secret-key"
			}
		`)
	} else if authMechanism == "profile" {
		assetConnectionBlock = fmt.Sprintf(`
			asset_connection {
				auth_mechanism = "` + authMechanism + `"
				reason = "default"
				region = "us-east-2"
				username = "dsfhubuser"
			}
		`)
	} else {
		assetConnectionBlock = fmt.Sprintf(`
			asset_connection {
				auth_mechanism = "` + authMechanism + `"
				reason = "default"
				region = "us-west-1"
			}
			`)
	}

	return fmt.Sprintf(`
resource "`+dsfCloudAccountResourceType+`" "%[1]s" {
	server_type = "AWS"
	
	admin_email = "`+testAdminEmail+`"
	arn = "%[3]s"
	asset_display_name = "%[3]s"
	asset_id = "%[3]s"
	gateway_id = "%[2]s"

	`+assetConnectionBlock+`
}`, resourceName, gatewayId, assetId)
}
