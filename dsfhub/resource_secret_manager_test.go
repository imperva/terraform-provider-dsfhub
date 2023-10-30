package dsfhub

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"log"
	"testing"
)

const secretManagerResourceName = "secret_manager"
const secretManagerType = "HASHICORP"
const secretManagerResourceTypeAndName = secretManagerResourceName + "." + secretManagerType

func TestAccSecretManager_basic(t *testing.T) {
	log.Printf("======================== BEGIN TEST ========================")
	log.Printf("[INFO] Running test TestAccSecretManager_basic \n")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccSecretManagerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckSecretManagerConfigBasic(t),
				Check: resource.ComposeTestCheckFunc(
					testCheckSecretManagerExists(secretManagerResourceName),
					resource.TestCheckResourceAttr(secretManagerResourceTypeAndName, secretManagerResourceName, secretManagerType),
				),
			},
			{
				ResourceName:      secretManagerResourceTypeAndName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: testAccSecretManagerId,
			},
		},
	})
}

func testAccSecretManagerId(state *terraform.State) (string, error) {
	log.Printf("[INFO] Running test testAccSecretManagerId \n")
	for _, rs := range state.RootModule().Resources {
		if rs.Type != dsfDataSourceType {
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

func testAccCheckSecretManagerConfigBasic(t *testing.T) string {
	log.Printf("[INFO] Running test testAccCheckSecretManagerConfigBasic \n")
	return fmt.Sprintf(`
resource "%s" "my_test_data_source" {
	admin_email = "%s"
	asset_display_name = "%s"
	asset_id = "%s"
	gateway_id = "%s""
	server_host_name = "%s"
	server_ip = "%s"
	server_port = "%s"
	server_type = "%s"
	sm_connection {
		reason         = "%s"
		auth_mechanism = "%s"
		role_name      = "%s"
	}
}`, secretManagerResourceName, testAdminEmail, testAssetDisplayName, testSMAssetId, testGatewayId, testServerHostName, testServerIP, testServerPort, testSMServerType, testSMConnectionReason, testSMAuthMechanism, testSMRoleName)
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
