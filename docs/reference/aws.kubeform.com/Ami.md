---
title: Ami
menu:
  docs_{{ .version }}:
    identifier: ami-aws.kubeform.com
    name: Ami
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Ami
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Ami` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AmiSpec](#amispec)***||
| `status` | ***[AmiStatus](#amistatus)***||
## AmiSpec

Appears on:[Ami](#ami), [AmiStatus](#amistatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `architecture` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `ebsBlockDevice` | ***[[]AmiSpecEbsBlockDevice](#amispecebsblockdevice)***| ***(Optional)*** |
| `enaSupport` | ***bool***| ***(Optional)*** |
| `ephemeralBlockDevice` | ***[[]AmiSpecEphemeralBlockDevice](#amispecephemeralblockdevice)***| ***(Optional)*** |
| `imageLocation` | ***string***| ***(Optional)*** |
| `kernelID` | ***string***| ***(Optional)*** |
| `manageEbsSnapshots` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `ramdiskID` | ***string***| ***(Optional)*** |
| `rootDeviceName` | ***string***| ***(Optional)*** |
| `rootSnapshotID` | ***string***| ***(Optional)*** |
| `sriovNetSupport` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `virtualizationType` | ***string***| ***(Optional)*** |
## AmiSpecEbsBlockDevice

Appears on:[AmiSpec](#amispec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `deleteOnTermination` | ***bool***| ***(Optional)*** |
| `deviceName` | ***string***||
| `encrypted` | ***bool***| ***(Optional)*** |
| `iops` | ***int***| ***(Optional)*** |
| `snapshotID` | ***string***| ***(Optional)*** |
| `volumeSize` | ***int***| ***(Optional)*** |
| `volumeType` | ***string***| ***(Optional)*** |
## AmiSpecEphemeralBlockDevice

Appears on:[AmiSpec](#amispec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `deviceName` | ***string***||
| `virtualName` | ***string***||
## AmiStatus

Appears on:[Ami](#ami)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AmiSpec](#amispec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
