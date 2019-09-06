## OrganizationsAccount
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `OrganizationsAccount` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[OrganizationsAccountSpec](#OrganizationsAccountSpec)***||
| `status` | ***[OrganizationsAccountStatus](#OrganizationsAccountStatus)***||
## OrganizationsAccountSpec
##### (Appears on:[OrganizationsAccount](#OrganizationsAccount), [OrganizationsAccountStatus](#OrganizationsAccountStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `email` | ***string***||
| `iamUserAccessToBilling` | ***string***| ***(Optional)*** |
| `joinedMethod` | ***string***| ***(Optional)*** |
| `joinedTimestamp` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `roleName` | ***string***| ***(Optional)*** |
| `status` | ***string***| ***(Optional)*** |
## OrganizationsAccountStatus
##### (Appears on:[OrganizationsAccount](#OrganizationsAccount))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[OrganizationsAccountSpec](#OrganizationsAccountSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
