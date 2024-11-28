package dsfhub

var ignoreLogAggregatorParamsByServerType = map[string]map[string]bool{
	"AWS KINESIS":   {"arn": true, "asset_display_name": true},
	"AWS LOG GROUP": {"arn": true, "asset_display_name": true},
	"AWS S3":        {"arn": true, "asset_display_name": true},
}

var requiredLogAggregatorJson = `{
	"ServerTypes": {
		"ALIBABA LOGSTORE": {
			"auth_mechanisms": {
				"default": [
					"reason"
				]
			},
			"required": [
				"gateway_id",
				"asset_display_name",
				"asset_id",
				"admin_email",
				"parent_asset_id",
				"project",
				"logstore",
				"endpoint"
			]
		},
		"AWS KINESIS": {
			"auth_mechanisms": {
				"default": [
					"reason",
					"region"
				]
			},
			"required": [
				"gateway_id",
				"asset_display_name",
				"arn",
				"admin_email"
			]
		},
		"AWS LOG GROUP": {
			"auth_mechanisms": {
				"default": [
					"reason",
					"region"
				]
			},
			"required": [
				"gateway_id",
				"asset_display_name",
				"arn",
				"admin_email",
				"parent_asset_id"
			]
		},
		"AWS S3": {
			"auth_mechanisms": {
				"default": [
					"reason",
					"reason"
				]
			},
			"required": [
				"gateway_id",
				"asset_display_name",
				"arn",
				"server_host_name",
				"admin_email"
			]
		},
		"AZURE EVENTHUB": {
			"auth_mechanisms": {
				"azure_ad": [
					"reason",
					"azure_storage_account",
					"azure_storage_container",
					"eventhub_name",
					"eventhub_namespace",
					"format"
				],
				"client_secret": [
					"reason",
					"application_id",
					"azure_storage_account",
					"azure_storage_container",
					"client_secret",
					"directory_id",
					"eventhub_name",
					"eventhub_namespace",
					"format",
					"subscription_id"
				],
				"default": [
					"reason",
					"azure_storage_account",
					"azure_storage_container",
					"azure_storage_secret_key",
					"eventhub_access_key",
					"eventhub_access_policy",
					"eventhub_name",
					"eventhub_namespace",
					"format"
				]
			},
			"required": [
				"gateway_id",
				"asset_display_name",
				"asset_id",
				"server_host_name",
				"server_port",
				"admin_email"
			]
		},
		"GCP CLOUD STORAGE BUCKET": {
			"auth_mechanisms": {
				"default": [
					"reason"
				],
				"service_account": [
					"reason",
					"key_file"
				]
			},
			"required": [
				"gateway_id",
				"asset_display_name",
				"asset_id",
				"server_host_name",
				"server_ip",
				"server_port",
				"admin_email"
			]
		},
		"GCP PUBSUB": {
			"auth_mechanisms": {
				"default": [
					"reason"
				],
				"service_account": [
					"reason",
					"key_file"
				]
			},
			"required": [
				"gateway_id",
				"asset_display_name",
				"asset_id",
				"server_host_name",
				"server_ip",
				"server_port",
				"admin_email",
				"pubsub_subscription"
			]
		},
		"SSH": {
			"auth_mechanisms": {
				"default": [
					"reason"
				],
				"kerberos": [
					"reason"
				]
			},
			"required": [
				"gateway_id",
				"asset_display_name",
				"asset_id",
				"server_host_name",
				"server_ip",
				"admin_email"
			]
		}
	}
}`
