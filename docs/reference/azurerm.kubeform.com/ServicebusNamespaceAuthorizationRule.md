---
title: ServicebusNamespaceAuthorizationRule
menu:
  docs_{{ .version }}:
    identifier: servicebusnamespaceauthorizationrule-azurerm.kubeform.com
    name: ServicebusNamespaceAuthorizationRule
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ServicebusNamespaceAuthorizationRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ServicebusNamespaceAuthorizationRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ServicebusNamespaceAuthorizationRuleSpec](#servicebusnamespaceauthorizationrulespec)***||
| `status` | ***[ServicebusNamespaceAuthorizationRuleStatus](#servicebusnamespaceauthorizationrulestatus)***||
## Phase(`string` alias)

Appears on:[ServicebusNamespaceAuthorizationRuleStatus](#servicebusnamespaceauthorizationrulestatus)

## ServicebusNamespaceAuthorizationRuleSpec

Appears on:[ServicebusNamespaceAuthorizationRule](#servicebusnamespaceauthorizationrule), [ServicebusNamespaceAuthorizationRuleStatus](#servicebusnamespaceauthorizationrulestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `listen` | ***bool***| ***(Optional)*** |
| `manage` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `namespaceName` | ***string***||
| `resourceGroupName` | ***string***||
| `send` | ***bool***| ***(Optional)*** |
## ServicebusNamespaceAuthorizationRuleStatus

Appears on:[ServicebusNamespaceAuthorizationRule](#servicebusnamespaceauthorizationrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ServicebusNamespaceAuthorizationRuleSpec](#servicebusnamespaceauthorizationrulespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `primary_connection_string` | ***string*** ||
| `primary_key` | ***string*** ||
| `secondary_connection_string` | ***string*** ||
| `secondary_key` | ***string*** ||
