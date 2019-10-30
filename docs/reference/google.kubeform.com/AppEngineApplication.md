---
title: AppEngineApplication
menu:
  docs_{{ .version }}:
    identifier: appengineapplication-google.kubeform.com
    name: AppEngineApplication
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AppEngineApplication
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `AppEngineApplication` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AppEngineApplicationSpec](#appengineapplicationspec)***||
| `status` | ***[AppEngineApplicationStatus](#appengineapplicationstatus)***||
## AppEngineApplicationSpec

Appears on:[AppEngineApplication](#appengineapplication), [AppEngineApplicationStatus](#appengineapplicationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `authDomain` | ***string***| ***(Optional)*** |
| `codeBucket` | ***string***| ***(Optional)*** |
| `defaultBucket` | ***string***| ***(Optional)*** |
| `defaultHostname` | ***string***| ***(Optional)*** |
| `featureSettings` | ***[[]AppEngineApplicationSpecFeatureSettings](#appengineapplicationspecfeaturesettings)***| ***(Optional)*** |
| `gcrDomain` | ***string***| ***(Optional)*** |
| `locationID` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `servingStatus` | ***string***| ***(Optional)*** |
| `urlDispatchRule` | ***[[]AppEngineApplicationSpecUrlDispatchRule](#appengineapplicationspecurldispatchrule)***| ***(Optional)*** |
## AppEngineApplicationSpecFeatureSettings

Appears on:[AppEngineApplicationSpec](#appengineapplicationspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `splitHealthChecks` | ***bool***| ***(Optional)*** |
## AppEngineApplicationSpecUrlDispatchRule

Appears on:[AppEngineApplicationSpec](#appengineapplicationspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `domain` | ***string***| ***(Optional)*** |
| `path` | ***string***| ***(Optional)*** |
| `service` | ***string***| ***(Optional)*** |
## AppEngineApplicationStatus

Appears on:[AppEngineApplication](#appengineapplication)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AppEngineApplicationSpec](#appengineapplicationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AppEngineApplicationStatus](#appengineapplicationstatus)

---
