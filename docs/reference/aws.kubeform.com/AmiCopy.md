---
title: AmiCopy
menu:
  docs_{{ .version }}:
    identifier: amicopy-aws.kubeform.com
    name: AmiCopy
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AmiCopy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AmiCopy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AmiCopySpec](#amicopyspec)***||
| `status` | ***[AmiCopyStatus](#amicopystatus)***||
## AmiCopySpec

Appears on:[AmiCopy](#amicopy), [AmiCopyStatus](#amicopystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `architecture` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `ebsBlockDevice` | ***[[]AmiCopySpecEbsBlockDevice](#amicopyspecebsblockdevice)***| ***(Optional)*** |
| `enaSupport` | ***bool***| ***(Optional)*** |
| `encrypted` | ***bool***| ***(Optional)*** |
| `ephemeralBlockDevice` | ***[[]AmiCopySpecEphemeralBlockDevice](#amicopyspecephemeralblockdevice)***| ***(Optional)*** |
| `imageLocation` | ***string***| ***(Optional)*** |
| `kernelID` | ***string***| ***(Optional)*** |
| `kmsKeyID` | ***string***| ***(Optional)*** |
| `manageEbsSnapshots` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `ramdiskID` | ***string***| ***(Optional)*** |
| `rootDeviceName` | ***string***| ***(Optional)*** |
| `rootSnapshotID` | ***string***| ***(Optional)*** |
| `sourceAmiID` | ***string***||
| `sourceAmiRegion` | ***string***||
| `sriovNetSupport` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `virtualizationType` | ***string***| ***(Optional)*** |
## AmiCopySpecEbsBlockDevice

Appears on:[AmiCopySpec](#amicopyspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `deleteOnTermination` | ***bool***| ***(Optional)*** |
| `deviceName` | ***string***| ***(Optional)*** |
| `encrypted` | ***bool***| ***(Optional)*** |
| `iops` | ***int***| ***(Optional)*** |
| `snapshotID` | ***string***| ***(Optional)*** |
| `volumeSize` | ***int***| ***(Optional)*** |
| `volumeType` | ***string***| ***(Optional)*** |
## AmiCopySpecEphemeralBlockDevice

Appears on:[AmiCopySpec](#amicopyspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `deviceName` | ***string***| ***(Optional)*** |
| `virtualName` | ***string***| ***(Optional)*** |
## AmiCopyStatus

Appears on:[AmiCopy](#amicopy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AmiCopySpec](#amicopyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
