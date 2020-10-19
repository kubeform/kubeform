---
title: StorageTransferJob
menu:
  docs_{{ .version }}:
    identifier: storagetransferjob-google.kubeform.com
    name: StorageTransferJob
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StorageTransferJob
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `StorageTransferJob` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StorageTransferJobSpec](#storagetransferjobspec)***||
| `status` | ***[StorageTransferJobStatus](#storagetransferjobstatus)***||
## Phase(`string` alias)

Appears on:[StorageTransferJobStatus](#storagetransferjobstatus)

## StorageTransferJobSpec

Appears on:[StorageTransferJob](#storagetransferjob), [StorageTransferJobStatus](#storagetransferjobstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `creationTime` | ***string***| ***(Optional)*** |
| `deletionTime` | ***string***| ***(Optional)*** |
| `description` | ***string***||
| `lastModificationTime` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `schedule` | ***[[]StorageTransferJobSpecSchedule](#storagetransferjobspecschedule)***||
| `status` | ***string***| ***(Optional)*** |
| `transferSpec` | ***[[]StorageTransferJobSpecTransferSpec](#storagetransferjobspectransferspec)***||
## StorageTransferJobSpecSchedule

Appears on:[StorageTransferJobSpec](#storagetransferjobspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `scheduleEndDate` | ***[[]StorageTransferJobSpecScheduleScheduleEndDate](#storagetransferjobspecschedulescheduleenddate)***| ***(Optional)*** |
| `scheduleStartDate` | ***[[]StorageTransferJobSpecScheduleScheduleStartDate](#storagetransferjobspecscheduleschedulestartdate)***||
| `startTimeOfDay` | ***[[]StorageTransferJobSpecScheduleStartTimeOfDay](#storagetransferjobspecschedulestarttimeofday)***| ***(Optional)*** |
## StorageTransferJobSpecScheduleScheduleEndDate

Appears on:[StorageTransferJobSpecSchedule](#storagetransferjobspecschedule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `day` | ***int64***||
| `month` | ***int64***||
| `year` | ***int64***||
## StorageTransferJobSpecScheduleScheduleStartDate

Appears on:[StorageTransferJobSpecSchedule](#storagetransferjobspecschedule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `day` | ***int64***||
| `month` | ***int64***||
| `year` | ***int64***||
## StorageTransferJobSpecScheduleStartTimeOfDay

Appears on:[StorageTransferJobSpecSchedule](#storagetransferjobspecschedule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `hours` | ***int64***||
| `minutes` | ***int64***||
| `nanos` | ***int64***||
| `seconds` | ***int64***||
## StorageTransferJobSpecTransferSpec

Appears on:[StorageTransferJobSpec](#storagetransferjobspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `awsS3DataSource` | ***[[]StorageTransferJobSpecTransferSpecAwsS3DataSource](#storagetransferjobspectransferspecawss3datasource)***| ***(Optional)*** |
| `gcsDataSink` | ***[[]StorageTransferJobSpecTransferSpecGcsDataSink](#storagetransferjobspectransferspecgcsdatasink)***| ***(Optional)*** |
| `gcsDataSource` | ***[[]StorageTransferJobSpecTransferSpecGcsDataSource](#storagetransferjobspectransferspecgcsdatasource)***| ***(Optional)*** |
| `httpDataSource` | ***[[]StorageTransferJobSpecTransferSpecHttpDataSource](#storagetransferjobspectransferspechttpdatasource)***| ***(Optional)*** |
| `objectConditions` | ***[[]StorageTransferJobSpecTransferSpecObjectConditions](#storagetransferjobspectransferspecobjectconditions)***| ***(Optional)*** |
| `transferOptions` | ***[[]StorageTransferJobSpecTransferSpecTransferOptions](#storagetransferjobspectransferspectransferoptions)***| ***(Optional)*** |
## StorageTransferJobSpecTransferSpecAwsS3DataSource

Appears on:[StorageTransferJobSpecTransferSpec](#storagetransferjobspectransferspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `awsAccessKey` | ***[[]StorageTransferJobSpecTransferSpecAwsS3DataSourceAwsAccessKey](#storagetransferjobspectransferspecawss3datasourceawsaccesskey)***||
| `bucketName` | ***string***||
## StorageTransferJobSpecTransferSpecAwsS3DataSourceAwsAccessKey

Appears on:[StorageTransferJobSpecTransferSpecAwsS3DataSource](#storagetransferjobspectransferspecawss3datasource)

| Field | Type | Description |
| ------ | ----- | ----------- |
## StorageTransferJobSpecTransferSpecGcsDataSink

Appears on:[StorageTransferJobSpecTransferSpec](#storagetransferjobspectransferspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucketName` | ***string***||
## StorageTransferJobSpecTransferSpecGcsDataSource

Appears on:[StorageTransferJobSpecTransferSpec](#storagetransferjobspectransferspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucketName` | ***string***||
## StorageTransferJobSpecTransferSpecHttpDataSource

Appears on:[StorageTransferJobSpecTransferSpec](#storagetransferjobspectransferspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `listURL` | ***string***||
## StorageTransferJobSpecTransferSpecObjectConditions

Appears on:[StorageTransferJobSpecTransferSpec](#storagetransferjobspectransferspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `excludePrefixes` | ***[]string***| ***(Optional)*** |
| `includePrefixes` | ***[]string***| ***(Optional)*** |
| `maxTimeElapsedSinceLastModification` | ***string***| ***(Optional)*** |
| `minTimeElapsedSinceLastModification` | ***string***| ***(Optional)*** |
## StorageTransferJobSpecTransferSpecTransferOptions

Appears on:[StorageTransferJobSpecTransferSpec](#storagetransferjobspectransferspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `deleteObjectsFromSourceAfterTransfer` | ***bool***| ***(Optional)*** |
| `deleteObjectsUniqueInSink` | ***bool***| ***(Optional)*** |
| `overwriteObjectsAlreadyExistingInSink` | ***bool***| ***(Optional)*** |
## StorageTransferJobStatus

Appears on:[StorageTransferJob](#storagetransferjob)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StorageTransferJobSpec](#storagetransferjobspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `transfer_spec.<index>.aws_s3_data_source.<index>.aws_access_key.<index>.access_key_id` | ***string*** ||
| `transfer_spec.<index>.aws_s3_data_source.<index>.aws_access_key.<index>.secret_access_key` | ***string*** ||
