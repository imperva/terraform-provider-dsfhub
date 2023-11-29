---
subcategory: "Example Assets"
layout: "dsfhub"
page_title: "AZURE MYSQL - Event Hub"
description: |-
  Provides a combined example of creating an AZURE MYSQL database, associated configurations for audit logs, onboarding to the DSFHUB with necessary configs for the DSF Agentless Gateway to access.
---

# Azure MYSQL Onboarding Template

Provides a module template for creating an AWS RDS MYSQL database, the associated audit log, firewall, event hub/namespace, diagnostic and storage account configurations, creating the [dsfhub_data_source](../r/data_source.md) and [dsfhub_log_aggregator](../r/log_aggregator.md) records to onboard to the DSFHUB with necessary access for the DSF Agentless Gateway.

<details>
<summary>AZURE MYSQL Variables</summary>

## AZURE MYSQL Variables

```hcl
# DSFHUB Provider Required Variables
variable "dsfhub_token" {} # TF_VAR_dsfhub_token env variable
variable "dsfhub_host" {} # TF_VAR_dsfhub_host env variable

# DSFHUB Asset Variables
variable "admin_email" {
  description = "The email address to notify about this asset"
  type = string
  default = "your@email.com"
}

variable "gateway_id" {
  description =  "The jsonarUid unique identifier of the agentless gateway. Example: '7a4af7cf-4292-89d9-46ec-183756ksdjd'"
  type = string
  default = "12345abcde-12345-abcde-12345-12345abcde"
}

# Azure Database Variables
variable "administrator_login" {
  description =  "The Administrator login for the MySQL Server."
  type = string
  default = "sonaradmin"
}

variable "administrator_login_password" {
  description =  "The Password associated with the administrator_login for the MySQL Server."
  type = string
  default = ""
}

variable "deployment-name" {
  description =  "Specifies deployment like eventhubnamespace, storage container etc"
  type = string
  default = "mydeployment"
}

variable "db_name" {
  description =  "Specifies the name of the MySQL Server. Changing this forces a new resource to be created. This needs to be globally unique within Azure."
  type = string
  default = "Azuremysqldb"
}

variable "db_version" {
  description =  " Specifies the version of MySQL to use. Valid values are 5.7, or 8.0."
  type = string
  default = "8.0"
}

variable "location" {
  description =  "Specifies the supported Azure location where the resource exists."
  type = string
  default = "East US"
}

variable "message_retention" {
  description =  "Specifies the number of days to retain the events for this Event Hub.."
  type = number
  default = 1
}

variable "partition_count" {
  description =  "Specifies the current number of shards on the Event Hub."
  type = number
  default = 2
}

variable "resource_group_name" {
  description =  " The name of the resource group in which to create the MySQL Server."
  type = string
  default = "My_Resource_group"
}

variable "sku" {
  description =  "Defines which tier to use. Valid options are Basic, Standard, and Premium. Please note that setting this field to Premium will force the creation of a new resource."
  type = string
  default = "Basic"
}

variable "sku_name" {
  description =  "Specifies the SKU Name for this MySQL Server. The name of the SKU, follows the tier + family + cores pattern (e.g. B_Gen4_1, GP_Gen5_8). For more information see the product documentation. Possible values are B_Gen4_1, B_Gen4_2, B_Gen5_1, B_Gen5_2, GP_Gen4_2, GP_Gen4_4, GP_Gen4_8, GP_Gen4_16, GP_Gen4_32, GP_Gen5_2, GP_Gen5_4, GP_Gen5_8, GP_Gen5_16, GP_Gen5_32, GP_Gen5_64, MO_Gen5_2, MO_Gen5_4, MO_Gen5_8, MO_Gen5_16 and MO_Gen5_32"
  type = string
  default = "B_Gen5_1"
}

variable "storage_mb" {
  description =  "Max storage allowed for a server. Possible values are between 5120 MB(5GB) and 1048576 MB(1TB) for the Basic SKU and between 5120 MB(5GB) and 16777216 MB(16TB) for General Purpose/Memory Optimized SKUs. "
  type = number
  default = 20480
}

variable "storage_account_tier" {
  description =  " Defines the Tier to use for this storage account. Valid options are Standard and Premium. For BlockBlobStorage and FileStorage accounts only Premium is valid."
  type = string
  default = "Standard"
}

# Azure Firewall Variables
variable "start_ip" {
  description =  "Specifies the Start IP Address associated with this Firewall Rule."
  type = string
  default = "127.0.0.1"
}

variable "end_ip" {
  description =  "Specifies the End IP Address associated with this Firewall Rule."
  type = string
  default = "127.0.0.1"
}
```
</details>

## Providers and Resources

```hcl
# Configure the Azure provider
terraform {
  required_providers {
    azurerm = {
      source  = "hashicorp/azurerm"
      version = "~> 3.0.2"
    }
    dsfhub = {
      source = "imperva/dsfhub"
    }
  }
  required_version = ">= 1.1.0"
}

provider "azurerm" {
  features {}
}

provider "dsfhub" {
	dsfhub_token = var.dsfhub_token # TF_VAR_dsfhub_token env variable
	dsfhub_host = var.dsfhub_host # TF_VAR_dsfhub_host env variable
	#insecure_ssl = false
}

# Configure Azure Mysql Server
resource "azurerm_mysql_server" "azure_mysql_server" {
  name                              = var.db_name
  location                          = var.location
  resource_group_name               = var.resource_group_name
  sku_name                          = var.sku_name
  administrator_login               = var.administrator_login
  administrator_login_password      = var.administrator_login_password
  storage_mb                        = var.storage_mb
  version                           = var.db_version
  public_network_access_enabled     = true
  ssl_enforcement_enabled           = true
  ssl_minimal_tls_version_enforced  = "TLS1_2"
}


# Enable Audit Log
resource "azurerm_mysql_configuration" "audit_log_enabled" {
    depends_on          = [azurerm_mysql_server.azure_mysql_server]
    name                = "audit_log_enabled"
    resource_group_name = var.resource_group_name
    server_name         = azurerm_mysql_server.azure_mysql_server.name
    value               = "ON"
    timeouts {
    	read = "10m"
  	}
}

resource "azurerm_mysql_configuration" "audit_log_events" {
    depends_on          = [azurerm_mysql_configuration.audit_log_enabled]
    name                = "audit_log_events"
    resource_group_name = var.resource_group_name
    server_name         = azurerm_mysql_server.azure_mysql_server.name
    value               = "CONNECTION,GENERAL,TABLE_ACCESS"
    timeouts {
    	read = "10m"
  	}
}

# Configure Azure Firewall Rule
resource "azurerm_mysql_firewall_rule" "example" {
  depends_on  = [azurerm_mysql_configuration.audit_log_events]
  name                = "${var.deployment-name}-mysql-firewallrule"
  resource_group_name = var.resource_group_name
  server_name         = azurerm_mysql_server.azure_mysql_server.name
  start_ip_address    = var.start_ip
  end_ip_address      = var.end_ip
}

# Configure Eventhub Namespace and Eventhub
resource "azurerm_eventhub_namespace" "azure_eventhub_namepaces" {
  name                = "${var.deployment-name}-eventhubnamespace"
  location            = var.location
  resource_group_name = var.resource_group_name
  sku                 = var.sku
  capacity            = 2
}

resource "azurerm_eventhub" "azure_eventhubs" {
  depends_on = [azurerm_eventhub_namespace.azure_eventhub_namepaces]
  name                = "${var.deployment-name}-mysql-eventhub"
  namespace_name      = azurerm_eventhub_namespace.azure_eventhub_namepaces.name
  resource_group_name = var.resource_group_name
  partition_count     = var.partition_count
  message_retention   = var.message_retention
}

# Setting up diagnostic setting
data "azurerm_eventhub_namespace_authorization_rule" "logging" {
  name                = "RootManageSharedAccessKey"
  namespace_name      = azurerm_eventhub_namespace.azure_eventhub_namepaces.name
  resource_group_name = var.resource_group_name
}

resource "azurerm_monitor_diagnostic_setting" "azurerm_monitor_diagnostic_settings" {
  depends_on                     = [azurerm_eventhub.azure_eventhubs]
  name                           = "${var.deployment-name}-mysql-diag-setting"
  target_resource_id             = azurerm_mysql_server.azure_mysql_server.id
  eventhub_name                  = azurerm_eventhub.azure_eventhubs.name
  eventhub_authorization_rule_id = data.azurerm_eventhub_namespace_authorization_rule.logging.id
  #eventhub_authorization_rule_id = "RootManageSharedAccessKey"
  log {
    category  = "MySqlAuditLogs"
    enabled = true

    retention_policy {
      enabled = false
    }
  } 
}

# Configure storage account and Storage container
resource "azurerm_storage_account" "azurerm_storage_accounts" {
  name                     = "${var.deployment-name}storageaccount"
  resource_group_name      = var.resource_group_name
  location                 = var.location
  account_tier             = var.storage_account_tier
  account_replication_type = "LRS"
}

resource "azurerm_storage_container" "azurerm_storage_containers" {
  name                  = "${var.deployment-name}-mysql-storagecontainer"
  storage_account_name  = azurerm_storage_account.azurerm_storage_accounts.name
  container_access_type = "private"
}

# Configure Azure Mysql asset for DSFHUB
resource "dsfhub_data_source" "example_azure_mysql" {
	server_type               = "AZURE MYSQL"
	admin_email               = var.admin_email	
	asset_display_name        = azurerm_mysql_server.azure_mysql_server.name
	asset_id                  = azurerm_mysql_server.azure_mysql_server.id
	gateway_id                = var.gateway_id
	location                  = var.location	
	server_host_name          = azurerm_mysql_server.azure_mysql_server.fqdn
    logs_destination_asset_id = azurerm_eventhub.azure_eventhubs.id
    audit_pull_enabled        = true
	asset_connection {
		auth_mechanism        = "password"
		password              = var.administrator_login_password
		reason                = "default"
		username              = var.administrator_login
	}
}

resource "dsfhub_log_aggregator" "example_azure_eventhub" {
	server_type        = "AZURE EVENTHUB"
	admin_email        = var.admin_email	
	asset_display_name = azurerm_eventhub.azure_eventhubs.name
	asset_id           = azurerm_eventhub.azure_eventhubs.id
	gateway_id         = var.gateway_id
	server_host_name   = "${azurerm_eventhub_namespace.azure_eventhub_namepaces.name}.servicebus.windows.net"
	server_port        = "443"
	asset_connection {
		auth_mechanism           = "default"
		azure_storage_account    = azurerm_storage_account.azurerm_storage_accounts.name
		azure_storage_container  = azurerm_storage_container.azurerm_storage_containers.name
		azure_storage_secret_key = azurerm_storage_account.azurerm_storage_accounts.primary_access_key
		eventhub_access_key      = azurerm_eventhub_namespace.azure_eventhub_namepaces.default_primary_key 
		eventhub_access_policy   = "RootManageSharedAccessKey" 
		eventhub_name            = azurerm_eventhub.azure_eventhubs.name 
		eventhub_namespace       = azurerm_eventhub_namespace.azure_eventhub_namepaces.name
		format                   = "Mysql"
		reason                   = "default"
	}
}
```

## Agentless Gateway Permission Dependencies:

The [DSF Agentless Gateway](https://registry.terraform.io/modules/imperva/dsf-agentless-gw/aws/latest) is required to have the following [Event Hub Authorization Rule](https://registry.terraform.io/providers/hashicorp/azurerm/latest/docs/resources/eventhub_authorization_rule) access to database audit.

<ul>
<li><a target="_blank" href="azure_eventhub_authorization.md">Event Hub Authorization Rule for EventHub Namespaces</a></li>
</ul>