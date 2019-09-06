## SesReceiptRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SesReceiptRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SesReceiptRuleSpec](#SesReceiptRuleSpec)***||
| `status` | ***[SesReceiptRuleStatus](#SesReceiptRuleStatus)***||
## SesReceiptRuleSpec
##### (Appears on:[SesReceiptRule](#SesReceiptRule), [SesReceiptRuleStatus](#SesReceiptRuleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `addHeaderAction` | ***[[]SesReceiptRuleSpecAddHeaderAction](#SesReceiptRuleSpecAddHeaderAction)***| ***(Optional)*** |
| `after` | ***string***| ***(Optional)*** |
| `bounceAction` | ***[[]SesReceiptRuleSpecBounceAction](#SesReceiptRuleSpecBounceAction)***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `lambdaAction` | ***[[]SesReceiptRuleSpecLambdaAction](#SesReceiptRuleSpecLambdaAction)***| ***(Optional)*** |
| `name` | ***string***||
| `recipients` | ***[]string***| ***(Optional)*** |
| `ruleSetName` | ***string***||
| `s3Action` | ***[[]SesReceiptRuleSpecS3Action](#SesReceiptRuleSpecS3Action)***| ***(Optional)*** |
| `scanEnabled` | ***bool***| ***(Optional)*** |
| `snsAction` | ***[[]SesReceiptRuleSpecSnsAction](#SesReceiptRuleSpecSnsAction)***| ***(Optional)*** |
| `stopAction` | ***[[]SesReceiptRuleSpecStopAction](#SesReceiptRuleSpecStopAction)***| ***(Optional)*** |
| `tlsPolicy` | ***string***| ***(Optional)*** |
| `workmailAction` | ***[[]SesReceiptRuleSpecWorkmailAction](#SesReceiptRuleSpecWorkmailAction)***| ***(Optional)*** |
## SesReceiptRuleSpecAddHeaderAction
##### (Appears on:[SesReceiptRuleSpec](#SesReceiptRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `headerName` | ***string***||
| `headerValue` | ***string***||
| `position` | ***int***||
## SesReceiptRuleSpecBounceAction
##### (Appears on:[SesReceiptRuleSpec](#SesReceiptRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `message` | ***string***||
| `position` | ***int***||
| `sender` | ***string***||
| `smtpReplyCode` | ***string***||
| `statusCode` | ***string***| ***(Optional)*** |
| `topicArn` | ***string***| ***(Optional)*** |
## SesReceiptRuleSpecLambdaAction
##### (Appears on:[SesReceiptRuleSpec](#SesReceiptRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `functionArn` | ***string***||
| `invocationType` | ***string***| ***(Optional)*** |
| `position` | ***int***||
| `topicArn` | ***string***| ***(Optional)*** |
## SesReceiptRuleSpecS3Action
##### (Appears on:[SesReceiptRuleSpec](#SesReceiptRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucketName` | ***string***||
| `kmsKeyArn` | ***string***| ***(Optional)*** |
| `objectKeyPrefix` | ***string***| ***(Optional)*** |
| `position` | ***int***||
| `topicArn` | ***string***| ***(Optional)*** |
## SesReceiptRuleSpecSnsAction
##### (Appears on:[SesReceiptRuleSpec](#SesReceiptRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `position` | ***int***||
| `topicArn` | ***string***||
## SesReceiptRuleSpecStopAction
##### (Appears on:[SesReceiptRuleSpec](#SesReceiptRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `position` | ***int***||
| `scope` | ***string***||
| `topicArn` | ***string***| ***(Optional)*** |
## SesReceiptRuleSpecWorkmailAction
##### (Appears on:[SesReceiptRuleSpec](#SesReceiptRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `organizationArn` | ***string***||
| `position` | ***int***||
| `topicArn` | ***string***| ***(Optional)*** |
## SesReceiptRuleStatus
##### (Appears on:[SesReceiptRule](#SesReceiptRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SesReceiptRuleSpec](#SesReceiptRuleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
