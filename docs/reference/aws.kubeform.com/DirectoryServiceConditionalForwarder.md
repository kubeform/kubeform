---
title: DirectoryServiceConditionalForwarder
menu:
  docs_{{ .version }}:
    identifier: directoryserviceconditionalforwarder-aws.kubeform.com
    name: DirectoryServiceConditionalForwarder
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DirectoryServiceConditionalForwarder
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DirectoryServiceConditionalForwarder` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DirectoryServiceConditionalForwarderSpec](#directoryserviceconditionalforwarderspec)***||
| `status` | ***[DirectoryServiceConditionalForwarderStatus](#directoryserviceconditionalforwarderstatus)***||
## DirectoryServiceConditionalForwarderSpec

Appears on:[DirectoryServiceConditionalForwarder](#directoryserviceconditionalforwarder), [DirectoryServiceConditionalForwarderStatus](#directoryserviceconditionalforwarderstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `directoryID` | ***string***||
| `dnsIPS` | ***[]string***||
| `remoteDomainName` | ***string***||
## DirectoryServiceConditionalForwarderStatus

Appears on:[DirectoryServiceConditionalForwarder](#directoryserviceconditionalforwarder)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DirectoryServiceConditionalForwarderSpec](#directoryserviceconditionalforwarderspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DirectoryServiceConditionalForwarderStatus](#directoryserviceconditionalforwarderstatus)

---
