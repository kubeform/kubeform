---
title: DatasyncLocationEfs
menu:
  docs_v0.0.1:
    identifier: datasynclocationefs-aws.kubeform.com
    name: DatasyncLocationEfs
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## DatasyncLocationEfs
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DatasyncLocationEfs` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DatasyncLocationEfsSpec](#DatasyncLocationEfsSpec)***||
| `status` | ***[DatasyncLocationEfsStatus](#DatasyncLocationEfsStatus)***||
## DatasyncLocationEfsSpec
##### (Appears on:[DatasyncLocationEfs](#DatasyncLocationEfs), [DatasyncLocationEfsStatus](#DatasyncLocationEfsStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `ec2Config` | ***[[]DatasyncLocationEfsSpecEc2Config](#DatasyncLocationEfsSpecEc2Config)***||
| `efsFileSystemArn` | ***string***||
| `subdirectory` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `uri` | ***string***| ***(Optional)*** |
## DatasyncLocationEfsSpecEc2Config
##### (Appears on:[DatasyncLocationEfsSpec](#DatasyncLocationEfsSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `securityGroupArns` | ***[]string***||
| `subnetArn` | ***string***||
## DatasyncLocationEfsStatus
##### (Appears on:[DatasyncLocationEfs](#DatasyncLocationEfs))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DatasyncLocationEfsSpec](#DatasyncLocationEfsSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
