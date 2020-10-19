---
title: ComputeInstanceGroupManager
menu:
  docs_{{ .version }}:
    identifier: computeinstancegroupmanager-google.kubeform.com
    name: ComputeInstanceGroupManager
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeInstanceGroupManager
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeInstanceGroupManager` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeInstanceGroupManagerSpec](#computeinstancegroupmanagerspec)***||
| `status` | ***[ComputeInstanceGroupManagerStatus](#computeinstancegroupmanagerstatus)***||
## ComputeInstanceGroupManagerSpec

Appears on:[ComputeInstanceGroupManager](#computeinstancegroupmanager), [ComputeInstanceGroupManagerStatus](#computeinstancegroupmanagerstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `baseInstanceName` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `fingerprint` | ***string***| ***(Optional)*** |
| `instanceGroup` | ***string***| ***(Optional)*** |
| `instanceTemplate` | ***string***||
| `name` | ***string***||
| `namedPort` | ***[[]ComputeInstanceGroupManagerSpecNamedPort](#computeinstancegroupmanagerspecnamedport)***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `targetPools` | ***[]string***| ***(Optional)*** |
| `targetSize` | ***int64***| ***(Optional)*** |
| `updateStrategy` | ***string***| ***(Optional)*** |
| `waitForInstances` | ***bool***| ***(Optional)*** |
| `zone` | ***string***| ***(Optional)*** |
## ComputeInstanceGroupManagerSpecNamedPort

Appears on:[ComputeInstanceGroupManagerSpec](#computeinstancegroupmanagerspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `port` | ***int64***||
## ComputeInstanceGroupManagerStatus

Appears on:[ComputeInstanceGroupManager](#computeinstancegroupmanager)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeInstanceGroupManagerSpec](#computeinstancegroupmanagerspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeInstanceGroupManagerStatus](#computeinstancegroupmanagerstatus)

---
