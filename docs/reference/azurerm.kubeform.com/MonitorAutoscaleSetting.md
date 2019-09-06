## MonitorAutoscaleSetting
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `MonitorAutoscaleSetting` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MonitorAutoscaleSettingSpec](#MonitorAutoscaleSettingSpec)***||
| `status` | ***[MonitorAutoscaleSettingStatus](#MonitorAutoscaleSettingStatus)***||
## MonitorAutoscaleSettingSpec
##### (Appears on:[MonitorAutoscaleSetting](#MonitorAutoscaleSetting), [MonitorAutoscaleSettingStatus](#MonitorAutoscaleSettingStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `enabled` | ***bool***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `notification` | ***[[]MonitorAutoscaleSettingSpecNotification](#MonitorAutoscaleSettingSpecNotification)***| ***(Optional)*** |
| `profile` | ***[[]MonitorAutoscaleSettingSpecProfile](#MonitorAutoscaleSettingSpecProfile)***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `targetResourceID` | ***string***||
## MonitorAutoscaleSettingSpecNotification
##### (Appears on:[MonitorAutoscaleSettingSpec](#MonitorAutoscaleSettingSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `email` | ***[[]MonitorAutoscaleSettingSpecNotificationEmail](#MonitorAutoscaleSettingSpecNotificationEmail)***| ***(Optional)*** |
| `webhook` | ***[[]MonitorAutoscaleSettingSpecNotificationWebhook](#MonitorAutoscaleSettingSpecNotificationWebhook)***| ***(Optional)*** |
## MonitorAutoscaleSettingSpecNotificationEmail
##### (Appears on:[MonitorAutoscaleSettingSpecNotification](#MonitorAutoscaleSettingSpecNotification))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `customEmails` | ***[]string***| ***(Optional)*** |
| `sendToSubscriptionAdministrator` | ***bool***| ***(Optional)*** |
| `sendToSubscriptionCoAdministrator` | ***bool***| ***(Optional)*** |
## MonitorAutoscaleSettingSpecNotificationWebhook
##### (Appears on:[MonitorAutoscaleSettingSpecNotification](#MonitorAutoscaleSettingSpecNotification))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `properties` | ***map[string]string***| ***(Optional)*** |
| `serviceURI` | ***string***||
## MonitorAutoscaleSettingSpecProfile
##### (Appears on:[MonitorAutoscaleSettingSpec](#MonitorAutoscaleSettingSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `capacity` | ***[[]MonitorAutoscaleSettingSpecProfileCapacity](#MonitorAutoscaleSettingSpecProfileCapacity)***||
| `fixedDate` | ***[[]MonitorAutoscaleSettingSpecProfileFixedDate](#MonitorAutoscaleSettingSpecProfileFixedDate)***| ***(Optional)*** |
| `name` | ***string***||
| `recurrence` | ***[[]MonitorAutoscaleSettingSpecProfileRecurrence](#MonitorAutoscaleSettingSpecProfileRecurrence)***| ***(Optional)*** |
| `rule` | ***[[]MonitorAutoscaleSettingSpecProfileRule](#MonitorAutoscaleSettingSpecProfileRule)***| ***(Optional)*** |
## MonitorAutoscaleSettingSpecProfileCapacity
##### (Appears on:[MonitorAutoscaleSettingSpecProfile](#MonitorAutoscaleSettingSpecProfile))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `default` | ***int***||
| `maximum` | ***int***||
| `minimum` | ***int***||
## MonitorAutoscaleSettingSpecProfileFixedDate
##### (Appears on:[MonitorAutoscaleSettingSpecProfile](#MonitorAutoscaleSettingSpecProfile))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `end` | ***string***||
| `start` | ***string***||
| `timezone` | ***string***| ***(Optional)*** |
## MonitorAutoscaleSettingSpecProfileRecurrence
##### (Appears on:[MonitorAutoscaleSettingSpecProfile](#MonitorAutoscaleSettingSpecProfile))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `days` | ***[]string***||
| `hours` | ***[]int64***||
| `minutes` | ***[]int64***||
| `timezone` | ***string***| ***(Optional)*** |
## MonitorAutoscaleSettingSpecProfileRule
##### (Appears on:[MonitorAutoscaleSettingSpecProfile](#MonitorAutoscaleSettingSpecProfile))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `metricTrigger` | ***[[]MonitorAutoscaleSettingSpecProfileRuleMetricTrigger](#MonitorAutoscaleSettingSpecProfileRuleMetricTrigger)***||
| `scaleAction` | ***[[]MonitorAutoscaleSettingSpecProfileRuleScaleAction](#MonitorAutoscaleSettingSpecProfileRuleScaleAction)***||
## MonitorAutoscaleSettingSpecProfileRuleMetricTrigger
##### (Appears on:[MonitorAutoscaleSettingSpecProfileRule](#MonitorAutoscaleSettingSpecProfileRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `metricName` | ***string***||
| `metricResourceID` | ***string***||
| `operator` | ***string***||
| `statistic` | ***string***||
| `threshold` | ***encoding/json.Number***||
| `timeAggregation` | ***string***||
| `timeGrain` | ***string***||
| `timeWindow` | ***string***||
## MonitorAutoscaleSettingSpecProfileRuleScaleAction
##### (Appears on:[MonitorAutoscaleSettingSpecProfileRule](#MonitorAutoscaleSettingSpecProfileRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `cooldown` | ***string***||
| `direction` | ***string***||
| `type` | ***string***||
| `value` | ***int***||
## MonitorAutoscaleSettingStatus
##### (Appears on:[MonitorAutoscaleSetting](#MonitorAutoscaleSetting))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MonitorAutoscaleSettingSpec](#MonitorAutoscaleSettingSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
