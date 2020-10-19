---
title: FsxLustreFileSystem
menu:
  docs_{{ .version }}:
    identifier: fsxlustrefilesystem-aws.kubeform.com
    name: FsxLustreFileSystem
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## FsxLustreFileSystem
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `FsxLustreFileSystem` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[FsxLustreFileSystemSpec](#fsxlustrefilesystemspec)***||
| `status` | ***[FsxLustreFileSystemStatus](#fsxlustrefilesystemstatus)***||
## FsxLustreFileSystemSpec

Appears on:[FsxLustreFileSystem](#fsxlustrefilesystem), [FsxLustreFileSystemStatus](#fsxlustrefilesystemstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `dnsName` | ***string***| ***(Optional)*** |
| `exportPath` | ***string***| ***(Optional)*** |
| `importPath` | ***string***| ***(Optional)*** |
| `importedFileChunkSize` | ***int64***| ***(Optional)*** |
| `networkInterfaceIDS` | ***[]string***| ***(Optional)*** |
| `ownerID` | ***string***| ***(Optional)*** |
| `securityGroupIDS` | ***[]string***| ***(Optional)*** |
| `storageCapacity` | ***int64***||
| `subnetIDS` | ***[]string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpcID` | ***string***| ***(Optional)*** |
| `weeklyMaintenanceStartTime` | ***string***| ***(Optional)*** |
## FsxLustreFileSystemStatus

Appears on:[FsxLustreFileSystem](#fsxlustrefilesystem)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[FsxLustreFileSystemSpec](#fsxlustrefilesystemspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[FsxLustreFileSystemStatus](#fsxlustrefilesystemstatus)

---
