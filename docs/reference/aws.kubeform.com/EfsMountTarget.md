---
title: EfsMountTarget
menu:
  docs_{{ .version }}:
    identifier: efsmounttarget-aws.kubeform.com
    name: EfsMountTarget
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## EfsMountTarget
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `EfsMountTarget` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[EfsMountTargetSpec](#efsmounttargetspec)***||
| `status` | ***[EfsMountTargetStatus](#efsmounttargetstatus)***||
## EfsMountTargetSpec

Appears on:[EfsMountTarget](#efsmounttarget), [EfsMountTargetStatus](#efsmounttargetstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `dnsName` | ***string***| ***(Optional)*** |
| `fileSystemArn` | ***string***| ***(Optional)*** |
| `fileSystemID` | ***string***||
| `ipAddress` | ***string***| ***(Optional)*** |
| `networkInterfaceID` | ***string***| ***(Optional)*** |
| `securityGroups` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***||
## EfsMountTargetStatus

Appears on:[EfsMountTarget](#efsmounttarget)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[EfsMountTargetSpec](#efsmounttargetspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[EfsMountTargetStatus](#efsmounttargetstatus)

---
