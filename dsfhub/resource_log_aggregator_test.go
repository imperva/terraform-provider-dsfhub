package dsfhub

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const dsfLogAggregatorResourceType = "dsfhub_log_aggregator"

func TestAccDSFLogAggregator_AwsLogGroup(t *testing.T) {
	gatewayId := os.Getenv("GATEWAY_ID")
	if gatewayId == "" {
		t.Fatal("GATEWAY_ID must be set")
	}

	const (
		assetId = "arn:aws:logs:us-east-2:123456789012:log-group:/aws/rds/instance/my-database/audit:*"
		resourceName = "my-database-log-group"
		serverHostName = "oracle-rds-db.xxxxx8rsfzja.us-east-2.rds.amazonaws.com"
		parentAssetId = "arn:aws:rds:us-east-2:123456789012:db:oracle-rds-db"
		parentResourceName = "my-oracle-db"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, resourceName)
	parentResourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, parentResourceName)

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		Steps: []resource.TestStep{
			// Failed: missing parent_asset_id
			{
				Config: testAccDSFLogAggregatorConfig_AwsLogGroup(resourceName, gatewayId, assetId, "", true),
				ExpectError: regexp.MustCompile("Error: missing required fields for dsfhub_data_source"),
			},
			// Onboard with AWS parent asset
			{
				Config: testAccDSFDataSourceConfig_AwsRdsOracle(parentResourceName, gatewayId, parentAssetId, "LOG_GROUP", "") + 
				testAccDSFLogAggregatorConfig_AwsLogGroup(resourceName, gatewayId, assetId, parentResourceTypeAndName + ".asset_id", true),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "gateway_service", "gateway-aws@oracle-rds.service"),
				),
			},
		},
	})
}

func testAccLogAggregatorId(state *terraform.State) (string, error) {
	log.Printf("[INFO] Running test testAccLogAggregatorId \n")
	for _, rs := range state.RootModule().Resources {
		if rs.Type != dsfLogAggregatorResourceType {
			continue
		}
		return fmt.Sprintf("%s", rs.Primary.ID), nil
	}
	return "", fmt.Errorf("error finding DSF dataSourceId")
}

func testCheckLogAggregatorExists(dataSourceId string) resource.TestCheckFunc {
	log.Printf("[INFO] Running test testCheckLogAggregatorExists \n")
	return func(state *terraform.State) error {
		res, ok := state.RootModule().Resources[dataSourceId]
		if !ok {
			return fmt.Errorf("DSF Log Aggregator Source resource not found by dataSourceId: %s", dataSourceId)
		}
		serverType, ok := res.Primary.Attributes["server_type"]
		if !ok || serverType == "" {
			return fmt.Errorf("DSF Log Aggregator Server Type does not exist for dataSourceId %s", dataSourceId)
		}
		client := testAccProvider.Meta().(*Client)
		_, err := client.ReadLogAggregator(res.Primary.ID)
		if err != nil {
			return fmt.Errorf("DSF Log Aggregator Server Type: %s (dataSourceId: %s) does not exist", serverType, dataSourceId)
		}
		return nil
	}
}

func testAccLogAggregatorDestroy(state *terraform.State) error {
	log.Printf("[INFO] Running test testAccLogAggregatorDestroy \n")
	client := testAccProvider.Meta().(*Client)
	for _, res := range state.RootModule().Resources {
		if res.Type != "dsfhub_log_aggregator" {
			continue
		}
		logAggregatorId := res.Primary.ID
		readLogAggregatorResponse, err := client.ReadLogAggregator(logAggregatorId)
		if readLogAggregatorResponse.Errors == nil {
			return fmt.Errorf("DSF Log Aggregator %s should have received an error in the response", logAggregatorId)
		}
		if err == nil {
			return fmt.Errorf("DSF Log Aggregator %s still exists for gatewayId: %s", logAggregatorId, testGatewayId)
		}
	}
	return nil
}

// Configs
func testAccDSFLogAggregatorConfig_AwsLogGroup(resourceName string, gatewayId string, assetId string, parentAssetId string, auditPullEnabled bool) string {	
	// handle reference to other assets
	parentAssetIdVal := testAccParseResourceReference(parentAssetId)

	return fmt.Sprintf(`
resource "` + dsfLogAggregatorResourceType + `" "%[1]s" {
	server_type = "AWS LOG GROUP"

	admin_email = "` + testAdminEmail + `"
	arn	= "%[3]s"
	asset_display_name = "%[3]s"
	asset_id = "%[3]s"
	audit_pull_enabled = %[5]t
	gateway_id = "%[2]s"
	parent_asset_id = ` + parentAssetIdVal + `

	asset_connection {
		auth_mechanism = "default"
		reason = "default"
		region = "us-east-2"
	}
}`, resourceName, gatewayId, assetId, parentAssetId, auditPullEnabled)
}
