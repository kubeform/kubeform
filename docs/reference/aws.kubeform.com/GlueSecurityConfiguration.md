## GlueSecurityConfiguration
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `GlueSecurityConfiguration` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[GlueSecurityConfigurationSpec](#GlueSecurityConfigurationSpec)***||
| `status` | ***[GlueSecurityConfigurationStatus](#GlueSecurityConfigurationStatus)***||
## GlueSecurityConfigurationSpec
##### (Appears on:[GlueSecurityConfiguration](#GlueSecurityConfiguration), [GlueSecurityConfigurationStatus](#GlueSecurityConfigurationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `encryptionConfiguration` | ***[[]GlueSecurityConfigurationSpecEncryptionConfiguration](#GlueSecurityConfigurationSpecEncryptionConfiguration)***||
| `name` | ***string***||
## GlueSecurityConfigurationSpecEncryptionConfiguration
##### (Appears on:[GlueSecurityConfigurationSpec](#GlueSecurityConfigurationSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `cloudwatchEncryption` | ***[[]GlueSecurityConfigurationSpecEncryptionConfigurationCloudwatchEncryption](#GlueSecurityConfigurationSpecEncryptionConfigurationCloudwatchEncryption)***||
| `jobBookmarksEncryption` | ***[[]GlueSecurityConfigurationSpecEncryptionConfigurationJobBookmarksEncryption](#GlueSecurityConfigurationSpecEncryptionConfigurationJobBookmarksEncryption)***||
| `s3Encryption` | ***[[]GlueSecurityConfigurationSpecEncryptionConfigurationS3Encryption](#GlueSecurityConfigurationSpecEncryptionConfigurationS3Encryption)***||
## GlueSecurityConfigurationSpecEncryptionConfigurationCloudwatchEncryption
##### (Appears on:[GlueSecurityConfigurationSpecEncryptionConfiguration](#GlueSecurityConfigurationSpecEncryptionConfiguration))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `cloudwatchEncryptionMode` | ***string***| ***(Optional)*** |
| `kmsKeyArn` | ***string***| ***(Optional)*** |
## GlueSecurityConfigurationSpecEncryptionConfigurationJobBookmarksEncryption
##### (Appears on:[GlueSecurityConfigurationSpecEncryptionConfiguration](#GlueSecurityConfigurationSpecEncryptionConfiguration))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `jobBookmarksEncryptionMode` | ***string***| ***(Optional)*** |
| `kmsKeyArn` | ***string***| ***(Optional)*** |
## GlueSecurityConfigurationSpecEncryptionConfigurationS3Encryption
##### (Appears on:[GlueSecurityConfigurationSpecEncryptionConfiguration](#GlueSecurityConfigurationSpecEncryptionConfiguration))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `kmsKeyArn` | ***string***| ***(Optional)*** |
| `s3EncryptionMode` | ***string***| ***(Optional)*** |
## GlueSecurityConfigurationStatus
##### (Appears on:[GlueSecurityConfiguration](#GlueSecurityConfiguration))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[GlueSecurityConfigurationSpec](#GlueSecurityConfigurationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
