---
title: MonitorActionGroup
menu:
  docs_{{ .version }}:
    identifier: monitoractiongroup-azurerm.kubeform.com
    name: MonitorActionGroup
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## MonitorActionGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `MonitorActionGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MonitorActionGroupSpec](#monitoractiongroupspec)***||
| `status` | ***[MonitorActionGroupStatus](#monitoractiongroupstatus)***||
## MonitorActionGroupSpec

Appears on:[MonitorActionGroup](#monitoractiongroup), [MonitorActionGroupStatus](#monitoractiongroupstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `emailReceiver` | ***[[]MonitorActionGroupSpecEmailReceiver](#monitoractiongroupspecemailreceiver)***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `shortName` | ***string***||
| `smsReceiver` | ***[[]MonitorActionGroupSpecSmsReceiver](#monitoractiongroupspecsmsreceiver)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `webhookReceiver` | ***[[]MonitorActionGroupSpecWebhookReceiver](#monitoractiongroupspecwebhookreceiver)***| ***(Optional)*** |
## MonitorActionGroupSpecEmailReceiver

Appears on:[MonitorActionGroupSpec](#monitoractiongroupspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `emailAddress` | ***string***||
| `name` | ***string***||
## MonitorActionGroupSpecSmsReceiver

Appears on:[MonitorActionGroupSpec](#monitoractiongroupspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `countryCode` | ***string***||
| `name` | ***string***||
| `phoneNumber` | ***string***||
## MonitorActionGroupSpecWebhookReceiver

Appears on:[MonitorActionGroupSpec](#monitoractiongroupspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `serviceURI` | ***string***||
## MonitorActionGroupStatus

Appears on:[MonitorActionGroup](#monitoractiongroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MonitorActionGroupSpec](#monitoractiongroupspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[MonitorActionGroupStatus](#monitoractiongroupstatus)

---
