---
title: LightsailStaticIP
menu:
  docs_{{ .version }}:
    identifier: lightsailstaticip-aws.kubeform.com
    name: LightsailStaticIP
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LightsailStaticIP
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LightsailStaticIP` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LightsailStaticIPSpec](#lightsailstaticipspec)***||
| `status` | ***[LightsailStaticIPStatus](#lightsailstaticipstatus)***||
## LightsailStaticIPSpec

Appears on:[LightsailStaticIP](#lightsailstaticip), [LightsailStaticIPStatus](#lightsailstaticipstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `ipAddress` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `supportCode` | ***string***| ***(Optional)*** |
## LightsailStaticIPStatus

Appears on:[LightsailStaticIP](#lightsailstaticip)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LightsailStaticIPSpec](#lightsailstaticipspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[LightsailStaticIPStatus](#lightsailstaticipstatus)

---
