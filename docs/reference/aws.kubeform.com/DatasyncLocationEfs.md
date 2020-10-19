---
title: DatasyncLocationEfs
menu:
  docs_{{ .version }}:
    identifier: datasynclocationefs-aws.kubeform.com
    name: DatasyncLocationEfs
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DatasyncLocationEfs
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DatasyncLocationEfs` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DatasyncLocationEfsSpec](#datasynclocationefsspec)***||
| `status` | ***[DatasyncLocationEfsStatus](#datasynclocationefsstatus)***||
## DatasyncLocationEfsSpec

Appears on:[DatasyncLocationEfs](#datasynclocationefs), [DatasyncLocationEfsStatus](#datasynclocationefsstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `ec2Config` | ***[[]DatasyncLocationEfsSpecEc2Config](#datasynclocationefsspecec2config)***||
| `efsFileSystemArn` | ***string***||
| `subdirectory` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `uri` | ***string***| ***(Optional)*** |
## DatasyncLocationEfsSpecEc2Config

Appears on:[DatasyncLocationEfsSpec](#datasynclocationefsspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `securityGroupArns` | ***[]string***||
| `subnetArn` | ***string***||
## DatasyncLocationEfsStatus

Appears on:[DatasyncLocationEfs](#datasynclocationefs)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DatasyncLocationEfsSpec](#datasynclocationefsspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DatasyncLocationEfsStatus](#datasynclocationefsstatus)

---
