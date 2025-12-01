---
subcategory: ""
layout: "dsfhub"
page_title: "DSFHUB Secret Manager Resource"
description: |-
  Provides a dsfhub_secret_manager terraform resource.
---

# Resource: dsfhub_secret_manager

Terraform resource for managing a DSFHub secret manager. 

A secret manager is a service used to store configuration information such as database passwords, API keys, or TLS certificates needed by an application at runtime.

The `dsfhub_secret_manager` resource supports the configuration parameters necessary to onboard a secret manager to DSF Hub. When configuring connections for assets (including data sources, cloud accounts, log aggregators and secret managers) in the DSF Hub, you have the option of using a secret manager as the source for fields such as user credentials or other secrets. If you use this option, certain configuration fields for a connection can be populated from an onboarded secret manager rather than from the configured fields in the asset itself. Documentation for the underlying API used in this resource can be found at [Onboarding and Managing Secret Managers](https://docs-cybersec.thalesgroup.com/bundle/v15.0-sonar-user-guide/page/85338.htm).

## Secret Manager Types

<ul>
	<li><a href="https://github.com/imperva/terraform-dsfhub-agentless-onboarding/blob/main/examples/dsfhub-aws-secret-manager/README.md">Amazon Web Services (AWS)</a></li>
	<li><a href="https://github.com/imperva/terraform-dsfhub-agentless-onboarding/blob/main/examples/dsfhub-cyberark-secret-manager/README.md">CyberArk</a></li>
	<li><a href="https://github.com/imperva/terraform-dsfhub-agentless-onboarding/blob/main/examples/dsfhub-hashicorp-secret-manager/README.md">HashiCorp</a></li>
</ul>

## Example Usage

For integrated examples with cloud resource configuration and audit setup, please see the [DSFHub Agentless-Onboarding modules](https://github.com/imperva/terraform-dsfhub-agentless-onboarding).

### Basic AWS Secret Manager

Example AWS Secret Manager with the `default` authentication mechanism:

```hcl
resource "dsfhub_secret_manager" "example_aws_secret_manager" {
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

Example AWS Secret Manager with the `iam_role` authentication mechanism:

```hcl
resource "dsfhub_secret_manager" "example_aws_secret_manager" {
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

Example AWS Secret Manager with the `key` authentication mechanism:

```hcl
resource "dsfhub_secret_manager" "example_aws_secret_manager" {
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

Example AWS Secret Manager with the `profile` authentication mechanism:

```hcl
resource "dsfhub_secret_manager" "example_aws_secret_manager" {
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

### Basic CyberArk Secret Manager

Example CyberArk Secret Manager with the `default` authentication mechanism:

```hcl
resource "dsfhub_secret_manager" "example_cyberark_secret_manager" {
  server_type        = "CYBERARK"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name" 
  asset_id           = "CyberArk.hostname:CYBERARK::13898"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  server_hostname    = "https://your_CyberArk_server_hostname.com"
  server_ip          = "https://your_CyberArk_server_hostname.com"
  asset_connection {
    auth_mechanism = "default"
    reason         = "default"
    query          = "AppID=<your_CyberArk_Application_ID>&Safe=<your_CyberArk_Safe_Name>;Folder=Root;"
  }
}
```

### Basic HashiCorp Secret Manager

Example HashiCorp Secret Manager with the `app_role` authentication mechanism:

```hcl
resource "dsfhub_secret_manager" "example_hashicorps_secret_manager" {
  server_type        = "HASHICORP"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name" 
  asset_id           = "HashiCorp.hostname:HASHICORP::8200"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  server_hostname    = "https://your_hashicorp_server_hostname.com"
  server_ip          = "1.2.3.4"
  server_port        = 8200
  asset_connection {
    auth_mechanism = "app_role"
    reason         = "default"
    role_name      = "your-role"
    secret_key     = "A1b2C3d4/H5i6J7k8/A1b2C3d4/H5i6J7k8"
  }
}
```

Example HashiCorp Secret Manager with the `ec2` authentication mechanism:

```hcl
resource "dsfhub_secret_manager" "example_hashicorps_secret_manager" {
  server_type        = "HASHICORP"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name" 
  asset_id           = "HashiCorp.hostname:HASHICORP::8200"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  server_hostname    = "https://your_hashicorp_server_hostname.com"
  server_ip          = "1.2.3.4"
  server_port        = 8200
  asset_connection {
    auth_mechanism = "ec2"
    reason         = "default"
    role_name      = "your-role"
  }
}
```

Example HashiCorp Secret Manager with the `iam_role` authentication mechanism:

```hcl
resource "dsfhub_secret_manager" "example_hashicorps_secret_manager" {
  server_type        = "HASHICORP"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name" 
  asset_id           = "HashiCorp.hostname:HASHICORP::8200"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  server_hostname    = "https://your_hashicorp_server_hostname.com"
  server_ip          = "1.2.3.4"
  server_port        = 8200
  asset_connection {
    auth_mechanism    = "iam_role"
    reason            = "default"
    access_id         = "ABCDE12345ABCDE12345"
    aws_iam_server_id = "vault.example.com"
    role_name         = "your-role"
    secret_key        = "A1b2C3d4/H5i6J7k8/A1b2C3d4/H5i6J7k8"
  }
}
```

Example HashiCorp Secret Manager with the `root_token` authentication mechanism:

```hcl
resource "dsfhub_secret_manager" "example_hashicorps_secret_manager" {
  server_type        = "HASHICORP"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-display-name" 
  asset_id           = "HashiCorp.hostname:HASHICORP::8200"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  server_hostname    = "https://your_hashicorp_server_hostname.com"
  server_ip          = "1.2.3.4"
  server_port        = 8200
  asset_connection {
    auth_mechanism = "root_token"
    reason         = "default"
    secret_key     = "A1b2C3d4/H5i6J7k8/A1b2C3d4/H5i6J7k8"
  }
}
```

### Integrated Usage with dsfhub_data_source

Example HashiCorp Secret Manager for a Sybase Data Source, with the Kerberos Authentication Mechanism.
* `audit_pull_enabled: true` on the data source asset triggers the Connect Gateway playbook to run against it.

```hcl
resource "dsfhub_secret_manager" "example_hashicorp" {
  server_type        = "HASHICORP"
  admin_email        = "somebody@company.com"
  asset_display_name = "example-hashicorp-asset" 
  asset_id           = "HashiCorp.hostname:HASHICORP::8200"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  server_host_name   = "https://your_hashicorp_server_hostname.com"
  server_ip          = "1.2.3.4"
  server_port        = 8200
  asset_connection {
    auth_mechanism    = "iam_role"
    access_id         = "ABCDE12345ABCDE12345"
    aws_iam_server_id = "your.vault.example.com"
    reason            = "default"
    role_name         = "your-role"
    secret_key        = "A1b2C3d4/H5i6J7k8/A1b2C3d4/H5i6J7k8"
  }
}

resource "dsfhub_data_source" "example_sybase_kerberos" {
  server_type = "SYBASE"

  admin_email        = "somebody@company.com"
  asset_display_name = "example-sybase-kerberos-asset"
  asset_id           = "example-server-host-name:SYBASE::5000"
  audit_pull_enabled = true
  database_name      = "sybsecurity"
  gateway_id         = "12345-abcde-12345-abcde-12345-abcde"
  server_host_name   = "example-server-host-name"
  server_ip          = "example-server-host-ip"
  server_port        = "5000"
  asset_version            = "16"

  asset_connection {
    auth_mechanism = "kerberos"

    reason       = "default"
    cache_file   = "/path/to/sybase_ticket"
    external     = false
    kerberos_kdc = "x.x.x.x"
    kerberos_spn = "sybase_service_principal@DOMAIN.COM"
    password     = "dummy_password_val"
    principal    = "dummy_principal_val"

    hashicorp_secret {
      secret_asset_id = dsfhub_secret_manager.example_hashicorp.asset_id
      path            = "your-secret-path"
      secret_name     = "your-secret-name"

      field_mapping {
        password  = "your-remote-password"
        principal = "your-remote-principal"
      }
    }
  }
}
```

## Argument Reference

The following arguments are required by all Secret Manager server types:

- `admin_email` - (String) The email address to notify about this asset
- `asset_display_name` - (String) User-friendly name of the asset, defined by user.
- `asset_id` - (String) The unique identifier of the asset.
- `gateway_id` - (String) The unique identifier of the Agentless Gateway that will own the asset. Example: "12345-abcde-12345-abcde-12345-abcde". You can find the value by connecting to SonarW and running 
```
db.getSiblingDB("lmrm__sonarg").asset.find(
  { "Server Type": "IMPERVA AGENTLESS GATEWAY", "Server Host Name": "your-hostname" },
  { jsonar_uid: 1, _id: 0 }
)
```
- `server_type` - (String) The type of cloud platform or service to be created as a secret manager. The available values are `AWS`, `CYBERARK` and `HASHICORP`.

The following arguments are optional, however some are only supported for certain server types. Please see the [asset specifications](https://docs-cybersec.thalesgroup.com/bundle/onboarding-databases-to-sonar-reference-guide/page/Asset-Specifications_35815461.html) for more details:

- `application` - (String) The Asset ID of the application asset that "owns" the asset.
- `asset_connection` - (Block) An `asset_connection` block as defined below.
- `asset_source` - (String) The source platform/vendor/system of the asset data. Usually the service responsible for creating that asset document
- `asset_version` - (Number) Denotes the database/service version of the asset
- `available_regions` - (List of string) A list of regions to iterate through while running the Discovery playbook actions.
- `aws_proxy_config` - (Block) An `aws_proxy_config` block as defined below for an AWS proxy configuration.
- `credentials_endpoint` - (String) A specific sts endpoint to use
- `criticality` - (Number) The asset's importance to the business. These values are measured on a scale from "Most critical" (1) to "Least critical" (4). Allowed values: 1, 2, 3, 4
- `jsonar_uid` - (String) Unique identifier (UID) attached to the Agentless Gateway controlling the asset
- `jsonar_uid_display_name` - (String) Unique identifier (UID) attached to the Agentless Gateway controlling the asset
- `location` - (String) Current human-readable description of the physical location of the asset, or region.
- `managed_by` - (String) Email of the person who maintains the asset; can be different from the owner specified in the owned_by field. Defaults to admin_email.
- `owned_by` - (String) Email of Owner / person responsible for the asset; can be different from the person in the managed_by field. Defaults to admin_email.
- `proxy` - (String) Proxy to use for AWS calls if aws_proxy_config is populated the proxy field will get populated from the http value there
- `region` - (String) For cloud systems with regions, the default region or region used with this asset
- `server_ip` - (String) IP address of the service where this asset is located. If no IP is available populate this field with other information that would identify the system e.g. hostname or AWS ARN, etc.
- `server_host_name` - (String) Hostname (or IP if name is unknown)
- `server_port` - (String) Port used by the source server
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

The following arguments are required:

- `auth_mechanism` - (String) Specifies the auth mechanism used by the connection
- `reason` - (String) Used to differentiate between connections belonging to the same asset. Use "default" or "sonargateway" for connections necessary for audit pull.

The following arguments are optional, however some are only supported for certain server types and authentication mechanism combinations. Please see the [asset specifications](https://docs-cybersec.thalesgroup.com/bundle/onboarding-databases-to-sonar-reference-guide/page/Asset-Specifications_35815461.html) for more details:

- `access_id` - (String) The Access key ID of AWS secret access key used for authentication
- `amazon_secret` - An `amazon_secret` block as defined below, to integrate the asset with AWS Secrets Manager.
- `aws_iam_server_id` - (String) e.g. vault.example.com
- `ca_certs_path` - (String) Certificate authority certificates path; what location should the sysetm look for certificate information from. Equivalent to --capath in a curl call
- `cert_file` - (String) Certificate used for access
- `cyberark_secret` - A `cyberark_secret` block as defined below, to integrate the asset with CyberArk.
- `external_id` - (String) External ID to use when assuming a role
- `hashicorp_secret` - A `hashicorp_secret` block as defined below, to integrate the asset with a HashiCorp Vault.
- `headers` - (List of string) Additional parameters to pass as HTTP headers when fetching credentials. Example: `["HEADER1: value1", "HEADER2: value2"]`
- `key_file` - (String) Path to Key used for accessing CyberArk
- `namespace` - (String) Specifies which HashiCorp namespace to fetch credentials from if not root.
- `nonce` - (String) Stored nonce, automatically added from first call if not already generated
- `protocol` - (String) A protocol prefix to use to connect if it isn't already specified in the host name e.g http:// or https://. Defaults to "http://"
- `query` - (String) Query parameters defining where the passwords, etc. should be retrieved from. Example: `AppID=<your_CyberArk_Application_ID>&Safe=<your_CyberArk_Safe_Name>;Folder=Root;` (This is everything to be included in the curl string from `Accounts?` and `Object=` in the call to retrieve data from CyberArk.)
- `region` - (String) Default AWS region for this asset
- `role_name` - (String) Role to use for authentication
- `secret_key` - (String) The Secret access key used for authentication
- `self_signed` - (String) Connection using the -k flag to accept self signed certificates
- `session_token` - (String) STS token used for session authentication
- `ssl` - (Boolean) If true, use SSL when connecting
- `store_aws_credentials` - (Boolean) Specifies whether this connection is used with the Hashicorp AWS integration (Dynamic Secret) to generate temporary key pairs for this system to use. Set to True when using Hashicorp Dynamic.
- `username` - (String) The name of a profile in `${JSONAR_LOCALDIR}/credentials/.aws/credentials` to use for authenticating. The value of `$JSONAR_LOCALDIR` can be found in `/etc/sysconfig/jsonar` on your DSF machine.
- `v2_key_engine` - (String) Indicates whether the HashiCorp Key/Value (KV) version 2 secrets engine is used.

The following secret manager blocks are optional (this relationship allows secret managers to store their secrets in another secret manager):
* amazon_secret
* cyberark_secret
* hashicorp_secret

#### AWS Secret Manager: `asset_connection.amazon_secret`

A maximum of one block is supported.

The following arguments are optional:

- `field_mapping`- (Map of string) Field mapping for AWS secret
- `secret_asset_id` - (String) AWS secret manager asset_id
- `secret_name` - (String) AWS secret name

#### CyberArk Secret Manager: `asset_connection.cyberark_secret`

A maximum of one block is supported.

The following arguments are optional:

- `field_mapping` - (Map of string) Field mapping for CyberArk secret
- `secret_asset_id` - (String) CyberArk secret manager asset_id
- `secret_name` - (String) CyberArk secret name

#### HashiCorp Secret Manager: `asset_connection.hashicorp_secret`

The following arguments are optional:

- `field_mapping` (Map of string) Field mapping for HashiCorp secret
- `path` - (String) HashiCorp secret path
- `secret_asset_id` - (String) HashiCorp secret manager asset_id
- `secret_name` - (String) HashiCorp secret name

## Import

In Terraform v1.5.0 and later, use an import block to import Secret Managers using the `asset_id`. For example:

```
import {
  to = dsfhub_secret_manager.example_secret_manager_aws
  id = "arn:partition:service:region:account-id"
}
```

Using terraform import, import Secret Managers using the `asset_id`. For example:

```
$ terraform import dsfhub_secret_manager.example_secret_manager_aws "arn:partition:service:region:account-id"
```

For detailed instructions on onboarding existing cloud resources to DSF using Terraform's import functionality, see [Importing and Onboarding Existing Data Sources with Terraform](https://docs-cybersec.thalesgroup.com/bundle/onboarding-databases-to-sonar-reference-guide/page/Importing-and-Onboarding-Existing-Data-Sources-with-Terraform_784990209.html).
