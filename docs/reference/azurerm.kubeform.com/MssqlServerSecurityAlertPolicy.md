---
title: MssqlServerSecurityAlertPolicy
menu:
  docs_{{ .version }}:
    identifier: mssqlserversecurityalertpolicy-azurerm.kubeform.com
    name: MssqlServerSecurityAlertPolicy
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## MssqlServerSecurityAlertPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `MssqlServerSecurityAlertPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MssqlServerSecurityAlertPolicySpec](#mssqlserversecurityalertpolicyspec)***||
| `status` | ***[MssqlServerSecurityAlertPolicyStatus](#mssqlserversecurityalertpolicystatus)***||
## MssqlServerSecurityAlertPolicySpec

Appears on:[MssqlServerSecurityAlertPolicy](#mssqlserversecurityalertpolicy), [MssqlServerSecurityAlertPolicyStatus](#mssqlserversecurityalertpolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `disabledAlerts` | ***[]string***| ***(Optional)*** |
| `emailAccountAdmins` | ***bool***| ***(Optional)*** |
| `emailAddresses` | ***[]string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `retentionDays` | ***int64***| ***(Optional)*** |
| `serverName` | ***string***||
| `state` | ***string***||
| `storageEndpoint` | ***string***| ***(Optional)*** |
## MssqlServerSecurityAlertPolicyStatus

Appears on:[MssqlServerSecurityAlertPolicy](#mssqlserversecurityalertpolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MssqlServerSecurityAlertPolicySpec](#mssqlserversecurityalertpolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[MssqlServerSecurityAlertPolicyStatus](#mssqlserversecurityalertpolicystatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `storage_account_access_key` | ***string*** ||
