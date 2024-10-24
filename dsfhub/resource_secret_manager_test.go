package dsfhub

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDSFSecretManager_Hashicorp(t *testing.T) {
	gatewayId := os.Getenv("GATEWAY_ID")
	if gatewayId == "" {
		t.Skip("GATEWAY_ID environment variable must be set")
	}

	const (
		serverPort   = "8200"
		assetId      = testOnPremServerHostName + ":HASHICORP::" + serverPort
		resourceName = "example-hashicorp"
	)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDSFSecretManagerConfig_Hashicorp(resourceName, gatewayId, assetId, testOnPremServerHostName, serverPort, "ec2", "vault-role-for-ec2"),
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

// Configs
func testAccSecretManagerConfig_Basic(t *testing.T) string {
	return fmt.Sprintf(`
resource "%s" "my_test_data_source" {
	admin_email = "%s"
	asset_display_name = "%s"
	asset_id = "%s"
	gateway_id = "%s"
	server_host_name = "%s"
	server_ip = "%s"
	server_port = "%s"
	server_type = "%s"
	asset_connection {
		reason         = "%s"
		auth_mechanism = "%s"
		role_name      = "%s"
	}
}`, dsfSecretManagerResourceType, testAdminEmail, testAssetDisplayName, testSMAssetId, testGatewayId, testServerHostName, testServerIP, testServerPort, testSMServerType, testSMConnectionReason, testSMAuthMechanism, testSMRoleName)
}

func testAccDSFSecretManagerConfig_Hashicorp(resourceName string, gatewayId string, assetId string, serverHostName string, serverPort string, authMechanism string, roleName string) string {
	return fmt.Sprintf(`
resource "`+dsfSecretManagerResourceType+`" "%[1]s" {
	server_type = "HASHICORP"
	
	admin_email = "`+testAdminEmail+`"
	asset_display_name = "%[3]s"
	asset_id = "%[3]s"
	gateway_id = "%[2]s"
	server_host_name = "%[4]s"
	server_ip = "%[4]s"
	server_port = "%[5]s"

	asset_connection {
		reason         = "default"
		auth_mechanism = "%[6]s"
		role_name      = "%[7]s"
	}
}`,
		resourceName, gatewayId, assetId, serverHostName, serverPort, authMechanism, roleName)
}
