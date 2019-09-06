## SqlElasticpool
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `SqlElasticpool` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SqlElasticpoolSpec](#SqlElasticpoolSpec)***||
| `status` | ***[SqlElasticpoolStatus](#SqlElasticpoolStatus)***||
## SqlElasticpoolSpec
##### (Appears on:[SqlElasticpool](#SqlElasticpool), [SqlElasticpoolStatus](#SqlElasticpoolStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `creationDate` | ***string***| ***(Optional)*** |
| `dbDtuMax` | ***int***| ***(Optional)*** |
| `dbDtuMin` | ***int***| ***(Optional)*** |
| `dtu` | ***int***||
| `edition` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `poolSize` | ***int***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `serverName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## SqlElasticpoolStatus
##### (Appears on:[SqlElasticpool](#SqlElasticpool))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SqlElasticpoolSpec](#SqlElasticpoolSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
