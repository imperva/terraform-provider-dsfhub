# ### Resource example for COCKROACHDB with default auth_mechanism ###
resource "dsfhub_data_source" "cockroachdb_default" {
	server_type = "COCKROACHDB"
	# ### required ### 
	admin_email = var.admin_email	# The email address to notify about this asset
	asset_display_name = var.asset_display_name	# User-friendly name of the asset, defined by user.
	asset_id = var.asset_id	# Asset ID
	gateway_id = var.gateway_id	# Gateway ID
	server_host_name = var.server_host_name	# Hostname (or IP if name is unknown)
	server_ip = var.server_ip	# IP address of the service where this asset is located. If no IP is available populate this field with other information that would identify the system e.g. hostname or AWS ARN, etc.

	# ### optional ### 
	# application = var.application	# The Asset ID of the application asset that \"owns\" the asset.
	# asset_source = var.asset_source	# The source platform/vendor/system of the asset data. Usually the service responsible for creating that asset document
	# audit_info = var.audit_info	# Normally auto-populated when enabling the audit policy, it is a sub-document in JSON format containing configuration information for audit management. See documentation for values that can be added manually depending on asset type. Editing this value does NOT enable the audit policy.
	# audit_pull_enabled = var.audit_pull_enabled	# If true, sonargateway will collect the audit logs for this system if it can.
	# audit_type = var.audit_type # Example Values: "COCKROACH_V21", "COCKROACH"	# Used to indicate what mechanism should be used to fetch logs on systems supporting multiple ways to get logs, see asset specific documentation for details
	# criticality = var.criticality # Example Values: "1", "2", "3", "4"	# The asset's importance to the business. These values are measured on a scale from \"Most critical\" (1) to \"Least critical\" (4). Allowed values: 1, 2, 3, 4
	# database_name = var.database_name	# Specifies the name of the database (or default DB) to connect to.
	# db_engine = var.db_engine	# Specifies the version of the engine being used by the database (e.g. oracle-ee, oracle-se, oracle-se1, oracle-se2)
	# enable_audit_management = var.enable_audit_management	# If true, Sonar is responsible for setting and updating the policies
	# enable_audit_monitoring = var.enable_audit_monitoring	# If true, Sonar sends emails/alerts when the audit policies change.
	# entitlement_enabled = var.entitlement_enabled	# If true, Entitlement Management system is enabled.
	# jsonar_uid = var.jsonar_uid	# Unique identifier (UID) attached to the Sonar machine controlling the asset
	# location = var.location	# Current human-readable description of the physical location of the asset, or region.
	# managed_by = var.managed_by	# Email of the person who maintains the asset; can be different from the owner specified in the owned_by field. Defaults to admin_email.
	# max_concurrent_conn = var.max_concurrent_conn	# Maximum number of concurrent connections that sensitive data management should use at once.
	# owned_by = var.owned_by	# Email of Owner / person responsible for the asset; can be different from the person in the managed_by field. Defaults to admin_email.
	# region = var.region	# For cloud systems with regions, the default region or region used with this asset
	# sdm_enabled = var.sdm_enabled	# Sensitive data management (SDM) is enabled if this parameter is set to True.
	# server_port = "26257"	# 
	# used_for = var.used_for # Example Values: "Production", "Test", "Development", "Demonstration", "QA", "Staging", "Training", "Disaster Recovery"	# Designates how this asset is used / the environment that the asset is supporting.
	# version = var.version	# Denotes the version of the asset
	asset_connection {
		auth_mechanism = "default"
		# ### required ### 
		reason = null # Example Values: "default", "sonargateway", "ad-hoc-query" # reason description: "What this connection is used for. Used to differentiate connections if multiple connections exist for this asset"
		# ### optional ### 
		# amazon_secret = null # amazon_secret description: "Configuration to integrate with AWS Secrets Manager"
		# autocommit = null # autocommit description: "If true, Commit automatically don't wait for transaction to be explicitly committed"
		# ca_file = null # ca_file description: "Path to a certificate authority file to use with the call. Equivalent to --cacert in a curl call"
		# cert_file = null # cert_file description: "Path to a certificate file to use with the call. Equivalent to --cert in a curl call"
		# cyberark_secret = null # cyberark_secret description: "Configuration to integrate with CyberArk Vault"
		# driver = null # driver description: "A path to a non-default driver location. If populated this driver will be used rather than the default"
		# hashicorp_secret = null # hashicorp_secret description: "Configuration to integrate with HashiCorp Vault"
		# key_file = null # key_file description: "Location on disk on the key to be used to authenticate"
		# odbc_connection_string = null # odbc_connection_string description: "Additional ODBC connection string parameters. This string will get added to the connection string"
		# passphrase = null # passphrase description: "Passphrase to access the Key if one is needed"
		# password = null # password description: "The password of the user being used to authenticate"
		# ssl = null # ssl description: "If true, use SSL when connecting"
		# ssl_server_cert = null # ssl_server_cert description: "Path to server certificate to use during authentication"
		# username = null # username description: "The user to use when connecting"
	}
}