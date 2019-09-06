## SecretsmanagerSecretVersion
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SecretsmanagerSecretVersion` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SecretsmanagerSecretVersionSpec](#SecretsmanagerSecretVersionSpec)***||
| `status` | ***[SecretsmanagerSecretVersionStatus](#SecretsmanagerSecretVersionStatus)***||
## SecretsmanagerSecretVersionSpec
##### (Appears on:[SecretsmanagerSecretVersion](#SecretsmanagerSecretVersion), [SecretsmanagerSecretVersionStatus](#SecretsmanagerSecretVersionStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `arn` | ***string***| ***(Optional)*** |
| `secretID` | ***string***||
| `versionID` | ***string***| ***(Optional)*** |
| `versionStages` | ***[]string***| ***(Optional)*** |
## SecretsmanagerSecretVersionStatus
##### (Appears on:[SecretsmanagerSecretVersion](#SecretsmanagerSecretVersion))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SecretsmanagerSecretVersionSpec](#SecretsmanagerSecretVersionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `secret_binary` | ***string*** ||
| `secret_string` | ***string*** ||
