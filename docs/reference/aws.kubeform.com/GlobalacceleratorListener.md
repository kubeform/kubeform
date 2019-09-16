---
title: GlobalacceleratorListener
menu:
  docs_{{ .version }}:
    identifier: globalacceleratorlistener-aws.kubeform.com
    name: GlobalacceleratorListener
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## GlobalacceleratorListener
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `GlobalacceleratorListener` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[GlobalacceleratorListenerSpec](#globalacceleratorlistenerspec)***||
| `status` | ***[GlobalacceleratorListenerStatus](#globalacceleratorlistenerstatus)***||
## GlobalacceleratorListenerSpec

Appears on:[GlobalacceleratorListener](#globalacceleratorlistener), [GlobalacceleratorListenerStatus](#globalacceleratorlistenerstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `acceleratorArn` | ***string***||
| `clientAffinity` | ***string***| ***(Optional)*** |
| `portRange` | ***[[]GlobalacceleratorListenerSpecPortRange](#globalacceleratorlistenerspecportrange)***||
| `protocol` | ***string***||
## GlobalacceleratorListenerSpecPortRange

Appears on:[GlobalacceleratorListenerSpec](#globalacceleratorlistenerspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `fromPort` | ***int***| ***(Optional)*** |
| `toPort` | ***int***| ***(Optional)*** |
## GlobalacceleratorListenerStatus

Appears on:[GlobalacceleratorListener](#globalacceleratorlistener)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[GlobalacceleratorListenerSpec](#globalacceleratorlistenerspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
