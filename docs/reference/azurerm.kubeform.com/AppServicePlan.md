## AppServicePlan
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `AppServicePlan` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AppServicePlanSpec](#AppServicePlanSpec)***||
| `status` | ***[AppServicePlanStatus](#AppServicePlanStatus)***||
## AppServicePlanSpec
##### (Appears on:[AppServicePlan](#AppServicePlan), [AppServicePlanStatus](#AppServicePlanStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `appServiceEnvironmentID` | ***string***| ***(Optional)*** |
| `isXenon` | ***bool***| ***(Optional)*** |
| `kind` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `maximumElasticWorkerCount` | ***int***| ***(Optional)*** |
| `maximumNumberOfWorkers` | ***int***| ***(Optional)*** |
| `name` | ***string***||
| `perSiteScaling` | ***bool***| ***(Optional)*** |
| `properties` | ***[[]AppServicePlanSpecProperties](#AppServicePlanSpecProperties)***| ***(Optional)*** Deprecated|
| `reserved` | ***bool***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `sku` | ***[[]AppServicePlanSpecSku](#AppServicePlanSpecSku)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## AppServicePlanSpecProperties
##### (Appears on:[AppServicePlanSpec](#AppServicePlanSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `appServiceEnvironmentID` | ***string***| ***(Optional)*** Deprecated|
| `perSiteScaling` | ***bool***| ***(Optional)*** Deprecated|
| `reserved` | ***bool***| ***(Optional)*** Deprecated|
## AppServicePlanSpecSku
##### (Appears on:[AppServicePlanSpec](#AppServicePlanSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `capacity` | ***int***| ***(Optional)*** |
| `size` | ***string***||
| `tier` | ***string***||
## AppServicePlanStatus
##### (Appears on:[AppServicePlan](#AppServicePlan))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AppServicePlanSpec](#AppServicePlanSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
