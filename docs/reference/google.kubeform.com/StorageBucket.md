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
| `spec` | ***[StorageBucketSpec](#StorageBucketSpec)***||
| `status` | ***[StorageBucketStatus](#StorageBucketStatus)***||
## StorageBucketSpec

Appears on:[StorageBucket](#StorageBucket), [StorageBucketStatus](#StorageBucketStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `cors` | ***[[]StorageBucketSpecCors](#StorageBucketSpecCors)***| ***(Optional)*** |
| `encryption` | ***[[]StorageBucketSpecEncryption](#StorageBucketSpecEncryption)***| ***(Optional)*** |
| `forceDestroy` | ***bool***| ***(Optional)*** |
| `labels` | ***map[string]string***| ***(Optional)*** |
| `lifecycleRule` | ***[[]StorageBucketSpecLifecycleRule](#StorageBucketSpecLifecycleRule)***| ***(Optional)*** |
| `location` | ***string***| ***(Optional)*** |
| `logging` | ***[[]StorageBucketSpecLogging](#StorageBucketSpecLogging)***| ***(Optional)*** |
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `storageClass` | ***string***| ***(Optional)*** |
| `url` | ***string***| ***(Optional)*** |
| `versioning` | ***[[]StorageBucketSpecVersioning](#StorageBucketSpecVersioning)***| ***(Optional)*** |
| `website` | ***[[]StorageBucketSpecWebsite](#StorageBucketSpecWebsite)***| ***(Optional)*** |
## StorageBucketSpecCors

Appears on:[StorageBucketSpec](#StorageBucketSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `maxAgeSeconds` | ***int***| ***(Optional)*** |
| `method` | ***[]string***| ***(Optional)*** |
| `origin` | ***[]string***| ***(Optional)*** |
| `responseHeader` | ***[]string***| ***(Optional)*** |
## StorageBucketSpecEncryption

Appears on:[StorageBucketSpec](#StorageBucketSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `defaultKmsKeyName` | ***string***||
## StorageBucketSpecLifecycleRule

Appears on:[StorageBucketSpec](#StorageBucketSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***[[]StorageBucketSpecLifecycleRuleAction](#StorageBucketSpecLifecycleRuleAction)***||
| `condition` | ***[[]StorageBucketSpecLifecycleRuleCondition](#StorageBucketSpecLifecycleRuleCondition)***||
## StorageBucketSpecLifecycleRuleAction

Appears on:[StorageBucketSpecLifecycleRule](#StorageBucketSpecLifecycleRule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `storageClass` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## StorageBucketSpecLifecycleRuleCondition

Appears on:[StorageBucketSpecLifecycleRule](#StorageBucketSpecLifecycleRule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `age` | ***int***| ***(Optional)*** |
| `createdBefore` | ***string***| ***(Optional)*** |
| `isLive` | ***bool***| ***(Optional)*** |
| `matchesStorageClass` | ***[]string***| ***(Optional)*** |
| `numNewerVersions` | ***int***| ***(Optional)*** |
## StorageBucketSpecLogging

Appears on:[StorageBucketSpec](#StorageBucketSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `logBucket` | ***string***||
| `logObjectPrefix` | ***string***| ***(Optional)*** |
## StorageBucketSpecVersioning

Appears on:[StorageBucketSpec](#StorageBucketSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
## StorageBucketSpecWebsite

Appears on:[StorageBucketSpec](#StorageBucketSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `mainPageSuffix` | ***string***| ***(Optional)*** |
| `notFoundPage` | ***string***| ***(Optional)*** |
## StorageBucketStatus

Appears on:[StorageBucket](#StorageBucket)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StorageBucketSpec](#StorageBucketSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
