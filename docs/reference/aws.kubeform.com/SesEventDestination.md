---
title: SesEventDestination
menu:
  docs_v0.0.1:
    identifier: seseventdestination-aws.kubeform.com
    name: SesEventDestination
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## SesEventDestination
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SesEventDestination` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SesEventDestinationSpec](#SesEventDestinationSpec)***||
| `status` | ***[SesEventDestinationStatus](#SesEventDestinationStatus)***||
## SesEventDestinationSpec
##### (Appears on:[SesEventDestination](#SesEventDestination), [SesEventDestinationStatus](#SesEventDestinationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `cloudwatchDestination` | ***[[]SesEventDestinationSpecCloudwatchDestination](#SesEventDestinationSpecCloudwatchDestination)***| ***(Optional)*** |
| `configurationSetName` | ***string***||
| `enabled` | ***bool***| ***(Optional)*** |
| `kinesisDestination` | ***[[]SesEventDestinationSpecKinesisDestination](#SesEventDestinationSpecKinesisDestination)***| ***(Optional)*** |
| `matchingTypes` | ***[]string***||
| `name` | ***string***||
| `snsDestination` | ***[[]SesEventDestinationSpecSnsDestination](#SesEventDestinationSpecSnsDestination)***| ***(Optional)*** |
## SesEventDestinationSpecCloudwatchDestination
##### (Appears on:[SesEventDestinationSpec](#SesEventDestinationSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `defaultValue` | ***string***||
| `dimensionName` | ***string***||
| `valueSource` | ***string***||
## SesEventDestinationSpecKinesisDestination
##### (Appears on:[SesEventDestinationSpec](#SesEventDestinationSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `roleArn` | ***string***||
| `streamArn` | ***string***||
## SesEventDestinationSpecSnsDestination
##### (Appears on:[SesEventDestinationSpec](#SesEventDestinationSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `topicArn` | ***string***||
## SesEventDestinationStatus
##### (Appears on:[SesEventDestination](#SesEventDestination))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SesEventDestinationSpec](#SesEventDestinationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
