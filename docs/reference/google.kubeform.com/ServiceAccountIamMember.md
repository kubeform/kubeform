## ServiceAccountIamMember
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ServiceAccountIamMember` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ServiceAccountIamMemberSpec](#ServiceAccountIamMemberSpec)***||
| `status` | ***[ServiceAccountIamMemberStatus](#ServiceAccountIamMemberStatus)***||
## ServiceAccountIamMemberSpec
##### (Appears on:[ServiceAccountIamMember](#ServiceAccountIamMember), [ServiceAccountIamMemberStatus](#ServiceAccountIamMemberStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `member` | ***string***||
| `role` | ***string***||
| `serviceAccountID` | ***string***||
## ServiceAccountIamMemberStatus
##### (Appears on:[ServiceAccountIamMember](#ServiceAccountIamMember))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ServiceAccountIamMemberSpec](#ServiceAccountIamMemberSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
