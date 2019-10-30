---
title: ConfigConfigurationAggregator
menu:
  docs_{{ .version }}:
    identifier: configconfigurationaggregator-aws.kubeform.com
    name: ConfigConfigurationAggregator
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ConfigConfigurationAggregator
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ConfigConfigurationAggregator` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ConfigConfigurationAggregatorSpec](#configconfigurationaggregatorspec)***||
| `status` | ***[ConfigConfigurationAggregatorStatus](#configconfigurationaggregatorstatus)***||
## ConfigConfigurationAggregatorSpec

Appears on:[ConfigConfigurationAggregator](#configconfigurationaggregator), [ConfigConfigurationAggregatorStatus](#configconfigurationaggregatorstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountAggregationSource` | ***[[]ConfigConfigurationAggregatorSpecAccountAggregationSource](#configconfigurationaggregatorspecaccountaggregationsource)***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `organizationAggregationSource` | ***[[]ConfigConfigurationAggregatorSpecOrganizationAggregationSource](#configconfigurationaggregatorspecorganizationaggregationsource)***| ***(Optional)*** |
## ConfigConfigurationAggregatorSpecAccountAggregationSource

Appears on:[ConfigConfigurationAggregatorSpec](#configconfigurationaggregatorspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `accountIDS` | ***[]string***||
| `allRegions` | ***bool***| ***(Optional)*** |
| `regions` | ***[]string***| ***(Optional)*** |
## ConfigConfigurationAggregatorSpecOrganizationAggregationSource

Appears on:[ConfigConfigurationAggregatorSpec](#configconfigurationaggregatorspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `allRegions` | ***bool***| ***(Optional)*** |
| `regions` | ***[]string***| ***(Optional)*** |
| `roleArn` | ***string***||
## ConfigConfigurationAggregatorStatus

Appears on:[ConfigConfigurationAggregator](#configconfigurationaggregator)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ConfigConfigurationAggregatorSpec](#configconfigurationaggregatorspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ConfigConfigurationAggregatorStatus](#configconfigurationaggregatorstatus)

---
