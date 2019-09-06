## GlueCrawler
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `GlueCrawler` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[GlueCrawlerSpec](#GlueCrawlerSpec)***||
| `status` | ***[GlueCrawlerStatus](#GlueCrawlerStatus)***||
## GlueCrawlerSpec
##### (Appears on:[GlueCrawler](#GlueCrawler), [GlueCrawlerStatus](#GlueCrawlerStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `classifiers` | ***[]string***| ***(Optional)*** |
| `configuration` | ***string***| ***(Optional)*** |
| `databaseName` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `dynamodbTarget` | ***[[]GlueCrawlerSpecDynamodbTarget](#GlueCrawlerSpecDynamodbTarget)***| ***(Optional)*** |
| `jdbcTarget` | ***[[]GlueCrawlerSpecJdbcTarget](#GlueCrawlerSpecJdbcTarget)***| ***(Optional)*** |
| `name` | ***string***||
| `role` | ***string***||
| `s3Target` | ***[[]GlueCrawlerSpecS3Target](#GlueCrawlerSpecS3Target)***| ***(Optional)*** |
| `schedule` | ***string***| ***(Optional)*** |
| `schemaChangePolicy` | ***[[]GlueCrawlerSpecSchemaChangePolicy](#GlueCrawlerSpecSchemaChangePolicy)***| ***(Optional)*** |
| `securityConfiguration` | ***string***| ***(Optional)*** |
| `tablePrefix` | ***string***| ***(Optional)*** |
## GlueCrawlerSpecDynamodbTarget
##### (Appears on:[GlueCrawlerSpec](#GlueCrawlerSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `path` | ***string***||
## GlueCrawlerSpecJdbcTarget
##### (Appears on:[GlueCrawlerSpec](#GlueCrawlerSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `connectionName` | ***string***||
| `exclusions` | ***[]string***| ***(Optional)*** |
| `path` | ***string***||
## GlueCrawlerSpecS3Target
##### (Appears on:[GlueCrawlerSpec](#GlueCrawlerSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `exclusions` | ***[]string***| ***(Optional)*** |
| `path` | ***string***||
## GlueCrawlerSpecSchemaChangePolicy
##### (Appears on:[GlueCrawlerSpec](#GlueCrawlerSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `deleteBehavior` | ***string***| ***(Optional)*** |
| `updateBehavior` | ***string***| ***(Optional)*** |
## GlueCrawlerStatus
##### (Appears on:[GlueCrawler](#GlueCrawler))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[GlueCrawlerSpec](#GlueCrawlerSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
