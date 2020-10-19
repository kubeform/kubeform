---
title: AppsyncDatasource
menu:
  docs_{{ .version }}:
    identifier: appsyncdatasource-aws.kubeform.com
    name: AppsyncDatasource
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AppsyncDatasource
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AppsyncDatasource` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AppsyncDatasourceSpec](#appsyncdatasourcespec)***||
| `status` | ***[AppsyncDatasourceStatus](#appsyncdatasourcestatus)***||
## AppsyncDatasourceSpec

Appears on:[AppsyncDatasource](#appsyncdatasource), [AppsyncDatasourceStatus](#appsyncdatasourcestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `apiID` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `dynamodbConfig` | ***[[]AppsyncDatasourceSpecDynamodbConfig](#appsyncdatasourcespecdynamodbconfig)***| ***(Optional)*** |
| `elasticsearchConfig` | ***[[]AppsyncDatasourceSpecElasticsearchConfig](#appsyncdatasourcespecelasticsearchconfig)***| ***(Optional)*** |
| `httpConfig` | ***[[]AppsyncDatasourceSpecHttpConfig](#appsyncdatasourcespechttpconfig)***| ***(Optional)*** |
| `lambdaConfig` | ***[[]AppsyncDatasourceSpecLambdaConfig](#appsyncdatasourcespeclambdaconfig)***| ***(Optional)*** |
| `name` | ***string***||
| `serviceRoleArn` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## AppsyncDatasourceSpecDynamodbConfig

Appears on:[AppsyncDatasourceSpec](#appsyncdatasourcespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `region` | ***string***| ***(Optional)*** |
| `tableName` | ***string***||
| `useCallerCredentials` | ***bool***| ***(Optional)*** |
## AppsyncDatasourceSpecElasticsearchConfig

Appears on:[AppsyncDatasourceSpec](#appsyncdatasourcespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `endpoint` | ***string***||
| `region` | ***string***| ***(Optional)*** |
## AppsyncDatasourceSpecHttpConfig

Appears on:[AppsyncDatasourceSpec](#appsyncdatasourcespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `endpoint` | ***string***||
## AppsyncDatasourceSpecLambdaConfig

Appears on:[AppsyncDatasourceSpec](#appsyncdatasourcespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `functionArn` | ***string***||
## AppsyncDatasourceStatus

Appears on:[AppsyncDatasource](#appsyncdatasource)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AppsyncDatasourceSpec](#appsyncdatasourcespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AppsyncDatasourceStatus](#appsyncdatasourcestatus)

---
