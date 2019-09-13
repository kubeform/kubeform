---
title: GuarddutyThreatintelset
menu:
  docs_{{ .version }}:
    identifier: guarddutythreatintelset-aws.kubeform.com
    name: GuarddutyThreatintelset
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## GuarddutyThreatintelset
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `GuarddutyThreatintelset` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[GuarddutyThreatintelsetSpec](#GuarddutyThreatintelsetSpec)***||
| `status` | ***[GuarddutyThreatintelsetStatus](#GuarddutyThreatintelsetStatus)***||
## GuarddutyThreatintelsetSpec
##### (Appears on:[GuarddutyThreatintelset](#GuarddutyThreatintelset), [GuarddutyThreatintelsetStatus](#GuarddutyThreatintelsetStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `activate` | ***bool***||
| `detectorID` | ***string***||
| `format` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
## GuarddutyThreatintelsetStatus
##### (Appears on:[GuarddutyThreatintelset](#GuarddutyThreatintelset))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[GuarddutyThreatintelsetSpec](#GuarddutyThreatintelsetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
