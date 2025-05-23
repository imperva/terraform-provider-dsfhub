package dsfhub

import (
	"fmt"
	"log"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDSFLogAggregator_AwsKinesis(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "aws_kinesis_basic"
		assetId      = testAwsKinesisPrefix + resourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, resourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// Failed: bad audit_type
			{
				Config:      testAccDSFLogAggregatorConfig_AwsKinesis(resourceName, gatewayId, assetId, "", false, "BAD_AUDIT_TYPE"),
				ExpectError: regexp.MustCompile("Asset attribute options mismatch: the value 'BAD_AUDIT_TYPE' for field 'audit_type' is invalid"),
			},
			// Test various audit types
			{Config: testAccDSFLogAggregatorConfig_AwsKinesis(resourceName, gatewayId, assetId, "", false, "KINESIS")},
			{Config: testAccDSFLogAggregatorConfig_AwsKinesis(resourceName, gatewayId, assetId, "", false, "KINESIS_AGGREGATED")},
			// Validate import
			{
				ResourceName:      resourceTypeAndName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDSFLogAggregator_AwsLogGroup(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		assetId            = testAwsLogGroupPrefix + "/aws/rds/instance/my-database/audit:*"
		resourceName       = "my-database-log-group"
		serverHostName     = "oracle-rds-db.xxxxx8rsfzja.us-east-2.rds.amazonaws.com"
		parentAssetId      = "arn:aws:rds:us-east-2:123456789012:db:oracle-rds-db"
		parentResourceName = "my-oracle-db"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, resourceName)
	parentResourceTypeAndName := fmt.Sprintf("%s.%s", dsfDataSourceResourceType, parentResourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// Failed: missing parent_asset_id
			{
				Config:      testAccDSFLogAggregatorConfig_AwsLogGroup(resourceName, gatewayId, assetId, "", true, "LOG_GROUP", ""),
				ExpectError: regexp.MustCompile("Error: missing required fields for dsfhub_data_source"),
			},
			// Onboard with AWS parent asset
			{
				Config: ConfigCompose(
					testAccDSFDataSourceConfig_AwsRdsOracle(parentResourceName, gatewayId, parentAssetId, "LOG_GROUP", ""),
					testAccDSFLogAggregatorConfig_AwsLogGroup(resourceName, gatewayId, assetId, parentResourceTypeAndName+".asset_id", true, "LOG_GROUP", ""),
				),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "gateway_service", "gateway-aws@oracle-rds.service"),
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

func TestAccDSFLogAggregator_AwsS3(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		assetId      = testAwsS3BucketPrefix + "my-s3-bucket"
		resourceName = "my-s3-bucket"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, resourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// Test various audit types
			{Config: testAccDSFLogAggregatorConfig_AwsS3(resourceName, gatewayId, assetId, "", "false", "")},
			{Config: testAccDSFLogAggregatorConfig_AwsS3(resourceName, gatewayId, assetId, "", "false", "LOG_GROUP")},
			{Config: testAccDSFLogAggregatorConfig_AwsS3(resourceName, gatewayId, assetId, "", "false", "REDSHIFT")},
			{Config: testAccDSFLogAggregatorConfig_AwsS3(resourceName, gatewayId, assetId, "", "false", "CLOUDWATCH")},
			{Config: testAccDSFLogAggregatorConfig_AwsS3(resourceName, gatewayId, assetId, "", "false", "DYNAMODB")},
			// validate import
			{
				ResourceName:      resourceTypeAndName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDSFLogAggregator_AzureEventhub(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		assetId      = "/subscriptions/ID/resourceGroups/someGroup/providers/Microsoft.EventHub/namespaces/somenamespace/eventhubs/someeventhub"
		resourceName = "my-azure-eventhub"
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, resourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// Failed: missing format
			{
				Config:      testAccDSFLogAggregatorConfig_AzureEventhub(resourceName, gatewayId, assetId, "default", "", "true", "", ""),
				ExpectError: regexp.MustCompile("Error: missing required fields for dsfhub_data_source with serverType 'AZURE EVENTHUB', missing fields: \"format\""),
			},
			// Failed: invalid format
			{
				Config:      testAccDSFLogAggregatorConfig_AzureEventhub(resourceName, gatewayId, assetId, "default", "", "true", "", "bad-format"),
				ExpectError: regexp.MustCompile("Asset attribute options mismatch: the value 'bad-format' for field 'format' is invalid"),
			},
			// Validate different auth_mechanisms
			{Config: testAccDSFLogAggregatorConfig_AzureEventhub(resourceName, gatewayId, assetId, "azure_ad", "", "false", "", "AzureSQL_Managed")},
			{Config: testAccDSFLogAggregatorConfig_AzureEventhub(resourceName, gatewayId, assetId, "client_secret", "", "false", "", "AzureSQL_Managed")},
			{Config: testAccDSFLogAggregatorConfig_AzureEventhub(resourceName, gatewayId, assetId, "default", "", "false", "", "AzureSQL_Managed")},
			// validate import
			{
				ResourceName:      resourceTypeAndName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDSFLogAggregator_GcpCloudStorageBucket(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "my-bucket-1"
		assetId      = "my-project:" + resourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, resourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// Failed: missing asset_display_name, asset_id, pubsub_subscription
			{
				Config:      testAccDSFLogAggregatorConfig_GcpCloudStorageBucket(resourceName, gatewayId, "", "", "false", ""),
				ExpectError: regexp.MustCompile("Error: missing required fields for dsfhub_data_source"),
			},
			// Onboard and connect/disconnect to gateway as standalone log aggregator
			{Config: testAccDSFLogAggregatorConfig_GcpCloudStorageBucket(resourceName, gatewayId, assetId, "", "true", "GCP MS SQL SERVER")},
			{Config: testAccDSFLogAggregatorConfig_GcpCloudStorageBucket(resourceName, gatewayId, assetId, "", "false", "GCP MS SQL SERVER")},
			// validate import
			{
				ResourceName:      resourceTypeAndName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccDSFLogAggregator_GcpPubsubBasic(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "my-basic-pubsub-subscription"
		assetId      = testPubsubSubscriptionPrefix + resourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, resourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// Failed: missing asset_display_name, asset_id, pubsub_subscription
			{
				Config:      testAccDSFLogAggregatorConfig_GcpPubsub(resourceName, gatewayId, "", "default", "", "false", "", ""),
				ExpectError: regexp.MustCompile("Error: missing required fields for dsfhub_data_source"),
			},
			// Test different auth_mechanisms
			{
				Config: testAccDSFLogAggregatorConfig_GcpPubsub(resourceName, gatewayId, assetId, "default", "", "false", "", ""),
			},
			{
				Config: testAccDSFLogAggregatorConfig_GcpPubsub(resourceName, gatewayId, assetId, "service_account", "", "false", "", ""),
			},
			// Test connect/disconnect w/ audit_type: "", content_type: ""
			{
				Config: testAccDSFLogAggregatorConfig_GcpPubsub(resourceName, gatewayId, assetId, "default", "", "true", "", ""),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "gateway_service", "gateway-gcp@postgresql.service"),
				),
			},
			{
				Config: testAccDSFLogAggregatorConfig_GcpPubsub(resourceName, gatewayId, assetId, "default", "", "false", "", ""),
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

func TestAccDSFLogAggregator_GcpPubsubAlloydbPostgresql(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "my-alloydb-pubsub-subscription"
		assetId      = testPubsubSubscriptionPrefix + resourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, resourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// Test connect/disconnect
			{
				Config: testAccDSFLogAggregatorConfig_GcpPubsub(resourceName, gatewayId, assetId, "default", "", "true", "ALLOYDB_POSTGRESQL", "GCP ALLOYDB POSTGRESQL"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "gateway_service", "gateway-gcp@alloydb-postgresql.service"),
				),
			},
			{
				Config: testAccDSFLogAggregatorConfig_GcpPubsub(resourceName, gatewayId, assetId, "default", "", "false", "ALLOYDB_POSTGRESQL", "GCP ALLOYDB POSTGRESQL"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "false"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "gateway_service", ""),
				),
			},
			{
				Config: testAccDSFLogAggregatorConfig_GcpPubsub(resourceName, gatewayId, assetId, "default", "", "true", "ALLOYDB_POSTGRESQL", "GCP ALLOYDB POSTGRESQL CLUSTER"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "gateway_service", "gateway-gcp@alloydb-postgresql.service"),
				),
			},
			{
				Config: testAccDSFLogAggregatorConfig_GcpPubsub(resourceName, gatewayId, assetId, "default", "", "false", "ALLOYDB_POSTGRESQL", "GCP ALLOYDB POSTGRESQL CLUSTER"),
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

func TestAccDSFLogAggregator_GcpPubsubBigQuery(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "my-bigquery-pubsub-subscription"
		assetId      = testPubsubSubscriptionPrefix + resourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, resourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// Test connect/disconnect
			{
				Config: testAccDSFLogAggregatorConfig_GcpPubsub(resourceName, gatewayId, assetId, "default", "", "true", "BIGQUERY", "GCP BIGQUERY"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "gateway_service", "gateway-gcp@bigquery.service"),
				),
			},
			{
				Config: testAccDSFLogAggregatorConfig_GcpPubsub(resourceName, gatewayId, assetId, "default", "", "false", "BIGQUERY", "GCP BIGQUERY"),
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

func TestAccDSFLogAggregator_GcpPubsubBigTable(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "my-bigtable-pubsub-subscription"
		assetId      = testPubsubSubscriptionPrefix + resourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, resourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// Test connect/disconnect
			{
				Config: testAccDSFLogAggregatorConfig_GcpPubsub(resourceName, gatewayId, assetId, "default", "", "true", "BIGTABLE", "GCP BIGTABLE"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "gateway_service", "gateway-gcp@bigtable.service"),
				),
			},
			{
				Config: testAccDSFLogAggregatorConfig_GcpPubsub(resourceName, gatewayId, assetId, "default", "", "false", "BIGTABLE", "GCP BIGTABLE"),
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

func TestAccDSFLogAggregator_GcpPubsubMssqlserver(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "my-mssql-server-pubsub-subscription"
		assetId      = testPubsubSubscriptionPrefix + resourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, resourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// Test connect/disconnect
			{
				Config: testAccDSFLogAggregatorConfig_GcpPubsub(resourceName, gatewayId, assetId, "default", "", "true", "MSSQL", "GCP MS SQL SERVER"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "gateway_service", "gateway-gcp@mssql.service"),
				),
			},
			{
				Config: testAccDSFLogAggregatorConfig_GcpPubsub(resourceName, gatewayId, assetId, "default", "", "false", "MSSQL", "GCP MS SQL SERVER"),
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

func TestAccDSFLogAggregator_GcpPubsubMysql(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "my-mysql-pubsub-subscription"
		assetId      = testPubsubSubscriptionPrefix + resourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, resourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// Test connect/disconnect
			{
				Config: testAccDSFLogAggregatorConfig_GcpPubsub(resourceName, gatewayId, assetId, "default", "", "true", "MYSQL", "GCP MYSQL"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "gateway_service", "gateway-gcp@mysql.service"),
				),
			},
			{
				Config: testAccDSFLogAggregatorConfig_GcpPubsub(resourceName, gatewayId, assetId, "default", "", "false", "MYSQL", "GCP MYSQL"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "false"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "gateway_service", ""),
				),
			},
			{
				Config: testAccDSFLogAggregatorConfig_GcpPubsub(resourceName, gatewayId, assetId, "default", "", "true", "GCP_MYSQL_SLOW", "GCP MYSQL"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "gateway_service", "gateway-gcp@mysql-slow-query.service"),
				),
			},
			{
				Config: testAccDSFLogAggregatorConfig_GcpPubsub(resourceName, gatewayId, assetId, "default", "", "false", "GCP_MYSQL_SLOW", "GCP MYSQL"),
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

func TestAccDSFLogAggregator_GcpPubsubPostgresql(t *testing.T) {
	gatewayId := checkGatewayId(t)

	const (
		resourceName = "my-postgresql-pubsub-subscription"
		assetId      = testPubsubSubscriptionPrefix + resourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, resourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// Test connect/disconnect
			{
				Config: testAccDSFLogAggregatorConfig_GcpPubsub(resourceName, gatewayId, assetId, "default", "", "true", "POSTGRESQL", "GCP POSTGRESQL"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "gateway_service", "gateway-gcp@postgresql.service"),
				),
			},
			{
				Config: testAccDSFLogAggregatorConfig_GcpPubsub(resourceName, gatewayId, assetId, "default", "", "false", "POSTGRESQL", "GCP POSTGRESQL"),
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

func TestAccDSFLogAggregator_GcpPubsubSpanner(t *testing.T) {
	gatewayId := checkGatewayId(t)

	skipTestForKnownIssue(t, "4.17", "SR-2063")

	const (
		resourceName = "my-spanner-pubsub-subscription"
		assetId      = testPubsubSubscriptionPrefix + resourceName
	)

	resourceTypeAndName := fmt.Sprintf("%s.%s", dsfLogAggregatorResourceType, resourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// Test connect/disconnect
			{
				Config: testAccDSFLogAggregatorConfig_GcpPubsub(resourceName, gatewayId, assetId, "default", "", "true", "SPANNER", "GCP SPANNER"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceTypeAndName, "audit_pull_enabled", "true"),
					resource.TestCheckResourceAttr(resourceTypeAndName, "gateway_service", "gateway-gcp@spanner.service"),
				),
			},
			{
				Config: testAccDSFLogAggregatorConfig_GcpPubsub(resourceName, gatewayId, assetId, "default", "", "false", "SPANNER", "GCP SPANNER"),
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

func testAccLogAggregatorId(state *terraform.State) (string, error) {
	log.Printf("[INFO] Running test testAccLogAggregatorId \n")
	for _, rs := range state.RootModule().Resources {
		if rs.Type != dsfLogAggregatorResourceType {
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
