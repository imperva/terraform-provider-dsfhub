package dsfhub

import (
	"fmt"
	"testing"
)

// Output a terraform config for a basic secret manager resource.
func testAccSecretManagerConfig_Basic(t *testing.T) string {
	return fmt.Sprintf(`
resource "%s" "my_test_data_source" {
  admin_email        = "%s"
  asset_display_name = "%s"
  asset_id           = "%s"
  gateway_id         = "%s"
  server_host_name   = "%s"
  server_ip          = "%s"
  server_port        = "%s"
  server_type        = "%s"
  
  asset_connection {
    reason         = "%s"
    auth_mechanism = "%s"
    role_name      = "%s"
  }
}`, dsfSecretManagerResourceType, testAdminEmail, testAssetDisplayName, testSMAssetId, testGatewayId, testServerHostName, testServerIP, testServerPort, testSMServerType, testSMConnectionReason, testSMAuthMechanism, testSMRoleName)
}

// Output a terraform config for a HASHICORP secret manager resource.
func testAccDSFSecretManagerConfig_Hashicorp(resourceName string, gatewayId string, assetId string, authMechanism string, roleName string) string {
	return fmt.Sprintf(`
resource "%[1]s" "%[2]s" {
  server_type        = "HASHICORP"
  
  admin_email        = "%[3]s"
  asset_display_name = "%[5]s"
  asset_id           = "%[5]s"
  gateway_id         = "%[4]s"
  server_host_name   = "%[8]s"
  server_ip          = "1.2.3.4"
  server_port        = "8200"

  asset_connection {
    reason         = "default"
    auth_mechanism = "%[6]s"
    role_name      = "%[7]s"
  }
}`, dsfSecretManagerResourceType, resourceName, testAdminEmail, gatewayId, assetId, authMechanism, roleName, testOnPremServerHostName)
}
