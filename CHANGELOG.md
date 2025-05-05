## 1.3.7 (Unreleased)

ENHANCEMENTS:

* resource/data_source: added support of the available_bucket_account_ids attribute
* resource/data_source: added support for AWS OPENSEARCH, AWS REDSHIFT SERVERLESS, AWS RDS POSTGRESQL CLUSTER, AZURE DATABRICKS WORKSPACE, AZURE DATA EXPLORER, AZURE MYSQL FLEXIBLE, GCP ALLOYDB POSTGRESQL CLUSTER and GCP ALLOYDB POSTGRESQL server types

## 1.3.6 (November 20, 2024)

ENHANCEMENTS:

* resource/data_source,log_aggregator: updated audit-collection logic to wait for asset to be synced to gateways before connecting/disconnecting

## 1.3.5 (November 12, 2024)

ENHANCEMENTS:

* resource/log_aggregator: added support for content_type attribute

## 1.3.4 (October 30, 2024)

ENHANCEMENTS:

* resource/data_source,log_aggregator: updated audit-collection logic to wait for audit_pull_enabled attribute to be set correctly

NOTES:

* removed vendor directory from source control

## 1.3.3 (October 7, 2024)

BUG FIXES:

* resource/cloud_account: changed region attribute to optional

## 1.3.2 (September 25, 2024)

ENHANCEMENTS:

* resource/data_source: added support for AZURE POSTGRESQL FLEXIBLE server type

## 1.3.1 (September 25, 2024)

BUG FIXES:

* resource/data_source: updated data type of availability_zones, db_instances_display_name, and enabled_logs_exports attributes from a string to a list of strings

## 1.3.0 (Septmeber 9, 2024)

ENHANCEMENTS:

* added support of the sync_type provider parameter

## 1.2.46 (July 5, 2024)

ENHANCEMENTS:

* resource/data_source: added support of the bucket_account_id attribute

## 1.2.45 (June 20, 2024)

ENHANCEMENTS:

* resource/data_source: added support of the arn attribute
* resource/log_aggregator: added support of the following attributes:
  * arn
  * bucket_account_id
* resource/secret_manager: added support of the arn attribute

BUG FIXES:

* resource/cloud_account,data_source,log_aggregator: updated data type of available_regions attribute from a string to a list of strings

## 1.2.44 (June 7, 2024)

BUG FIXES:

* resource/log_aggregator: added support of audit_types

## 1.2.43 (May 16, 2024)

BUG FIXES:

* resource/data_source: removed ad_provider attribute
* resource/log_aggregator: added s3_provider attribute

## 1.2.42 (February 28, 2024)

NOTES:

* fixed typos in READMEs

## 1.2.41 (February 26, 2024)

NOTES:

* added example module for AWS RDS Aurora PostgreSQL Cluster onboarding

## 1.2.40 (February 5, 2024)

BUG FIXES:

* resource/cloud_account: support server_ip attribute

## 1.2.39 (December 5, 2023)

NOTES:

* updated AWS RDS MS SQL Server example module to include IAM role permissions required for writing to S3 buckets

## 1.2.38 (December 5, 2023)

ENHANCEMENTS:

* added support of data sources for the following resources:
  * cloud_account 
  * data_source 
  * log_aggregator
  * secret_manager resources

## 1.2.37 (November 29, 2023)

NOTES:

* updated secret_manager documentation to show integration with data_source resource

## 1.2.36 (November 29, 2023)

NOTES:

* added Azure authorization example guide

## 1.2.35 (November 29, 2023)

NOTES:

* updated formatting of guide documentation page

## 1.2.34 (November 29, 2023)

NOTES:

* added guide for bulk management of assets from csv file

## 1.2.33 (November 28, 2023)

NOTES:

* reorganized agentless permission guides
* updated AWS RDS PostgreSQL log aggregator guide to include SQL scripts

## 1.2.32 (November 27, 2023)

NOTES:

* added example module for GCP MySQL, and GCP PubSub resources

## 1.2.31 (November 24, 2023)

BUG FIXES:

* resource/log_aggregator: support audit_type, and pubsub_subscription attributes

## 1.2.30 (November 24, 2023)

BUG FIXES:

* resource/secret_manager: fixed bug with "path" field not being mapped correctly

## 1.2.29 (November 24, 2023)

NOTES:

* added example module for Azure MySQL data_source resource

ENHANCEMENTS:

* updated default value for audit_pull_enabled from true to false

## 1.2.28 (November 22, 2023)

BUG FIXES:

* resource/log_aggregator: support asset_connection fields for Azure Eventhub assets

## 1.2.27 (November 21, 2023)

BUG FIXES:

* escape special characters in API paths to support values such as ":" within asset_ids

## 1.2.26 (November 17, 2023)

NOTES:

* reorganized documentation side navigation structure
* added example module for AWS RDS Oracle via ODBC 

BUG FIXES:

* resource/data_source: support audit_type values for resources other than Azure CosmosDB

## 1.2.25 (November 17, 2023)

NOTES:

* added log_aggregator resource to AWS RDS MS SQL Server example module

BUG FIXES:

* resource/data_source: re-enable audit collection on update 

## 1.2.24 (November 17, 2023)

NOTES:

* added example module for AWS RDS MS SQL Server resource
* added example module for AWS RDS PostgreSQL resource

## 1.2.22 (November 16, 2023)

ENHANCEMENTS:

* resource/data_source,log_aggregator: added support for enabling audit via enable-audit-collection/disable-audit-collection API calls based on audit_pull_enabled flag in resources

## 1.2.21 (November 15, 2023)

NOTES:

* refactored documentation

## 1.2.20 (November 14, 2023)

NOTES:

* refactored documentation

## 1.2.19 (November 14, 2023)

NOTES:

* refactored documentation and navigation

## 1.2.18 (November 14, 2023)

BUG FIXES:

* handle "Server Port" field returning as int

## 1.2.17 (November 14, 2023)

BUG FIXES:

* handle "Server Port" field returning as float

## 1.2.16 (November 14, 2023)

NOTES:

* added links to local guide pages for IAM references per module

## 1.2.15 (November 14, 2023)

NOTES:

* updated documentation structure to introduce guides

## 1.2.14 (November 14, 2023)

NOTES:

* added permission examples to subcategory list
* added AWS RDS MYSQL template example

## 1.2.13 (November 14, 2023)

NOTES:

* restructured documentation with submenus
* added agentless gateway permissions

## 1.2.12 (November 14, 2023)

BUG FIXES:

* handle "Server Port" field returning as both string and int

## 1.2.11 (November 6, 2023)

NOTES:

* updated documentation examples to split out generic and resource specific variables

## 1.2.10 (November 2, 2023)

BUG FIXES:

* resource/secret_manager: fixed incorrect data types

## 1.2.9 (November 2, 2023)

NOTES:

* added required_providers block to documentation

## 1.2.8 (November 2, 2023)

ENHANCEMENTS:

* updated asset_connection field to be sensitive and hidden in terraform state file

## 1.2.6 (November 1, 2023)

NOTES:

* updated documentation for cloud_account resource to show use of parent_asset_id in the creation of a data_source resource

BUG FIXES:

* updated resource files to return error in the event the asset_id is invalid

## 1.2.5 (October 31, 2023)

NOTES:

* updated individual resource example documentation from .tf to .md
* added parameter definitions

## 1.2.4 (October 31, 2023)

NOTES:

* updated individual resource example documentation from .tf to .md
* added parameter definitions

## 1.2.3 (October 31, 2023)

NOTES:

* updated individual resource example documentation from .tf to .md
* added parameter definitions

## 1.2.1 (October 31, 2023)
 
NOTES:

* updated data_source syntax links

## 1.2.0 (October 31, 2023)

NOTES:

* updated cloud_account links

## 1.0.14 (October 31, 2023)

NOTES:

* updated data_source examples, and added usage examples with multiple asset_connections

## 1.0.13 (October 31, 2023)

NOTES:

* added examples for data sources and updated examples in index

## 1.0.12 (October 30, 2023)

NOTES:

* updated documentation links back to repo to reduce size of readme files

## 1.0.11 (October 30, 2023)

NOTES:

* migrated all documentation to expandable inline page examples

## 1.0.10 (October 30, 2023)

NOTES:

* updated documentation to fix resource syntax, and reorganized file locations

## 1.0.9 (October 30, 2023)


## 1.0.8 (October 30, 2023)

NOTES:

* updated release.yml, make file
* updated go.mod
* added vendor files

## 1.0.1 (October 30, 2023)

FEATURES:

* Initial release of dsfhub terraform provider, supporting the following resources:
  * cloud_account
  * data_source
  * log_aggregator
  * secret_manager
