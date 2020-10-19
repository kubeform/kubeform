---
title: AppConfiguration
menu:
  docs_{{ .version }}:
    identifier: appconfiguration-azurerm.kubeform.com
    name: AppConfiguration
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AppConfiguration
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `AppConfiguration` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AppConfigurationSpec](#appconfigurationspec)***||
| `status` | ***[AppConfigurationStatus](#appconfigurationstatus)***||
## AppConfigurationSpec

Appears on:[AppConfiguration](#appconfiguration), [AppConfigurationStatus](#appconfigurationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `endpoint` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `primaryReadKey` | ***[[]AppConfigurationSpecPrimaryReadKey](#appconfigurationspecprimaryreadkey)***| ***(Optional)*** |
| `primaryWriteKey` | ***[[]AppConfigurationSpecPrimaryWriteKey](#appconfigurationspecprimarywritekey)***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `secondaryReadKey` | ***[[]AppConfigurationSpecSecondaryReadKey](#appconfigurationspecsecondaryreadkey)***| ***(Optional)*** |
| `secondaryWriteKey` | ***[[]AppConfigurationSpecSecondaryWriteKey](#appconfigurationspecsecondarywritekey)***| ***(Optional)*** |
| `sku` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## AppConfigurationSpecPrimaryReadKey

Appears on:[AppConfigurationSpec](#appconfigurationspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
## AppConfigurationSpecPrimaryWriteKey

Appears on:[AppConfigurationSpec](#appconfigurationspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
## AppConfigurationSpecSecondaryReadKey

Appears on:[AppConfigurationSpec](#appconfigurationspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
## AppConfigurationSpecSecondaryWriteKey

Appears on:[AppConfigurationSpec](#appconfigurationspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
## AppConfigurationStatus

Appears on:[AppConfiguration](#appconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AppConfigurationSpec](#appconfigurationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AppConfigurationStatus](#appconfigurationstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `primary_read_key.<index>.connection_string` | ***string*** ||
| `primary_read_key.<index>.id` | ***string*** ||
| `primary_read_key.<index>.secret` | ***string*** ||
| `primary_write_key.<index>.connection_string` | ***string*** ||
| `primary_write_key.<index>.id` | ***string*** ||
| `primary_write_key.<index>.secret` | ***string*** ||
| `secondary_read_key.<index>.connection_string` | ***string*** ||
| `secondary_read_key.<index>.id` | ***string*** ||
| `secondary_read_key.<index>.secret` | ***string*** ||
| `secondary_write_key.<index>.connection_string` | ***string*** ||
| `secondary_write_key.<index>.id` | ***string*** ||
| `secondary_write_key.<index>.secret` | ***string*** ||
