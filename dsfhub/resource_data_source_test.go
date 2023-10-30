package dsfhub

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"log"
	"testing"
)

const dsfDataSourceResourceName = "dsf_data_source"
const dsfDataSourceType = "aws-rds-mysql"
const dsfDataSourceResourceTypeAndName = dsfDataSourceResourceName + "." + dsfDataSourceType

func TestAccDSFDataSource_basic(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestAccDSFDataSource_basic \n")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccDSFDataSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDSFDataSourceConfigBasic(t),
				Check: resource.ComposeTestCheckFunc(
					testCheckDSFDataSourceExists(dsfDataSourceResourceName),
					resource.TestCheckResourceAttr(dsfDataSourceResourceTypeAndName, dsfDataSourceResourceName, dsfDataSourceType),
				),
			},
			{
				ResourceName:      dsfDataSourceResourceTypeAndName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccDSFDataSourceId,
			},
		},
	})
}

func testAccDSFDataSourceId(state *terraform.State) (string, error) {
	log.Printf("[INFO] Running test testAccDSFDataSourceId \n")
	for _, rs := range state.RootModule().Resources {
		if rs.Type != dsfDataSourceType {
			continue
		}
		return fmt.Sprintf("%s", rs.Primary.ID), nil
	}
	return "", fmt.Errorf("error finding DSF dataSourceId")
}

func testCheckDSFDataSourceExists(dataSourceId string) resource.TestCheckFunc {
	log.Printf("[INFO] Running test testCheckDSFDataSourceExists \n")
	return func(state *terraform.State) error {
		res, ok := state.RootModule().Resources[dataSourceId]
		if !ok {
			return fmt.Errorf("DSF Data Source resource not found by dataSourceId: %s", dataSourceId)
		}
		serverType, ok := res.Primary.Attributes["server_type"]
		if !ok || serverType == "" {
			return fmt.Errorf("DSF Data Source Server Type does not exist for dataSourceId %s", dataSourceId)
		}
		client := testAccProvider.Meta().(*Client)
		_, err := client.ReadDSFDataSource(res.Primary.ID)
		if err != nil {
			return fmt.Errorf("DSF Data Source Server Type: %s (dataSourceId: %s) does not exist", serverType, dataSourceId)
		}
		return nil
	}
}

func testAccCheckDSFDataSourceConfigBasic(t *testing.T) string {
	log.Printf("[INFO] Running test testAccCheckDSFDataSourceConfigBasic \n")
	return fmt.Sprintf(`
resource "%s" "my_test_data_source" {
	admin_email = "%s"
	arn = "%s"
	asset_display_name = "%s"
	gateway_id = %s
	server_host_name = "%s"
	server_type = "%s"
}`, dsfDataSourceResourceName, testAdminEmail, testArn, testAssetDisplayName, testGatewayId, testServerHostName, testDSServerType)
}

func testAccDSFDataSourceDestroy(state *terraform.State) error {
	log.Printf("[INFO] Running test testAccDSFDataSourceDestroy \n")
	client := testAccProvider.Meta().(*Client)
	for _, res := range state.RootModule().Resources {
		if res.Type != "dsf_data_source" {
			continue
		}
		dsfDataSourceId := res.Primary.ID
		readDSFDataSourceResponse, err := client.ReadDSFDataSource(dsfDataSourceId)
		if readDSFDataSourceResponse.Errors == nil {
			return fmt.Errorf("DSF Data Source %s should have received an error in the response", dsfDataSourceId)
		}
		if err == nil {
			return fmt.Errorf("DSF Data Source %s still exists for gatewayId: %s", dsfDataSourceId, testGatewayId)
		}
	}
	return nil
}
