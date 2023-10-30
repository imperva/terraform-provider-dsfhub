package dsfhub

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const cloudAccountResourceName = "cloud_account"
const cloudAccountType = "aws"
const cloudAccountResourceTypeAndName = cloudAccountResourceName + "." + cloudAccountType

func TestAccCloudAccount_basic(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestAccCloudAccount_basic \n")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCloudAccountDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCloudAccountConfigBasic(t),
				Check: resource.ComposeTestCheckFunc(
					testCheckCloudAccountExists(cloudAccountResourceName),
					resource.TestCheckResourceAttr(cloudAccountResourceTypeAndName, cloudAccountResourceName, cloudAccountType),
				),
			},
			{
				ResourceName:      cloudAccountResourceTypeAndName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccCloudAccountId,
			},
		},
	})
}

func testAccCloudAccountId(state *terraform.State) (string, error) {
	log.Printf("[INFO] Running test testAccCloudAccountId \n")
	for _, rs := range state.RootModule().Resources {
		if rs.Type != cloudAccountType {
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

func testAccCheckCloudAccountConfigBasic(t *testing.T) string {
	log.Printf("[INFO] Running test testAccCheckCloudAccountConfigBasic \n")
	return fmt.Sprintf(`
resource "%s" "my_test_data_source" {
	admin_email = "%s"
	arn = "%s"
	asset_display_name = "%s"
	gateway_id = %s
	server_host_name = "%s"
	server_type = "%s"
}`, cloudAccountResourceName, testAdminEmail, testArn, testAssetDisplayName, testGatewayId, testServerHostName, testDSServerType)
}

func testAccCloudAccountDestroy(state *terraform.State) error {
	log.Printf("[INFO] Running test testAccCloudAccountDestroy \n")
	client := testAccProvider.Meta().(*Client)
	for _, res := range state.RootModule().Resources {
		if res.Type != "dsf_data_source" {
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
