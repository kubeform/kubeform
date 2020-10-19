---
title: RDS
menu:
  docs_{{ .version }}:
    identifier: rds-modules.kubeform.com
    name: RDS
    parent: modules.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## RDS
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `modules.kubeform.com/v1alpha1` |
|    `kind` | string | `RDS` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RDSSpec](#rdsspec)***||
| `status` | ***[RDSStatus](#rdsstatus)***||
## Phase(`string` alias)

Appears on:[RDSStatus](#rdsstatus)

## RDSSpec

Appears on:[RDS](#rds)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***| ***(Optional)*** |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `source` | ***string***| ***(Optional)*** |
| `allocatedStorage` | ***string***| ***(Optional)*** The allocated storage in gigabytes|
| `allowMajorVersionUpgrade` | ***bool***| ***(Optional)*** Indicates that major version upgrades are allowed. Changing this parameter does not result in an outage and the change is asynchronously applied as soon as possible|
| `applyImmediately` | ***bool***| ***(Optional)*** Specifies whether any database modifications are applied immediately, or during the next maintenance window|
| `autoMinorVersionUpgrade` | ***bool***| ***(Optional)*** Indicates that minor engine upgrades will be applied automatically to the DB instance during the maintenance window|
| `availabilityZone` | ***string***| ***(Optional)*** The Availability Zone of the RDS instance|
| `backupRetentionPeriod` | ***encoding/json.Number***| ***(Optional)*** The days to retain backups for|
| `backupWindow` | ***string***| ***(Optional)*** The daily time range (in UTC) during which automated backups are created if they are enabled. Example: '09:46-10:16'. Must not overlap with maintenance_window|
| `caCertIdentifier` | ***string***| ***(Optional)*** Specifies the identifier of the CA certificate for the DB instance|
| `characterSetName` | ***string***| ***(Optional)*** (Optional) The character set name to use for DB encoding in Oracle instances. This can't be changed. See Oracle Character Sets Supported in Amazon RDS for more information|
| `copyTagsToSnapshot` | ***bool***| ***(Optional)*** On delete, copy all Instance tags to the final snapshot (if final_snapshot_identifier is specified)|
| `createDbInstance` | ***bool***| ***(Optional)*** Whether to create a database instance|
| `createDbOptionGroup` | ***bool***| ***(Optional)*** (Optional) Create a database option group|
| `createDbParameterGroup` | ***bool***| ***(Optional)*** Whether to create a database parameter group|
| `createDbSubnetGroup` | ***bool***| ***(Optional)*** Whether to create a database subnet group|
| `createMonitoringRole` | ***bool***| ***(Optional)*** Create IAM role with a defined name that permits RDS to send enhanced monitoring metrics to CloudWatch Logs.|
| `dbSubnetGroupName` | ***string***| ***(Optional)*** Name of DB subnet group. DB instance will be created in the VPC associated with the DB subnet group. If unspecified, will be created in the default VPC|
| `deleteAutomatedBackups` | ***bool***| ***(Optional)*** Specifies whether to remove automated backups immediately after the DB instance is deleted|
| `deletionProtection` | ***bool***| ***(Optional)*** The database can't be deleted when this value is set to true.|
| `domain` | ***string***| ***(Optional)*** The ID of the Directory Service Active Directory domain to create the instance in|
| `domainIamRoleName` | ***string***| ***(Optional)*** (Required if domain is provided) The name of the IAM role to be used when making API calls to the Directory Service|
| `enabledCloudwatchLogsExports` | ***[]string***| ***(Optional)*** List of log types to enable for exporting to CloudWatch logs. If omitted, no logs will be exported. Valid values (depending on engine): alert, audit, error, general, listener, slowquery, trace, postgresql (PostgreSQL), upgrade (PostgreSQL).|
| `engine` | ***string***| ***(Optional)*** The database engine to use|
| `engineVersion` | ***string***| ***(Optional)*** The engine version to use|
| `family` | ***string***| ***(Optional)*** The family of the DB parameter group|
| `finalSnapshotIdentifier` | ***string***| ***(Optional)*** The name of your final DB snapshot when this DB instance is deleted.|
| `iamDatabaseAuthenticationEnabled` | ***bool***| ***(Optional)*** Specifies whether or mappings of AWS Identity and Access Management (IAM) accounts to database accounts is enabled|
| `identifier` | ***string***| ***(Optional)*** The name of the RDS instance, if omitted, Terraform will assign a random, unique identifier|
| `instanceClass` | ***string***| ***(Optional)*** The instance type of the RDS instance|
| `iops` | ***encoding/json.Number***| ***(Optional)*** The amount of provisioned IOPS. Setting this implies a storage_type of 'io1'|
| `kmsKeyID` | ***string***| ***(Optional)*** The ARN for the KMS encryption key. If creating an encrypted replica, set this to the destination KMS ARN. If storage_encrypted is set to true and kms_key_id is not specified the default KMS key created in your account will be used|
| `licenseModel` | ***string***| ***(Optional)*** License model information for this DB instance. Optional, but required for some DB engines, i.e. Oracle SE1|
| `maintenanceWindow` | ***string***| ***(Optional)*** The window to perform maintenance in. Syntax: 'ddd:hh24:mi-ddd:hh24:mi'. Eg: 'Mon:00:00-Mon:03:00'|
| `majorEngineVersion` | ***string***| ***(Optional)*** Specifies the major version of the engine that this option group should be associated with|
| `maxAllocatedStorage` | ***encoding/json.Number***| ***(Optional)*** Specifies the value for Storage Autoscaling|
| `monitoringInterval` | ***encoding/json.Number***| ***(Optional)*** The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance. To disable collecting Enhanced Monitoring metrics, specify 0. The default is 0. Valid Values: 0, 1, 5, 10, 15, 30, 60.|
| `monitoringRoleArn` | ***string***| ***(Optional)*** The ARN for the IAM role that permits RDS to send enhanced monitoring metrics to CloudWatch Logs. Must be specified if monitoring_interval is non-zero.|
| `monitoringRoleName` | ***string***| ***(Optional)*** Name of the IAM role which will be created when create_monitoring_role is enabled.|
| `multiAz` | ***bool***| ***(Optional)*** Specifies if the RDS instance is multi-AZ|
| `name` | ***string***| ***(Optional)*** The DB name to create. If omitted, no database is created initially|
| `optionGroupDescription` | ***string***| ***(Optional)*** The description of the option group|
| `optionGroupName` | ***string***| ***(Optional)*** Name of the DB option group to associate|
| `optionGroupTimeouts` | ***map[string]string***| ***(Optional)*** Define maximum timeout for deletion of `aws_db_option_group` resource|
| `options` | ***encoding/json.RawMessage***| ***(Optional)*** A list of Options to apply.|
| `parameterGroupDescription` | ***string***| ***(Optional)*** Description of the DB parameter group to create|
| `parameterGroupName` | ***string***| ***(Optional)*** Name of the DB parameter group to associate or create|
| `parameters` | ***encoding/json.RawMessage***| ***(Optional)*** A list of DB parameters (map) to apply|
| `password` | ***string***| ***(Optional)*** Password for the master DB user. Note that this may show up in logs, and it will be stored in the state file|
| `performanceInsightsEnabled` | ***bool***| ***(Optional)*** Specifies whether Performance Insights are enabled|
| `performanceInsightsRetentionPeriod` | ***encoding/json.Number***| ***(Optional)*** The amount of time in days to retain Performance Insights data. Either 7 (7 days) or 731 (2 years).|
| `port` | ***string***| ***(Optional)*** The port on which the DB accepts connections|
| `publiclyAccessible` | ***bool***| ***(Optional)*** Bool to control if instance is publicly accessible|
| `replicateSourceDb` | ***string***| ***(Optional)*** Specifies that this resource is a Replicate database, and to use this value as the source database. This correlates to the identifier of another Amazon RDS Database to replicate.|
| `skipFinalSnapshot` | ***bool***| ***(Optional)*** Determines whether a final DB snapshot is created before the DB instance is deleted. If true is specified, no DBSnapshot is created. If false is specified, a DB snapshot is created before the DB instance is deleted, using the value from final_snapshot_identifier|
| `snapshotIdentifier` | ***string***| ***(Optional)*** Specifies whether or not to create this database from a snapshot. This correlates to the snapshot ID you'd find in the RDS console, e.g: rds:production-2015-06-26-06-05.|
| `storageEncrypted` | ***bool***| ***(Optional)*** Specifies whether the DB instance is encrypted|
| `storageType` | ***string***| ***(Optional)*** One of 'standard' (magnetic), 'gp2' (general purpose SSD), or 'io1' (provisioned IOPS SSD). The default is 'io1' if iops is specified, 'standard' if not. Note that this behaviour is different from the AWS web console, where the default is 'gp2'.|
| `subnetIDS` | ***[]string***| ***(Optional)*** A list of VPC subnet IDs|
| `tags` | ***map[string]string***| ***(Optional)*** A mapping of tags to assign to all resources|
| `timeouts` | ***map[string]string***| ***(Optional)*** (Optional) Updated Terraform resource management timeouts. Applies to `aws_db_instance` in particular to permit resource management times|
| `timezone` | ***string***| ***(Optional)*** (Optional) Time zone of the DB instance. timezone is currently only supported by Microsoft SQL Server. The timezone can only be set on creation. See MSSQL User Guide for more information.|
| `useParameterGroupNamePrefix` | ***bool***| ***(Optional)*** Whether to use the parameter group name prefix or not|
| `username` | ***string***| ***(Optional)*** Username for the master DB user|
| `vpcSecurityGroupIDS` | ***[]string***| ***(Optional)*** List of VPC security groups to associate|
## RDSStatus

Appears on:[RDS](#rds)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***kubeform.dev/kubeform/apis/modules/v1alpha1.RDSOutput***| ***(Optional)*** |
| `state` | ***string***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
