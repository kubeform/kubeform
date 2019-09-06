## IamRole
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IamRole` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IamRoleSpec](#IamRoleSpec)***||
| `status` | ***[IamRoleStatus](#IamRoleStatus)***||
## IamRoleSpec
##### (Appears on:[IamRole](#IamRole), [IamRoleStatus](#IamRoleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `assumeRolePolicy` | ***string***||
| `createDate` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `forceDetachPolicies` | ***bool***| ***(Optional)*** |
| `maxSessionDuration` | ***int***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `path` | ***string***| ***(Optional)*** |
| `permissionsBoundary` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `uniqueID` | ***string***| ***(Optional)*** |
## IamRoleStatus
##### (Appears on:[IamRole](#IamRole))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IamRoleSpec](#IamRoleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
