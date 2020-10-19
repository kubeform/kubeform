---
title: SesEventDestination
menu:
  docs_{{ .version }}:
    identifier: seseventdestination-aws.kubeform.com
    name: SesEventDestination
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SesEventDestination
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SesEventDestination` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SesEventDestinationSpec](#seseventdestinationspec)***||
| `status` | ***[SesEventDestinationStatus](#seseventdestinationstatus)***||
## Phase(`string` alias)

Appears on:[SesEventDestinationStatus](#seseventdestinationstatus)

## SesEventDestinationSpec

Appears on:[SesEventDestination](#seseventdestination), [SesEventDestinationStatus](#seseventdestinationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `cloudwatchDestination` | ***[[]SesEventDestinationSpecCloudwatchDestination](#seseventdestinationspeccloudwatchdestination)***| ***(Optional)*** |
| `configurationSetName` | ***string***||
| `enabled` | ***bool***| ***(Optional)*** |
| `kinesisDestination` | ***[[]SesEventDestinationSpecKinesisDestination](#seseventdestinationspeckinesisdestination)***| ***(Optional)*** |
| `matchingTypes` | ***[]string***||
| `name` | ***string***||
| `snsDestination` | ***[[]SesEventDestinationSpecSnsDestination](#seseventdestinationspecsnsdestination)***| ***(Optional)*** |
## SesEventDestinationSpecCloudwatchDestination

Appears on:[SesEventDestinationSpec](#seseventdestinationspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `defaultValue` | ***string***||
| `dimensionName` | ***string***||
| `valueSource` | ***string***||
## SesEventDestinationSpecKinesisDestination

Appears on:[SesEventDestinationSpec](#seseventdestinationspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `roleArn` | ***string***||
| `streamArn` | ***string***||
## SesEventDestinationSpecSnsDestination

Appears on:[SesEventDestinationSpec](#seseventdestinationspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `topicArn` | ***string***||
## SesEventDestinationStatus

Appears on:[SesEventDestination](#seseventdestination)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SesEventDestinationSpec](#seseventdestinationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
