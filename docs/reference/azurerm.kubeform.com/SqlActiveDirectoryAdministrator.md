## SqlActiveDirectoryAdministrator
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `SqlActiveDirectoryAdministrator` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SqlActiveDirectoryAdministratorSpec](#SqlActiveDirectoryAdministratorSpec)***||
| `status` | ***[SqlActiveDirectoryAdministratorStatus](#SqlActiveDirectoryAdministratorStatus)***||
## SqlActiveDirectoryAdministratorSpec
##### (Appears on:[SqlActiveDirectoryAdministrator](#SqlActiveDirectoryAdministrator), [SqlActiveDirectoryAdministratorStatus](#SqlActiveDirectoryAdministratorStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `login` | ***string***||
| `objectID` | ***string***||
| `resourceGroupName` | ***string***||
| `serverName` | ***string***||
| `tenantID` | ***string***||
## SqlActiveDirectoryAdministratorStatus
##### (Appears on:[SqlActiveDirectoryAdministrator](#SqlActiveDirectoryAdministrator))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SqlActiveDirectoryAdministratorSpec](#SqlActiveDirectoryAdministratorSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
