---
title: DynamodbTable
menu:
  docs_{{ .version }}:
    identifier: dynamodbtable-aws.kubeform.com
    name: DynamodbTable
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DynamodbTable
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DynamodbTable` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DynamodbTableSpec](#DynamodbTableSpec)***||
| `status` | ***[DynamodbTableStatus](#DynamodbTableStatus)***||
## DynamodbTableSpec
##### (Appears on:[DynamodbTable](#DynamodbTable), [DynamodbTableStatus](#DynamodbTableStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `attribute` | ***[[]DynamodbTableSpecAttribute](#DynamodbTableSpecAttribute)***||
| `billingMode` | ***string***| ***(Optional)*** |
| `globalSecondaryIndex` | ***[[]DynamodbTableSpecGlobalSecondaryIndex](#DynamodbTableSpecGlobalSecondaryIndex)***| ***(Optional)*** |
| `hashKey` | ***string***||
| `localSecondaryIndex` | ***[[]DynamodbTableSpecLocalSecondaryIndex](#DynamodbTableSpecLocalSecondaryIndex)***| ***(Optional)*** |
| `name` | ***string***||
| `pointInTimeRecovery` | ***[[]DynamodbTableSpecPointInTimeRecovery](#DynamodbTableSpecPointInTimeRecovery)***| ***(Optional)*** |
| `rangeKey` | ***string***| ***(Optional)*** |
| `readCapacity` | ***int***| ***(Optional)*** |
| `serverSideEncryption` | ***[[]DynamodbTableSpecServerSideEncryption](#DynamodbTableSpecServerSideEncryption)***| ***(Optional)*** |
| `streamArn` | ***string***| ***(Optional)*** |
| `streamEnabled` | ***bool***| ***(Optional)*** |
| `streamLabel` | ***string***| ***(Optional)*** |
| `streamViewType` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `ttl` | ***[[]DynamodbTableSpecTtl](#DynamodbTableSpecTtl)***| ***(Optional)*** |
| `writeCapacity` | ***int***| ***(Optional)*** |
## DynamodbTableSpecAttribute
##### (Appears on:[DynamodbTableSpec](#DynamodbTableSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `type` | ***string***||
## DynamodbTableSpecGlobalSecondaryIndex
##### (Appears on:[DynamodbTableSpec](#DynamodbTableSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `hashKey` | ***string***||
| `name` | ***string***||
| `nonKeyAttributes` | ***[]string***| ***(Optional)*** |
| `projectionType` | ***string***||
| `rangeKey` | ***string***| ***(Optional)*** |
| `readCapacity` | ***int***| ***(Optional)*** |
| `writeCapacity` | ***int***| ***(Optional)*** |
## DynamodbTableSpecLocalSecondaryIndex
##### (Appears on:[DynamodbTableSpec](#DynamodbTableSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `nonKeyAttributes` | ***[]string***| ***(Optional)*** |
| `projectionType` | ***string***||
| `rangeKey` | ***string***||
## DynamodbTableSpecPointInTimeRecovery
##### (Appears on:[DynamodbTableSpec](#DynamodbTableSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
## DynamodbTableSpecServerSideEncryption
##### (Appears on:[DynamodbTableSpec](#DynamodbTableSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
## DynamodbTableSpecTtl
##### (Appears on:[DynamodbTableSpec](#DynamodbTableSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `attributeName` | ***string***||
| `enabled` | ***bool***| ***(Optional)*** |
## DynamodbTableStatus
##### (Appears on:[DynamodbTable](#DynamodbTable))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DynamodbTableSpec](#DynamodbTableSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
