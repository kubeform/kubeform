---
title: EventhubNamespaceAuthorizationRule
menu:
  docs_{{ .version }}:
    identifier: eventhubnamespaceauthorizationrule-azurerm.kubeform.com
    name: EventhubNamespaceAuthorizationRule
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## EventhubNamespaceAuthorizationRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `EventhubNamespaceAuthorizationRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[EventhubNamespaceAuthorizationRuleSpec](#eventhubnamespaceauthorizationrulespec)***||
| `status` | ***[EventhubNamespaceAuthorizationRuleStatus](#eventhubnamespaceauthorizationrulestatus)***||
## EventhubNamespaceAuthorizationRuleSpec

Appears on:[EventhubNamespaceAuthorizationRule](#eventhubnamespaceauthorizationrule), [EventhubNamespaceAuthorizationRuleStatus](#eventhubnamespaceauthorizationrulestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `listen` | ***bool***| ***(Optional)*** |
| `location` | ***string***| ***(Optional)*** Deprecated|
| `manage` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `namespaceName` | ***string***||
| `resourceGroupName` | ***string***||
| `send` | ***bool***| ***(Optional)*** |
## EventhubNamespaceAuthorizationRuleStatus

Appears on:[EventhubNamespaceAuthorizationRule](#eventhubnamespaceauthorizationrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[EventhubNamespaceAuthorizationRuleSpec](#eventhubnamespaceauthorizationrulespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[EventhubNamespaceAuthorizationRuleStatus](#eventhubnamespaceauthorizationrulestatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `primary_connection_string` | ***string*** ||
| `primary_key` | ***string*** ||
| `secondary_connection_string` | ***string*** ||
| `secondary_key` | ***string*** ||
