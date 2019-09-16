---
title: NotificationHubAuthorizationRule
menu:
  docs_{{ .version }}:
    identifier: notificationhubauthorizationrule-azurerm.kubeform.com
    name: NotificationHubAuthorizationRule
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## NotificationHubAuthorizationRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `NotificationHubAuthorizationRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NotificationHubAuthorizationRuleSpec](#notificationhubauthorizationrulespec)***||
| `status` | ***[NotificationHubAuthorizationRuleStatus](#notificationhubauthorizationrulestatus)***||
## NotificationHubAuthorizationRuleSpec

Appears on:[NotificationHubAuthorizationRule](#notificationhubauthorizationrule), [NotificationHubAuthorizationRuleStatus](#notificationhubauthorizationrulestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `listen` | ***bool***| ***(Optional)*** |
| `manage` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `namespaceName` | ***string***||
| `notificationHubName` | ***string***||
| `primaryAccessKey` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `secondaryAccessKey` | ***string***| ***(Optional)*** |
| `send` | ***bool***| ***(Optional)*** |
## NotificationHubAuthorizationRuleStatus

Appears on:[NotificationHubAuthorizationRule](#notificationhubauthorizationrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NotificationHubAuthorizationRuleSpec](#notificationhubauthorizationrulespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
