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

func TestAccDSFDataSource_Basic(t *testing.T) {
	gatewayId := os.Getenv("GATEWAY_ID")
	if gatewayId == "" {
		t.Skip("GATEWAY_ID environment variable must be set")
	}

	const resourceName = "basic_test_data_source"
	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)

	resource.ParallelTest(t, resource.TestCase{
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

	resource.ParallelTest(t, resource.TestCase{
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

func TestAccDSFDataSource_AwsRdsAuroraPostgresqlClusterCloudWatch(t *testing.T) {
	gatewayId := os.Getenv("GATEWAY_ID")
	if gatewayId == "" {
		t.Skip("GATEWAY_ID environment variable must be set")
	}

	const (
		assetId      = "arn:aws:rds:us-east-2:123456789012:cluster:my-aurorapostgresql-cluster"
		resourceName = "aurora_postgresql_cluster_onboarding"

		instanceAssetId      = assetId + "-writer"
		instanceResourceName = resourceName + "_instance"

		logGroupAssetId      = "arn:aws:logs:us-east-2:123456789012:log-group:/aws/rds/cluster/my-cluster/postgresql:*"
		logGroupResourceName = resourceName + "_log_group"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	instanceResourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, instanceResourceName)
	logGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, logGroupResourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// onboard and connect to gateway, check that the log group is connected
			{
				Config: ConfigCompose(
					testAccDSFDataSourceConfig_AwsRdsAuroraPostgresqlCluster(resourceName, gatewayId, assetId, "LOG_GROUP", resourceName),
					testAccDSFDataSourceConfig_AwsRdsAuroraPostgresql(instanceResourceName, gatewayId, instanceAssetId, resourceName),
					testAccDSFLogAggregatorConfig_AwsLogGroup(logGroupResourceName, gatewayId, logGroupAssetId, resourceTypeAndName+".asset_id", true, "LOG_GROUP", ""),
				),
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

func TestAccDSFDataSource_AwsRdsAuroraMysqlClusterCloudWatchSlowQuery(t *testing.T) {
	gatewayId := os.Getenv("GATEWAY_ID")
	if gatewayId == "" {
		t.Skip("GATEWAY_ID environment variable must be set")
	}

	const (
		assetId      = "arn:aws:rds:us-east-2:123456789012:cluster:my-auroramysql-cluster"
		resourceName = "aurora_mysql_cluster_onboarding"

		instanceAssetId      = assetId + "-writer"
		instanceResourceName = resourceName + "_instance"

		logGroupAssetId      = "arn:aws:logs:us-east-2:123456789012:log-group:/aws/rds/cluster/my-aurora-cluster/audit:*"
		logGroupResourceName = resourceName + "_log_group"

		slowLogGroupAssetId      = "arn:aws:logs:us-east-2:123456789012:log-group:/aws/rds/cluster/my-aurora-cluster/slowquery:*"
		slowLogGroupResourceName = resourceName + "_slow_log_group"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	//TODO: check that instance asset is connected once fixed: https://onejira.imperva.com/browse/SR-2046
	// instanceResourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, instanceResourceName)
	logGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, logGroupResourceName)
	slowLogGroupResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, slowLogGroupResourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// onboard and connect to gateway
			{
				Config: ConfigCompose(
					testAccDSFDataSourceConfig_AwsRdsAuroraMysqlCluster(resourceName, gatewayId, assetId, "", resourceName),
					testAccDSFDataSourceConfig_AwsRdsAuroraMysql(instanceResourceName, gatewayId, instanceAssetId, resourceName),
					testAccDSFLogAggregatorConfig_AwsLogGroup(logGroupResourceName, gatewayId, logGroupAssetId, resourceTypeAndName+".asset_id", true, "LOG_GROUP", ""),
					testAccDSFLogAggregatorConfig_AwsLogGroup(slowLogGroupResourceName, gatewayId, slowLogGroupAssetId, resourceTypeAndName+".asset_id", true, "AWS_RDS_AURORA_MYSQL_SLOW", logGroupResourceTypeAndName),
				),
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
			// disconnect assets
			{
				Config: ConfigCompose(
					testAccDSFDataSourceConfig_AwsRdsAuroraMysqlCluster(resourceName, gatewayId, assetId, "", resourceName),
					testAccDSFDataSourceConfig_AwsRdsAuroraMysql(instanceResourceName, gatewayId, instanceAssetId, resourceName),
					testAccDSFLogAggregatorConfig_AwsLogGroup(logGroupResourceName, gatewayId, logGroupAssetId, resourceTypeAndName+".asset_id", false, "LOG_GROUP", ""),
					testAccDSFLogAggregatorConfig_AwsLogGroup(slowLogGroupResourceName, gatewayId, slowLogGroupAssetId, resourceTypeAndName+".asset_id", false, "AWS_RDS_AURORA_MYSQL_SLOW", logGroupResourceTypeAndName),
				),
				// verify log group assets are disconnected
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(logGroupResourceTypeAndName, "audit_pull_enabled", "false"),
					resource.TestCheckResourceAttr(logGroupResourceTypeAndName, "gateway_service", ""),
					resource.TestCheckResourceAttr(slowLogGroupResourceTypeAndName, "audit_pull_enabled", "false"),
					resource.TestCheckResourceAttr(slowLogGroupResourceTypeAndName, "gateway_service", ""),
				),
			},
			// refresh and verify DB assets are disconnected
			{
				RefreshState: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "false"),
					// resource.TestCheckResourceAttr(instanceResourceTypeAndName, "audit_pull_enabled", "false"),
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

func TestAccDSFDataSource_GcpMysql(t *testing.T) {
	gatewayId := os.Getenv("GATEWAY_ID")
	if gatewayId == "" {
		t.Skip("GATEWAY_ID environment variable must be set")
	}

	const (
		resourceName = "gcp_mysql_connect_disconnect_gateway"
		assetId      = "my-project:us-west-1:mysql-instance-1"

		pubsubResourceName = "mysql-instance-1-subscription"
		pubsubAssetId      = testPubsubSubscriptionPrefix + pubsubResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	pubsubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, pubsubResourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// onboard and connect to gateway, check that the DB asset is connected
			{
				Config: ConfigCompose(
					testAccDSFDataSourceConfig_GcpMysql(resourceName, gatewayId, assetId, "true", pubsubResourceTypeAndName+".asset_id"),
					testAccDSFLogAggregatorConfig_GcpPubsub(pubsubResourceName, gatewayId, pubsubAssetId, "default", "", "", "MYSQL", ""),
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
				),
			},
			// refresh and verify pubsub aset is connected
			{
				RefreshState: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(pubsubResourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(pubsubResourceTypeAndName, "gateway_service", "gateway-gcp@mysql.service"),
				),
			},
			// disconnect asset, check that the DB asset is disconnected
			{
				Config: ConfigCompose(
					testAccDSFDataSourceConfig_GcpMysql(resourceName, gatewayId, assetId, "false", pubsubResourceTypeAndName+".asset_id"),
					testAccDSFLogAggregatorConfig_GcpPubsub(pubsubResourceName, gatewayId, pubsubAssetId, "default", "", "", "MYSQL", ""),
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "false"),
				),
			},
			// refresh and verify pubsub aset is disconnected
			{
				RefreshState: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "false"),
					resource.TestCheckResourceAttr(pubsubResourceTypeAndName, "audit_pull_enabled", "false"),
					resource.TestCheckResourceAttr(pubsubResourceTypeAndName, "gateway_service", ""),
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

func TestAccDSFDataSource_GcpMysqlSlowQuery(t *testing.T) {
	gatewayId := os.Getenv("GATEWAY_ID")
	if gatewayId == "" {
		t.Skip("GATEWAY_ID environment variable must be set")
	}

	const (
		resourceName = "gcp_mysql_slow_query"
		assetId      = "my-project:us-west-1:mysql-instance-2"

		auditPubsubResourceName = "mysql-instance-2-audit-subscription"
		auditPubsubAssetId      = testPubsubSubscriptionPrefix + auditPubsubResourceName

		slowQueryPubsubResourceName = "mysql-instance-2-slow-query-subscription"
		slowQueryPubsubAssetId      = testPubsubSubscriptionPrefix + slowQueryPubsubResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	auditPubsubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, auditPubsubResourceName)
	slowQueryPubsubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, slowQueryPubsubResourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// onboard and connect to gateway, check that the DB asset is connected
			{
				Config: ConfigCompose(
					testAccDSFDataSourceConfig_GcpMysql(resourceName, gatewayId, assetId, "true", auditPubsubResourceTypeAndName+".asset_id"),
					testAccDSFLogAggregatorConfig_GcpPubsub(auditPubsubResourceName, gatewayId, auditPubsubAssetId, "default", "", "", "MYSQL", ""),
					testAccDSFLogAggregatorConfig_GcpPubsub(slowQueryPubsubResourceName, gatewayId, slowQueryPubsubAssetId, "default", "", "true", "GCP_MYSQL_SLOW", "GCP MYSQL"),
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
				),
			},
			// refresh and verify pubsub asets are connected
			{
				RefreshState: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(auditPubsubResourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(auditPubsubResourceTypeAndName, "gateway_service", "gateway-gcp@mysql.service"),
					resource.TestCheckResourceAttr(slowQueryPubsubResourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(slowQueryPubsubResourceTypeAndName, "gateway_service", "gateway-gcp@mysql-slow-query.service"),
				),
			},
			// disconnect asset, check that the DB asset is disconnected
			{
				Config: ConfigCompose(
					testAccDSFDataSourceConfig_GcpMysql(resourceName, gatewayId, assetId, "false", auditPubsubResourceTypeAndName+".asset_id"),
					testAccDSFLogAggregatorConfig_GcpPubsub(auditPubsubResourceName, gatewayId, auditPubsubAssetId, "default", "", "", "MYSQL", ""),
					testAccDSFLogAggregatorConfig_GcpPubsub(slowQueryPubsubResourceName, gatewayId, slowQueryPubsubAssetId, "default", "", "false", "GCP_MYSQL_SLOW", "GCP MYSQL"),
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "false"),
				),
			},
			// refresh and verify pubsub asets are disconnected
			{
				RefreshState: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "false"),
					resource.TestCheckResourceAttr(auditPubsubResourceTypeAndName, "audit_pull_enabled", "false"),
					resource.TestCheckResourceAttr(auditPubsubResourceTypeAndName, "gateway_service", ""),
					resource.TestCheckResourceAttr(slowQueryPubsubResourceTypeAndName, "audit_pull_enabled", "false"),
					resource.TestCheckResourceAttr(slowQueryPubsubResourceTypeAndName, "gateway_service", ""),
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

func TestAccDSFDataSource_GcpPostgresql(t *testing.T) {
	gatewayId := os.Getenv("GATEWAY_ID")
	if gatewayId == "" {
		t.Skip("GATEWAY_ID environment variable must be set")
	}

	const (
		resourceName = "gcp_postgresql_connect_disconnect_gateway"
		assetId      = "my-project:us-west-1:postgresql-instance-1"

		pubsubResourceName = "postgresql-instance-1-subscription"
		pubsubAssetId      = testPubsubSubscriptionPrefix + pubsubResourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, resourceName)
	pubsubResourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, pubsubResourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// onboard and connect to gateway, check that the DB asset is connected
			{
				Config: ConfigCompose(
					testAccDSFDataSourceConfig_GcpPostgresql(resourceName, gatewayId, assetId, "true", pubsubResourceTypeAndName+".asset_id"),
					testAccDSFLogAggregatorConfig_GcpPubsub(pubsubResourceName, gatewayId, pubsubAssetId, "default", "", "", "POSTGRESQL", ""),
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
				),
			},
			// refresh and verify pubsub aset is connected
			{
				RefreshState: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(pubsubResourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(pubsubResourceTypeAndName, "gateway_service", "gateway-gcp@postgresql.service"),
				),
			},
			// disconnect asset, check that the DB asset is disconnected
			{
				Config: ConfigCompose(
					testAccDSFDataSourceConfig_GcpPostgresql(resourceName, gatewayId, assetId, "false", pubsubResourceTypeAndName+".asset_id"),
					testAccDSFLogAggregatorConfig_GcpPubsub(pubsubResourceName, gatewayId, pubsubAssetId, "default", "", "", "POSTGRESQL", ""),
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "false"),
				),
			},
			// refresh and verify pubsub aset is disconnected
			{
				RefreshState: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "false"),
					resource.TestCheckResourceAttr(pubsubResourceTypeAndName, "audit_pull_enabled", "false"),
					resource.TestCheckResourceAttr(pubsubResourceTypeAndName, "gateway_service", ""),
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
