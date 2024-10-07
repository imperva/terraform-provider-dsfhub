package dsfhub

//type RequiredFieldsMap struct {
//	ServerType map[string]RequiredFields `json:"ServerTypes"`
//}
//
//type RequiredFields struct {
//	Required       []string            `json:"required"`
//	AuthMechanisms map[string][]string `json:"auth_mechanisms"`
//}
//
//type AssetSchema struct {
//	Connections map[string]SchemaField `json:"connections"`
//	Details     map[string]SchemaField `json:"details"`
//}
//
//type SchemaField struct {
//	DefaultValue interface{} `json:"defaultValue"`
//	Description  string      `json:"description"`
//	DisplayName  string      `json:"displayName"`
//	Example      interface{} `json:"example"`
//	Optional     bool        `json:"optional"`
//	Required     bool        `json:"required"`
//	Type         string      `json:"type"`
//	Values       interface{} `json:"values"`
//	ID           string      `json:"id"`
//}

var ignoreCloudAccountParamsByServerType = map[string]map[string]bool{
	"AWS": {"arn": true, "asset_display_name": true},
}

var requiredCloudAccountJson = `{
    "ServerTypes": {
        "ALIBABA": {
            "auth_mechanisms": {
                "key": [
                    "reason",
                    "reason",
                    "access_id",
                    "access_key"
                ],
                "machine_role": [
                    "reason",
                    "reason"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "admin_email"
            ]
        },
        "AWS": {
            "auth_mechanisms": {
                "default": [
                    "region",
                    "reason"
                ],
                "iam_role": [
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
        "AZURE": {
            "auth_mechanisms": {
                "auth_file": [
                    "reason",
                    "reason",
                    "key_file"
                ],
                "client_secret": [
                    "reason",
                    "directory_id",
                    "application_id",
                    "client_secret",
                    "subscription_id",
                    "reason"
                ],
                "managed_identity": [
                    "reason",
                    "subscription_id",
                    "reason"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "admin_email"
            ]
        },
        "GCP": {
            "auth_mechanisms": {
                "default": [
                    "reason",
                    "reason"
                ],
                "service_account": [
                    "reason",
                    "reason",
                    "key_file"
                ]
            },
            "required": [
                "gateway_id",
                "asset_display_name",
                "asset_id",
                "server_host_name",
                "admin_email"
            ]
        }
    }
}`
