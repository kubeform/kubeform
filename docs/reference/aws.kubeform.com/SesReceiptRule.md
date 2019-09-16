---
title: SesReceiptRule
menu:
  docs_{{ .version }}:
    identifier: sesreceiptrule-aws.kubeform.com
    name: SesReceiptRule
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SesReceiptRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SesReceiptRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SesReceiptRuleSpec](#sesreceiptrulespec)***||
| `status` | ***[SesReceiptRuleStatus](#sesreceiptrulestatus)***||
## SesReceiptRuleSpec

Appears on:[SesReceiptRule](#sesreceiptrule), [SesReceiptRuleStatus](#sesreceiptrulestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `addHeaderAction` | ***[[]SesReceiptRuleSpecAddHeaderAction](#sesreceiptrulespecaddheaderaction)***| ***(Optional)*** |
| `after` | ***string***| ***(Optional)*** |
| `bounceAction` | ***[[]SesReceiptRuleSpecBounceAction](#sesreceiptrulespecbounceaction)***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `lambdaAction` | ***[[]SesReceiptRuleSpecLambdaAction](#sesreceiptrulespeclambdaaction)***| ***(Optional)*** |
| `name` | ***string***||
| `recipients` | ***[]string***| ***(Optional)*** |
| `ruleSetName` | ***string***||
| `s3Action` | ***[[]SesReceiptRuleSpecS3Action](#sesreceiptrulespecs3action)***| ***(Optional)*** |
| `scanEnabled` | ***bool***| ***(Optional)*** |
| `snsAction` | ***[[]SesReceiptRuleSpecSnsAction](#sesreceiptrulespecsnsaction)***| ***(Optional)*** |
| `stopAction` | ***[[]SesReceiptRuleSpecStopAction](#sesreceiptrulespecstopaction)***| ***(Optional)*** |
| `tlsPolicy` | ***string***| ***(Optional)*** |
| `workmailAction` | ***[[]SesReceiptRuleSpecWorkmailAction](#sesreceiptrulespecworkmailaction)***| ***(Optional)*** |
## SesReceiptRuleSpecAddHeaderAction

Appears on:[SesReceiptRuleSpec](#sesreceiptrulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `headerName` | ***string***||
| `headerValue` | ***string***||
| `position` | ***int***||
## SesReceiptRuleSpecBounceAction

Appears on:[SesReceiptRuleSpec](#sesreceiptrulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `message` | ***string***||
| `position` | ***int***||
| `sender` | ***string***||
| `smtpReplyCode` | ***string***||
| `statusCode` | ***string***| ***(Optional)*** |
| `topicArn` | ***string***| ***(Optional)*** |
## SesReceiptRuleSpecLambdaAction

Appears on:[SesReceiptRuleSpec](#sesreceiptrulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `functionArn` | ***string***||
| `invocationType` | ***string***| ***(Optional)*** |
| `position` | ***int***||
| `topicArn` | ***string***| ***(Optional)*** |
## SesReceiptRuleSpecS3Action

Appears on:[SesReceiptRuleSpec](#sesreceiptrulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucketName` | ***string***||
| `kmsKeyArn` | ***string***| ***(Optional)*** |
| `objectKeyPrefix` | ***string***| ***(Optional)*** |
| `position` | ***int***||
| `topicArn` | ***string***| ***(Optional)*** |
## SesReceiptRuleSpecSnsAction

Appears on:[SesReceiptRuleSpec](#sesreceiptrulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `position` | ***int***||
| `topicArn` | ***string***||
## SesReceiptRuleSpecStopAction

Appears on:[SesReceiptRuleSpec](#sesreceiptrulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `position` | ***int***||
| `scope` | ***string***||
| `topicArn` | ***string***| ***(Optional)*** |
## SesReceiptRuleSpecWorkmailAction

Appears on:[SesReceiptRuleSpec](#sesreceiptrulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `organizationArn` | ***string***||
| `position` | ***int***||
| `topicArn` | ***string***| ***(Optional)*** |
## SesReceiptRuleStatus

Appears on:[SesReceiptRule](#sesreceiptrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SesReceiptRuleSpec](#sesreceiptrulespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
