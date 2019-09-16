---
title: InspectorResourceGroup
menu:
  docs_{{ .version }}:
    identifier: inspectorresourcegroup-aws.kubeform.com
    name: InspectorResourceGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## InspectorResourceGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `InspectorResourceGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[InspectorResourceGroupSpec](#InspectorResourceGroupSpec)***||
| `status` | ***[InspectorResourceGroupStatus](#InspectorResourceGroupStatus)***||
## InspectorResourceGroupSpec

Appears on:[InspectorResourceGroup](#InspectorResourceGroup), [InspectorResourceGroupStatus](#InspectorResourceGroupStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***||
## InspectorResourceGroupStatus

Appears on:[InspectorResourceGroup](#InspectorResourceGroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[InspectorResourceGroupSpec](#InspectorResourceGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
