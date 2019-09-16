---
title: StoragegatewayCachedIscsiVolume
menu:
  docs_{{ .version }}:
    identifier: storagegatewaycachediscsivolume-aws.kubeform.com
    name: StoragegatewayCachedIscsiVolume
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StoragegatewayCachedIscsiVolume
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `StoragegatewayCachedIscsiVolume` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StoragegatewayCachedIscsiVolumeSpec](#StoragegatewayCachedIscsiVolumeSpec)***||
| `status` | ***[StoragegatewayCachedIscsiVolumeStatus](#StoragegatewayCachedIscsiVolumeStatus)***||
## StoragegatewayCachedIscsiVolumeSpec

Appears on:[StoragegatewayCachedIscsiVolume](#StoragegatewayCachedIscsiVolume), [StoragegatewayCachedIscsiVolumeStatus](#StoragegatewayCachedIscsiVolumeStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `chapEnabled` | ***bool***| ***(Optional)*** |
| `gatewayArn` | ***string***||
| `lunNumber` | ***int***| ***(Optional)*** |
| `networkInterfaceID` | ***string***||
| `networkInterfacePort` | ***int***| ***(Optional)*** |
| `snapshotID` | ***string***| ***(Optional)*** |
| `sourceVolumeArn` | ***string***| ***(Optional)*** |
| `targetArn` | ***string***| ***(Optional)*** |
| `targetName` | ***string***||
| `volumeArn` | ***string***| ***(Optional)*** |
| `volumeID` | ***string***| ***(Optional)*** |
| `volumeSizeInBytes` | ***int***||
## StoragegatewayCachedIscsiVolumeStatus

Appears on:[StoragegatewayCachedIscsiVolume](#StoragegatewayCachedIscsiVolume)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StoragegatewayCachedIscsiVolumeSpec](#StoragegatewayCachedIscsiVolumeSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
