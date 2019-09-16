---
title: LambdaAlias
menu:
  docs_{{ .version }}:
    identifier: lambdaalias-aws.kubeform.com
    name: LambdaAlias
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LambdaAlias
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LambdaAlias` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LambdaAliasSpec](#LambdaAliasSpec)***||
| `status` | ***[LambdaAliasStatus](#LambdaAliasStatus)***||
## LambdaAliasSpec

Appears on:[LambdaAlias](#LambdaAlias), [LambdaAliasStatus](#LambdaAliasStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `functionName` | ***string***||
| `functionVersion` | ***string***||
| `invokeArn` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `routingConfig` | ***[[]LambdaAliasSpecRoutingConfig](#LambdaAliasSpecRoutingConfig)***| ***(Optional)*** |
## LambdaAliasSpecRoutingConfig

Appears on:[LambdaAliasSpec](#LambdaAliasSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `additionalVersionWeights` | ***map[string]encoding/json.Number***| ***(Optional)*** |
## LambdaAliasStatus

Appears on:[LambdaAlias](#LambdaAlias)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LambdaAliasSpec](#LambdaAliasSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
