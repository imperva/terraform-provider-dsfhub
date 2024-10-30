package dsfhub

import "fmt"

// Output a terraform config for an AWS cloud account resource.
//
// Supports all authentication mechanisms: "key", "profile", "iam_role", and
// "default".
func testAccDSFCloudAccountConfig_Aws(resourceName string, gatewayId string, assetId string, authMechanism string) string {
	var assetConnectionBlock string

	if authMechanism == "key" {
		assetConnectionBlock = fmt.Sprintf(`
      asset_connection {
        access_id = "my-access-id"
        auth_mechanism = "` + authMechanism + `"
        reason = "default"
        region = "us-east-1"
        secret_key = "my-secret-key"
      }
    `)
	} else if authMechanism == "profile" {
		assetConnectionBlock = fmt.Sprintf(`
      asset_connection {
        auth_mechanism = "` + authMechanism + `"
        reason = "default"
        region = "us-east-2"
        username = "dsfhubuser"
      }
    `)
	} else {
		assetConnectionBlock = fmt.Sprintf(`
      asset_connection {
        auth_mechanism = "` + authMechanism + `"
        reason = "default"
        region = "us-west-1"
      }
      `)
	}

	return fmt.Sprintf(`
resource "`+dsfCloudAccountResourceType+`" "%[1]s" {
  server_type = "AWS"
  
  admin_email = "`+testAdminEmail+`"
  arn = "%[3]s"
  asset_display_name = "%[3]s"
  asset_id = "%[3]s"
  gateway_id = "%[2]s"

  `+assetConnectionBlock+`
}`, resourceName, gatewayId, assetId)
}
