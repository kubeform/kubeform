---
title: StorageBucket
menu:
  docs_{{ .version }}:
    identifier: storagebucket-google.kubeform.com
    name: StorageBucket
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StorageBucket
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `StorageBucket` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StorageBucketSpec](#storagebucketspec)***||
| `status` | ***[StorageBucketStatus](#storagebucketstatus)***||
## StorageBucketSpec

Appears on:[StorageBucket](#storagebucket), [StorageBucketStatus](#storagebucketstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `cors` | ***[[]StorageBucketSpecCors](#storagebucketspeccors)***| ***(Optional)*** |
| `encryption` | ***[[]StorageBucketSpecEncryption](#storagebucketspecencryption)***| ***(Optional)*** |
| `forceDestroy` | ***bool***| ***(Optional)*** |
| `labels` | ***map[string]string***| ***(Optional)*** |
| `lifecycleRule` | ***[[]StorageBucketSpecLifecycleRule](#storagebucketspeclifecyclerule)***| ***(Optional)*** |
| `location` | ***string***| ***(Optional)*** |
| `logging` | ***[[]StorageBucketSpecLogging](#storagebucketspeclogging)***| ***(Optional)*** |
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `storageClass` | ***string***| ***(Optional)*** |
| `url` | ***string***| ***(Optional)*** |
| `versioning` | ***[[]StorageBucketSpecVersioning](#storagebucketspecversioning)***| ***(Optional)*** |
| `website` | ***[[]StorageBucketSpecWebsite](#storagebucketspecwebsite)***| ***(Optional)*** |
## StorageBucketSpecCors

Appears on:[StorageBucketSpec](#storagebucketspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `maxAgeSeconds` | ***int***| ***(Optional)*** |
| `method` | ***[]string***| ***(Optional)*** |
| `origin` | ***[]string***| ***(Optional)*** |
| `responseHeader` | ***[]string***| ***(Optional)*** |
## StorageBucketSpecEncryption

Appears on:[StorageBucketSpec](#storagebucketspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `defaultKmsKeyName` | ***string***||
## StorageBucketSpecLifecycleRule

Appears on:[StorageBucketSpec](#storagebucketspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***[[]StorageBucketSpecLifecycleRuleAction](#storagebucketspeclifecycleruleaction)***||
| `condition` | ***[[]StorageBucketSpecLifecycleRuleCondition](#storagebucketspeclifecyclerulecondition)***||
## StorageBucketSpecLifecycleRuleAction

Appears on:[StorageBucketSpecLifecycleRule](#storagebucketspeclifecyclerule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `storageClass` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## StorageBucketSpecLifecycleRuleCondition

Appears on:[StorageBucketSpecLifecycleRule](#storagebucketspeclifecyclerule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `age` | ***int***| ***(Optional)*** |
| `createdBefore` | ***string***| ***(Optional)*** |
| `isLive` | ***bool***| ***(Optional)*** |
| `matchesStorageClass` | ***[]string***| ***(Optional)*** |
| `numNewerVersions` | ***int***| ***(Optional)*** |
## StorageBucketSpecLogging

Appears on:[StorageBucketSpec](#storagebucketspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `logBucket` | ***string***||
| `logObjectPrefix` | ***string***| ***(Optional)*** |
## StorageBucketSpecVersioning

Appears on:[StorageBucketSpec](#storagebucketspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
## StorageBucketSpecWebsite

Appears on:[StorageBucketSpec](#storagebucketspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `mainPageSuffix` | ***string***| ***(Optional)*** |
| `notFoundPage` | ***string***| ***(Optional)*** |
## StorageBucketStatus

Appears on:[StorageBucket](#storagebucket)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StorageBucketSpec](#storagebucketspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
