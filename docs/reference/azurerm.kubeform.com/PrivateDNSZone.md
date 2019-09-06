## PrivateDNSZone
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `PrivateDNSZone` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PrivateDNSZoneSpec](#PrivateDNSZoneSpec)***||
| `status` | ***[PrivateDNSZoneStatus](#PrivateDNSZoneStatus)***||
## PrivateDNSZoneSpec
##### (Appears on:[PrivateDNSZone](#PrivateDNSZone), [PrivateDNSZoneStatus](#PrivateDNSZoneStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `maxNumberOfRecordSets` | ***int***| ***(Optional)*** |
| `maxNumberOfVirtualNetworkLinks` | ***int***| ***(Optional)*** |
| `maxNumberOfVirtualNetworkLinksWithRegistration` | ***int***| ***(Optional)*** |
| `name` | ***string***||
| `numberOfRecordSets` | ***int***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## PrivateDNSZoneStatus
##### (Appears on:[PrivateDNSZone](#PrivateDNSZone))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PrivateDNSZoneSpec](#PrivateDNSZoneSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
