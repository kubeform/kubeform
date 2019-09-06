## NotificationHubNamespace_
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `NotificationHubNamespace_` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NotificationHubNamespace_Spec](#NotificationHubNamespace_Spec)***||
| `status` | ***[NotificationHubNamespace_Status](#NotificationHubNamespace_Status)***||
## NotificationHubNamespace_Spec
##### (Appears on:[NotificationHubNamespace_](#NotificationHubNamespace_), [NotificationHubNamespace_Status](#NotificationHubNamespace_Status))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `enabled` | ***bool***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `namespaceType` | ***string***||
| `resourceGroupName` | ***string***||
| `servicebusEndpoint` | ***string***| ***(Optional)*** |
| `sku` | ***[[]NotificationHubNamespace_SpecSku](#NotificationHubNamespace_SpecSku)***| ***(Optional)*** Deprecated|
| `skuName` | ***string***| ***(Optional)*** |
## NotificationHubNamespace_SpecSku
##### (Appears on:[NotificationHubNamespace_Spec](#NotificationHubNamespace_Spec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
## NotificationHubNamespace_Status
##### (Appears on:[NotificationHubNamespace_](#NotificationHubNamespace_))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NotificationHubNamespace_Spec](#NotificationHubNamespace_Spec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
