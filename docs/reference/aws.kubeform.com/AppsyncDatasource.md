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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AppsyncDatasourceSpec](#AppsyncDatasourceSpec)***||
| `status` | ***[AppsyncDatasourceStatus](#AppsyncDatasourceStatus)***||
## AppsyncDatasourceSpec

Appears on:[AppsyncDatasource](#AppsyncDatasource), [AppsyncDatasourceStatus](#AppsyncDatasourceStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `apiID` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `dynamodbConfig` | ***[[]AppsyncDatasourceSpecDynamodbConfig](#AppsyncDatasourceSpecDynamodbConfig)***| ***(Optional)*** |
| `elasticsearchConfig` | ***[[]AppsyncDatasourceSpecElasticsearchConfig](#AppsyncDatasourceSpecElasticsearchConfig)***| ***(Optional)*** |
| `httpConfig` | ***[[]AppsyncDatasourceSpecHttpConfig](#AppsyncDatasourceSpecHttpConfig)***| ***(Optional)*** |
| `lambdaConfig` | ***[[]AppsyncDatasourceSpecLambdaConfig](#AppsyncDatasourceSpecLambdaConfig)***| ***(Optional)*** |
| `name` | ***string***||
| `serviceRoleArn` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## AppsyncDatasourceSpecDynamodbConfig

Appears on:[AppsyncDatasourceSpec](#AppsyncDatasourceSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `region` | ***string***| ***(Optional)*** |
| `tableName` | ***string***||
| `useCallerCredentials` | ***bool***| ***(Optional)*** |
## AppsyncDatasourceSpecElasticsearchConfig

Appears on:[AppsyncDatasourceSpec](#AppsyncDatasourceSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `endpoint` | ***string***||
| `region` | ***string***| ***(Optional)*** |
## AppsyncDatasourceSpecHttpConfig

Appears on:[AppsyncDatasourceSpec](#AppsyncDatasourceSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `endpoint` | ***string***||
## AppsyncDatasourceSpecLambdaConfig

Appears on:[AppsyncDatasourceSpec](#AppsyncDatasourceSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `functionArn` | ***string***||
## AppsyncDatasourceStatus

Appears on:[AppsyncDatasource](#AppsyncDatasource)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AppsyncDatasourceSpec](#AppsyncDatasourceSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
