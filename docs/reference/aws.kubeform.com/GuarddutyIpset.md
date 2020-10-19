---
title: GuarddutyIpset
menu:
  docs_{{ .version }}:
    identifier: guarddutyipset-aws.kubeform.com
    name: GuarddutyIpset
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## GuarddutyIpset
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `GuarddutyIpset` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[GuarddutyIpsetSpec](#guarddutyipsetspec)***||
| `status` | ***[GuarddutyIpsetStatus](#guarddutyipsetstatus)***||
## GuarddutyIpsetSpec

Appears on:[GuarddutyIpset](#guarddutyipset), [GuarddutyIpsetStatus](#guarddutyipsetstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `activate` | ***bool***||
| `detectorID` | ***string***||
| `format` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
## GuarddutyIpsetStatus

Appears on:[GuarddutyIpset](#guarddutyipset)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[GuarddutyIpsetSpec](#guarddutyipsetspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[GuarddutyIpsetStatus](#guarddutyipsetstatus)

---
