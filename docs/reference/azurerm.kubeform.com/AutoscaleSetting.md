## AutoscaleSetting
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `AutoscaleSetting` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AutoscaleSettingSpec](#AutoscaleSettingSpec)***||
| `status` | ***[AutoscaleSettingStatus](#AutoscaleSettingStatus)***||
## AutoscaleSettingSpec
##### (Appears on:[AutoscaleSetting](#AutoscaleSetting), [AutoscaleSettingStatus](#AutoscaleSettingStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `enabled` | ***bool***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `notification` | ***[[]AutoscaleSettingSpecNotification](#AutoscaleSettingSpecNotification)***| ***(Optional)*** |
| `profile` | ***[[]AutoscaleSettingSpecProfile](#AutoscaleSettingSpecProfile)***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `targetResourceID` | ***string***||
## AutoscaleSettingSpecNotification
##### (Appears on:[AutoscaleSettingSpec](#AutoscaleSettingSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `email` | ***[[]AutoscaleSettingSpecNotificationEmail](#AutoscaleSettingSpecNotificationEmail)***| ***(Optional)*** |
| `webhook` | ***[[]AutoscaleSettingSpecNotificationWebhook](#AutoscaleSettingSpecNotificationWebhook)***| ***(Optional)*** |
## AutoscaleSettingSpecNotificationEmail
##### (Appears on:[AutoscaleSettingSpecNotification](#AutoscaleSettingSpecNotification))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `customEmails` | ***[]string***| ***(Optional)*** |
| `sendToSubscriptionAdministrator` | ***bool***| ***(Optional)*** |
| `sendToSubscriptionCoAdministrator` | ***bool***| ***(Optional)*** |
## AutoscaleSettingSpecNotificationWebhook
##### (Appears on:[AutoscaleSettingSpecNotification](#AutoscaleSettingSpecNotification))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `properties` | ***map[string]string***| ***(Optional)*** |
| `serviceURI` | ***string***||
## AutoscaleSettingSpecProfile
##### (Appears on:[AutoscaleSettingSpec](#AutoscaleSettingSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `capacity` | ***[[]AutoscaleSettingSpecProfileCapacity](#AutoscaleSettingSpecProfileCapacity)***||
| `fixedDate` | ***[[]AutoscaleSettingSpecProfileFixedDate](#AutoscaleSettingSpecProfileFixedDate)***| ***(Optional)*** |
| `name` | ***string***||
| `recurrence` | ***[[]AutoscaleSettingSpecProfileRecurrence](#AutoscaleSettingSpecProfileRecurrence)***| ***(Optional)*** |
| `rule` | ***[[]AutoscaleSettingSpecProfileRule](#AutoscaleSettingSpecProfileRule)***| ***(Optional)*** |
## AutoscaleSettingSpecProfileCapacity
##### (Appears on:[AutoscaleSettingSpecProfile](#AutoscaleSettingSpecProfile))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `default` | ***int***||
| `maximum` | ***int***||
| `minimum` | ***int***||
## AutoscaleSettingSpecProfileFixedDate
##### (Appears on:[AutoscaleSettingSpecProfile](#AutoscaleSettingSpecProfile))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `end` | ***string***||
| `start` | ***string***||
| `timezone` | ***string***| ***(Optional)*** |
## AutoscaleSettingSpecProfileRecurrence
##### (Appears on:[AutoscaleSettingSpecProfile](#AutoscaleSettingSpecProfile))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `days` | ***[]string***||
| `hours` | ***[]int64***||
| `minutes` | ***[]int64***||
| `timezone` | ***string***| ***(Optional)*** |
## AutoscaleSettingSpecProfileRule
##### (Appears on:[AutoscaleSettingSpecProfile](#AutoscaleSettingSpecProfile))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `metricTrigger` | ***[[]AutoscaleSettingSpecProfileRuleMetricTrigger](#AutoscaleSettingSpecProfileRuleMetricTrigger)***||
| `scaleAction` | ***[[]AutoscaleSettingSpecProfileRuleScaleAction](#AutoscaleSettingSpecProfileRuleScaleAction)***||
## AutoscaleSettingSpecProfileRuleMetricTrigger
##### (Appears on:[AutoscaleSettingSpecProfileRule](#AutoscaleSettingSpecProfileRule))
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
## AutoscaleSettingSpecProfileRuleScaleAction
##### (Appears on:[AutoscaleSettingSpecProfileRule](#AutoscaleSettingSpecProfileRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `cooldown` | ***string***||
| `direction` | ***string***||
| `type` | ***string***||
| `value` | ***int***||
## AutoscaleSettingStatus
##### (Appears on:[AutoscaleSetting](#AutoscaleSetting))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AutoscaleSettingSpec](#AutoscaleSettingSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
