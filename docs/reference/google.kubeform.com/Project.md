---
title: Project
menu:
  docs_v0.0.1:
    identifier: project-google.kubeform.com
    name: Project
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## Project
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `Project` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ProjectSpec](#ProjectSpec)***||
| `status` | ***[ProjectStatus](#ProjectStatus)***||
## ProjectSpec
##### (Appears on:[Project](#Project), [ProjectStatus](#ProjectStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `appEngine` | ***[[]ProjectSpecAppEngine](#ProjectSpecAppEngine)***| ***(Optional)*** Deprecated|
| `autoCreateNetwork` | ***bool***| ***(Optional)*** |
| `billingAccount` | ***string***| ***(Optional)*** |
| `folderID` | ***string***| ***(Optional)*** |
| `labels` | ***map[string]string***| ***(Optional)*** |
| `name` | ***string***||
| `number` | ***string***| ***(Optional)*** |
| `orgID` | ***string***| ***(Optional)*** |
| `projectID` | ***string***||
| `skipDelete` | ***bool***| ***(Optional)*** |
## ProjectSpecAppEngine
##### (Appears on:[ProjectSpec](#ProjectSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `authDomain` | ***string***| ***(Optional)*** |
| `codeBucket` | ***string***| ***(Optional)*** |
| `defaultBucket` | ***string***| ***(Optional)*** |
| `defaultHostname` | ***string***| ***(Optional)*** |
| `featureSettings` | ***[[]ProjectSpecAppEngineFeatureSettings](#ProjectSpecAppEngineFeatureSettings)***| ***(Optional)*** |
| `gcrDomain` | ***string***| ***(Optional)*** |
| `locationID` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `servingStatus` | ***string***| ***(Optional)*** |
| `urlDispatchRule` | ***[[]ProjectSpecAppEngineUrlDispatchRule](#ProjectSpecAppEngineUrlDispatchRule)***| ***(Optional)*** |
## ProjectSpecAppEngineFeatureSettings
##### (Appears on:[ProjectSpecAppEngine](#ProjectSpecAppEngine))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `splitHealthChecks` | ***bool***| ***(Optional)*** |
## ProjectSpecAppEngineUrlDispatchRule
##### (Appears on:[ProjectSpecAppEngine](#ProjectSpecAppEngine))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `domain` | ***string***| ***(Optional)*** |
| `path` | ***string***| ***(Optional)*** |
| `service` | ***string***| ***(Optional)*** |
## ProjectStatus
##### (Appears on:[Project](#Project))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ProjectSpec](#ProjectSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
