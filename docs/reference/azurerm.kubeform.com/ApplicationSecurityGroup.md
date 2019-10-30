---
title: ApplicationSecurityGroup
menu:
  docs_{{ .version }}:
    identifier: applicationsecuritygroup-azurerm.kubeform.com
    name: ApplicationSecurityGroup
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApplicationSecurityGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ApplicationSecurityGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApplicationSecurityGroupSpec](#applicationsecuritygroupspec)***||
| `status` | ***[ApplicationSecurityGroupStatus](#applicationsecuritygroupstatus)***||
## ApplicationSecurityGroupSpec

Appears on:[ApplicationSecurityGroup](#applicationsecuritygroup), [ApplicationSecurityGroupStatus](#applicationsecuritygroupstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## ApplicationSecurityGroupStatus

Appears on:[ApplicationSecurityGroup](#applicationsecuritygroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApplicationSecurityGroupSpec](#applicationsecuritygroupspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ApplicationSecurityGroupStatus](#applicationsecuritygroupstatus)

---
