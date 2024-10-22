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

const dsfDataSourceResourceType = "dsfhub_data_source"

func TestAccDSFDataSource_Basic(t *testing.T) {
	gatewayId := os.Getenv("GATEWAY_ID")
	if gatewayId == "" {
		t.Fatal("GATEWAY_ID must be set")
	}

	const resourceName = "basic_test_data_source"
	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDSFDataSourceConfig_Basic(
					resourceName,
					testAdminEmail,
					testArn,
					gatewayId,
					testServerHostName,
					testDSServerType,
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "false"),
				),
			},
			{
				ResourceName:      resourceTypeAndName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDSFDataSource_AwsRdsOracleConnectDisconnectGateway(t *testing.T) {
	gatewayId := os.Getenv("GATEWAY_ID")
	if gatewayId == "" {
		t.Fatal("GATEWAY_ID must be set")
	}

	const resourceName = "rds_oracle_connect_disconnect_gateway"
	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// onboard and connect to gateway
			{
				Config: testAccDSFDataSourceConfig_AwsRdsOracle(resourceName, gatewayId, resourceName, "UNIFIED", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "gateway_service", "gateway-odbc@oracle_unified.service"),
				),
			},
			// update audit_type -> reconnect asset to gateway
			{
				Config: testAccDSFDataSourceConfig_AwsRdsOracle(resourceName, gatewayId, resourceName, "UNIFIED_AGGREGATED", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "gateway_service", "gateway-odbc@oracle_unified_aggregated.service"),
				),
			},
			// disconnect asset
			{
				Config: testAccDSFDataSourceConfig_AwsRdsOracle(resourceName, gatewayId, resourceName, "UNIFIED_AGGREGATED", false),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "false"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "gateway_service", ""),
				),
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

func testAccDSFDataSourceId(state *terraform.State) (string, error) {
	log.Printf("[INFO] Running test testAccDSFDataSourceId \n")
	for _, rs := range state.RootModule().Resources {
		if rs.Type != dsfDataSourceResourceType {
			continue
		}
		return fmt.Sprintf("%s", rs.Primary.ID), nil
	}
	return "", fmt.Errorf("error finding DSF dataSourceId")
}

// Confirm assets are destroyed after an acceptance test run
func testAccDSFDataSourceDestroy(state *terraform.State) error {
	log.Printf("[INFO] Running test testAccDSFDataSourceDestroy")
	// allow "disableAsset" playbook enough time to run
	time.Sleep(5 + time.Second)

	// check if asset still exists on hub
	client := testAccProvider.Meta().(*Client)
	for _, res := range state.RootModule().Resources {
		if res.Type != dsfDataSourceResourceType {
			continue
		}
		assetId := res.Primary.ID
		readDSFDataSourceResponse, err := client.ReadDSFDataSource(assetId)
		if readDSFDataSourceResponse.Errors == nil {
			return fmt.Errorf("DSF Data Source %s should have received an error in the response", assetId)
		}
		if err == nil {
			return fmt.Errorf("DSF Data Source %s still exists", assetId)
		}
	}
	return nil
}

// Configs
func testAccDSFDataSourceConfig_Basic(resourceName string, adminEmail string, assetId string, gatewayId string, serverHostName string, serverType string) string {
	return fmt.Sprintf(`
resource "` + dsfDataSourceResourceType + `" "%[1]s" {
	admin_email = "%[2]s"
	asset_id = "%[3]s"
	asset_display_name = "%[3]s"
	gateway_id = "%[4]s"
	server_host_name = "%[5]s"
	server_type = "%[6]s"
}`,
		resourceName, adminEmail, assetId, gatewayId, serverHostName, serverType)
}

func testAccDSFDataSourceConfig_AwsRdsOracle(resourceName string, gatewayId string, assetId string, auditType string, auditPullEnabled bool) string {
	return fmt.Sprintf(`
resource "` + dsfDataSourceResourceType + `" "%[1]s" {
	server_type					= "AWS RDS ORACLE"

	admin_email					= "` + testAdminEmail + `"
	asset_display_name	= "%[3]s"
	asset_id						= "%[3]s"
	audit_pull_enabled	= %[5]t
	audit_type					= "%[4]s"
	gateway_id					= "%[2]s"
	server_host_name		= "test.com"
	server_port					= "1521"
	service_name				= "ORCL"

	asset_connection {
		auth_mechanism	= "password"
		password				= "password"
		reason					= "default"
		username				= "username"
	}
}
`,
		resourceName, gatewayId, assetId, auditType, auditPullEnabled)
}
