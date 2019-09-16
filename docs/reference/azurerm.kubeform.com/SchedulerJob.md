---
title: SchedulerJob
menu:
  docs_{{ .version }}:
    identifier: schedulerjob-azurerm.kubeform.com
    name: SchedulerJob
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SchedulerJob
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `SchedulerJob` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SchedulerJobSpec](#schedulerjobspec)***||
| `status` | ***[SchedulerJobStatus](#schedulerjobstatus)***||
## SchedulerJobSpec

Appears on:[SchedulerJob](#schedulerjob), [SchedulerJobStatus](#schedulerjobstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `actionStorageQueue` | ***[[]SchedulerJobSpecActionStorageQueue](#schedulerjobspecactionstoragequeue)***| ***(Optional)*** |
| `actionWeb` | ***[[]SchedulerJobSpecActionWeb](#schedulerjobspecactionweb)***| ***(Optional)*** |
| `errorActionStorageQueue` | ***[[]SchedulerJobSpecErrorActionStorageQueue](#schedulerjobspecerroractionstoragequeue)***| ***(Optional)*** |
| `errorActionWeb` | ***[[]SchedulerJobSpecErrorActionWeb](#schedulerjobspecerroractionweb)***| ***(Optional)*** |
| `jobCollectionName` | ***string***||
| `name` | ***string***||
| `recurrence` | ***[[]SchedulerJobSpecRecurrence](#schedulerjobspecrecurrence)***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `retry` | ***[[]SchedulerJobSpecRetry](#schedulerjobspecretry)***| ***(Optional)*** |
| `startTime` | ***string***| ***(Optional)*** |
| `state` | ***string***| ***(Optional)*** |
## SchedulerJobSpecActionStorageQueue

Appears on:[SchedulerJobSpec](#schedulerjobspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `message` | ***string***||
| `sasToken` | ***string***||
| `storageAccountName` | ***string***||
| `storageQueueName` | ***string***||
## SchedulerJobSpecActionWeb

Appears on:[SchedulerJobSpec](#schedulerjobspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `authenticationActiveDirectory` | ***[[]SchedulerJobSpecActionWebAuthenticationActiveDirectory](#schedulerjobspecactionwebauthenticationactivedirectory)***| ***(Optional)*** |
| `authenticationBasic` | ***[[]SchedulerJobSpecActionWebAuthenticationBasic](#schedulerjobspecactionwebauthenticationbasic)***| ***(Optional)*** |
| `authenticationCertificate` | ***[[]SchedulerJobSpecActionWebAuthenticationCertificate](#schedulerjobspecactionwebauthenticationcertificate)***| ***(Optional)*** |
| `body` | ***string***| ***(Optional)*** |
| `headers` | ***map[string]string***| ***(Optional)*** |
| `method` | ***string***||
| `url` | ***string***||
## SchedulerJobSpecActionWebAuthenticationActiveDirectory

Appears on:[SchedulerJobSpecActionWeb](#schedulerjobspecactionweb)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `audience` | ***string***| ***(Optional)*** |
| `clientID` | ***string***||
| `tenantID` | ***string***||
## SchedulerJobSpecActionWebAuthenticationBasic

Appears on:[SchedulerJobSpecActionWeb](#schedulerjobspecactionweb)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `username` | ***string***||
## SchedulerJobSpecActionWebAuthenticationCertificate

Appears on:[SchedulerJobSpecActionWeb](#schedulerjobspecactionweb)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `expiration` | ***string***| ***(Optional)*** |
| `subjectName` | ***string***| ***(Optional)*** |
| `thumbprint` | ***string***| ***(Optional)*** |
## SchedulerJobSpecErrorActionStorageQueue

Appears on:[SchedulerJobSpec](#schedulerjobspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `message` | ***string***||
| `sasToken` | ***string***||
| `storageAccountName` | ***string***||
| `storageQueueName` | ***string***||
## SchedulerJobSpecErrorActionWeb

Appears on:[SchedulerJobSpec](#schedulerjobspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `authenticationActiveDirectory` | ***[[]SchedulerJobSpecErrorActionWebAuthenticationActiveDirectory](#schedulerjobspecerroractionwebauthenticationactivedirectory)***| ***(Optional)*** |
| `authenticationBasic` | ***[[]SchedulerJobSpecErrorActionWebAuthenticationBasic](#schedulerjobspecerroractionwebauthenticationbasic)***| ***(Optional)*** |
| `authenticationCertificate` | ***[[]SchedulerJobSpecErrorActionWebAuthenticationCertificate](#schedulerjobspecerroractionwebauthenticationcertificate)***| ***(Optional)*** |
| `body` | ***string***| ***(Optional)*** |
| `headers` | ***map[string]string***| ***(Optional)*** |
| `method` | ***string***||
| `url` | ***string***||
## SchedulerJobSpecErrorActionWebAuthenticationActiveDirectory

Appears on:[SchedulerJobSpecErrorActionWeb](#schedulerjobspecerroractionweb)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `audience` | ***string***| ***(Optional)*** |
| `clientID` | ***string***||
| `tenantID` | ***string***||
## SchedulerJobSpecErrorActionWebAuthenticationBasic

Appears on:[SchedulerJobSpecErrorActionWeb](#schedulerjobspecerroractionweb)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `username` | ***string***||
## SchedulerJobSpecErrorActionWebAuthenticationCertificate

Appears on:[SchedulerJobSpecErrorActionWeb](#schedulerjobspecerroractionweb)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `expiration` | ***string***| ***(Optional)*** |
| `subjectName` | ***string***| ***(Optional)*** |
| `thumbprint` | ***string***| ***(Optional)*** |
## SchedulerJobSpecRecurrence

Appears on:[SchedulerJobSpec](#schedulerjobspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int***| ***(Optional)*** |
| `endTime` | ***string***| ***(Optional)*** |
| `frequency` | ***string***||
| `hours` | ***[]int64***| ***(Optional)*** |
| `interval` | ***int***| ***(Optional)*** |
| `minutes` | ***[]int64***| ***(Optional)*** |
| `monthDays` | ***[]int64***| ***(Optional)*** |
| `monthlyOccurrences` | ***[[]SchedulerJobSpecRecurrenceMonthlyOccurrences](#schedulerjobspecrecurrencemonthlyoccurrences)***| ***(Optional)*** |
| `weekDays` | ***[]string***| ***(Optional)*** |
## SchedulerJobSpecRecurrenceMonthlyOccurrences

Appears on:[SchedulerJobSpecRecurrence](#schedulerjobspecrecurrence)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `day` | ***string***||
| `occurrence` | ***int***||
## SchedulerJobSpecRetry

Appears on:[SchedulerJobSpec](#schedulerjobspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int***| ***(Optional)*** |
| `interval` | ***string***| ***(Optional)*** |
## SchedulerJobStatus

Appears on:[SchedulerJob](#schedulerjob)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SchedulerJobSpec](#schedulerjobspec)***| ***(Optional)*** |
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
