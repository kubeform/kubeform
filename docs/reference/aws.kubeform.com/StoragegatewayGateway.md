---
title: StoragegatewayGateway
menu:
  docs_v0.0.1:
    identifier: storagegatewaygateway-aws.kubeform.com
    name: StoragegatewayGateway
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## StoragegatewayGateway
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `StoragegatewayGateway` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StoragegatewayGatewaySpec](#StoragegatewayGatewaySpec)***||
| `status` | ***[StoragegatewayGatewayStatus](#StoragegatewayGatewayStatus)***||
## StoragegatewayGatewaySpec
##### (Appears on:[StoragegatewayGateway](#StoragegatewayGateway), [StoragegatewayGatewayStatus](#StoragegatewayGatewayStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `activationKey` | ***string***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `gatewayID` | ***string***| ***(Optional)*** |
| `gatewayIPAddress` | ***string***| ***(Optional)*** |
| `gatewayName` | ***string***||
| `gatewayTimezone` | ***string***||
| `gatewayType` | ***string***| ***(Optional)*** |
| `mediumChangerType` | ***string***| ***(Optional)*** |
| `smbActiveDirectorySettings` | ***[[]StoragegatewayGatewaySpecSmbActiveDirectorySettings](#StoragegatewayGatewaySpecSmbActiveDirectorySettings)***| ***(Optional)*** |
| `tapeDriveType` | ***string***| ***(Optional)*** |
## StoragegatewayGatewaySpecSmbActiveDirectorySettings
##### (Appears on:[StoragegatewayGatewaySpec](#StoragegatewayGatewaySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `domainName` | ***string***||
| `username` | ***string***||
## StoragegatewayGatewayStatus
##### (Appears on:[StoragegatewayGateway](#StoragegatewayGateway))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StoragegatewayGatewaySpec](#StoragegatewayGatewaySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `smb_active_directory_settings.<index>.password` | ***string*** ||
| `smb_guest_password` | ***string*** ||
