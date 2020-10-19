---
title: Cloudtrail
menu:
  docs_{{ .version }}:
    identifier: cloudtrail-aws.kubeform.com
    name: Cloudtrail
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Cloudtrail
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Cloudtrail` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CloudtrailSpec](#cloudtrailspec)***||
| `status` | ***[CloudtrailStatus](#cloudtrailstatus)***||
## CloudtrailSpec

Appears on:[Cloudtrail](#cloudtrail), [CloudtrailStatus](#cloudtrailstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `cloudWatchLogsGroupArn` | ***string***| ***(Optional)*** |
| `cloudWatchLogsRoleArn` | ***string***| ***(Optional)*** |
| `enableLogFileValidation` | ***bool***| ***(Optional)*** |
| `enableLogging` | ***bool***| ***(Optional)*** |
| `eventSelector` | ***[[]CloudtrailSpecEventSelector](#cloudtrailspeceventselector)***| ***(Optional)*** |
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

Appears on:[CloudtrailSpec](#cloudtrailspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `dataResource` | ***[[]CloudtrailSpecEventSelectorDataResource](#cloudtrailspeceventselectordataresource)***| ***(Optional)*** |
| `includeManagementEvents` | ***bool***| ***(Optional)*** |
| `readWriteType` | ***string***| ***(Optional)*** |
## CloudtrailSpecEventSelectorDataResource

Appears on:[CloudtrailSpecEventSelector](#cloudtrailspeceventselector)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
| `values` | ***[]string***||
## CloudtrailStatus

Appears on:[Cloudtrail](#cloudtrail)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CloudtrailSpec](#cloudtrailspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[CloudtrailStatus](#cloudtrailstatus)

---
