---
title: EventhubNamespace 
menu:
  docs_{{ .version }}:
    identifier: eventhubnamespace--azurerm.kubeform.com
    name: EventhubNamespace 
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## EventhubNamespace_
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `EventhubNamespace_` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[EventhubNamespace_Spec](#eventhubnamespace_spec)***||
| `status` | ***[EventhubNamespace_Status](#eventhubnamespace_status)***||
## EventhubNamespace_Spec

Appears on:[EventhubNamespace_](#eventhubnamespace_), [EventhubNamespace_Status](#eventhubnamespace_status)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `autoInflateEnabled` | ***bool***| ***(Optional)*** |
| `capacity` | ***int64***| ***(Optional)*** |
| `kafkaEnabled` | ***bool***| ***(Optional)*** Deprecated|
| `location` | ***string***||
| `maximumThroughputUnits` | ***int64***| ***(Optional)*** |
| `name` | ***string***||
| `networkRulesets` | ***[[]EventhubNamespace_SpecNetworkRulesets](#eventhubnamespace_specnetworkrulesets)***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `sku` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## EventhubNamespace_SpecNetworkRulesets

Appears on:[EventhubNamespace_Spec](#eventhubnamespace_spec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `defaultAction` | ***string***||
| `ipRule` | ***[[]EventhubNamespace_SpecNetworkRulesetsIpRule](#eventhubnamespace_specnetworkrulesetsiprule)***| ***(Optional)*** |
| `virtualNetworkRule` | ***[[]EventhubNamespace_SpecNetworkRulesetsVirtualNetworkRule](#eventhubnamespace_specnetworkrulesetsvirtualnetworkrule)***| ***(Optional)*** |
## EventhubNamespace_SpecNetworkRulesetsIpRule

Appears on:[EventhubNamespace_SpecNetworkRulesets](#eventhubnamespace_specnetworkrulesets)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***string***| ***(Optional)*** |
| `ipMask` | ***string***||
## EventhubNamespace_SpecNetworkRulesetsVirtualNetworkRule

Appears on:[EventhubNamespace_SpecNetworkRulesets](#eventhubnamespace_specnetworkrulesets)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ignoreMissingVirtualNetworkServiceEndpoint` | ***bool***| ***(Optional)*** |
| `subnetID` | ***string***||
## EventhubNamespace_Status

Appears on:[EventhubNamespace_](#eventhubnamespace_)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[EventhubNamespace_Spec](#eventhubnamespace_spec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[EventhubNamespace_Status](#eventhubnamespace_status)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `default_primary_connection_string` | ***string*** ||
| `default_primary_key` | ***string*** ||
| `default_secondary_connection_string` | ***string*** ||
| `default_secondary_key` | ***string*** ||
