package dsfhub

import "fmt"

const awsConnectionDefault = `
  asset_connection {
    auth_mechanism  = "default"
    reason          = "default"
    region          = "us-west-1"
  }
`

const awsConnectionKey = `
  asset_connection {
    access_id       = "my-access-id"
    auth_mechanism  = "key"
    reason          = "default"
    region          = "us-east-1"
    secret_key      = "my-secret-key"
  }
`

const awsConnectionProfile = `
  asset_connection {
    auth_mechanism  = "profile"
    reason          = "default"
    region          = "us-east-2"
    username        = "dsfhubuser"
  }
`

// Output a terraform config for an AWS cloud account resource.
//
// Supports all authentication mechanisms: "key", "profile", "iam_role", and
// "default".
func testAccDSFCloudAccountConfig_Aws(resourceName string, gatewayId string, assetId string, authMechanism string) string {
	var assetConnectionBlock string

  switch authMechanism {
    case "key":
      assetConnectionBlock = awsConnectionKey
    case "profile":
      assetConnectionBlock = awsConnectionProfile
    default:
      assetConnectionBlock = awsConnectionDefault
  }

	return fmt.Sprintf(`
resource "`+dsfCloudAccountResourceType+`" "%[1]s" {
  server_type = "AWS"
  
  admin_email        = "`+testAdminEmail+`"
  arn                = "%[3]s"
  asset_display_name = "%[3]s"
  asset_id           = "%[3]s"
  gateway_id         = "%[2]s"

  `+assetConnectionBlock+`
}`, resourceName, gatewayId, assetId)
}
