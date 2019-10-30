---
title: CloudformationStackSetInstance
menu:
  docs_{{ .version }}:
    identifier: cloudformationstacksetinstance-aws.kubeform.com
    name: CloudformationStackSetInstance
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CloudformationStackSetInstance
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CloudformationStackSetInstance` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CloudformationStackSetInstanceSpec](#cloudformationstacksetinstancespec)***||
| `status` | ***[CloudformationStackSetInstanceStatus](#cloudformationstacksetinstancestatus)***||
## CloudformationStackSetInstanceSpec

Appears on:[CloudformationStackSetInstance](#cloudformationstacksetinstance), [CloudformationStackSetInstanceStatus](#cloudformationstacksetinstancestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountID` | ***string***| ***(Optional)*** |
| `parameterOverrides` | ***map[string]string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `retainStack` | ***bool***| ***(Optional)*** |
| `stackID` | ***string***| ***(Optional)*** |
| `stackSetName` | ***string***||
## CloudformationStackSetInstanceStatus

Appears on:[CloudformationStackSetInstance](#cloudformationstacksetinstance)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CloudformationStackSetInstanceSpec](#cloudformationstacksetinstancespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[CloudformationStackSetInstanceStatus](#cloudformationstacksetinstancestatus)

---
