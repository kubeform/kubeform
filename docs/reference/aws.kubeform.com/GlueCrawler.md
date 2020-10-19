---
title: GlueCrawler
menu:
  docs_{{ .version }}:
    identifier: gluecrawler-aws.kubeform.com
    name: GlueCrawler
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## GlueCrawler
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `GlueCrawler` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[GlueCrawlerSpec](#gluecrawlerspec)***||
| `status` | ***[GlueCrawlerStatus](#gluecrawlerstatus)***||
## GlueCrawlerSpec

Appears on:[GlueCrawler](#gluecrawler), [GlueCrawlerStatus](#gluecrawlerstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `catalogTarget` | ***[[]GlueCrawlerSpecCatalogTarget](#gluecrawlerspeccatalogtarget)***| ***(Optional)*** |
| `classifiers` | ***[]string***| ***(Optional)*** |
| `configuration` | ***string***| ***(Optional)*** |
| `databaseName` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `dynamodbTarget` | ***[[]GlueCrawlerSpecDynamodbTarget](#gluecrawlerspecdynamodbtarget)***| ***(Optional)*** |
| `jdbcTarget` | ***[[]GlueCrawlerSpecJdbcTarget](#gluecrawlerspecjdbctarget)***| ***(Optional)*** |
| `name` | ***string***||
| `role` | ***string***||
| `s3Target` | ***[[]GlueCrawlerSpecS3Target](#gluecrawlerspecs3target)***| ***(Optional)*** |
| `schedule` | ***string***| ***(Optional)*** |
| `schemaChangePolicy` | ***[[]GlueCrawlerSpecSchemaChangePolicy](#gluecrawlerspecschemachangepolicy)***| ***(Optional)*** |
| `securityConfiguration` | ***string***| ***(Optional)*** |
| `tablePrefix` | ***string***| ***(Optional)*** |
## GlueCrawlerSpecCatalogTarget

Appears on:[GlueCrawlerSpec](#gluecrawlerspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `databaseName` | ***string***||
| `tables` | ***[]string***||
## GlueCrawlerSpecDynamodbTarget

Appears on:[GlueCrawlerSpec](#gluecrawlerspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `path` | ***string***||
## GlueCrawlerSpecJdbcTarget

Appears on:[GlueCrawlerSpec](#gluecrawlerspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `connectionName` | ***string***||
| `exclusions` | ***[]string***| ***(Optional)*** |
| `path` | ***string***||
## GlueCrawlerSpecS3Target

Appears on:[GlueCrawlerSpec](#gluecrawlerspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `exclusions` | ***[]string***| ***(Optional)*** |
| `path` | ***string***||
## GlueCrawlerSpecSchemaChangePolicy

Appears on:[GlueCrawlerSpec](#gluecrawlerspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `deleteBehavior` | ***string***| ***(Optional)*** |
| `updateBehavior` | ***string***| ***(Optional)*** |
## GlueCrawlerStatus

Appears on:[GlueCrawler](#gluecrawler)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[GlueCrawlerSpec](#gluecrawlerspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[GlueCrawlerStatus](#gluecrawlerstatus)

---
