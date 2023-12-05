---
subcategory: ""
layout: "dsfhub"
page_title: "Manage Bulk Assets from CSV"
description: |-
  Manage Bulk Assets from CSV
---

# Manage Bulk Assets from CSV

Terraform can natively import csv data using the [csvdecode](https://www.terraform.io/docs/language/functions/csvdecode.html) function. The following example shows how to use the csvdecode function to manage [dsf_data_source](https://registry.terraform.io/modules/imperva/dsf-agentless-gw/aws/latest) resources in bulk from a csv file.

<details>
<summary>Example CSV file format</summary>

Create a csv file with the following format.  The first row is the header row and the remaining rows are the asset data.  The header row is used to map the column data to the asset attributes.

```csv
id,asset_id,jsonar_uid,asset_display_name,Server Type,Server IP,Server Host Name,Service Name,Server Port,version,audit_type,auth_mechanism,username,password,reason,admin_email
1,my.hostname1:ORACLE:ORA19C:1521,ABCDE-12345-ABCDE-12345,my.hostname1:ORACLE:ORA19C:1521,ORACLE,0.0.0.0,my.hostname1:ORACLE:ORA19C:1521,my-ora-service-name,3202,19,UNIFIED,kerberos,test,test,sonargateway,your@email.com
2,my.hostname2:ORACLE:ORA19C:1521,ABCDE-12345-ABCDE-12345,my.hostname2:ORACLE:ORA19C:1521,ORACLE,0.0.0.0,my.hostname2:ORACLE:ORA19C:1521,my-ora-service-name,3202,19,UNIFIED,password,admin,password,sonargateway,your@email.com
```
</details>

## Example Bulk Import Usage

<details>
<summary>Example Variables for Bulk Import</summary>

## Example Variables for Bulk Import

```
# DSFHUB Provider Required Variables
variable "dsfhub_token" {}
variable "dsfhub_host" {}

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

variable "csv_file_path" {
	description =  "Path to the csv file to import"
	type = string
	default = "sample_assets.csv"
}
```
</details>

## Proviers and Resources for Bulk Import

```hcl
# Specify path for provider
terraform {
  required_providers {
    dsfhub = {
      source = "imperva/dsfhub"
    }
  }
}
provider "dsfhub" {
	dsfhub_token = var.dsfhub_token
	dsfhub_host = var.dsfhub_host
}

locals {
	asset_csv = csvdecode(file("${path.module}/${var.csv_file_path}"))
}

# ### Resource example for bulk import ###
resource "dsfhub_data_source" "bulk-database-import" {
	for_each = { for asset in local.asset_csv : asset.asset_id => asset }
	server_type = each.value["Server Type"]

	admin_email         = each.value.admin_email
	asset_display_name  = each.value.asset_display_name
	asset_id            = each.value.asset_id
	audit_pull_enabled  = true
	audit_type			= each.value.audit_type
	gateway_id          = each.value.jsonar_uid
	server_host_name    = each.value["Server Host Name"]
	server_ip           = each.value["Server IP"]
	server_port         = each.value["Server Port"]
	service_name		= each.value["Service Name"]
	version             = each.value.version

	dynamic "asset_connection" {
    	for_each = each.value.auth_mechanism=="password" ? [1] : []
    	content {
			auth_mechanism  = each.value.auth_mechanism
			password        = each.value.password
			reason          = each.value.reason
			username        = each.value.username
    	}
  	}

	dynamic "asset_connection" {
    	for_each = each.value.auth_mechanism=="kerberos" ? [1] : []
    	content {
			auth_mechanism  = each.value.auth_mechanism
			reason          = each.value.reason
    	}
  	}
}
```

## Bulk Import of Exising Assets into Terraform State

For exising assets that are already managed by DSF, you can import them using the [DSF-CLI](https://github.com/imperva/dsf-cli) into Terraform state using the following command.  Execute the following command in the same directory where you have the bulk import tf file shown above.

```
for ASSET_ID in $(dsf data_source read | jq -r '.data[].id'); do terraform import "dsfhub_data_source.bulk-database-import[\"$ASSET_ID\"]" "${ASSET_ID}"; done
```


