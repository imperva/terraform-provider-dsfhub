package dsfhub

import (
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

const logAggregatorResourceName = "log_aggregator"
const logAggregatorType = "aws"
const logAggregatorResourceTypeAndName = logAggregatorResourceName + "." + logAggregatorType

func TestAccLogAggregator_basic(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestAccLogAggregator_basic \n")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccLogAggregatorDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckLogAggregatorConfigBasic(t),
				Check: resource.ComposeTestCheckFunc(
					testCheckLogAggregatorExists(logAggregatorResourceName),
					resource.TestCheckResourceAttr(logAggregatorResourceTypeAndName, logAggregatorResourceName, logAggregatorType),
				),
			},
			{
				ResourceName:      logAggregatorResourceTypeAndName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccLogAggregatorId,
			},
		},
	})
}

func testAccLogAggregatorId(state *terraform.State) (string, error) {
	log.Printf("[INFO] Running test testAccLogAggregatorId \n")
	for _, rs := range state.RootModule().Resources {
		if rs.Type != logAggregatorType {
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

func testAccCheckLogAggregatorConfigBasic(t *testing.T) string {
	log.Printf("[INFO] Running test testAccCheckLogAggregatorConfigBasic \n")
	return fmt.Sprintf(`
resource "%s" "my_test_data_source" {
	admin_email = "%s"
	arn = "%s"
	asset_display_name = "%s"
	gateway_id = %s
	server_host_name = "%s"
	server_type = "%s"
}`, logAggregatorResourceName, testAdminEmail, testArn, testAssetDisplayName, testGatewayId, testServerHostName, testDSServerType)
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
