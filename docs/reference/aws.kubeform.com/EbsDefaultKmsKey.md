---
title: EbsDefaultKmsKey
menu:
  docs_{{ .version }}:
    identifier: ebsdefaultkmskey-aws.kubeform.com
    name: EbsDefaultKmsKey
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## EbsDefaultKmsKey
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `EbsDefaultKmsKey` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[EbsDefaultKmsKeySpec](#ebsdefaultkmskeyspec)***||
| `status` | ***[EbsDefaultKmsKeyStatus](#ebsdefaultkmskeystatus)***||
## EbsDefaultKmsKeySpec


Appears on:[EbsDefaultKmsKey](#ebsdefaultkmskey), [EbsDefaultKmsKeyStatus](#ebsdefaultkmskeystatus)


| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `keyArn` | ***string***||
## EbsDefaultKmsKeyStatus


Appears on:[EbsDefaultKmsKey](#ebsdefaultkmskey)


| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[EbsDefaultKmsKeySpec](#ebsdefaultkmskeyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
