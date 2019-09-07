---
title: DirectoryServiceConditionalForwarder
menu:
  docs_v0.0.1:
    identifier: directoryserviceconditionalforwarder-aws.kubeform.com
    name: DirectoryServiceConditionalForwarder
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## DirectoryServiceConditionalForwarder
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DirectoryServiceConditionalForwarder` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DirectoryServiceConditionalForwarderSpec](#DirectoryServiceConditionalForwarderSpec)***||
| `status` | ***[DirectoryServiceConditionalForwarderStatus](#DirectoryServiceConditionalForwarderStatus)***||
## DirectoryServiceConditionalForwarderSpec
##### (Appears on:[DirectoryServiceConditionalForwarder](#DirectoryServiceConditionalForwarder), [DirectoryServiceConditionalForwarderStatus](#DirectoryServiceConditionalForwarderStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `directoryID` | ***string***||
| `dnsIPS` | ***[]string***||
| `remoteDomainName` | ***string***||
## DirectoryServiceConditionalForwarderStatus
##### (Appears on:[DirectoryServiceConditionalForwarder](#DirectoryServiceConditionalForwarder))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DirectoryServiceConditionalForwarderSpec](#DirectoryServiceConditionalForwarderSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
