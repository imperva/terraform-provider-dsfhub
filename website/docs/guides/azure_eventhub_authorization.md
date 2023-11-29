---
subcategory: "Agentless Gateway Permissions"
layout: "dsfhub"
page_title: "Azure Authorization - EventHub"
description: |-
  Azure Authorization - EventHub
---

# Agentless Gateway Authorization Permissions - Azure EventHub

The DSF Agentless Gateway requires the following Authorization permissions to access an Azure EventHub Namespace.

## Variable Reference

```
# Variables
variable "deployment-name" {
  description =  "Specifies deployment like eventhubnamespace, storage container etc"
  type = string
  default = "mydeployment"
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
```

## Example Usage

```hcl
### Azure Provider ###
provider "azurerm" {
  features {}
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
  
  log {
    category  = "MySqlAuditLogs"
    enabled = true

    retention_policy {
      enabled = false
    }
  } 
}
```
