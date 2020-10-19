---
title: GameliftBuild
menu:
  docs_{{ .version }}:
    identifier: gameliftbuild-aws.kubeform.com
    name: GameliftBuild
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## GameliftBuild
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `GameliftBuild` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[GameliftBuildSpec](#gameliftbuildspec)***||
| `status` | ***[GameliftBuildStatus](#gameliftbuildstatus)***||
## GameliftBuildSpec

Appears on:[GameliftBuild](#gameliftbuild), [GameliftBuildStatus](#gameliftbuildstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `operatingSystem` | ***string***||
| `storageLocation` | ***[[]GameliftBuildSpecStorageLocation](#gameliftbuildspecstoragelocation)***||
| `version` | ***string***| ***(Optional)*** |
## GameliftBuildSpecStorageLocation

Appears on:[GameliftBuildSpec](#gameliftbuildspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucket` | ***string***||
| `key` | ***string***||
| `roleArn` | ***string***||
## GameliftBuildStatus

Appears on:[GameliftBuild](#gameliftbuild)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[GameliftBuildSpec](#gameliftbuildspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[GameliftBuildStatus](#gameliftbuildstatus)

---
