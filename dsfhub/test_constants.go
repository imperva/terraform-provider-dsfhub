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

	testPubsubSubscriptionPrefix = "projects/my-project/subscriptions/"
)
