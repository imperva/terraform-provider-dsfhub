package dsfhub

type RequiredFieldsMap struct {
	ServerType map[string]RequiredFields `json:"ServerTypes"`
}

type RequiredFields struct {
	Required       []string            `json:"required"`
	AuthMechanisms map[string][]string `json:"auth_mechanisms"`
}

type AssetSchema struct {
	Connections map[string]SchemaField `json:"connections"`
	Details     map[string]SchemaField `json:"details"`
}

type SchemaField struct {
	DefaultValue interface{} `json:"defaultValue"`
	Description  string      `json:"description"`
	DisplayName  string      `json:"displayName"`
	Example      interface{} `json:"example"`
	Optional     bool        `json:"optional"`
	Required     bool        `json:"required"`
	Type         string      `json:"type"`
	Values       interface{} `json:"values"`
	ID           string      `json:"id"`
}

var assetSchemaJson = `{
    "connections": {
        "AccessID": {
            "defaultValue": null,
            "description": "",
            "displayName": "Access ID",
            "example": "",
            "id": "access_id",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "AccessKey": {
            "defaultValue": null,
            "description": "",
            "displayName": "Access Key",
            "example": "",
            "id": "access_key",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "AccessMethod": {
            "defaultValue": "VSAM",
            "description": "The z/OS file system access method to be used",
            "displayName": "Access Method",
            "example": "VSAM",
            "id": "access_method",
            "required": false,
            "type": "string"
        },
        "AccountName": {
            "defaultValue": null,
            "description": "The cloudant account name required when connecting a resource with IAM role",
            "displayName": "Account Name",
            "example": "",
            "id": "account_name",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "AmazonSecret": {
            "defaultValue": null,
            "description": "Configuration to integrate with AWS Secrets Manager",
            "displayName": "Amazon Secret",
            "example": {
                "field_mapping": {
                    "<local_field1>": "<remote_field1>",
                    "<local_field2>": "<remote_field2>"
                },
                "secret_asset_id": "<AWS_asset_id>",
                "secret_name": "<secret_name>"
            },
            "id": "amazon_secret",
            "required": false,
            "type": "map"
        },
        "ApiKey": {
            "defaultValue": null,
            "description": "IAM authentication API key",
            "displayName": "Api Key",
            "example": "",
            "id": "api_key",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "ApplicationID": {
            "defaultValue": null,
            "description": "This is also referred to as the Client ID and it\u2019s the unique identifier for the registered application being used to execute Python SDK commands against Azure\u2019s API services. You can find this number under Azure Active Directory -> App Registrations -> Owned Applications",
            "displayName": "Application ID",
            "example": "a1b2c3de-123c-1234-ab12-ab12c2de3fg4",
            "id": "application_id",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "AuthMechanism": {
            "defaultValue": null,
            "description": "",
            "displayName": "Auth Mechanism",
            "example": "",
            "id": "auth_mechanism",
            "required": true,
            "type": "string"
        },
        "Autocommit": {
            "defaultValue": null,
            "description": "If true, Commit automatically don't wait for transaction to be explicitly committed",
            "displayName": "Autocommit",
            "example": "",
            "id": "autocommit",
            "required": false,
            "type": "bool"
        },
        "AwsConnectionID": {
            "defaultValue": null,
            "description": "The parent AWS connection document_id",
            "displayName": "AWS Connection ID",
            "example": "63336fbbad8bb10f000243bd",
            "id": "aws_connection_id",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "AwsIamServerID": {
            "defaultValue": null,
            "description": "e.g. vault.example.com",
            "displayName": "AWS IAM Server ID",
            "example": "",
            "id": "aws_iam_server_id",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "AzureStorageAccount": {
            "defaultValue": null,
            "description": "The name of the unique namespace where the EventHub is located. The field can contain only lowercase letters and numbers. Name must be between 3 and 24 characters.",
            "displayName": "Azure Storage Account",
            "example": "sonarn7qjsvi8qrgm4vhcs1m",
            "id": "azure_storage_account",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "AzureStorageContainer": {
            "defaultValue": null,
            "description": "Location where a given EventHub\u2019s processing is stored (One storage container per EventHub). This name may only contain lowercase letters, numbers, and hyphens, and must begin with a letter or a number. Each hyphen must be preceded and followed by a non-hyphen character. The name must also be between 3 and 63 characters long.",
            "displayName": "Azure Storage Container",
            "example": "container-for-sonar-events-sql-eastus",
            "id": "azure_storage_container",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "AzureStorageSecretKey": {
            "defaultValue": null,
            "description": "",
            "displayName": "Azure Storage Secret Key",
            "example": "",
            "id": "azure_storage_secret_key",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "BaseDn": {
            "defaultValue": null,
            "description": "",
            "displayName": "Base DN",
            "example": "",
            "id": "base_dn",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "Bucket": {
            "defaultValue": null,
            "description": "",
            "displayName": "Bucket",
            "example": "",
            "id": "bucket",
            "required": false,
            "type": "string"
        },
        "CaCertsPath": {
            "defaultValue": null,
            "description": "Certificate authority certificates path; what location should the sysetm look for certificate information from. Equivalent to --capath in a curl call",
            "displayName": "CA Certs Path",
            "example": "",
            "id": "ca_certs_path",
            "required": false,
            "type": "string"
        },
        "CaFile": {
            "defaultValue": null,
            "description": "Path to a certificate authority file to use with the call. Equivalent to --cacert in a curl call",
            "displayName": "CA File",
            "example": "",
            "id": "ca_file",
            "required": false,
            "type": "string"
        },
        "CacheFile": {
            "defaultValue": null,
            "description": "",
            "displayName": "Cache File",
            "example": "",
            "id": "cache_file",
            "required": false,
            "type": "string"
        },
        "CertFile": {
            "defaultValue": null,
            "description": "Path to a certificate file to use with the call. Equivalent to --cert in a curl call",
            "displayName": "Cert File",
            "example": "",
            "id": "cert_file",
            "required": false,
            "type": "string"
        },
        "ClientID": {
            "defaultValue": null,
            "description": "The unique Application (client) ID assigned to your app when it was registered (same as the Application ID)",
            "displayName": "Client Id",
            "example": "bd51d56c-e744-4a58-91e1-9bbbbb7e821c",
            "id": "client_id",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "ClientSecret": {
            "defaultValue": null,
            "description": "Secret used to authenticate to the token endpoint",
            "displayName": "Client Secret",
            "example": "",
            "id": "client_secret",
            "required": false,
            "type": "string"
        },
        "ClusterID": {
            "defaultValue": null,
            "description": "",
            "displayName": "Cluster ID",
            "example": "",
            "id": "cluster_id",
            "required": false,
            "type": "string"
        },
        "ClusterMemberID": {
            "defaultValue": null,
            "description": "",
            "displayName": "Cluster Member ID",
            "example": "",
            "id": "cluster_member_id",
            "required": false,
            "type": "string"
        },
        "ClusterName": {
            "defaultValue": null,
            "description": "",
            "displayName": "Cluster Name",
            "example": "",
            "id": "cluster_name",
            "required": false,
            "type": "string"
        },
        "ContentType": {
            "defaultValue": null,
            "description": "Content-Type to append to the HTTP headers in the curl call",
            "displayName": "Content-Type",
            "example": "",
            "id": "content_type",
            "required": false,
            "type": "string"
        },
        "CredentialExpiry": {
            "defaultValue": null,
            "description": "",
            "displayName": "Credential Expiry",
            "example": "",
            "id": "credential_expiry",
            "required": false,
            "type": "string"
        },
        "CredentialFields": {
            "defaultValue": null,
            "description": "Document containing values to build a profile from. Filling this will create a profile using the given profile name",
            "displayName": "Credential Fields",
            "example": {
                "credential_source": "Ec2InstanceMetadata",
                "role_arn": "arn:aws:iam::111777333222:role/other_role"
            },
            "id": "credential_fields",
            "required": false,
            "type": "map"
        },
        "Crn": {
            "defaultValue": null,
            "description": "The CRN unique identifier of the resource",
            "displayName": "Crn",
            "example": "crn:version:cname:ctype:service-name:location:scope:service-instance:resource-type:resource",
            "id": "crn",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "CyberarkSecret": {
            "defaultValue": null,
            "description": "Configuration to integrate with CyberArk Vault",
            "displayName": "CyberArk Secret",
            "example": {
                "field_mapping": {
                    "<local_field1>": "<remote_field1>",
                    "<local_field2>": "<remote_field2>"
                },
                "secret_asset_id ": "Cyberark_asset_id",
                "secret_name": "<secret_name>"
            },
            "id": "cyberark_secret",
            "required": false,
            "type": "map"
        },
        "DatabaseName": {
            "defaultValue": null,
            "description": "Specifies the name of the database (or default DB) to connect to.",
            "displayName": "Database Name",
            "example": "master",
            "id": "database_name",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "DbRole": {
            "defaultValue": null,
            "description": "",
            "displayName": "DB Role",
            "example": "",
            "id": "db_role",
            "required": false,
            "type": "string"
        },
        "DirectoryID": {
            "defaultValue": null,
            "description": "This is also referred to as the Tenant ID and is a GUID representing the Active Directory Tenant. It can be found in the Azure Active Directory page under the Azure portal",
            "displayName": "Directory ID",
            "example": "a1b2c3de-123c-1234-ab12-ab12c2de3fg4",
            "id": "directory_id",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "Dn": {
            "defaultValue": null,
            "description": "The distinguished name of a particular PKI certificate",
            "displayName": "Distinguished Name",
            "example": "",
            "id": "dn",
            "required": false,
            "type": "string"
        },
        "DnsSrv": {
            "defaultValue": null,
            "description": "",
            "displayName": "DNS SRV",
            "example": "",
            "id": "dns_srv",
            "required": false,
            "type": "bool"
        },
        "Driver": {
            "defaultValue": null,
            "description": "A path to a non-default driver location. If populated this driver will be used rather than the default",
            "displayName": "Driver",
            "example": "${JSONAR_BASEDIR}/lib/libsqora.so",
            "id": "driver",
            "required": false,
            "type": "string"
        },
        "Dsn": {
            "defaultValue": null,
            "description": "Data Source Name",
            "displayName": "DSN",
            "example": "mynetalias",
            "id": "dsn",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "EventhubAccessKey": {
            "defaultValue": null,
            "description": "",
            "displayName": "Eventhub Access Key",
            "example": "",
            "id": "eventhub_access_key",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "EventhubAccessPolicy": {
            "defaultValue": null,
            "description": "Authorization policy that will allow Sonar to access this specific EventHub.",
            "displayName": "EventHub Access Policy",
            "example": "RootManageSharedAccessKey",
            "id": "eventhub_access_policy",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "EventhubName": {
            "defaultValue": null,
            "description": "EventHub name without additional resource ID information.",
            "displayName": "EventHub Name",
            "example": "sonar-events-sql",
            "id": "eventhub_name",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "EventhubNamespace": {
            "defaultValue": null,
            "description": "The name for the management container that the EventHub belongs to, one namespace can contain multiple EventHubs. The namespace can contain only letters, numbers, and hyphens. The namespace must start with a letter, and it must end with a letter or number. The value must be between 6 and 50 characters long.",
            "displayName": "EventHub Namespace",
            "example": "sonar-eastus-bc33f0f89304c57c389cfa4530127fb6",
            "id": "eventhub_namespace",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "External": {
            "defaultValue": null,
            "description": "",
            "displayName": "External",
            "example": "",
            "id": "external",
            "required": false,
            "type": "bool"
        },
        "ExternalID": {
            "defaultValue": null,
            "description": "External ID to use when assuming a role",
            "displayName": "External ID",
            "example": "sonarid",
            "id": "external_id",
            "required": false,
            "type": "string"
        },
        "ExtraKinitParameters": {
            "defaultValue": null,
            "description": "",
            "displayName": "Extra Kinit Parameters",
            "example": "",
            "id": "extra_kinit_parameters",
            "required": false,
            "type": "string"
        },
        "Format": {
            "defaultValue": null,
            "description": "The type of audit data being sent to the EventHub, for example Postgresql or Cosmos_SQL",
            "displayName": "Format",
            "example": "Select...",
            "id": "format",
            "optional": false,
            "required": true,
            "type": "options",
            "values": [
                "AzureSQL_Managed",
                "Blob",
                "Cosmos_Mongo",
                "Cosmos_SQL",
                "Cosmos_Table",
                "Databricks_Workspace",
                "File",
                "Mariadb",
                "Mysql",
                "Postgresql",
                "Queue",
                "Sql",
                "Synapse",
                "Table"
            ]
        },
        "HashicorpSecret": {
            "defaultValue": null,
            "description": "Configuration to integrate with HashiCorp Vault",
            "displayName": "HashiCorp Secret",
            "example": {
                "field_mapping": {
                    "<local_field1>": "<remote_field1>",
                    "<local_field2>": "<remote_field2>"
                },
                "path": "secret/",
                "secret_asset_id": "<hashicorp_asset_id>",
                "secret_name": "<secret_name>"
            },
            "id": "hashicorp_secret",
            "required": false,
            "type": "map"
        },
        "HiveServerType": {
            "defaultValue": null,
            "description": "",
            "displayName": "Hive Server Type",
            "example": "",
            "id": "hive_server_type",
            "required": false,
            "type": "string"
        },
        "HostNameMismatch": {
            "defaultValue": null,
            "description": "",
            "displayName": "Host Name Mismatch",
            "example": "",
            "id": "host_name_mismatch",
            "required": false,
            "type": "bool"
        },
        "Hosts": {
            "defaultValue": "cloudantnosqldb",
            "description": "Required for quering the logdna url. cloudantnosqldb in the case of a cloudant DB",
            "displayName": "Hosts",
            "example": "cloudantnosqldb",
            "id": "hosts",
            "required": false,
            "type": "string"
        },
        "Httppath": {
            "defaultValue": null,
            "description": "",
            "displayName": "HTTP Path",
            "example": "",
            "id": "httppath",
            "required": false,
            "type": "string"
        },
        "IsCluster": {
            "defaultValue": null,
            "description": "",
            "displayName": "Is Cluster",
            "example": "",
            "id": "is_cluster",
            "required": false,
            "type": "bool"
        },
        "JdbcSslTrustServerCertificate": {
            "defaultValue": false,
            "description": "",
            "displayName": "JDBC SSL Trust Server Certificate",
            "example": false,
            "id": "jdbc_ssl_trust_server_certificate",
            "required": false,
            "type": "bool"
        },
        "JdbcSslTrustStoreLocation": {
            "defaultValue": null,
            "description": "",
            "displayName": "JDBC SSL Trust Store Location",
            "example": "",
            "id": "jdbc_ssl_trust_store_location",
            "required": false,
            "type": "string"
        },
        "JdbcSslTrustStorePassword": {
            "defaultValue": null,
            "description": "",
            "displayName": "JDBC SSL Trust Store Password",
            "example": "",
            "id": "jdbc_ssl_trust_store_password",
            "required": false,
            "type": "string"
        },
        "KerberosHostFqdn": {
            "defaultValue": null,
            "description": "",
            "displayName": "Kerberos Host Fqdn",
            "example": "",
            "id": "kerberos_host_fqdn",
            "required": false,
            "type": "string"
        },
        "KerberosKdc": {
            "defaultValue": null,
            "description": "",
            "displayName": "Kerberos KDC",
            "example": "",
            "id": "kerberos_kdc",
            "required": false,
            "type": "string"
        },
        "KerberosRetryCount": {
            "defaultValue": null,
            "description": "",
            "displayName": "Kerberos Retry Count",
            "example": "",
            "id": "kerberos_retry_count",
            "required": false,
            "type": "string"
        },
        "KerberosServiceKdc": {
            "defaultValue": null,
            "description": "",
            "displayName": "Kerberos Service KDC",
            "example": "",
            "id": "kerberos_service_kdc",
            "required": false,
            "type": "string"
        },
        "KerberosServiceRealm": {
            "defaultValue": null,
            "description": "",
            "displayName": "Kerberos Service Realm",
            "example": "",
            "id": "kerberos_service_realm",
            "required": false,
            "type": "string"
        },
        "KerberosSpn": {
            "defaultValue": null,
            "description": "",
            "displayName": "Kerberos SPN",
            "example": "",
            "id": "kerberos_spn",
            "required": false,
            "type": "string"
        },
        "KeyFile": {
            "defaultValue": null,
            "description": "Location on disk on the key to be used to authenticate",
            "displayName": "Key File",
            "example": "",
            "id": "key_file",
            "required": false,
            "type": "string"
        },
        "KeytabFile": {
            "defaultValue": null,
            "description": "Specify a non-default keytab location",
            "displayName": "Keytab File",
            "example": "/etc/krb5.keytab",
            "id": "keytab_file",
            "required": false,
            "type": "string"
        },
        "KinitProgramPath": {
            "defaultValue": null,
            "description": "",
            "displayName": "Kinit Program Path",
            "example": "",
            "id": "kinit_program_path",
            "required": false,
            "type": "string"
        },
        "NetServiceName": {
            "defaultValue": null,
            "description": "Alias in tnsnames.ora replaces hostname, service name, and port in connection string",
            "displayName": "Net Service Name",
            "example": "",
            "id": "net_service_name",
            "required": false,
            "type": "string"
        },
        "Nonce": {
            "defaultValue": null,
            "description": "",
            "displayName": "Nonce",
            "example": "",
            "id": "nonce",
            "required": false,
            "type": "string"
        },
        "Ntlm": {
            "defaultValue": false,
            "description": "",
            "displayName": "NTLM",
            "example": false,
            "id": "ntlm",
            "required": false,
            "type": "bool"
        },
        "OauthParameters": {
            "defaultValue": null,
            "description": "Additional parameters to pass when requesting a token",
            "displayName": "Oauth Parameters",
            "example": {
                "parameter": "value"
            },
            "id": "oauth_parameters",
            "optional": false,
            "required": true,
            "type": "map"
        },
        "OdbcConnectionString": {
            "defaultValue": null,
            "description": "Additional ODBC connection string parameters. This string will get added to the connection string",
            "displayName": "ODBC Connection String",
            "example": ";MaxVarchar=65535",
            "id": "odbc_connection_string",
            "required": false,
            "type": "string"
        },
        "PageSize": {
            "defaultValue": null,
            "description": "",
            "displayName": "Page Size",
            "example": "",
            "id": "page_size",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "Passphrase": {
            "defaultValue": null,
            "description": "",
            "displayName": "Passphrase",
            "example": "",
            "id": "passphrase",
            "required": false,
            "type": "string"
        },
        "Password": {
            "defaultValue": null,
            "description": "The password of the user being used to authenticate",
            "displayName": "Password",
            "example": "",
            "id": "password",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "Port": {
            "defaultValue": null,
            "description": "",
            "displayName": "Port",
            "example": "",
            "id": "port",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "Principal": {
            "defaultValue": null,
            "description": "Azure username",
            "displayName": "Principal",
            "example": "",
            "id": "principal",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "Protocol": {
            "defaultValue": null,
            "description": "",
            "displayName": "Protocol",
            "example": "",
            "id": "protocol",
            "required": false,
            "type": "string"
        },
        "ProxyAutoDetect": {
            "defaultValue": null,
            "description": "",
            "displayName": "Proxy Auto Detect",
            "example": "",
            "id": "proxy_auto_detect",
            "required": false,
            "type": "string"
        },
        "ProxyPassword": {
            "defaultValue": null,
            "description": "",
            "displayName": "Proxy Password",
            "example": "",
            "id": "proxy_password",
            "required": false,
            "type": "string"
        },
        "ProxyPort": {
            "defaultValue": null,
            "description": "",
            "displayName": "Proxy Port",
            "example": "",
            "id": "proxy_port",
            "required": false,
            "type": "string"
        },
        "ProxyServer": {
            "defaultValue": null,
            "description": "",
            "displayName": "Proxy Server",
            "example": "",
            "id": "proxy_server",
            "required": false,
            "type": "string"
        },
        "ProxySslType": {
            "defaultValue": null,
            "description": "",
            "displayName": "Proxy SSL Type",
            "example": "",
            "id": "proxy_ssl_type",
            "required": false,
            "type": "string"
        },
        "Query": {
            "defaultValue": null,
            "description": "",
            "displayName": "Query",
            "example": "",
            "id": "query",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "Reason": {
            "defaultValue": null,
            "description": "What this connection is used for. Used to differentiate connections if multiple connections exist for this asset",
            "displayName": "Reason",
            "example": "audit management",
            "id": "reason",
            "required": false,
            "type": "string",
            "values": [
                "default"
            ]
        },
        "RedirectUri": {
            "defaultValue": null,
            "description": "",
            "displayName": "Redirect URI",
            "example": "",
            "id": "redirect_uri",
            "required": false,
            "type": "string"
        },
        "Region": {
            "defaultValue": null,
            "description": "The IMS control region. The z/OS address space initiated for the IMS subsystem",
            "displayName": "Region",
            "example": "IGBK",
            "id": "region",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "ReplicaSet": {
            "defaultValue": null,
            "description": "",
            "displayName": "Replica Set",
            "example": "",
            "id": "replica_set",
            "required": false,
            "type": "string"
        },
        "ResourceID": {
            "defaultValue": null,
            "description": "Azure resource application ID",
            "displayName": "Resource Id",
            "example": "",
            "id": "resource_id",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "RoleName": {
            "defaultValue": null,
            "description": "Role to use for authentication",
            "displayName": "Role Name",
            "example": "",
            "id": "role_name",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "Schema": {
            "defaultValue": null,
            "description": "Schema name. A schema is a logical grouping of database objects",
            "displayName": "Schema",
            "example": "",
            "id": "schema",
            "required": false,
            "type": "string"
        },
        "SecretKey": {
            "defaultValue": null,
            "description": "A secret string that the application uses to prove its identity when when requesting a token (same as Client Secret)",
            "displayName": "Secret Key",
            "example": "LP99Q~i.lWR9pDjYQWPDsYB_5bVPNzwrpNIdmcTf",
            "id": "secret_key",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "SecureConnection": {
            "defaultValue": false,
            "description": "",
            "displayName": "Secure Connection",
            "example": false,
            "id": "secure_connection",
            "required": false,
            "type": "bool"
        },
        "SelfSigned": {
            "defaultValue": null,
            "description": "",
            "displayName": "Self Signed",
            "example": "",
            "id": "self_signed",
            "required": false,
            "type": "bool"
        },
        "SelfSignedCert": {
            "defaultValue": null,
            "description": "",
            "displayName": "Self Signed Cert",
            "example": "",
            "id": "self_signed_cert",
            "required": false,
            "type": "bool"
        },
        "ServerIP": {
            "defaultValue": null,
            "description": "",
            "displayName": "Server IP",
            "example": "",
            "id": "server_ip",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "ServerPort": {
            "defaultValue": null,
            "description": "",
            "displayName": "Server Port",
            "example": "",
            "id": "server_port",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "ServiceKey": {
            "defaultValue": null,
            "description": "The service key required in the logdna url query to connect to logdna and pull the logs",
            "displayName": "Service Key",
            "example": "",
            "id": "service_key",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "SnowflakeRole": {
            "defaultValue": null,
            "description": "Role with which to log into Snowflake",
            "displayName": "Snowflake Role",
            "example": "",
            "id": "snowflake_role",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "Ssl": {
            "defaultValue": null,
            "description": "If true, use SSL when connecting",
            "displayName": "SSL",
            "example": "",
            "id": "ssl",
            "required": false,
            "type": "bool"
        },
        "SslServerCert": {
            "defaultValue": null,
            "description": "",
            "displayName": "SSL Server Cert",
            "example": "",
            "id": "ssl_server_cert",
            "required": false,
            "type": "string"
        },
        "StoreAwsCredentials": {
            "defaultValue": null,
            "description": "",
            "displayName": "Store AWS Credentials",
            "example": "",
            "id": "store_aws_credentials",
            "required": false,
            "type": "bool"
        },
        "SubscriptionID": {
            "defaultValue": null,
            "description": "This is the Azure account subscription ID. You can find this number under the Subscriptions page on the Azure portal",
            "displayName": "Subscription ID",
            "example": "a1b2c3de-123c-1234-ab12-ab12c2de3fg4",
            "id": "subscription_id",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "TenantID": {
            "defaultValue": null,
            "description": "The globally unique identifier (GUID) that is different than your organization name or domain (same as the Directory ID)",
            "displayName": "Tenant Id",
            "example": "72faaaaa-86f1-41af-91aa-2d7cd071db47",
            "id": "tenant_id",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "ThriftTransport": {
            "defaultValue": null,
            "description": "",
            "displayName": "Thrift Transport",
            "example": "",
            "id": "thrift_transport",
            "required": false,
            "type": "string"
        },
        "TmpUser": {
            "defaultValue": false,
            "description": "If true create a temporary user",
            "displayName": "Tmp User",
            "example": false,
            "id": "tmp_user",
            "required": false,
            "type": "bool"
        },
        "Token": {
            "defaultValue": null,
            "description": "Saved token to use to authenticate",
            "displayName": "Token",
            "example": "",
            "id": "token",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "TokenEndpoint": {
            "defaultValue": null,
            "description": "URL of endpoint to query when requesting a token",
            "displayName": "Token Endpoint",
            "example": "",
            "id": "token_endpoint",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "Transportmode": {
            "defaultValue": null,
            "description": "",
            "displayName": "Transportmode",
            "example": "",
            "id": "transportmode",
            "required": false,
            "type": "string"
        },
        "Url": {
            "defaultValue": null,
            "description": "",
            "displayName": "Url",
            "example": "",
            "id": "url",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "UseKeytab": {
            "defaultValue": null,
            "description": "If true, authenticate using a key tab",
            "displayName": "Use Keytab",
            "example": "",
            "id": "use_keytab",
            "required": false,
            "type": "bool"
        },
        "Username": {
            "defaultValue": null,
            "description": "The user to use when connecting",
            "displayName": "Username",
            "example": "bob",
            "id": "username",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "V2KeyEngine": {
            "defaultValue": null,
            "description": "Use a KV2 secret engine",
            "displayName": "V2 Key Engine",
            "example": "",
            "id": "v2_key_engine",
            "required": false,
            "type": "bool"
        },
        "VirtualHostname": {
            "defaultValue": null,
            "description": "",
            "displayName": "Virtual Hostname",
            "example": "",
            "id": "virtual_hostname",
            "required": false,
            "type": "string"
        },
        "VirtualIP": {
            "defaultValue": null,
            "description": "",
            "displayName": "Virtual IP",
            "example": "",
            "id": "virtual_ip",
            "required": false,
            "type": "string"
        },
        "WalletDir": {
            "defaultValue": null,
            "description": "Path to the Oracle wallet directory",
            "displayName": "Wallet Dir",
            "example": "${JSONAR_LOCALDIR}/oracle/wallet",
            "id": "wallet_dir",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "Warehouse": {
            "defaultValue": null,
            "description": "The name of the warehouse to connect to",
            "displayName": "Warehouse",
            "example": "",
            "id": "warehouse",
            "required": false,
            "type": "string"
        }
    },
    "details": {
        "AdminEmail": {
            "defaultValue": null,
            "description": "The email address to notify about this asset",
            "displayName": "Admin Email",
            "example": "person@company.com",
            "id": "admin_email",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "Application": {
            "defaultValue": null,
            "description": "The Asset ID of the application asset that \"owns\" the asset.",
            "displayName": "Application",
            "example": "",
            "id": "application",
            "required": false,
            "type": "string"
        },
        "Archive": {
            "defaultValue": null,
            "description": "If True archive files in the asset after being processed by sonargd. Defaults to True if field isn't present",
            "displayName": "Archive",
            "example": "",
            "id": "archive",
            "required": false,
            "type": "bool"
        },
        "Arn": {
            "defaultValue": "",
            "description": "Amazon Resource Name - format is arn:partition:service:region:account-id and used as the asset_id",
            "displayName": "ARN (Asset ID)",
            "example": "arn:aws:iam::123456789012/role",
            "id": "arn",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "AssetDisplayName": {
            "defaultValue": "",
            "description": "User-friendly name of the asset, defined by user.",
            "displayName": "Asset Display Name",
            "example": "",
            "id": "asset_display_name",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "AssetID": {
            "defaultValue": null,
            "description": "",
            "displayName": "Asset ID",
            "example": "",
            "id": "asset_id",
            "required": true,
            "type": "string"
        },
        "AssetSource": {
            "defaultValue": "",
            "description": "The source platform/vendor/system of the asset data. Usually the service responsible for creating that asset document",
            "displayName": "Asset Source",
            "example": "USC",
            "id": "asset_source",
            "required": false,
            "type": "string"
        },
        "AuditDataType": {
            "defaultValue": null,
            "description": "",
            "displayName": "Audit Data Type",
            "example": "",
            "id": "audit_data_type",
            "required": false,
            "type": "string"
        },
        "AuditInfo": {
            "defaultValue": null,
            "description": "Normally auto-populated when enabling the audit policy, it is a sub-document in JSON format containing configuration information for audit management. See documentation for values that can be added manually depending on asset type. Editing this value does NOT enable the audit policy.",
            "displayName": "Audit Info",
            "example": {
                "policy_template_name": "<template name>"
            },
            "id": "audit_info",
            "required": false,
            "type": "map"
        },
        "AuditPullEnabled": {
            "defaultValue": null,
            "description": "If true, sonargateway will collect the audit logs for this system if it can.",
            "displayName": "Audit Pull Enabled",
            "example": "",
            "id": "audit_pull_enabled",
            "required": false,
            "type": "bool"
        },
        "AuditType": {
            "defaultValue": null,
            "description": "Used to indicate what mechanism should be used to fetch logs on systems supporting multiple ways to get logs, see asset specific documentation for details",
            "displayName": "Audit Type",
            "example": "Select...",
            "id": "audit_type",
            "required": false,
            "type": "options",
            "values": [
                "SSH"
            ]
        },
        "AvailabilityZones": {
            "defaultValue": null,
            "description": "",
            "displayName": "Availability Zones",
            "example": "",
            "id": "availability_zones",
            "required": false,
            "type": "string"
        },
        "AvailableRegions": {
            "defaultValue": null,
            "description": "A list of regions to use in discovery actions that iterate through region",
            "displayName": "Available Regions",
            "example": "us-east-1, us-east-2",
            "id": "available_regions",
            "required": false,
            "type": "string"
        },
        "AwsProxyConfig": {
            "defaultValue": null,
            "description": "AWS specific proxy configuration",
            "displayName": "AWS Proxy Config",
            "example": {
                "http": "< proxy >",
                "https": "< proxy >"
            },
            "id": "aws_proxy_config",
            "required": false,
            "type": "map"
        },
        "CaCertsPath": {
            "defaultValue": null,
            "description": "Certificate authority certificates path; what location should the sysetm look for certificate information from. Equivalent to --capath in a curl call",
            "displayName": "CA Certs Path",
            "example": "",
            "id": "ca_certs_path",
            "required": false,
            "type": "string"
        },
        "CaFile": {
            "defaultValue": null,
            "description": "Path to a certificate authority file to use with the call. Equivalent to --cacert in a curl call",
            "displayName": "CA File",
            "example": "",
            "id": "ca_file",
            "required": false,
            "type": "string"
        },
        "ClusterEngine": {
            "defaultValue": null,
            "description": "",
            "displayName": "Cluster Engine",
            "example": "",
            "id": "cluster_engine",
            "required": false,
            "type": "string"
        },
        "ClusterID": {
            "defaultValue": null,
            "description": "",
            "displayName": "Cluster ID",
            "example": "",
            "id": "cluster_id",
            "required": false,
            "type": "string"
        },
        "ClusterMemberID": {
            "defaultValue": null,
            "description": "",
            "displayName": "Cluster Member ID",
            "example": "",
            "id": "cluster_member_id",
            "required": false,
            "type": "string"
        },
        "ClusterName": {
            "defaultValue": null,
            "description": "",
            "displayName": "Cluster Name",
            "example": "",
            "id": "cluster_name",
            "required": false,
            "type": "string"
        },
        "ConsumerGroup": {
            "defaultValue": null,
            "description": "The Consumer Group the EventHub Consumer Client will use to fetch events. Defaults to '$Default'",
            "displayName": "Consumer Group",
            "example": "$Default",
            "id": "consumer_group",
            "required": false,
            "type": "string"
        },
        "ConsumerGroupWorkers": {
            "defaultValue": 2,
            "description": "Only applies if pull_type is consumer_group. The number of consumers that will be part of the consumer group. For best performance this should match the number of shards in your logstore.",
            "displayName": "Consumer Group Workers",
            "example": 2,
            "id": "consumer_group_workers",
            "required": false,
            "type": "string"
        },
        "ContentType": {
            "defaultValue": null,
            "description": "Content type should be set to the desired <'parent' asset 'Server Type'>, which is the one that uses this asset as a destination for logs. NOTE: The content_type field will take precedence on the lookup for parent_asset_id field when checking which server is sending logs to this asset.",
            "displayName": "Content Type",
            "example": "",
            "id": "content_type",
            "required": false,
            "type": "string"
        },
        "CredentialsEndpoint": {
            "defaultValue": null,
            "description": "A specific sts endpoint to use",
            "displayName": "Credentials Endpoint",
            "example": "https://sts.amazonaws.com",
            "id": "credentials_endpoint",
            "required": false,
            "type": "string"
        },
        "Criticality": {
            "defaultValue": null,
            "description": "The asset's importance to the business. These values are measured on a scale from \"Most critical\" (1) to \"Least critical\" (4). Allowed values: 1, 2, 3, 4",
            "displayName": "Criticality",
            "example": "Select...",
            "id": "criticality",
            "required": false,
            "type": "options",
            "values": [
                1,
                2,
                3,
                4
            ]
        },
        "DatabaseName": {
            "defaultValue": null,
            "description": "Specifies the name of the database (or default DB) to connect to.",
            "displayName": "Database Name",
            "example": "master",
            "id": "database_name",
            "required": false,
            "type": "string"
        },
        "DbEngine": {
            "defaultValue": null,
            "description": "Specifies the version of the engine being used by the database (e.g. oracle-ee, oracle-se, oracle-se1, oracle-se2)",
            "displayName": "DB Engine",
            "example": "oracle-se2",
            "id": "db_engine",
            "required": false,
            "type": "string"
        },
        "DbInstancesDisplayName": {
            "defaultValue": null,
            "description": "",
            "displayName": "DB Instances Display Name",
            "example": "",
            "id": "db_instances_display_name",
            "required": false,
            "type": "string"
        },
        "DurationThreshold": {
            "defaultValue": null,
            "description": "",
            "displayName": "Duration Threshold",
            "example": "",
            "id": "duration_threshold",
            "required": false,
            "type": "string"
        },
        "EnableAuditManagement": {
            "defaultValue": null,
            "description": "If true, Sonar is responsible for setting and updating the policies",
            "displayName": "Audit Management Enabled",
            "example": "",
            "id": "enable_audit_management",
            "required": false,
            "type": "bool"
        },
        "EnableAuditMonitoring": {
            "defaultValue": null,
            "description": "If true, Sonar sends emails/alerts when the audit policies change.",
            "displayName": "Audit Monitoring Enabled",
            "example": "",
            "id": "enable_audit_monitoring",
            "required": false,
            "type": "bool"
        },
        "EnabledLogsExports": {
            "defaultValue": null,
            "description": "",
            "displayName": "Enabled Logs Exports",
            "example": "",
            "id": "enabled_logs_exports",
            "required": false,
            "type": "string"
        },
        "Endpoint": {
            "defaultValue": null,
            "description": "Logstore's endpoint",
            "displayName": "Endpoint",
            "example": "cn-qingdao.log.aliyuncs.com",
            "id": "endpoint",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "EntitlementEnabled": {
            "defaultValue": null,
            "description": "If true, Entitlement Management system is enabled.",
            "displayName": "Entitlement Enabled",
            "example": "",
            "id": "entitlement_enabled",
            "required": false,
            "type": "bool"
        },
        "GatewayID": {
            "defaultValue": null,
            "description": "",
            "displayName": "Gateway ID",
            "example": "",
            "id": "gateway_id",
            "required": true,
            "type": "string"
        },
        "GatewayService": {
            "defaultValue": null,
            "description": "The name of the gateway pull service (if any) used to retrieve logs for this source. Usually set by the connect gateway playbook.",
            "displayName": "Gateway Service",
            "example": "gateway-<cloud_type>@<server_type>.service",
            "id": "gateway_service",
            "required": false,
            "type": "string"
        },
        "HostTimezoneOffset": {
            "defaultValue": null,
            "description": "The offset value string is in the format \"-/+hh:mm\"",
            "displayName": "Timezone Offset",
            "example": "\"-06:00\", \"+02:00\"",
            "id": "host_timezone_offset",
            "required": false,
            "type": "string"
        },
        "IgnoreLatestOf": {
            "defaultValue": null,
            "description": "A regex defining a group. From all the files with the same group, the latest one will be ignored, so that it isn't archived until server is done writing",
            "displayName": "Ignore Latest Of",
            "example": "(.+)_[^_]+$",
            "id": "ignore_latest_of",
            "required": false,
            "type": "string"
        },
        "IsCluster": {
            "defaultValue": null,
            "description": "",
            "displayName": "Is Cluster",
            "example": "",
            "id": "is_cluster",
            "required": false,
            "type": "bool"
        },
        "IsMultiZones": {
            "defaultValue": null,
            "description": "",
            "displayName": "Is Multi Zones",
            "example": "",
            "id": "is_multi_zones",
            "required": false,
            "type": "bool"
        },
        "JsonarUID": {
            "defaultValue": null,
            "description": "Unique identifier (UID) attached to the Sonar machine controlling the asset",
            "displayName": "jSonar UID",
            "example": "",
            "id": "jsonar_uid",
            "required": false,
            "type": "string"
        },
        "Location": {
            "defaultValue": null,
            "description": "Current human-readable description of the physical location of the asset, or region.",
            "displayName": "Location",
            "example": "",
            "id": "location",
            "required": false,
            "type": "string"
        },
        "LogBucketID": {
            "defaultValue": null,
            "description": "Asset ID of the S3 bucket which stores the logs for this server",
            "displayName": "Log Bucket ID",
            "example": "",
            "id": "log_bucket_id",
            "required": false,
            "type": "string"
        },
        "LogsDestinationAssetID": {
            "defaultValue": null,
            "description": "The asset name of the log aggregator that stores this asset's logs.",
            "displayName": "Logs Destination Asset Name",
            "example": "",
            "id": "logs_destination_asset_id",
            "required": false,
            "type": "string"
        },
        "Logstore": {
            "defaultValue": null,
            "description": "Unit that is used to collect, store and query logs",
            "displayName": "Logstore",
            "example": "actiontrail_log",
            "id": "logstore",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "ManagedBy": {
            "defaultValue": null,
            "description": "Email of the person who maintains the asset; can be different from the owner specified in the owned_by field. Defaults to admin_email.",
            "displayName": "Managed By",
            "example": "person@company.com",
            "id": "managed_by",
            "required": false,
            "type": "string"
        },
        "MaxConcurrentConn": {
            "defaultValue": null,
            "description": "Maximum number of concurrent connections that sensitive data management should use at once.",
            "displayName": "Max Concurrent Connections",
            "example": "",
            "id": "max_concurrent_conn",
            "required": false,
            "type": "string"
        },
        "OwnedBy": {
            "defaultValue": null,
            "description": "Email of Owner / person responsible for the asset; can be different from the person in the managed_by field. Defaults to admin_email.",
            "displayName": "Owned By",
            "example": "person@company.com",
            "id": "owned_by",
            "required": false,
            "type": "string"
        },
        "ParentAssetID": {
            "defaultValue": null,
            "description": "The name of an asset that this asset is part of (/related to). E.g. an AWS resource will generally have an AWS account asset as its parent. Also used to connect some log aggregating asset with the sources of their logs. E.g. An AWS LOG GROUP asset can have an AWS RDS as its parent, indicating that that is the log group for that RDS.",
            "displayName": "Parent Asset ID",
            "example": "",
            "id": "parent_asset_id",
            "required": false,
            "type": "string"
        },
        "Project": {
            "defaultValue": null,
            "description": "Project separates different resources of multiple users and control access to specific resources",
            "displayName": "Project",
            "example": "slsaudit-center-1234567891234567-cn-qingdao",
            "id": "project",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "S3Provider": {
            "defaultValue": "aws-rds-mssql",
            "description": "The type of AWS RDS instance that the S3 asset is receiving audit logs from",
            "displayName": "Provider",
            "example": "aws-rds-mssql",
            "id": "s3_provider",
            "required": false,
            "type": "string"
        },
        "ProviderUrl": {
            "defaultValue": null,
            "description": "",
            "displayName": "Provider URL",
            "example": "",
            "id": "provider_url",
            "required": false,
            "type": "string"
        },
        "Proxy": {
            "defaultValue": null,
            "description": "",
            "displayName": "Proxy",
            "example": "",
            "id": "proxy",
            "required": false,
            "type": "string"
        },
        "PubsubSubscription": {
            "defaultValue": null,
            "description": "",
            "displayName": "Pub/Sub Subscription",
            "example": "",
            "id": "pubsub_subscription",
            "required": false,
            "type": "string"
        },
        "PullType": {
            "defaultValue": "log_client",
            "description": "The method used to pull data from the logstore.",
            "displayName": "Pull Type",
            "example": "log_client, consumer_group",
            "id": "pull_type",
            "required": false,
            "type": "string",
            "values": [
                "log_client",
                "consumer_group"
            ]
        },
        "Region": {
            "defaultValue": null,
            "description": "For cloud systems with regions, the default region or region used with this asset",
            "displayName": "Region",
            "example": "us-east-1",
            "id": "region",
            "required": false,
            "type": "string"
        },
        "SdmEnabled": {
            "defaultValue": null,
            "description": "Sensitive data management (SDM) is enabled if this parameter is set to True.",
            "displayName": "SDM Enabled",
            "example": "",
            "id": "sdm_enabled",
            "required": false,
            "type": "bool"
        },
        "Searches": {
            "defaultValue": null,
            "description": "",
            "displayName": "Searches",
            "example": "",
            "id": "searches",
            "required": false,
            "type": "string"
        },
        "ServerHostName": {
            "defaultValue": null,
            "description": "Hostname (or IP if name is unknown)",
            "displayName": "Server Host Name",
            "example": "server.company.com",
            "id": "server_host_name",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "ServerIP": {
            "defaultValue": null,
            "description": "IP address of the service where this asset is located. If no IP is available populate this field with other information that would identify the system e.g. hostname or AWS ARN, etc.",
            "displayName": "Server IP",
            "example": "192.168.0.1",
            "id": "server_ip",
            "optional": false,
            "required": true,
            "type": "string"
        },
        "ServerPort": {
            "defaultValue": "8083",
            "description": "",
            "displayName": "Server Port",
            "example": 8083,
            "id": "server_port",
            "required": false,
            "type": "string"
        },
        "ServerType": {
            "defaultValue": null,
            "description": "",
            "displayName": "Server Type",
            "example": "",
            "id": "server_type",
            "required": true,
            "type": "string"
        },
        "ServiceEndpoint": {
            "defaultValue": null,
            "description": "Specify a particular endpoint for a given service",
            "displayName": "Service Endpoint",
            "example": "https://logs.us-east-1.amazonaws.com",
            "id": "service_endpoint",
            "required": false,
            "type": "string"
        },
        "ServiceEndpoints": {
            "defaultValue": null,
            "description": "Specify particular endpoints for a given service in the form of <service name>: \"endpoint\"",
            "displayName": "Service Endpoints",
            "example": {
                "logs": "https://logs.us-east-1.amazonaws.com"
            },
            "id": "service_endpoints",
            "required": false,
            "type": "map"
        },
        "ServiceName": {
            "defaultValue": "MSSQLSERVER",
            "description": "",
            "displayName": "Service Name",
            "example": "MSSQLSERVER",
            "id": "service_name",
            "required": false,
            "type": "string"
        },
        "SmtpTimeout": {
            "defaultValue": null,
            "description": "",
            "displayName": "SMTP Timeout",
            "example": "",
            "id": "smtp_timeout",
            "required": false,
            "type": "string"
        },
        "Ssl": {
            "defaultValue": null,
            "description": "",
            "displayName": "Ssl",
            "example": "",
            "id": "ssl",
            "required": false,
            "type": "bool"
        },
        "SubscriptionID": {
            "defaultValue": null,
            "description": "This is the Azure account subscription ID. You can find this number under the Subscriptions page on the Azure portal",
            "displayName": "Subscription ID",
            "example": "a1b2c3de-123c-1234-ab12-ab12c2de3fg4",
            "id": "subscription_id",
            "required": false,
            "type": "string"
        },
        "UsedFor": {
            "defaultValue": null,
            "description": "Designates how this asset is used / the environment that the asset is supporting.",
            "displayName": "Used For",
            "example": "Select...",
            "id": "used_for",
            "required": false,
            "type": "options",
            "values": [
                "Production",
                "Test",
                "Development",
                "Demonstration",
                "QA",
                "Staging",
                "Training",
                "Disaster Recovery"
            ]
        },
        "Version": {
            "defaultValue": null,
            "description": "Denotes the version of the asset",
            "displayName": "Version",
            "example": 5.7,
            "id": "version",
            "required": false,
            "type": "float"
        },
        "VirtualHostname": {
            "defaultValue": null,
            "description": "",
            "displayName": "Virtual Hostname",
            "example": "",
            "id": "virtual_hostname",
            "required": false,
            "type": "string"
        },
        "VirtualIP": {
            "defaultValue": null,
            "description": "",
            "displayName": "Virtual IP",
            "example": "",
            "id": "virtual_ip",
            "required": false,
            "type": "string"
        },
        "XelDirectory": {
            "defaultValue": null,
            "description": "Absolute path of the XEL files location",
            "displayName": "XEL Directory",
            "example": "",
            "id": "xel_directory",
            "required": false,
            "type": "string"
        }
    }
}`
