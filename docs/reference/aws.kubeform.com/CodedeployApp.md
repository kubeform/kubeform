---
title: CodedeployApp
menu:
  docs_{{ .version }}:
    identifier: codedeployapp-aws.kubeform.com
    name: CodedeployApp
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CodedeployApp
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CodedeployApp` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CodedeployAppSpec](#CodedeployAppSpec)***||
| `status` | ***[CodedeployAppStatus](#CodedeployAppStatus)***||
## CodedeployAppSpec

Appears on:[CodedeployApp](#CodedeployApp), [CodedeployAppStatus](#CodedeployAppStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `computePlatform` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `uniqueID` | ***string***| ***(Optional)*** |
## CodedeployAppStatus

Appears on:[CodedeployApp](#CodedeployApp)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CodedeployAppSpec](#CodedeployAppSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
