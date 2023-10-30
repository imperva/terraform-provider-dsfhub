package dsfhub

var ignoreSecretManagerParamsByServerType = map[string]map[string]bool{
	"AWS": {"arn": true, "asset_display_name": true},
}

var requiredSecretManagerFieldsJson = `{
	"ServerTypes": {
		"AWS": {
			"auth_mechanisms": {
				"default": [
					"reason",
					"region"
				],
				"iam_role": [
					"reason",
					"region"
				],
				"key": [
					"reason",
					"region",
					"access_id",
					"secret_key"
				],
				"profile": [
					"reason",
					"username",
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
		"CYBERARK": {
			"auth_mechanisms": {
				"default": [
					"reason",
					"query"
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
		},
		"HASHICORP": {
			"auth_mechanisms": {
				"app_role": [
					"reason",
					"role_name",
					"secret_key"
				],
				"ec2": [
					"reason",
					"role_name"
				],
				"iam_role": [
					"reason",
					"access_id",
					"aws_iam_server_id",
					"role_name",
					"secret_key"
				],
				"root_token": [
					"reason",
					"secret_key"
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
		}
	}
}`
