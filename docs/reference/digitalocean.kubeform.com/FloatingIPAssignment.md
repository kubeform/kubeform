---
title: FloatingIPAssignment
menu:
  docs_{{ .version }}:
    identifier: floatingipassignment-digitalocean.kubeform.com
    name: FloatingIPAssignment
    parent: digitalocean.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## FloatingIPAssignment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
|    `kind` | string | `FloatingIPAssignment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[FloatingIPAssignmentSpec](#floatingipassignmentspec)***||
| `status` | ***[FloatingIPAssignmentStatus](#floatingipassignmentstatus)***||
## FloatingIPAssignmentSpec

Appears on:[FloatingIPAssignment](#floatingipassignment), [FloatingIPAssignmentStatus](#floatingipassignmentstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `dropletID` | ***int***||
| `ipAddress` | ***string***||
## FloatingIPAssignmentStatus

Appears on:[FloatingIPAssignment](#floatingipassignment)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[FloatingIPAssignmentSpec](#floatingipassignmentspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[FloatingIPAssignmentStatus](#floatingipassignmentstatus)

---
