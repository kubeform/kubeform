---
title: CloudhsmV2Hsm
menu:
  docs_{{ .version }}:
    identifier: cloudhsmv2hsm-aws.kubeform.com
    name: CloudhsmV2Hsm
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CloudhsmV2Hsm
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CloudhsmV2Hsm` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CloudhsmV2HsmSpec](#CloudhsmV2HsmSpec)***||
| `status` | ***[CloudhsmV2HsmStatus](#CloudhsmV2HsmStatus)***||
## CloudhsmV2HsmSpec
##### (Appears on:[CloudhsmV2Hsm](#CloudhsmV2Hsm), [CloudhsmV2HsmStatus](#CloudhsmV2HsmStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `availabilityZone` | ***string***| ***(Optional)*** |
| `clusterID` | ***string***||
| `hsmEniID` | ***string***| ***(Optional)*** |
| `hsmID` | ***string***| ***(Optional)*** |
| `hsmState` | ***string***| ***(Optional)*** |
| `ipAddress` | ***string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
## CloudhsmV2HsmStatus
##### (Appears on:[CloudhsmV2Hsm](#CloudhsmV2Hsm))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CloudhsmV2HsmSpec](#CloudhsmV2HsmSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
