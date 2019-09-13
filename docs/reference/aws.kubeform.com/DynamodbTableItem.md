---
title: DynamodbTableItem
menu:
  docs_{{ .version }}:
    identifier: dynamodbtableitem-aws.kubeform.com
    name: DynamodbTableItem
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DynamodbTableItem
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DynamodbTableItem` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DynamodbTableItemSpec](#DynamodbTableItemSpec)***||
| `status` | ***[DynamodbTableItemStatus](#DynamodbTableItemStatus)***||
## DynamodbTableItemSpec
##### (Appears on:[DynamodbTableItem](#DynamodbTableItem), [DynamodbTableItemStatus](#DynamodbTableItemStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `hashKey` | ***string***||
| `item` | ***string***||
| `rangeKey` | ***string***| ***(Optional)*** |
| `tableName` | ***string***||
## DynamodbTableItemStatus
##### (Appears on:[DynamodbTableItem](#DynamodbTableItem))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DynamodbTableItemSpec](#DynamodbTableItemSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
