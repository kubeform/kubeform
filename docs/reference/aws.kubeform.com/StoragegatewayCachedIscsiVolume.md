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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StoragegatewayCachedIscsiVolumeSpec](#storagegatewaycachediscsivolumespec)***||
| `status` | ***[StoragegatewayCachedIscsiVolumeStatus](#storagegatewaycachediscsivolumestatus)***||
## Phase(`string` alias)

Appears on:[StoragegatewayCachedIscsiVolumeStatus](#storagegatewaycachediscsivolumestatus)

## StoragegatewayCachedIscsiVolumeSpec

Appears on:[StoragegatewayCachedIscsiVolume](#storagegatewaycachediscsivolume), [StoragegatewayCachedIscsiVolumeStatus](#storagegatewaycachediscsivolumestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `chapEnabled` | ***bool***| ***(Optional)*** |
| `gatewayArn` | ***string***||
| `lunNumber` | ***int64***| ***(Optional)*** |
| `networkInterfaceID` | ***string***||
| `networkInterfacePort` | ***int64***| ***(Optional)*** |
| `snapshotID` | ***string***| ***(Optional)*** |
| `sourceVolumeArn` | ***string***| ***(Optional)*** |
| `targetArn` | ***string***| ***(Optional)*** |
| `targetName` | ***string***||
| `volumeArn` | ***string***| ***(Optional)*** |
| `volumeID` | ***string***| ***(Optional)*** |
| `volumeSizeInBytes` | ***int64***||
## StoragegatewayCachedIscsiVolumeStatus

Appears on:[StoragegatewayCachedIscsiVolume](#storagegatewaycachediscsivolume)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StoragegatewayCachedIscsiVolumeSpec](#storagegatewaycachediscsivolumespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
