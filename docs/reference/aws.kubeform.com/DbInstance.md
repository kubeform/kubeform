---
title: DbInstance
menu:
  docs_{{ .version }}:
    identifier: dbinstance-aws.kubeform.com
    name: DbInstance
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DbInstance
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DbInstance` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DbInstanceSpec](#dbinstancespec)***||
| `status` | ***[DbInstanceStatus](#dbinstancestatus)***||
## DbInstanceSpec

Appears on:[DbInstance](#dbinstance), [DbInstanceStatus](#dbinstancestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `address` | ***string***| ***(Optional)*** |
| `allocatedStorage` | ***int64***| ***(Optional)*** |
| `allowMajorVersionUpgrade` | ***bool***| ***(Optional)*** |
| `applyImmediately` | ***bool***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `autoMinorVersionUpgrade` | ***bool***| ***(Optional)*** |
| `availabilityZone` | ***string***| ***(Optional)*** |
| `backupRetentionPeriod` | ***int64***| ***(Optional)*** |
| `backupWindow` | ***string***| ***(Optional)*** |
| `caCertIdentifier` | ***string***| ***(Optional)*** |
| `characterSetName` | ***string***| ***(Optional)*** |
| `copyTagsToSnapshot` | ***bool***| ***(Optional)*** |
| `dbSubnetGroupName` | ***string***| ***(Optional)*** |
| `deletionProtection` | ***bool***| ***(Optional)*** |
| `domain` | ***string***| ***(Optional)*** |
| `domainIamRoleName` | ***string***| ***(Optional)*** |
| `enabledCloudwatchLogsExports` | ***[]string***| ***(Optional)*** |
| `endpoint` | ***string***| ***(Optional)*** |
| `engine` | ***string***| ***(Optional)*** |
| `engineVersion` | ***string***| ***(Optional)*** |
| `finalSnapshotIdentifier` | ***string***| ***(Optional)*** |
| `hostedZoneID` | ***string***| ***(Optional)*** |
| `iamDatabaseAuthenticationEnabled` | ***bool***| ***(Optional)*** |
| `identifier` | ***string***| ***(Optional)*** |
| `identifierPrefix` | ***string***| ***(Optional)*** |
| `instanceClass` | ***string***||
| `iops` | ***int64***| ***(Optional)*** |
| `kmsKeyID` | ***string***| ***(Optional)*** |
| `licenseModel` | ***string***| ***(Optional)*** |
| `maintenanceWindow` | ***string***| ***(Optional)*** |
| `maxAllocatedStorage` | ***int64***| ***(Optional)*** |
| `monitoringInterval` | ***int64***| ***(Optional)*** |
| `monitoringRoleArn` | ***string***| ***(Optional)*** |
| `multiAz` | ***bool***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `optionGroupName` | ***string***| ***(Optional)*** |
| `parameterGroupName` | ***string***| ***(Optional)*** |
| `performanceInsightsEnabled` | ***bool***| ***(Optional)*** |
| `performanceInsightsKmsKeyID` | ***string***| ***(Optional)*** |
| `performanceInsightsRetentionPeriod` | ***int64***| ***(Optional)*** |
| `port` | ***int64***| ***(Optional)*** |
| `publiclyAccessible` | ***bool***| ***(Optional)*** |
| `replicas` | ***[]string***| ***(Optional)*** |
| `replicateSourceDb` | ***string***| ***(Optional)*** |
| `resourceID` | ***string***| ***(Optional)*** |
| `s3Import` | ***[[]DbInstanceSpecS3Import](#dbinstancespecs3import)***| ***(Optional)*** |
| `securityGroupNames` | ***[]string***| ***(Optional)*** |
| `skipFinalSnapshot` | ***bool***| ***(Optional)*** |
| `snapshotIdentifier` | ***string***| ***(Optional)*** |
| `status` | ***string***| ***(Optional)*** |
| `storageEncrypted` | ***bool***| ***(Optional)*** |
| `storageType` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `timezone` | ***string***| ***(Optional)*** |
| `username` | ***string***| ***(Optional)*** |
| `vpcSecurityGroupIDS` | ***[]string***| ***(Optional)*** |
## DbInstanceSpecS3Import

Appears on:[DbInstanceSpec](#dbinstancespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucketName` | ***string***||
| `bucketPrefix` | ***string***| ***(Optional)*** |
| `ingestionRole` | ***string***||
| `sourceEngine` | ***string***||
| `sourceEngineVersion` | ***string***||
## DbInstanceStatus

Appears on:[DbInstance](#dbinstance)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DbInstanceSpec](#dbinstancespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DbInstanceStatus](#dbinstancestatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `password` | ***string*** ||
