---
title: GlueSecurityConfiguration
menu:
  docs_{{ .version }}:
    identifier: gluesecurityconfiguration-aws.kubeform.com
    name: GlueSecurityConfiguration
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## GlueSecurityConfiguration
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `GlueSecurityConfiguration` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[GlueSecurityConfigurationSpec](#gluesecurityconfigurationspec)***||
| `status` | ***[GlueSecurityConfigurationStatus](#gluesecurityconfigurationstatus)***||
## GlueSecurityConfigurationSpec

Appears on:[GlueSecurityConfiguration](#gluesecurityconfiguration), [GlueSecurityConfigurationStatus](#gluesecurityconfigurationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `encryptionConfiguration` | ***[[]GlueSecurityConfigurationSpecEncryptionConfiguration](#gluesecurityconfigurationspecencryptionconfiguration)***||
| `name` | ***string***||
## GlueSecurityConfigurationSpecEncryptionConfiguration

Appears on:[GlueSecurityConfigurationSpec](#gluesecurityconfigurationspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cloudwatchEncryption` | ***[[]GlueSecurityConfigurationSpecEncryptionConfigurationCloudwatchEncryption](#gluesecurityconfigurationspecencryptionconfigurationcloudwatchencryption)***||
| `jobBookmarksEncryption` | ***[[]GlueSecurityConfigurationSpecEncryptionConfigurationJobBookmarksEncryption](#gluesecurityconfigurationspecencryptionconfigurationjobbookmarksencryption)***||
| `s3Encryption` | ***[[]GlueSecurityConfigurationSpecEncryptionConfigurationS3Encryption](#gluesecurityconfigurationspecencryptionconfigurations3encryption)***||
## GlueSecurityConfigurationSpecEncryptionConfigurationCloudwatchEncryption

Appears on:[GlueSecurityConfigurationSpecEncryptionConfiguration](#gluesecurityconfigurationspecencryptionconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cloudwatchEncryptionMode` | ***string***| ***(Optional)*** |
| `kmsKeyArn` | ***string***| ***(Optional)*** |
## GlueSecurityConfigurationSpecEncryptionConfigurationJobBookmarksEncryption

Appears on:[GlueSecurityConfigurationSpecEncryptionConfiguration](#gluesecurityconfigurationspecencryptionconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `jobBookmarksEncryptionMode` | ***string***| ***(Optional)*** |
| `kmsKeyArn` | ***string***| ***(Optional)*** |
## GlueSecurityConfigurationSpecEncryptionConfigurationS3Encryption

Appears on:[GlueSecurityConfigurationSpecEncryptionConfiguration](#gluesecurityconfigurationspecencryptionconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `kmsKeyArn` | ***string***| ***(Optional)*** |
| `s3EncryptionMode` | ***string***| ***(Optional)*** |
## GlueSecurityConfigurationStatus

Appears on:[GlueSecurityConfiguration](#gluesecurityconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[GlueSecurityConfigurationSpec](#gluesecurityconfigurationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
