---
title: Project
menu:
  docs_{{ .version }}:
    identifier: project-google.kubeform.com
    name: Project
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Project
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `Project` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ProjectSpec](#projectspec)***||
| `status` | ***[ProjectStatus](#projectstatus)***||
## Phase(`string` alias)

Appears on:[ProjectStatus](#projectstatus)

## ProjectSpec

Appears on:[Project](#project), [ProjectStatus](#projectstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `appEngine` | ***[[]ProjectSpecAppEngine](#projectspecappengine)***| ***(Optional)*** Deprecated|
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

Appears on:[ProjectSpec](#projectspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `authDomain` | ***string***| ***(Optional)*** |
| `codeBucket` | ***string***| ***(Optional)*** |
| `defaultBucket` | ***string***| ***(Optional)*** |
| `defaultHostname` | ***string***| ***(Optional)*** |
| `featureSettings` | ***[[]ProjectSpecAppEngineFeatureSettings](#projectspecappenginefeaturesettings)***| ***(Optional)*** |
| `gcrDomain` | ***string***| ***(Optional)*** |
| `locationID` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `servingStatus` | ***string***| ***(Optional)*** |
| `urlDispatchRule` | ***[[]ProjectSpecAppEngineUrlDispatchRule](#projectspecappengineurldispatchrule)***| ***(Optional)*** |
## ProjectSpecAppEngineFeatureSettings

Appears on:[ProjectSpecAppEngine](#projectspecappengine)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `splitHealthChecks` | ***bool***| ***(Optional)*** |
## ProjectSpecAppEngineUrlDispatchRule

Appears on:[ProjectSpecAppEngine](#projectspecappengine)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `domain` | ***string***| ***(Optional)*** |
| `path` | ***string***| ***(Optional)*** |
| `service` | ***string***| ***(Optional)*** |
## ProjectStatus

Appears on:[Project](#project)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ProjectSpec](#projectspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
