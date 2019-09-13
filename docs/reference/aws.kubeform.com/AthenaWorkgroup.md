---
title: AthenaWorkgroup
menu:
  docs_{{ .version }}:
    identifier: athenaworkgroup-aws.kubeform.com
    name: AthenaWorkgroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AthenaWorkgroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AthenaWorkgroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AthenaWorkgroupSpec](#AthenaWorkgroupSpec)***||
| `status` | ***[AthenaWorkgroupStatus](#AthenaWorkgroupStatus)***||
## AthenaWorkgroupSpec
##### (Appears on:[AthenaWorkgroup](#AthenaWorkgroup), [AthenaWorkgroupStatus](#AthenaWorkgroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `configuration` | ***[[]AthenaWorkgroupSpecConfiguration](#AthenaWorkgroupSpecConfiguration)***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `state` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## AthenaWorkgroupSpecConfiguration
##### (Appears on:[AthenaWorkgroupSpec](#AthenaWorkgroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `bytesScannedCutoffPerQuery` | ***int***| ***(Optional)*** |
| `enforceWorkgroupConfiguration` | ***bool***| ***(Optional)*** |
| `publishCloudwatchMetricsEnabled` | ***bool***| ***(Optional)*** |
| `resultConfiguration` | ***[[]AthenaWorkgroupSpecConfigurationResultConfiguration](#AthenaWorkgroupSpecConfigurationResultConfiguration)***| ***(Optional)*** |
## AthenaWorkgroupSpecConfigurationResultConfiguration
##### (Appears on:[AthenaWorkgroupSpecConfiguration](#AthenaWorkgroupSpecConfiguration))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `encryptionConfiguration` | ***[[]AthenaWorkgroupSpecConfigurationResultConfigurationEncryptionConfiguration](#AthenaWorkgroupSpecConfigurationResultConfigurationEncryptionConfiguration)***| ***(Optional)*** |
| `outputLocation` | ***string***| ***(Optional)*** |
## AthenaWorkgroupSpecConfigurationResultConfigurationEncryptionConfiguration
##### (Appears on:[AthenaWorkgroupSpecConfigurationResultConfiguration](#AthenaWorkgroupSpecConfigurationResultConfiguration))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `encryptionOption` | ***string***| ***(Optional)*** |
| `kmsKeyArn` | ***string***| ***(Optional)*** |
## AthenaWorkgroupStatus
##### (Appears on:[AthenaWorkgroup](#AthenaWorkgroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AthenaWorkgroupSpec](#AthenaWorkgroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
