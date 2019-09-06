## IamServiceLinkedRole
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IamServiceLinkedRole` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IamServiceLinkedRoleSpec](#IamServiceLinkedRoleSpec)***||
| `status` | ***[IamServiceLinkedRoleStatus](#IamServiceLinkedRoleStatus)***||
## IamServiceLinkedRoleSpec
##### (Appears on:[IamServiceLinkedRole](#IamServiceLinkedRole), [IamServiceLinkedRoleStatus](#IamServiceLinkedRoleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `awsServiceName` | ***string***||
| `createDate` | ***string***| ***(Optional)*** |
| `customSuffix` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `path` | ***string***| ***(Optional)*** |
| `uniqueID` | ***string***| ***(Optional)*** |
## IamServiceLinkedRoleStatus
##### (Appears on:[IamServiceLinkedRole](#IamServiceLinkedRole))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IamServiceLinkedRoleSpec](#IamServiceLinkedRoleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
