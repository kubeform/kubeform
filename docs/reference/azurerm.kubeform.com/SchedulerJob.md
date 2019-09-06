## SchedulerJob
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `SchedulerJob` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SchedulerJobSpec](#SchedulerJobSpec)***||
| `status` | ***[SchedulerJobStatus](#SchedulerJobStatus)***||
## SchedulerJobSpec
##### (Appears on:[SchedulerJob](#SchedulerJob), [SchedulerJobStatus](#SchedulerJobStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `actionStorageQueue` | ***[[]SchedulerJobSpecActionStorageQueue](#SchedulerJobSpecActionStorageQueue)***| ***(Optional)*** |
| `actionWeb` | ***[[]SchedulerJobSpecActionWeb](#SchedulerJobSpecActionWeb)***| ***(Optional)*** |
| `errorActionStorageQueue` | ***[[]SchedulerJobSpecErrorActionStorageQueue](#SchedulerJobSpecErrorActionStorageQueue)***| ***(Optional)*** |
| `errorActionWeb` | ***[[]SchedulerJobSpecErrorActionWeb](#SchedulerJobSpecErrorActionWeb)***| ***(Optional)*** |
| `jobCollectionName` | ***string***||
| `name` | ***string***||
| `recurrence` | ***[[]SchedulerJobSpecRecurrence](#SchedulerJobSpecRecurrence)***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `retry` | ***[[]SchedulerJobSpecRetry](#SchedulerJobSpecRetry)***| ***(Optional)*** |
| `startTime` | ***string***| ***(Optional)*** |
| `state` | ***string***| ***(Optional)*** |
## SchedulerJobSpecActionStorageQueue
##### (Appears on:[SchedulerJobSpec](#SchedulerJobSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `message` | ***string***||
| `sasToken` | ***string***||
| `storageAccountName` | ***string***||
| `storageQueueName` | ***string***||
## SchedulerJobSpecActionWeb
##### (Appears on:[SchedulerJobSpec](#SchedulerJobSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `authenticationActiveDirectory` | ***[[]SchedulerJobSpecActionWebAuthenticationActiveDirectory](#SchedulerJobSpecActionWebAuthenticationActiveDirectory)***| ***(Optional)*** |
| `authenticationBasic` | ***[[]SchedulerJobSpecActionWebAuthenticationBasic](#SchedulerJobSpecActionWebAuthenticationBasic)***| ***(Optional)*** |
| `authenticationCertificate` | ***[[]SchedulerJobSpecActionWebAuthenticationCertificate](#SchedulerJobSpecActionWebAuthenticationCertificate)***| ***(Optional)*** |
| `body` | ***string***| ***(Optional)*** |
| `headers` | ***map[string]string***| ***(Optional)*** |
| `method` | ***string***||
| `url` | ***string***||
## SchedulerJobSpecActionWebAuthenticationActiveDirectory
##### (Appears on:[SchedulerJobSpecActionWeb](#SchedulerJobSpecActionWeb))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `audience` | ***string***| ***(Optional)*** |
| `clientID` | ***string***||
| `tenantID` | ***string***||
## SchedulerJobSpecActionWebAuthenticationBasic
##### (Appears on:[SchedulerJobSpecActionWeb](#SchedulerJobSpecActionWeb))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `username` | ***string***||
## SchedulerJobSpecActionWebAuthenticationCertificate
##### (Appears on:[SchedulerJobSpecActionWeb](#SchedulerJobSpecActionWeb))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `expiration` | ***string***| ***(Optional)*** |
| `subjectName` | ***string***| ***(Optional)*** |
| `thumbprint` | ***string***| ***(Optional)*** |
## SchedulerJobSpecErrorActionStorageQueue
##### (Appears on:[SchedulerJobSpec](#SchedulerJobSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `message` | ***string***||
| `sasToken` | ***string***||
| `storageAccountName` | ***string***||
| `storageQueueName` | ***string***||
## SchedulerJobSpecErrorActionWeb
##### (Appears on:[SchedulerJobSpec](#SchedulerJobSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `authenticationActiveDirectory` | ***[[]SchedulerJobSpecErrorActionWebAuthenticationActiveDirectory](#SchedulerJobSpecErrorActionWebAuthenticationActiveDirectory)***| ***(Optional)*** |
| `authenticationBasic` | ***[[]SchedulerJobSpecErrorActionWebAuthenticationBasic](#SchedulerJobSpecErrorActionWebAuthenticationBasic)***| ***(Optional)*** |
| `authenticationCertificate` | ***[[]SchedulerJobSpecErrorActionWebAuthenticationCertificate](#SchedulerJobSpecErrorActionWebAuthenticationCertificate)***| ***(Optional)*** |
| `body` | ***string***| ***(Optional)*** |
| `headers` | ***map[string]string***| ***(Optional)*** |
| `method` | ***string***||
| `url` | ***string***||
## SchedulerJobSpecErrorActionWebAuthenticationActiveDirectory
##### (Appears on:[SchedulerJobSpecErrorActionWeb](#SchedulerJobSpecErrorActionWeb))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `audience` | ***string***| ***(Optional)*** |
| `clientID` | ***string***||
| `tenantID` | ***string***||
## SchedulerJobSpecErrorActionWebAuthenticationBasic
##### (Appears on:[SchedulerJobSpecErrorActionWeb](#SchedulerJobSpecErrorActionWeb))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `username` | ***string***||
## SchedulerJobSpecErrorActionWebAuthenticationCertificate
##### (Appears on:[SchedulerJobSpecErrorActionWeb](#SchedulerJobSpecErrorActionWeb))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `expiration` | ***string***| ***(Optional)*** |
| `subjectName` | ***string***| ***(Optional)*** |
| `thumbprint` | ***string***| ***(Optional)*** |
## SchedulerJobSpecRecurrence
##### (Appears on:[SchedulerJobSpec](#SchedulerJobSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int***| ***(Optional)*** |
| `endTime` | ***string***| ***(Optional)*** |
| `frequency` | ***string***||
| `hours` | ***[]int64***| ***(Optional)*** |
| `interval` | ***int***| ***(Optional)*** |
| `minutes` | ***[]int64***| ***(Optional)*** |
| `monthDays` | ***[]int64***| ***(Optional)*** |
| `monthlyOccurrences` | ***[[]SchedulerJobSpecRecurrenceMonthlyOccurrences](#SchedulerJobSpecRecurrenceMonthlyOccurrences)***| ***(Optional)*** |
| `weekDays` | ***[]string***| ***(Optional)*** |
## SchedulerJobSpecRecurrenceMonthlyOccurrences
##### (Appears on:[SchedulerJobSpecRecurrence](#SchedulerJobSpecRecurrence))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `day` | ***string***||
| `occurrence` | ***int***||
## SchedulerJobSpecRetry
##### (Appears on:[SchedulerJobSpec](#SchedulerJobSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int***| ***(Optional)*** |
| `interval` | ***string***| ***(Optional)*** |
## SchedulerJobStatus
##### (Appears on:[SchedulerJob](#SchedulerJob))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SchedulerJobSpec](#SchedulerJobSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `action_web.<index>.authentication_active_directory.<index>.secret` | ***string*** ||
| `action_web.<index>.authentication_basic.<index>.password` | ***string*** ||
| `action_web.<index>.authentication_certificate.<index>.password` | ***string*** ||
| `action_web.<index>.authentication_certificate.<index>.pfx` | ***string*** ||
| `error_action_web.<index>.authentication_active_directory.<index>.secret` | ***string*** ||
| `error_action_web.<index>.authentication_basic.<index>.password` | ***string*** ||
| `error_action_web.<index>.authentication_certificate.<index>.password` | ***string*** ||
| `error_action_web.<index>.authentication_certificate.<index>.pfx` | ***string*** ||
