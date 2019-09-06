## PostgresqlConfiguration
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `PostgresqlConfiguration` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PostgresqlConfigurationSpec](#PostgresqlConfigurationSpec)***||
| `status` | ***[PostgresqlConfigurationStatus](#PostgresqlConfigurationStatus)***||
## PostgresqlConfigurationSpec
##### (Appears on:[PostgresqlConfiguration](#PostgresqlConfiguration), [PostgresqlConfigurationStatus](#PostgresqlConfigurationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `serverName` | ***string***||
| `value` | ***string***||
## PostgresqlConfigurationStatus
##### (Appears on:[PostgresqlConfiguration](#PostgresqlConfiguration))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PostgresqlConfigurationSpec](#PostgresqlConfigurationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
