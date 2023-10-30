---
layout: "dsfhub"
page_title: "dsfhub_log_aggregator Resource"
sidebar_current: "docs-dsfhub-resource-dsfhub-log-aggregator"
description: |-
Provides a dsfhub_log_aggregator resource for AZURE EVENTHUB.
---

# dsfhub_log_aggregator (Resource)

Provides a dsfhub_log_aggregator resource for AZURE EVENTHUB.

## Example usage:

```hcl
# ### DSF Provider ###
provider "dsfhub" {
	dsfhub_token = var.dsfhub_token # TF_VAR_dsfhub_token env variable
	dsfhub_host = var.dsfhub_host # TF_VAR_dsfhub_host env variable
	#insecure_ssl = false
}

# ### Resource example for AZURE EVENTHUBresource "dsfhub_log_aggregator" "example_azure_eventhub" {
	server_type = "AZURE EVENTHUB"
	# ### required ### 
	admin_email = var.admin_email	# The email address to notify about this asset
	asset_display_name = var.asset_display_name	# User-friendly name of the asset, defined by user.
	asset_id = var.asset_id	# Asset ID
	gateway_id = var.gateway_id	# Gateway ID
	server_host_name = var.server_host_name	# Hostname (or IP if name is unknown)
	server_port = var.server_port	# Port used by the source server

	# ### optional ### 
	# application = var.application	# The Asset ID of the application asset that \"owns\" the asset.
	# asset_source = var.asset_source	# The source platform/vendor/system of the asset data. Usually the service responsible for creating that asset document
	# audit_info = var.audit_info	# Normally auto-populated when enabling the audit policy, it is a sub-document in JSON format containing configuration information for audit management. See documentation for values that can be added manually depending on asset type. Editing this value does NOT enable the audit policy.
	# audit_pull_enabled = var.audit_pull_enabled	# If true, sonargateway will collect the audit logs for this system if it can.
	# audit_type = var.audit_type # Example Values: "COSMOS_TABLE", "BLOB"	# Used to indicate what mechanism should be used to fetch logs on systems supporting multiple ways to get logs, see asset specific documentation for details
	# consumer_group = var.consumer_group	# The Consumer Group the EventHub Consumer Client will use to fetch events. Defaults to '$Default'
	# content_type = var.content_type	# Content type should be set to the desired <'parent' asset 'Server Type'>, which is the one that uses this asset as a destination for logs. NOTE: The content_type field will take precedence on the lookup for parent_asset_id field when checking which server is sending logs to this asset.
	# criticality = var.criticality # Example Values: "1", "2", "3", "4"	# The asset's importance to the business. These values are measured on a scale from \"Most critical\" (1) to \"Least critical\" (4). Allowed values: 1, 2, 3, 4
	# database_name = var.database_name	# Specifies the name of the database (or default DB) to connect to.
	# db_engine = var.db_engine	# Specifies the version of the engine being used by the database (e.g. oracle-ee, oracle-se, oracle-se1, oracle-se2)
	# enable_audit_management = var.enable_audit_management	# If true, Sonar is responsible for setting and updating the policies
	# enable_audit_monitoring = var.enable_audit_monitoring	# If true, Sonar sends emails/alerts when the audit policies change.
	# entitlement_enabled = var.entitlement_enabled	# If true, Entitlement Management system is enabled.
	# gateway_service = var.gateway_service	# The name of the gateway pull service (if any) used to retrieve logs for this source. Usually set by the connect gateway playbook.
	# jsonar_uid = var.jsonar_uid	# Unique identifier (UID) attached to the Sonar machine controlling the asset
	# location = var.location	# Current human-readable description of the physical location of the asset, or region.
	# managed_by = var.managed_by	# Email of the person who maintains the asset; can be different from the owner specified in the owned_by field. Defaults to admin_email.
	# owned_by = var.owned_by	# Email of Owner / person responsible for the asset; can be different from the person in the managed_by field. Defaults to admin_email.
	# parent_asset_id = var.parent_asset_id	# The name of an asset that this asset is part of (/related to). E.g. an AWS resource will generally have an AWS account asset as its parent. Also used to connect some log aggregating asset with the sources of their logs. E.g. An AWS LOG GROUP asset can have an AWS RDS as its parent, indicating that that is the log group for that RDS.
	# region = var.region	# For cloud systems with regions, the default region or region used with this asset
	# used_for = var.used_for # Example Values: "Production", "Test", "Development", "Demonstration", "QA", "Staging", "Training", "Disaster Recovery"	# Designates how this asset is used / the environment that the asset is supporting.
	# version = var.version	# Denotes the version of the asset
	asset_connection {
		auth_mechanism = "default"
		# ### required ### 
		azure_storage_account = null # azure_storage_account description: "The name of the unique namespace where the EventHub is located. The field can contain only lowercase letters and numbers. Name must be between 3 and 24 characters."
		azure_storage_container = null # azure_storage_container description: "Location where a given EventHub’s processing is stored (One storage container per EventHub). This name may only contain lowercase letters, numbers, and hyphens, and must begin with a letter or a number. Each hyphen must be preceded and followed by a non-hyphen character. The name must also be between 3 and 63 characters long."
		azure_storage_secret_key = null # azure_storage_secret_key description: ""
		eventhub_access_key = null # eventhub_access_key description: ""
		eventhub_access_policy = null # eventhub_access_policy description: "Authorization policy that will allow Sonar to access this specific EventHub."
		eventhub_name = null # eventhub_name description: "EventHub name without additional resource ID information."
		eventhub_namespace = null # eventhub_namespace description: "The name for the management container that the EventHub belongs to, one namespace can contain multiple EventHubs. The namespace can contain only letters, numbers, and hyphens. The namespace must start with a letter, and it must end with a letter or number. The value must be between 6 and 50 characters long."
		format = null # Example Values: "AzureSQL_Managed", "Blob", "Cosmos_Mongo", "Cosmos_SQL", "Cosmos_Table", "Databricks_Workspace", "File", "Mariadb", "Mysql", "Postgresql", "Queue", "Sql", "Synapse", "Table" # format description: "The type of audit data being sent to the EventHub, for example Postgresql or Cosmos_SQL"
		reason = null # Example Values: "default", "sonargateway" # reason description: "What this connection is used for. Used to differentiate connections if multiple connections exist for this asset"
		# ### optional ### 
		# amazon_secret = null # amazon_secret description: "Configuration to integrate with AWS Secrets Manager"
		# cyberark_secret = null # cyberark_secret description: "Configuration to integrate with CyberArk Vault"
		# hashicorp_secret = null # hashicorp_secret description: "Configuration to integrate with HashiCorp Vault"
		# ssl = null # ssl description: "If true, use SSL when connecting"
	}
	asset_connection {
		auth_mechanism = "azure_ad"
		# ### required ### 
		azure_storage_account = null # azure_storage_account description: "The name of the unique namespace where the EventHub is located. The field can contain only lowercase letters and numbers. Name must be between 3 and 24 characters."
		azure_storage_container = null # azure_storage_container description: "Location where a given EventHub’s processing is stored (One storage container per EventHub). This name may only contain lowercase letters, numbers, and hyphens, and must begin with a letter or a number. Each hyphen must be preceded and followed by a non-hyphen character. The name must also be between 3 and 63 characters long."
		eventhub_name = null # eventhub_name description: "EventHub name without additional resource ID information."
		eventhub_namespace = null # eventhub_namespace description: "The name for the management container that the EventHub belongs to, one namespace can contain multiple EventHubs. The namespace can contain only letters, numbers, and hyphens. The namespace must start with a letter, and it must end with a letter or number. The value must be between 6 and 50 characters long."
		format = null # Example Values: "AzureSQL_Managed", "Blob", "Cosmos_Mongo", "Cosmos_SQL", "Cosmos_Table", "Databricks_Workspace", "File", "Mariadb", "Mysql", "Postgresql", "Queue", "Sql", "Synapse", "Table" # format description: "The type of audit data being sent to the EventHub, for example Postgresql or Cosmos_SQL"
		reason = null # Example Values: "default", "sonargateway" # reason description: "What this connection is used for. Used to differentiate connections if multiple connections exist for this asset"
		# ### optional ### 
		# amazon_secret = null # amazon_secret description: "Configuration to integrate with AWS Secrets Manager"
		# cyberark_secret = null # cyberark_secret description: "Configuration to integrate with CyberArk Vault"
		# hashicorp_secret = null # hashicorp_secret description: "Configuration to integrate with HashiCorp Vault"
		# ssl = null # ssl description: "If true, use SSL when connecting"
	}
}
```


