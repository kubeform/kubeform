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
| `spec` | ***[ConfigConfigurationAggregatorSpec](#ConfigConfigurationAggregatorSpec)***||
| `status` | ***[ConfigConfigurationAggregatorStatus](#ConfigConfigurationAggregatorStatus)***||
## ConfigConfigurationAggregatorSpec
##### (Appears on:[ConfigConfigurationAggregator](#ConfigConfigurationAggregator), [ConfigConfigurationAggregatorStatus](#ConfigConfigurationAggregatorStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountAggregationSource` | ***[[]ConfigConfigurationAggregatorSpecAccountAggregationSource](#ConfigConfigurationAggregatorSpecAccountAggregationSource)***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `organizationAggregationSource` | ***[[]ConfigConfigurationAggregatorSpecOrganizationAggregationSource](#ConfigConfigurationAggregatorSpecOrganizationAggregationSource)***| ***(Optional)*** |
## ConfigConfigurationAggregatorSpecAccountAggregationSource
##### (Appears on:[ConfigConfigurationAggregatorSpec](#ConfigConfigurationAggregatorSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `accountIDS` | ***[]string***||
| `allRegions` | ***bool***| ***(Optional)*** |
| `regions` | ***[]string***| ***(Optional)*** |
## ConfigConfigurationAggregatorSpecOrganizationAggregationSource
##### (Appears on:[ConfigConfigurationAggregatorSpec](#ConfigConfigurationAggregatorSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `allRegions` | ***bool***| ***(Optional)*** |
| `regions` | ***[]string***| ***(Optional)*** |
| `roleArn` | ***string***||
## ConfigConfigurationAggregatorStatus
##### (Appears on:[ConfigConfigurationAggregator](#ConfigConfigurationAggregator))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ConfigConfigurationAggregatorSpec](#ConfigConfigurationAggregatorSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
