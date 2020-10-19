---
title: ComputeProjectDefaultNetworkTier
menu:
  docs_{{ .version }}:
    identifier: computeprojectdefaultnetworktier-google.kubeform.com
    name: ComputeProjectDefaultNetworkTier
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeProjectDefaultNetworkTier
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeProjectDefaultNetworkTier` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeProjectDefaultNetworkTierSpec](#computeprojectdefaultnetworktierspec)***||
| `status` | ***[ComputeProjectDefaultNetworkTierStatus](#computeprojectdefaultnetworktierstatus)***||
## ComputeProjectDefaultNetworkTierSpec

Appears on:[ComputeProjectDefaultNetworkTier](#computeprojectdefaultnetworktier), [ComputeProjectDefaultNetworkTierStatus](#computeprojectdefaultnetworktierstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `networkTier` | ***string***||
| `project` | ***string***| ***(Optional)*** |
## ComputeProjectDefaultNetworkTierStatus

Appears on:[ComputeProjectDefaultNetworkTier](#computeprojectdefaultnetworktier)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeProjectDefaultNetworkTierSpec](#computeprojectdefaultnetworktierspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeProjectDefaultNetworkTierStatus](#computeprojectdefaultnetworktierstatus)

---
