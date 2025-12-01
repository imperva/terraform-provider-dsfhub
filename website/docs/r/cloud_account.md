---
subcategory: ""
layout: "dsfhub"
page_title: "DSFHUB Cloud Account - Resource"
description: |-
  Provides a dsfhub_cloud_account terraform resource.
---

# Resource: dsfhub_cloud_account

Terraform resource for managing a DSFHub cloud account. 

Documentation for the underlying APIs used in this resource can be found at
[Onboarding and Managing Cloud Accounts](https://docs-cybersec.thalesgroup.com/bundle/v15.0-sonar-user-guide/page/84557.htm).

## Cloud Account Types

<ul>
	<li><a href="https://github.com/imperva/terraform-dsfhub-agentless-onboarding/blob/main/examples/dsfhub-alibaba-cloud-account/README.md">Alibaba</li>
  <li><a href="https://github.com/imperva/terraform-dsfhub-agentless-onboarding/blob/main/examples/dsfhub-aws-cloud-account/README.md">Amazon Web Services (AWS)</a></li>
	<li><a href="https://github.com/imperva/terraform-dsfhub-agentless-onboarding/blob/main/examples/dsfhub-azure-cloud-account/README.md">Azure</a></li>
	<li><a href="https://github.com/imperva/terraform-dsfhub-agentless-onboarding/blob/main/examples/dsfhub-gcp-cloud-account/README.md">Google Cloud Platform (GCP)</a></li>
</ul>

## Example Usage

For integrated examples with cloud resource configuration and audit setup, please see the [DSFHub Agentless-Onboarding modules](https://github.com/imperva/terraform-dsfhub-agentless-onboarding).

### Basic AWS Cloud Account

Example AWS Cloud Account with the `default` authentication mechanism:

```hcl
resource "dsfhub_cloud_account" "example_aws_cloud_account" {
  server_type        = "AWS"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name" 
  asset_id           = "arn:partition:service:region:account-id"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  asset_connection {
    auth_mechanism = "default"
    reason         = "default"
    region         = "us-east-2"
  }
}
```

Example AWS Cloud Account with the `iam_role` authentication mechanism:

```hcl
resource "dsfhub_cloud_account" "example_aws_cloud_account" {
  server_type        = "AWS"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name" 
  asset_id           = "arn:partition:service:region:account-id"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  asset_connection {
    auth_mechanism = "iam_role"
    reason         = "default"
    region         = "us-east-2"
  }
}
```

Example AWS Cloud Account with the `key` authentication mechanism:

```hcl
resource "dsfhub_cloud_account" "example_aws_cloud_account" {
  server_type        = "AWS"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name" 
  asset_id           = "arn:partition:service:region:account-id"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  asset_connection {
    auth_mechanism = "key"
    reason         = "default"
    region         = "us-east-2"
    access_id      = "your-access-id"
    access_key     = "your-access-key"
  }
}
```

Example AWS Cloud Account with the `profile` authentication mechanism:

```hcl
resource "dsfhub_cloud_account" "example_aws_cloud_account" {
  server_type        = "AWS"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name" 
  asset_id           = "arn:partition:service:region:account-id"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  asset_connection {
    auth_mechanism = "profile"
    reason         = "default"
    region         = "us-east-2"
    username       = "your-username"
  }
}
```

### Basic Azure Cloud Account

Example Azure Cloud Account with the `auth_file` authentication mechanism:

```hcl
resource "dsfhub_cloud_account" "example_azure_cloud_account" {
  server_type        = "AZURE"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name" 
  asset_id           = "/subscriptions/11111111-2222-3333-4444-123456789012/asset"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  asset_connection {
    auth_mechanism = "auth_file"
    reason         = "default"
    key_file       = "your-key-file"
  }
}
```

Example Azure Cloud Account with the `client_secret` authentication mechanism:

```hcl
resource "dsfhub_cloud_account" "example_azure_cloud_account" {
  server_type        = "AZURE"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name" 
  asset_id           = "/subscriptions/11111111-2222-3333-4444-123456789012/asset"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  asset_connection {
    auth_mechanism  = "client_secret"
    reason          = "default"
    directory_id    = "11111111-2222-3333-4444-123456789012"
    application_id  = "12345678-1234-1234-1234-123456789012"
    client_secret   = "your-secret"
    subscription_id = "11111111-2222-3333-4444-123456789012"
  }
}
```

Example Azure Cloud Account with the `managed_identity` authentication mechanism:

```hcl
resource "dsfhub_cloud_account" "example_azure_cloud_account" {
  server_type        = "AZURE"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name" 
  asset_id           = "/subscriptions/11111111-2222-3333-4444-123456789012/asset"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  asset_connection {
    auth_mechanism  = "managed_identity"
    reason          = "default"
    subscription_id = "11111111-2222-3333-4444-123456789012"
  }
}
```

### Basic GCP Cloud Account

Example GCP Cloud Account with the `default` authentication mechanism:

```hcl
resource "dsfhub_cloud_account" "example_gcp_cloud_account" {
  server_type        = "GCP"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name" 
  asset_id           = "my_service_account@project-name.iam.gserviceaccount.com:project-name"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  asset_connection {
    auth_mechanism = "default"
    reason         = "default"
  }
}
```

Example GCP Cloud Account with the `service_account` authentication mechanism:

```hcl
resource "dsfhub_cloud_account" "example_gcp_cloud_account" {
  server_type        = "GCP"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name" 
  asset_id           = "my_service_account@project-name.iam.gserviceaccount.com:project-name"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  asset_connection {
    auth_mechanism = "service_account"
    reason         = "default"
    key_file       = "/path/to/gcp/credentials/service_account.json"
  }
}
```

### Basic Alibaba Cloud Account

Example Alibaba Cloud Account with the `key` authentication mechanism:

```hcl
resource "dsfhub_cloud_account" "example_alibaba_cloud_account" {
  server_type        = "ALIBABA"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name" 
  asset_id           = "arn:acs:123456789012"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  asset_connection {
    auth_mechanism = "key"
    reason         = "default"
    access_id      = "your-access-id"
    access_key     = "your-access-key"
  }
}
```

Example Alibaba Cloud Account with the `machine_role` authentication mechanism:

```hcl
resource "dsfhub_cloud_account" "example_alibaba_cloud_account" {
  server_type        = "ALIBABA"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name" 
  asset_id           = "arn:acs:123456789012"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  asset_connection {
    auth_mechanism = "machine_role"
    reason         = "default"
  }
}
```

## Argument Reference

The following arguments are required by all Cloud Account server types:

- `admin_email` - (String) The email address to notify about this asset.
- `asset_connection` - (Block) An `asset_connection` block as defined below.
- `asset_display_name` - (String) User-friendly name of the asset, defined by user.
- `asset_id` - (String) The unique identifier of the asset.
- `gateway_id` - (String) The unique identifier of the Agentless Gateway that will own the asset. Example: "12345-abcde-12345-abcde-12345-abcde". You can find the value by connecting to SonarW and running 
```
db.getSiblingDB("lmrm__sonarg").asset.find(
  { "Server Type": "IMPERVA AGENTLESS GATEWAY", "Server Host Name": "your-hostname" },
  { jsonar_uid: 1, _id: 0 }
)
```
- `server_type` - (String) The type of cloud platform to be created as a clound account. The available values are `AWS`, `ALIBABA`, `AZURE`, and `GCP`.

The following arguments are optional, however some are only supported for certain server types. Please see the [asset specifications](https://docs-cybersec.thalesgroup.com/bundle/onboarding-databases-to-sonar-reference-guide/page/Asset-Specifications_35815461.html) for more details:

- `application` - (String) The Asset ID of the application asset that "owns" the asset.
- `asset_source` - (String) The source platform/vendor/system of the asset data. Usually the service responsible for creating that asset document
- `asset_version` - (Number) Denotes the database/service version of the asset
- `available_regions` - (List of string) A list of regions to iterate through while running the Discovery playbook actions.
- `aws_proxy_config` - (Block) An `aws_proxy_config` block as defined below for an AWS proxy configuration.
- `credentials_endpoint` - (String) A specific sts endpoint to use.
- `criticality` - (Number) The asset's importance to the business. These values are measured on a scale from "Most critical" (1) to "Least critical" (4). Allowed values: 1, 2, 3, 4.
- `gateway_service` - (String) `gateway-aws@<DB type>.service` Not necessary to be set manually on the asset. Will be set by the Connect Gateway playbook.
- `jsonar_uid` - (String) Unique identifier (UID) attached to the Agentless Gateway controlling the asset.
- `location` - (String) Current human-readable description of the physical location of the asset, or region.
- `managed_by` - (String) Email of the person who maintains the asset; can be different from the owner specified in the owned_by field. Defaults to admin_email.
- `owned_by` - (String) Email of Owner / person responsible for the asset; can be different from the person in the managed_by field. Defaults to admin_email.
- `proxy` - (String) Proxy to use for AWS calls if aws_proxy_config is populated the proxy field will get populated from the http value there.
- `region` - (String) For cloud systems with regions, the default region or region used with this asset.
- `service_endpoints` - (Block) A `service_endpoints` block as defined below that specifies particular endpoints for a given service in the form of `<service name>: "endpoint"`.
- `used_for` - (String) Designates how this asset is used / the environment that the asset is supporting.

### aws_proxy_config

The following arguments are optional:

- `http` - (String) HTTP endpoint for AWS proxy config
- `https` - (String) HTTPS endpoint for AWS proxy config

### service_endpoints

The following argument is optional:

- `logs` - (String) The log endpoint for a given service

### asset_connection

A minimum of one `asset_connection` blocks is required.

The following arguments are required:

- `auth_mechanism` - (String) Specifies the auth mechanism used by the connection
- `reason` - (String) Used to differentiate between connections belonging to the same asset. Use "default" or "sonargateway" for connections necessary for audit pull.
- `region` - (String) Default AWS region for this asset

The following arguments are optional, however some are only supported for certain server types and authentication mechanism combinations. Please see the [asset specifications](https://docs-cybersec.thalesgroup.com/bundle/onboarding-databases-to-sonar-reference-guide/page/Asset-Specifications_35815461.html) for more details:

- `access_id` - (String) The Access key ID of AWS secret access key used for authentication
- `access_key` - (String) The Secret access key used for authentication
- `amazon_secret` - An `amazon_secret` block as defined below, to integrate the asset with AWS Secrets Manager.
- `application_id` - (String) This is also referred to as the Client ID and it’s the unique identifier for the registered application being used to execute Python SDK commands against Azure’s API services. You can find this number under Azure Active Directory -> App Registrations -> Owned Applications
- `ca_certs_path` - (String) Certificate authority certificates path; what location should the sysetm look for certificate information from. Equivalent to --capath in a curl call
- `client_secret` - (String) This a string containing a secret used by the application to prove its identity when requesting a token. You can get a secret by going to Azure Active Directory -> App Registrations -> Owned Applications, selecting the desired application and then going to Certificates & secrets -> Client secrets -> + New client secret
- `cyberark_secret` - A `cyberark_secret` block as defined below, to integrate the asset with CyberArk.
- `directory_id` - (String) This is also referred to as the Tenant ID and is a GUID representing the Active Directory Tenant. It can be found in the Azure Active Directory page under the Azure portal
- `external_id` - (String) External ID to use when assuming a role
- `hashicorp_secret` - A `hashicorp_secret` block as defined below, to integrate the asset with a HashiCorp Vault.
- `key_file` - (String) Location on disk on the key to be used for authentication
- `project_id` - (String) Used when running Sonar on a GCP hosted environment that doesn't have a service account linked to it
- `role_name` - (String) What role is used to get credentials from.
- `secret_key` - (String) The Secret access key used for authentication
- `session_token` - (String) STS token used for session authentication
- `ssl` (Boolean) If true, use SSL when connecting
- `subscription_id` - (String) This is the Azure account subscription ID. You can find this number under the Subscriptions page on the Azure portal
- `username` - (String) The name of a profile in `${JSONAR_LOCALDIR}/credentials/.aws/credentials` to use for authenticating. The value of `$JSONAR_LOCALDIR` can be found in `/etc/sysconfig/jsonar` on your DSF machine.

#### AWS Secret Manager: `asset_connection.amazon_secret`

A maximum of one block is supported.

The following arguments are optional:

- `field_mapping` (Map of string) Field mapping for AWS secret
- `secret_asset_id` - (String) AWS secret manager asset_id
- `secret_name` - (String) AWS secret name

#### CyberArk Secret Manager: `asset_connection.cyberark_secret`

A maximum of one block is supported.

The following arguments are optional:

- `field_mapping` (Map of string) Field mapping for CyberArk secret
- `secret_asset_id` - (String) CyberArk secret manager asset id
- `secret_name` - (String) CyberArk secret name

#### HashiCorp Secret Manager: `asset_connection.hashicorp_secret`

The following arguments are optional:

- `field_mapping` (Map of string) Field mapping for HashiCorp secret
- `path` - (String) HashiCorp secret path
- `secret_asset_id` - (String) HashiCorp secret manager asset id
- `secret_name` - (String) HashiCorp secret name

## Import

In Terraform v1.5.0 and later, use an import block to import Cloud Accounts using the `asset_id`. For example:

```
import {
  to = dsfhub_cloud_account.example
  id = "arn:partition:service:region:account-id"
}
```

Using terraform import, import Cloud Accounts using the `asset_id`. For example:

```
$ terraform import dsfhub_cloud_account.example_aws_cloud_account "arn:partition:service:region:account-id"
```

For detailed instructions on onboarding existing cloud resources to DSF using Terraform's import functionality, see [Importing and Onboarding Existing Data Sources with Terraform](https://docs-cybersec.thalesgroup.com/bundle/onboarding-databases-to-sonar-reference-guide/page/Importing-and-Onboarding-Existing-Data-Sources-with-Terraform_784990209.html).