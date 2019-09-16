---
title: GlacierVault
menu:
  docs_{{ .version }}:
    identifier: glaciervault-aws.kubeform.com
    name: GlacierVault
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## GlacierVault
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `GlacierVault` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[GlacierVaultSpec](#GlacierVaultSpec)***||
| `status` | ***[GlacierVaultStatus](#GlacierVaultStatus)***||
## GlacierVaultSpec

Appears on:[GlacierVault](#GlacierVault), [GlacierVaultStatus](#GlacierVaultStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accessPolicy` | ***string***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `location` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `notification` | ***[[]GlacierVaultSpecNotification](#GlacierVaultSpecNotification)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## GlacierVaultSpecNotification

Appears on:[GlacierVaultSpec](#GlacierVaultSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `events` | ***[]string***||
| `snsTopic` | ***string***||
## GlacierVaultStatus

Appears on:[GlacierVault](#GlacierVault)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[GlacierVaultSpec](#GlacierVaultSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
