## Project
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
|    `kind` | string | `Project` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ProjectSpec](#ProjectSpec)***||
| `status` | ***[ProjectStatus](#ProjectStatus)***||
## ProjectSpec
##### (Appears on:[Project](#Project), [ProjectStatus](#ProjectStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `createdAt` | ***string***| ***(Optional)*** the date and time when the project was created, (ISO8601)|
| `description` | ***string***| ***(Optional)*** the descirption of the project|
| `environment` | ***string***| ***(Optional)*** the environment of the project's resources|
| `name` | ***string***|the human-readable name for the project|
| `ownerID` | ***int***| ***(Optional)*** the id of the project owner.|
| `ownerUUID` | ***string***| ***(Optional)*** the unique universal identifier of the project owner.|
| `purpose` | ***string***| ***(Optional)*** the purpose of the project|
| `resources` | ***[]string***| ***(Optional)*** the resources associated with the project|
| `updatedAt` | ***string***| ***(Optional)*** the date and time when the project was last updated, (ISO8601)|
## ProjectStatus
##### (Appears on:[Project](#Project))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ProjectSpec](#ProjectSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
