---
title: NetworkInterface
menu:
  docs_{{ .version }}:
    identifier: networkinterface-aws.kubeform.com
    name: NetworkInterface
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## NetworkInterface
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `NetworkInterface` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NetworkInterfaceSpec](#networkinterfacespec)***||
| `status` | ***[NetworkInterfaceStatus](#networkinterfacestatus)***||
## NetworkInterfaceSpec

Appears on:[NetworkInterface](#networkinterface), [NetworkInterfaceStatus](#networkinterfacestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `attachment` | ***[[]NetworkInterfaceSpecAttachment](#networkinterfacespecattachment)***| ***(Optional)*** |
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

Appears on:[NetworkInterfaceSpec](#networkinterfacespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `attachmentID` | ***string***| ***(Optional)*** |
| `deviceIndex` | ***int***||
| `instance` | ***string***||
## NetworkInterfaceStatus

Appears on:[NetworkInterface](#networkinterface)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NetworkInterfaceSpec](#networkinterfacespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[NetworkInterfaceStatus](#networkinterfacestatus)

---
