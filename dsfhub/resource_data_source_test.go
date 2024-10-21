package dsfhub

import (
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const dsfDataSourceResourceName = "dsfhub_data_source"
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
				Config: testAccCheckDSFDataSourceConfigBasic(),
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

func TestAccDSFDataSource_connectDisconnectGateway(t *testing.T) {
	gatewayId := os.Getenv("GATEWAY_ID")
	if gatewayId == "" {
		t.Fatal("GATEWAY_ID must be set")
	}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// onboard and connect to gateway
			{Config: testAccDSFDataSource_AwsRdsOracle(
				"rds_oracle_connect_disconnect_gateway",
				gatewayId,
				"rds_oracle_connect_disconnect_gateway",
				"UNIFIED",
				true,
			)},
			// update audit_type -> reconnect asset to gateway
			{Config: testAccDSFDataSource_AwsRdsOracle(
				"rds_oracle_connect_disconnect_gateway",
				gatewayId,
				"rds_oracle_connect_disconnect_gateway",
				"UNIFIED_AGGREGATED",
				true,
			)},
			// disconnect asset
			{Config: testAccDSFDataSource_AwsRdsOracle(
				"rds_oracle_connect_disconnect_gateway",
				gatewayId,
				"rds_oracle_connect_disconnect_gateway",
				"UNIFIED_AGGREGATED",
				false,
			)},
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

func testAccCheckDSFDataSourceConfigBasic() string {
	// log.Printf("[INFO] Running test testAccCheckDSFDataSourceConfigBasic \n")
	return fmt.Sprintf(`
resource "%s" "my_test_data_source" {
	admin_email = "%s"
	arn = "%s"
	asset_id = "%s"
	asset_display_name = "%s"
	gateway_id = "%s"
	server_host_name = "%s"
	server_type = "%s"
}`, dsfDataSourceResourceName, testAdminEmail, testArn, testArn, testAssetDisplayName, testGatewayId, testServerHostName, testDSServerType)
}

// Confirm assets are destroyed after an acceptance test run
func testAccDSFDataSourceDestroy(state *terraform.State) error {
	log.Printf("[INFO] Running test testAccDSFDataSourceDestroy \n")
	// allow "disableAsset" playbook enough time to run
	time.Sleep(5 + time.Second)

	// check if asset still exists on hub
	client := testAccProvider.Meta().(*Client)
	for _, res := range state.RootModule().Resources {
		if res.Type != "dsfhub_data_source" {
			continue
		}
		dsfDataSourceId := res.Primary.ID
		readDSFDataSourceResponse, err := client.ReadDSFDataSource(dsfDataSourceId)
		if readDSFDataSourceResponse.Errors == nil {
			return fmt.Errorf("DSF Data Source %s should have received an error in the response", dsfDataSourceId)
		}
		if err == nil {
			return fmt.Errorf("DSF Data Source %s still exists", dsfDataSourceId)
		}
	}
	return nil
}

func testAccDSFDataSource_AwsRdsOracle(resourceName string, gatewayId string, assetId string, auditType string, auditPullEnabled bool) string {
	return fmt.Sprintf(`
resource "dsfhub_data_source" "%[1]s" {
	server_type 				= "AWS RDS ORACLE"

	admin_email 				= "test@example.com"
	asset_display_name 	= "%[3]s"
	asset_id 						= "%[3]s"
	audit_pull_enabled 	= %[5]t
	audit_type          = "%[4]s"
	gateway_id 					= "%[2]s"
	server_host_name 		= "test.com"
	server_ip 					= "test.com"
	server_port 				= "1521"
	service_name 				= "ORCL"

	asset_connection {
		auth_mechanism 	= "password"
		password 				= "password"
		reason 					= "default"
		username				= "username"
	}
}
`,
		resourceName, gatewayId, assetId, auditType, auditPullEnabled)
}
