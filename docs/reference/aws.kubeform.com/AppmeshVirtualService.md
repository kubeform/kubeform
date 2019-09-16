---
title: AppmeshVirtualService
menu:
  docs_{{ .version }}:
    identifier: appmeshvirtualservice-aws.kubeform.com
    name: AppmeshVirtualService
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AppmeshVirtualService
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AppmeshVirtualService` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AppmeshVirtualServiceSpec](#appmeshvirtualservicespec)***||
| `status` | ***[AppmeshVirtualServiceStatus](#appmeshvirtualservicestatus)***||
## AppmeshVirtualServiceSpec

Appears on:[AppmeshVirtualService](#appmeshvirtualservice), [AppmeshVirtualServiceStatus](#appmeshvirtualservicestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `createdDate` | ***string***| ***(Optional)*** |
| `lastUpdatedDate` | ***string***| ***(Optional)*** |
| `meshName` | ***string***||
| `name` | ***string***||
| `spec` | ***[[]AppmeshVirtualServiceSpecSpec](#appmeshvirtualservicespecspec)***||
## AppmeshVirtualServiceSpecSpec

Appears on:[AppmeshVirtualServiceSpec](#appmeshvirtualservicespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `provider` | ***[[]AppmeshVirtualServiceSpecSpecProvider](#appmeshvirtualservicespecspecprovider)***| ***(Optional)*** |
## AppmeshVirtualServiceSpecSpecProvider

Appears on:[AppmeshVirtualServiceSpecSpec](#appmeshvirtualservicespecspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `virtualNode` | ***[[]AppmeshVirtualServiceSpecSpecProviderVirtualNode](#appmeshvirtualservicespecspecprovidervirtualnode)***| ***(Optional)*** |
| `virtualRouter` | ***[[]AppmeshVirtualServiceSpecSpecProviderVirtualRouter](#appmeshvirtualservicespecspecprovidervirtualrouter)***| ***(Optional)*** |
## AppmeshVirtualServiceSpecSpecProviderVirtualNode

Appears on:[AppmeshVirtualServiceSpecSpecProvider](#appmeshvirtualservicespecspecprovider)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `virtualNodeName` | ***string***||
## AppmeshVirtualServiceSpecSpecProviderVirtualRouter

Appears on:[AppmeshVirtualServiceSpecSpecProvider](#appmeshvirtualservicespecspecprovider)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `virtualRouterName` | ***string***||
## AppmeshVirtualServiceStatus

Appears on:[AppmeshVirtualService](#appmeshvirtualservice)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AppmeshVirtualServiceSpec](#appmeshvirtualservicespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
