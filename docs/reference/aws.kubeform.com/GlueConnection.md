---
title: GlueConnection
menu:
  docs_{{ .version }}:
    identifier: glueconnection-aws.kubeform.com
    name: GlueConnection
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## GlueConnection
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `GlueConnection` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[GlueConnectionSpec](#glueconnectionspec)***||
| `status` | ***[GlueConnectionStatus](#glueconnectionstatus)***||
## GlueConnectionSpec

Appears on:[GlueConnection](#glueconnection), [GlueConnectionStatus](#glueconnectionstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `catalogID` | ***string***| ***(Optional)*** |
| `connectionProperties` | ***map[string]string***||
| `connectionType` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `matchCriteria` | ***[]string***| ***(Optional)*** |
| `name` | ***string***||
| `physicalConnectionRequirements` | ***[[]GlueConnectionSpecPhysicalConnectionRequirements](#glueconnectionspecphysicalconnectionrequirements)***| ***(Optional)*** |
## GlueConnectionSpecPhysicalConnectionRequirements

Appears on:[GlueConnectionSpec](#glueconnectionspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `availabilityZone` | ***string***| ***(Optional)*** |
| `securityGroupIDList` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
## GlueConnectionStatus

Appears on:[GlueConnection](#glueconnection)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[GlueConnectionSpec](#glueconnectionspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
