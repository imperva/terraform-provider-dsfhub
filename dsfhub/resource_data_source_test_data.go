package dsfhub

import "fmt"

// Output a terraform config for a basic data source resource.
func testAccDSFDataSourceConfig_Basic(resourceName string, adminEmail string, assetId string, gatewayId string, serverHostName string, serverType string) string {
	return fmt.Sprintf(`
resource "`+dsfDataSourceResourceType+`" "%[1]s" {
  admin_email        = "%[2]s"
  asset_id           = "%[3]s"
  asset_display_name = "%[3]s"
  gateway_id         = "%[4]s"
  server_host_name   = "%[5]s"
  server_type        = "%[6]s"
}`, resourceName, adminEmail, assetId, gatewayId, serverHostName, serverType)
}

const commonBasicConnectionPassword = `
  asset_connection {
    auth_mechanism = "password"
    password       = "password"
    reason         = "default"
    username       = "username"
  }
`

// Output a terraform config for an AWS RDS ORACLE data source resource.
func testAccDSFDataSourceConfig_AwsRdsOracle(resourceName string, gatewayId string, assetId string, auditType string, auditPullEnabled string) string {
	// convert audit_pull_enabled to "null" if empty
	if auditPullEnabled == "" {
		auditPullEnabled = "null"
	}

	return fmt.Sprintf(`
resource "`+dsfDataSourceResourceType+`" "%[1]s" {
  server_type        = "AWS RDS ORACLE"

  admin_email        = "`+testAdminEmail+`"
  asset_display_name = "%[3]s"
  asset_id           = "%[3]s"
  audit_pull_enabled = %[5]s
  audit_type         = "%[4]s"
  gateway_id         = "%[2]s"
  server_host_name   = "test.com"
  server_port	       = "1521"
  service_name       = "ORCL"

  `+commonBasicConnectionPassword+`
}
`, resourceName, gatewayId, assetId, auditType, auditPullEnabled)
}

// Output a terraform config for an AWS RDS AURORA POSTGRESQL CLUSTER data
// source resource.
func testAccDSFDataSourceConfig_AwsRdsAuroraPostgresqlCluster(resourceName string, gatewayId string, assetId string, auditType string, clusterId string) string {
	return fmt.Sprintf(`
resource "`+dsfDataSourceResourceType+`" "%[1]s" {
  server_type        = "AWS RDS AURORA POSTGRESQL CLUSTER"

  admin_email	       = "`+testAdminEmail+`"
  asset_display_name = "%[3]s"
  asset_id           = "%[3]s"
  audit_type         = "%[4]s"
  cluster_id         = "%[5]s"
  cluster_name       = "%[5]s"
  gateway_id         = "%[2]s"
  region             = "us-east-2"
  server_host_name   = "my-cluster.cluster-xxxxk8rsfzja.us-east-2.rds.amazonaws.com"
  server_port        = "5432"

  `+commonBasicConnectionPassword+`
}	
`, resourceName, gatewayId, assetId, auditType, clusterId)
}

// Output a terraform config for an AWS RDS AURORA POSTGRESQL data source
// reource.
func testAccDSFDataSourceConfig_AwsRdsAuroraPostgresql(resourceName string, gatewayId string, assetId string, clusterId string) string {
	return fmt.Sprintf(`
resource "`+dsfDataSourceResourceType+`" "%[1]s" {
  server_type        = "AWS RDS AURORA POSTGRESQL"

  admin_email        = "`+testAdminEmail+`"
  asset_display_name = "%[3]s"
  asset_id           = "%[3]s"
  cluster_id         = "%[4]s"
  cluster_name       = "%[4]s"
  gateway_id         = "%[2]s"
  region             = "us-east-2"
  server_host_name   = "my-cluster.cluster-xxxxk8rsfzja.us-east-2.rds.amazonaws.com"
  server_port        = "5432"

  `+commonBasicConnectionPassword+`
}	
`, resourceName, gatewayId, assetId, clusterId)
}

// Output a terraform config for an AWS RDS AURORA MYSQL CLUSTER data source
// resource.
func testAccDSFDataSourceConfig_AwsRdsAuroraMysqlCluster(resourceName string, gatewayId string, assetId string, auditType string, clusterId string) string {
	return fmt.Sprintf(`
resource "`+dsfDataSourceResourceType+`" "%[1]s" {
  server_type        = "AWS RDS AURORA MYSQL CLUSTER"

  admin_email        = "`+testAdminEmail+`"
  asset_display_name = "%[3]s"
  asset_id           = "%[3]s"
  audit_type         = "%[4]s"
  cluster_id         = "%[5]s"
  cluster_name       = "%[5]s"
  gateway_id         = "%[2]s"
  region             = "us-east-2"
  server_host_name   = "my-cluster.cluster-xxxxk8rsfzja.us-east-2.rds.amazonaws.com"
  server_port        = "3306"

  `+commonBasicConnectionPassword+`
}	
`, resourceName, gatewayId, assetId, auditType, clusterId)
}

// Output a terraform config for an AWS RDS AURORA MYSQL data source resource.
func testAccDSFDataSourceConfig_AwsRdsAuroraMysql(resourceName string, gatewayId string, assetId string, clusterId string) string {
	return fmt.Sprintf(`
resource "`+dsfDataSourceResourceType+`" "%[1]s" {
  server_type        = "AWS RDS AURORA MYSQL"

  admin_email        = "`+testAdminEmail+`"
  asset_display_name = "%[3]s"
  asset_id           = "%[3]s"
  #TODO: re-add cluster fields when supported by USC: https://onejira.imperva.com/browse/USC-2389
  #cluster_id         = "%[4]s"
  #cluster_name       = "%[4]s"
  gateway_id          = "%[2]s"
  region              = "us-east-2"
  server_host_name    = "my-cluster.cluster-xxxxk8rsfzja.us-east-2.rds.amazonaws.com"
  server_port         = "5432"

  `+commonBasicConnectionPassword+`
}	
`, resourceName, gatewayId, assetId, clusterId)
}

// Output a terraform config for a GCP MYSQL data source resource.
func testAccDSFDataSourceConfig_GcpMysql(resourceName string, gatewayId string, assetId string, auditPullEnabled string, logsDestinationAssetId string) string {
	// handle reference to other assets
	logsDestinationAssetIdVal := testAccParseResourceAttributeReference(logsDestinationAssetId)

	// convert audit_pull_enabled to "null" if empty
	if auditPullEnabled == "" {
		auditPullEnabled = "null"
	}

	return fmt.Sprintf(`
resource "`+dsfDataSourceResourceType+`" "%[1]s" {
  server_type               = "GCP MYSQL"

  admin_email               = "`+testAdminEmail+`"
  asset_display_name        = "%[3]s"
  asset_id                  = "%[3]s"
  audit_pull_enabled        = %[4]s
  gateway_id                = "%[2]s"
  logs_destination_asset_id = `+logsDestinationAssetIdVal+`
  server_host_name          = "4.3.2.1"
  server_ip                 = "1.2.3.4"
  server_port               = "3306"
}
  `, resourceName, gatewayId, assetId, auditPullEnabled)
}
