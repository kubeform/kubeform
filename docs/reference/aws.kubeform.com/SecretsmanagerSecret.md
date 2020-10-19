---
title: SecretsmanagerSecret
menu:
  docs_{{ .version }}:
    identifier: secretsmanagersecret-aws.kubeform.com
    name: SecretsmanagerSecret
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SecretsmanagerSecret
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SecretsmanagerSecret` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SecretsmanagerSecretSpec](#secretsmanagersecretspec)***||
| `status` | ***[SecretsmanagerSecretStatus](#secretsmanagersecretstatus)***||
## Phase(`string` alias)

Appears on:[SecretsmanagerSecretStatus](#secretsmanagersecretstatus)

## SecretsmanagerSecretSpec

Appears on:[SecretsmanagerSecret](#secretsmanagersecret), [SecretsmanagerSecretStatus](#secretsmanagersecretstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `kmsKeyID` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `policy` | ***string***| ***(Optional)*** |
| `recoveryWindowInDays` | ***int64***| ***(Optional)*** |
| `rotationEnabled` | ***bool***| ***(Optional)*** |
| `rotationLambdaArn` | ***string***| ***(Optional)*** |
| `rotationRules` | ***[[]SecretsmanagerSecretSpecRotationRules](#secretsmanagersecretspecrotationrules)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## SecretsmanagerSecretSpecRotationRules

Appears on:[SecretsmanagerSecretSpec](#secretsmanagersecretspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `automaticallyAfterDays` | ***int64***||
## SecretsmanagerSecretStatus

Appears on:[SecretsmanagerSecret](#secretsmanagersecret)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SecretsmanagerSecretSpec](#secretsmanagersecretspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
