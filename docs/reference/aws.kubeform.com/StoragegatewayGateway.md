---
title: StoragegatewayGateway
menu:
  docs_{{ .version }}:
    identifier: storagegatewaygateway-aws.kubeform.com
    name: StoragegatewayGateway
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StoragegatewayGateway
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `StoragegatewayGateway` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StoragegatewayGatewaySpec](#storagegatewaygatewayspec)***||
| `status` | ***[StoragegatewayGatewayStatus](#storagegatewaygatewaystatus)***||
## Phase(`string` alias)

Appears on:[StoragegatewayGatewayStatus](#storagegatewaygatewaystatus)

## StoragegatewayGatewaySpec

Appears on:[StoragegatewayGateway](#storagegatewaygateway), [StoragegatewayGatewayStatus](#storagegatewaygatewaystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `activationKey` | ***string***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `gatewayID` | ***string***| ***(Optional)*** |
| `gatewayIPAddress` | ***string***| ***(Optional)*** |
| `gatewayName` | ***string***||
| `gatewayTimezone` | ***string***||
| `gatewayType` | ***string***| ***(Optional)*** |
| `mediumChangerType` | ***string***| ***(Optional)*** |
| `smbActiveDirectorySettings` | ***[[]StoragegatewayGatewaySpecSmbActiveDirectorySettings](#storagegatewaygatewayspecsmbactivedirectorysettings)***| ***(Optional)*** |
| `tapeDriveType` | ***string***| ***(Optional)*** |
## StoragegatewayGatewaySpecSmbActiveDirectorySettings

Appears on:[StoragegatewayGatewaySpec](#storagegatewaygatewayspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `domainName` | ***string***||
| `username` | ***string***||
## StoragegatewayGatewayStatus

Appears on:[StoragegatewayGateway](#storagegatewaygateway)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StoragegatewayGatewaySpec](#storagegatewaygatewayspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `smb_active_directory_settings.<index>.password` | ***string*** ||
| `smb_guest_password` | ***string*** ||
