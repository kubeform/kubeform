---
title: TransferServer
menu:
  docs_{{ .version }}:
    identifier: transferserver-aws.kubeform.com
    name: TransferServer
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## TransferServer
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `TransferServer` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[TransferServerSpec](#TransferServerSpec)***||
| `status` | ***[TransferServerStatus](#TransferServerStatus)***||
## TransferServerSpec

Appears on:[TransferServer](#TransferServer), [TransferServerStatus](#TransferServerStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `endpoint` | ***string***| ***(Optional)*** |
| `endpointDetails` | ***[[]TransferServerSpecEndpointDetails](#TransferServerSpecEndpointDetails)***| ***(Optional)*** |
| `endpointType` | ***string***| ***(Optional)*** |
| `forceDestroy` | ***bool***| ***(Optional)*** |
| `identityProviderType` | ***string***| ***(Optional)*** |
| `invocationRole` | ***string***| ***(Optional)*** |
| `loggingRole` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `url` | ***string***| ***(Optional)*** |
## TransferServerSpecEndpointDetails

Appears on:[TransferServerSpec](#TransferServerSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `vpcEndpointID` | ***string***||
## TransferServerStatus

Appears on:[TransferServer](#TransferServer)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[TransferServerSpec](#TransferServerSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
