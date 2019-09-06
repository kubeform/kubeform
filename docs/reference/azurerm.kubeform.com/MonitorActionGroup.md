## MonitorActionGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `MonitorActionGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MonitorActionGroupSpec](#MonitorActionGroupSpec)***||
| `status` | ***[MonitorActionGroupStatus](#MonitorActionGroupStatus)***||
## MonitorActionGroupSpec
##### (Appears on:[MonitorActionGroup](#MonitorActionGroup), [MonitorActionGroupStatus](#MonitorActionGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `emailReceiver` | ***[[]MonitorActionGroupSpecEmailReceiver](#MonitorActionGroupSpecEmailReceiver)***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `shortName` | ***string***||
| `smsReceiver` | ***[[]MonitorActionGroupSpecSmsReceiver](#MonitorActionGroupSpecSmsReceiver)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `webhookReceiver` | ***[[]MonitorActionGroupSpecWebhookReceiver](#MonitorActionGroupSpecWebhookReceiver)***| ***(Optional)*** |
## MonitorActionGroupSpecEmailReceiver
##### (Appears on:[MonitorActionGroupSpec](#MonitorActionGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `emailAddress` | ***string***||
| `name` | ***string***||
## MonitorActionGroupSpecSmsReceiver
##### (Appears on:[MonitorActionGroupSpec](#MonitorActionGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `countryCode` | ***string***||
| `name` | ***string***||
| `phoneNumber` | ***string***||
## MonitorActionGroupSpecWebhookReceiver
##### (Appears on:[MonitorActionGroupSpec](#MonitorActionGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `serviceURI` | ***string***||
## MonitorActionGroupStatus
##### (Appears on:[MonitorActionGroup](#MonitorActionGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MonitorActionGroupSpec](#MonitorActionGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
