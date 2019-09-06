## GameliftAlias
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `GameliftAlias` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[GameliftAliasSpec](#GameliftAliasSpec)***||
| `status` | ***[GameliftAliasStatus](#GameliftAliasStatus)***||
## GameliftAliasSpec
##### (Appears on:[GameliftAlias](#GameliftAlias), [GameliftAliasStatus](#GameliftAliasStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `routingStrategy` | ***[[]GameliftAliasSpecRoutingStrategy](#GameliftAliasSpecRoutingStrategy)***||
## GameliftAliasSpecRoutingStrategy
##### (Appears on:[GameliftAliasSpec](#GameliftAliasSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `fleetID` | ***string***| ***(Optional)*** |
| `message` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## GameliftAliasStatus
##### (Appears on:[GameliftAlias](#GameliftAlias))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[GameliftAliasSpec](#GameliftAliasSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
