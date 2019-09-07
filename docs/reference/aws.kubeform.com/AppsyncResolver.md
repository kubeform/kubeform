---
title: AppsyncResolver
menu:
  docs_v0.0.1:
    identifier: appsyncresolver-aws.kubeform.com
    name: AppsyncResolver
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## AppsyncResolver
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AppsyncResolver` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AppsyncResolverSpec](#AppsyncResolverSpec)***||
| `status` | ***[AppsyncResolverStatus](#AppsyncResolverStatus)***||
## AppsyncResolverSpec
##### (Appears on:[AppsyncResolver](#AppsyncResolver), [AppsyncResolverStatus](#AppsyncResolverStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `apiID` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `dataSource` | ***string***||
| `field` | ***string***||
| `requestTemplate` | ***string***||
| `responseTemplate` | ***string***||
| `type` | ***string***||
## AppsyncResolverStatus
##### (Appears on:[AppsyncResolver](#AppsyncResolver))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AppsyncResolverSpec](#AppsyncResolverSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
