---
title: MonitoringGroup
menu:
  docs_{{ .version }}:
    identifier: monitoringgroup-google.kubeform.com
    name: MonitoringGroup
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## MonitoringGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `MonitoringGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MonitoringGroupSpec](#monitoringgroupspec)***||
| `status` | ***[MonitoringGroupStatus](#monitoringgroupstatus)***||
## MonitoringGroupSpec

Appears on:[MonitoringGroup](#monitoringgroup), [MonitoringGroupStatus](#monitoringgroupstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `displayName` | ***string***||
| `filter` | ***string***||
| `isCluster` | ***bool***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `parentName` | ***string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
## MonitoringGroupStatus

Appears on:[MonitoringGroup](#monitoringgroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MonitoringGroupSpec](#monitoringgroupspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[MonitoringGroupStatus](#monitoringgroupstatus)

---
