---
title: SnsSmsPreferences
menu:
  docs_{{ .version }}:
    identifier: snssmspreferences-aws.kubeform.com
    name: SnsSmsPreferences
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SnsSmsPreferences
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SnsSmsPreferences` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SnsSmsPreferencesSpec](#SnsSmsPreferencesSpec)***||
| `status` | ***[SnsSmsPreferencesStatus](#SnsSmsPreferencesStatus)***||
## SnsSmsPreferencesSpec
##### (Appears on:[SnsSmsPreferences](#SnsSmsPreferences), [SnsSmsPreferencesStatus](#SnsSmsPreferencesStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `defaultSenderID` | ***string***| ***(Optional)*** |
| `defaultSmsType` | ***string***| ***(Optional)*** |
| `deliveryStatusIamRoleArn` | ***string***| ***(Optional)*** |
| `deliveryStatusSuccessSamplingRate` | ***string***| ***(Optional)*** |
| `monthlySpendLimit` | ***string***| ***(Optional)*** |
| `usageReportS3Bucket` | ***string***| ***(Optional)*** |
## SnsSmsPreferencesStatus
##### (Appears on:[SnsSmsPreferences](#SnsSmsPreferences))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SnsSmsPreferencesSpec](#SnsSmsPreferencesSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
