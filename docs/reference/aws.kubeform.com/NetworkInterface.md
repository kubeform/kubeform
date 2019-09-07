---
title: NetworkInterface
menu:
  docs_v0.0.1:
    identifier: networkinterface-aws.kubeform.com
    name: NetworkInterface
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## NetworkInterface
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `NetworkInterface` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NetworkInterfaceSpec](#NetworkInterfaceSpec)***||
| `status` | ***[NetworkInterfaceStatus](#NetworkInterfaceStatus)***||
## NetworkInterfaceSpec
##### (Appears on:[NetworkInterface](#NetworkInterface), [NetworkInterfaceStatus](#NetworkInterfaceStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `attachment` | ***[[]NetworkInterfaceSpecAttachment](#NetworkInterfaceSpecAttachment)***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `privateDNSName` | ***string***| ***(Optional)*** |
| `privateIP` | ***string***| ***(Optional)*** |
| `privateIPS` | ***[]string***| ***(Optional)*** |
| `privateIPSCount` | ***int***| ***(Optional)*** |
| `securityGroups` | ***[]string***| ***(Optional)*** |
| `sourceDestCheck` | ***bool***| ***(Optional)*** |
| `subnetID` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## NetworkInterfaceSpecAttachment
##### (Appears on:[NetworkInterfaceSpec](#NetworkInterfaceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `attachmentID` | ***string***| ***(Optional)*** |
| `deviceIndex` | ***int***||
| `instance` | ***string***||
## NetworkInterfaceStatus
##### (Appears on:[NetworkInterface](#NetworkInterface))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NetworkInterfaceSpec](#NetworkInterfaceSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
