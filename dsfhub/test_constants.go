package dsfhub

const (
	testAdminEmail           = "test@email.com"
	testArn                  = "arn:aws:rds:us-east-2:123456789:db:your-db"
	testServerHostName       = "your-db-name.abcde12345.us-east-2.rds.amazonaws.com"
	testOnPremServerHostName = "server.company.com"
	testServerIP             = "1.2.3.4"
	testServerPort           = "8200"

	testInvalidDSFHUBHost = "https://invalid.host.com"
	testDSServerType      = "AWS RDS MYSQL"
	testGatewayId         = "e33bfbe4-a93a-c4e5-8e9c-6e5558c2e2cd"

	testAssetDisplayName = "arn:aws:rds:us-east-2:123456789:db:your-db"

	testSMConnectionReason = "default"
	testSMRoleName         = "vault-role-for-ec2"
	testSMAuthMechanism    = "ec2"
	testSMAssetId          = "your-host-name-here"

	testSMServerType = "HASHICORP"

	testAwsAccountId             = "123456789012"
	testAwsKinesisPrefix         = "arn:aws:kinesis:us-east-1:" + testAwsAccountId + ":stream/"
	testAwsLogGroupPrefix        = "arn:aws:logs:us-east-1:" + testAwsAccountId + ":log-group:"
	testAwsRdsPrefix             = "arn:aws:rds:us-east-2:123456789012"
	testAwsRdsClusterPrefix      = testAwsRdsPrefix + ":cluster:"
	testAwsRdsDbPrefix           = testAwsRdsPrefix + ":db:"
	testAwsS3BucketPrefix        = "arn:aws:s3:::"
	testAzurePrefix              = "/subscriptions/my-subscription-id/resourceGroups/my-resource-group/providers/"
	testEventhubPrefix           = testAzurePrefix + "Microsoft.EventHub/namespaces/my-namespace/eventhubs/"
	testPubsubSubscriptionPrefix = "projects/my-project/subscriptions/"
)
