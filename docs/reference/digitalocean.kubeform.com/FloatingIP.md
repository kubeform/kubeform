---
title: FloatingIP
menu:
  docs_v0.0.1:
    identifier: floatingip-digitalocean.kubeform.com
    name: FloatingIP
    parent: digitalocean.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## FloatingIP
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
|    `kind` | string | `FloatingIP` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[FloatingIPSpec](#FloatingIPSpec)***||
| `status` | ***[FloatingIPStatus](#FloatingIPStatus)***||
## FloatingIPSpec
##### (Appears on:[FloatingIP](#FloatingIP), [FloatingIPStatus](#FloatingIPStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `dropletID` | ***int***| ***(Optional)*** |
| `ipAddress` | ***string***| ***(Optional)*** |
| `region` | ***string***||
| `urn` | ***string***| ***(Optional)*** the uniform resource name for the floating ip|
## FloatingIPStatus
##### (Appears on:[FloatingIP](#FloatingIP))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[FloatingIPSpec](#FloatingIPSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
