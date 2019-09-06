## Token
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `linode.kubeform.com/v1alpha1` |
|    `kind` | string | `Token` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[TokenSpec](#TokenSpec)***||
| `status` | ***[TokenStatus](#TokenStatus)***||
## TokenSpec
##### (Appears on:[Token](#Token), [TokenStatus](#TokenStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `created` | ***string***| ***(Optional)*** The date and time this token was created.|
| `expiry` | ***string***| ***(Optional)*** When this token will expire. Personal Access Tokens cannot be renewed, so after this time the token will be completely unusable and a new token will need to be generated. Tokens may be created with 'null' as their expiry and will never expire unless revoked.|
| `label` | ***string***| ***(Optional)*** The label of the Linode Token.|
| `scopes` | ***string***|The scopes this token was created with. These define what parts of the Account the token can be used to access. Many command-line tools, such as the Linode CLI, require tokens with access to *. Tokens with more restrictive scopes are generally more secure.|
| `token` | ***string***| ***(Optional)*** The token used to access the API.|
## TokenStatus
##### (Appears on:[Token](#Token))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[TokenSpec](#TokenSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
