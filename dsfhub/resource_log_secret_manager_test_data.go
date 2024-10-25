package dsfhub

import (
	"fmt"
	"testing"
)

// Output a terraform config for a basic secret manager resource.
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
    reason = "%s"
    auth_mechanism = "%s"
    role_name = "%s"
  }
}`, dsfSecretManagerResourceType, testAdminEmail, testAssetDisplayName, testSMAssetId, testGatewayId, testServerHostName, testServerIP, testServerPort, testSMServerType, testSMConnectionReason, testSMAuthMechanism, testSMRoleName)
}

// Output a terraform config for a HASHICORP secret manager resource.
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
    reason = "default"
    auth_mechanism = "%[6]s"
    role_name = "%[7]s"
  }
}`,
		resourceName, gatewayId, assetId, serverHostName, serverPort, authMechanism, roleName)
}
