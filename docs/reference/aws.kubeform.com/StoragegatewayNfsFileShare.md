---
title: StoragegatewayNfsFileShare
menu:
  docs_{{ .version }}:
    identifier: storagegatewaynfsfileshare-aws.kubeform.com
    name: StoragegatewayNfsFileShare
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StoragegatewayNfsFileShare
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `StoragegatewayNfsFileShare` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StoragegatewayNfsFileShareSpec](#storagegatewaynfsfilesharespec)***||
| `status` | ***[StoragegatewayNfsFileShareStatus](#storagegatewaynfsfilesharestatus)***||
## Phase(`string` alias)

Appears on:[StoragegatewayNfsFileShareStatus](#storagegatewaynfsfilesharestatus)

## StoragegatewayNfsFileShareSpec

Appears on:[StoragegatewayNfsFileShare](#storagegatewaynfsfileshare), [StoragegatewayNfsFileShareStatus](#storagegatewaynfsfilesharestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `clientList` | ***[]string***||
| `defaultStorageClass` | ***string***| ***(Optional)*** |
| `fileshareID` | ***string***| ***(Optional)*** |
| `gatewayArn` | ***string***||
| `guessMimeTypeEnabled` | ***bool***| ***(Optional)*** |
| `kmsEncrypted` | ***bool***| ***(Optional)*** |
| `kmsKeyArn` | ***string***| ***(Optional)*** |
| `locationArn` | ***string***||
| `nfsFileShareDefaults` | ***[[]StoragegatewayNfsFileShareSpecNfsFileShareDefaults](#storagegatewaynfsfilesharespecnfsfilesharedefaults)***| ***(Optional)*** |
| `objectACL` | ***string***| ***(Optional)*** |
| `readOnly` | ***bool***| ***(Optional)*** |
| `requesterPays` | ***bool***| ***(Optional)*** |
| `roleArn` | ***string***||
| `squash` | ***string***| ***(Optional)*** |
## StoragegatewayNfsFileShareSpecNfsFileShareDefaults

Appears on:[StoragegatewayNfsFileShareSpec](#storagegatewaynfsfilesharespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `directoryMode` | ***string***| ***(Optional)*** |
| `fileMode` | ***string***| ***(Optional)*** |
| `groupID` | ***int64***| ***(Optional)*** |
| `ownerID` | ***int64***| ***(Optional)*** |
## StoragegatewayNfsFileShareStatus

Appears on:[StoragegatewayNfsFileShare](#storagegatewaynfsfileshare)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StoragegatewayNfsFileShareSpec](#storagegatewaynfsfilesharespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
