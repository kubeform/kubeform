---
title: Cloudtrail
menu:
  docs_v0.0.1:
    identifier: cloudtrail-aws.kubeform.com
    name: Cloudtrail
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## Cloudtrail
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Cloudtrail` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CloudtrailSpec](#CloudtrailSpec)***||
| `status` | ***[CloudtrailStatus](#CloudtrailStatus)***||
## CloudtrailSpec
##### (Appears on:[Cloudtrail](#Cloudtrail), [CloudtrailStatus](#CloudtrailStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `cloudWatchLogsGroupArn` | ***string***| ***(Optional)*** |
| `cloudWatchLogsRoleArn` | ***string***| ***(Optional)*** |
| `enableLogFileValidation` | ***bool***| ***(Optional)*** |
| `enableLogging` | ***bool***| ***(Optional)*** |
| `eventSelector` | ***[[]CloudtrailSpecEventSelector](#CloudtrailSpecEventSelector)***| ***(Optional)*** |
| `homeRegion` | ***string***| ***(Optional)*** |
| `includeGlobalServiceEvents` | ***bool***| ***(Optional)*** |
| `isMultiRegionTrail` | ***bool***| ***(Optional)*** |
| `isOrganizationTrail` | ***bool***| ***(Optional)*** |
| `kmsKeyID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `s3BucketName` | ***string***||
| `s3KeyPrefix` | ***string***| ***(Optional)*** |
| `snsTopicName` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## CloudtrailSpecEventSelector
##### (Appears on:[CloudtrailSpec](#CloudtrailSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `dataResource` | ***[[]CloudtrailSpecEventSelectorDataResource](#CloudtrailSpecEventSelectorDataResource)***| ***(Optional)*** |
| `includeManagementEvents` | ***bool***| ***(Optional)*** |
| `readWriteType` | ***string***| ***(Optional)*** |
## CloudtrailSpecEventSelectorDataResource
##### (Appears on:[CloudtrailSpecEventSelector](#CloudtrailSpecEventSelector))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
| `values` | ***[]string***||
## CloudtrailStatus
##### (Appears on:[Cloudtrail](#Cloudtrail))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CloudtrailSpec](#CloudtrailSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
