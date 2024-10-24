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
		t.Skip("GATEWAY_ID environment variable must be set")
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
		t.Skip("GATEWAY_ID environment variable must be set")
	}

	const resourceName = "rds_oracle_connect_disconnect_gateway"
	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// onboard and connect to gateway
			{
				Config: testAccDSFDataSourceConfig_AwsRdsOracle(resourceName, gatewayId, resourceName, "UNIFIED", "true"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "gateway_service", "gateway-odbc@oracle_unified.service"),
				),
			},
			// update audit_type -> reconnect asset to gateway
			{
				Config: testAccDSFDataSourceConfig_AwsRdsOracle(resourceName, gatewayId, resourceName, "UNIFIED_AGGREGATED", "true"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "gateway_service", "gateway-odbc@oracle_unified_aggregated.service"),
				),
			},
			// disconnect asset
			{
				Config: testAccDSFDataSourceConfig_AwsRdsOracle(resourceName, gatewayId, resourceName, "UNIFIED_AGGREGATED", "false"),
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

func TestAccDSFDataSource_AwsRdsPostgresqlClusterCloudWatch(t *testing.T) {
	gatewayId := os.Getenv("GATEWAY_ID")
	if gatewayId == "" {
		t.Skip("GATEWAY_ID environment variable must be set")
	}

	const (
		assetId = "arn:aws:rds:us-east-2:123456789012:cluster:my-aurorapostgresql-cluster"
		resourceName = "aurora_postgresql_cluster_onboarding"

		instanceAssetId = assetId + "-writer"
		instanceResourceName = resourceName + "_instance"

		logGroupAssetId = "arn:aws:logs:us-east-2:123456789012:log-group:/aws/rds/cluster/my-cluster/postgresql:*"
		logGroupResourceName = resourceName + "_log_group"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	instanceResourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, instanceResourceName)
	logGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, logGroupResourceName)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// onboard and connect to gateway
			{
				Config: testAccDSFDataSourceConfig_AwsRdsAuroraPostgresqlCluster(resourceName, gatewayId, assetId, "LOG_GROUP", resourceName) + 
					testAccDSFDataSourceConfig_AwsRdsAuroraPostgresql(instanceResourceName, gatewayId, instanceAssetId, resourceName) + 
					testAccDSFLogAggregatorConfig_AwsLogGroup(logGroupResourceName, gatewayId, logGroupAssetId, resourceTypeAndName + ".asset_id", true, "LOG_GROUP", ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(logGroupResourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(logGroupResourceTypeAndName, "gateway_service", "gateway-aws@aurora-postgresql.service"),
				),
			},
			// refresh and verify DB assets are connected
			{
				RefreshState: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(instanceResourceTypeAndName, "audit_pull_enabled", "true"),
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

func TestAccDSFDataSource_AwsRdsMysqlClusterCloudWatchSlowQuery(t *testing.T) {
	gatewayId := os.Getenv("GATEWAY_ID")
	if gatewayId == "" {
		t.Skip("GATEWAY_ID environment variable must be set")
	}

	const (
		assetId = "arn:aws:rds:us-east-2:123456789012:cluster:my-auroramysql-cluster"
		resourceName = "aurora_mysql_cluster_onboarding"

		instanceAssetId = assetId + "-writer"
		instanceResourceName = resourceName + "_instance"

		logGroupAssetId = "arn:aws:logs:us-east-2:123456789012:log-group:/aws/rds/cluster/my-aurora-cluster/audit:*"
		logGroupResourceName = resourceName + "_log_group"

		slowLogGroupAssetId = "arn:aws:logs:us-east-2:123456789012:log-group:/aws/rds/cluster/my-aurora-cluster/slowquery:*"
		slowLogGroupResourceName = resourceName + "_slow_log_group"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	// instanceResourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, instanceResourceName)
	logGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, logGroupResourceName)
	slowLogGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, slowLogGroupResourceName)

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// onboard and connect to gateway
			{
				Config: testAccDSFDataSourceConfig_AwsRdsAuroraMysqlCluster(resourceName, gatewayId, assetId, "", resourceName) + 
					testAccDSFDataSourceConfig_AwsRdsAuroraMysql(instanceResourceName, gatewayId, instanceAssetId, resourceName) + 
					testAccDSFLogAggregatorConfig_AwsLogGroup(logGroupResourceName, gatewayId, logGroupAssetId, resourceTypeAndName + ".asset_id", true, "LOG_GROUP", "") +
					testAccDSFLogAggregatorConfig_AwsLogGroup(slowLogGroupResourceName, gatewayId, slowLogGroupAssetId, resourceTypeAndName + ".asset_id", true, "AWS_RDS_AURORA_MYSQL_SLOW", logGroupResourceTypeAndName),
				// verify log group assets are connected
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(logGroupResourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(logGroupResourceTypeAndName, "gateway_service", "gateway-aws@aurora-mysql.service"),
					resource.TestCheckResourceAttr(slowLogGroupResourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(slowLogGroupResourceTypeAndName, "gateway_service", "gateway-aws@aurora-mysql-slow-query.service"),
				),
			},
			// refresh and verify DB assets are connected
			{
				RefreshState: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					// resource.TestCheckResourceAttr(instanceResourceTypeAndName, "audit_pull_enabled", "true"),
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

func testAccDSFDataSourceConfig_AwsRdsOracle(resourceName string, gatewayId string, assetId string, auditType string, auditPullEnabled string) string {
	// convert audit_pull_enabled to "null" if empty
	if auditPullEnabled == "" {
		auditPullEnabled = "null"
	}
	
	return fmt.Sprintf(`
resource "` + dsfDataSourceResourceType + `" "%[1]s" {
	server_type					= "AWS RDS ORACLE"

	admin_email					= "` + testAdminEmail + `"
	asset_display_name	= "%[3]s"
	asset_id						= "%[3]s"
	audit_pull_enabled	= %[5]s
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

func testAccDSFDataSourceConfig_AwsRdsAuroraPostgresqlCluster(resourceName string, gatewayId string, assetId string, auditType string, clusterId string) string {
	return fmt.Sprintf(`
resource "` + dsfDataSourceResourceType + `" "%[1]s" {
	server_type = "AWS RDS AURORA POSTGRESQL CLUSTER"

	admin_email					= "` + testAdminEmail + `"
	asset_display_name = "%[3]s"
  asset_id           = "%[3]s"
  audit_type         = "%[4]s"
	cluster_id         = "%[5]s"
	cluster_name       = "%[5]s"
  gateway_id         = "%[2]s"
  region             = "us-east-2"
  server_host_name   = "my-cluster.cluster-xxxxk8rsfzja.us-east-2.rds.amazonaws.com"
  server_port        = "5432"

	asset_connection {
		auth_mechanism = "password"
		password = "my-password"
		reason = "default"
		username = "my-user"
	}
}	
`,
	resourceName, gatewayId, assetId, auditType, clusterId)
}

func testAccDSFDataSourceConfig_AwsRdsAuroraPostgresql(resourceName string, gatewayId string, assetId string, clusterId string) string {
	return fmt.Sprintf(`
resource "` + dsfDataSourceResourceType + `" "%[1]s" {
	server_type = "AWS RDS AURORA POSTGRESQL"

	admin_email					= "` + testAdminEmail + `"
	asset_display_name = "%[3]s"
  asset_id           = "%[3]s"
	cluster_id         = "%[4]s"
	cluster_name       = "%[4]s"
  gateway_id         = "%[2]s"
  region             = "us-east-2"
  server_host_name   = "my-cluster.cluster-xxxxk8rsfzja.us-east-2.rds.amazonaws.com"
  server_port        = "5432"

	asset_connection {
		auth_mechanism = "password"
		password = "my-password"
		reason = "default"
		username = "my-user"
	}
}	
`,
	resourceName, gatewayId, assetId, clusterId)
}

func testAccDSFDataSourceConfig_AwsRdsAuroraMysqlCluster(resourceName string, gatewayId string, assetId string, auditType string, clusterId string) string {
	return fmt.Sprintf(`
resource "` + dsfDataSourceResourceType + `" "%[1]s" {
	server_type = "AWS RDS AURORA MYSQL CLUSTER"

	admin_email					= "` + testAdminEmail + `"
	asset_display_name = "%[3]s"
  asset_id           = "%[3]s"
  audit_type         = "%[4]s"
	cluster_id         = "%[5]s"
	cluster_name       = "%[5]s"
  gateway_id         = "%[2]s"
  region             = "us-east-2"
  server_host_name   = "my-cluster.cluster-xxxxk8rsfzja.us-east-2.rds.amazonaws.com"
  server_port        = "3306"

	asset_connection {
		auth_mechanism = "password"
		password = "my-password"
		reason = "default"
		username = "my-user"
	}
}	
`,
	resourceName, gatewayId, assetId, auditType, clusterId)
}

func testAccDSFDataSourceConfig_AwsRdsAuroraMysql(resourceName string, gatewayId string, assetId string, clusterId string) string {
	return fmt.Sprintf(`
resource "` + dsfDataSourceResourceType + `" "%[1]s" {
	server_type = "AWS RDS AURORA MYSQL"

	admin_email					= "` + testAdminEmail + `"
	asset_display_name = "%[3]s"
  asset_id           = "%[3]s"
	#TODO: re-add cluster fields when supported by USC: https://onejira.imperva.com/browse/USC-2389
	#cluster_id         = "%[4]s"
	#cluster_name       = "%[4]s"
  gateway_id         = "%[2]s"
  region             = "us-east-2"
  server_host_name   = "my-cluster.cluster-xxxxk8rsfzja.us-east-2.rds.amazonaws.com"
  server_port        = "5432"

	asset_connection {
		auth_mechanism = "password"
		password = "my-password"
		reason = "default"
		username = "my-user"
	}
}	
`,
	resourceName, gatewayId, assetId, clusterId)
}
