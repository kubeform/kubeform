---
title: EventgridDomain
menu:
  docs_{{ .version }}:
    identifier: eventgriddomain-azurerm.kubeform.com
    name: EventgridDomain
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## EventgridDomain
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `EventgridDomain` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[EventgridDomainSpec](#eventgriddomainspec)***||
| `status` | ***[EventgridDomainStatus](#eventgriddomainstatus)***||
## EventgridDomainSpec

Appears on:[EventgridDomain](#eventgriddomain), [EventgridDomainStatus](#eventgriddomainstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `endpoint` | ***string***| ***(Optional)*** |
| `inputMappingDefaultValues` | ***[[]EventgridDomainSpecInputMappingDefaultValues](#eventgriddomainspecinputmappingdefaultvalues)***| ***(Optional)*** |
| `inputMappingFields` | ***[[]EventgridDomainSpecInputMappingFields](#eventgriddomainspecinputmappingfields)***| ***(Optional)*** |
| `inputSchema` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## EventgridDomainSpecInputMappingDefaultValues

Appears on:[EventgridDomainSpec](#eventgriddomainspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `dataVersion` | ***string***| ***(Optional)*** |
| `eventType` | ***string***| ***(Optional)*** |
| `subject` | ***string***| ***(Optional)*** |
## EventgridDomainSpecInputMappingFields

Appears on:[EventgridDomainSpec](#eventgriddomainspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `dataVersion` | ***string***| ***(Optional)*** |
| `eventTime` | ***string***| ***(Optional)*** |
| `eventType` | ***string***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `subject` | ***string***| ***(Optional)*** |
| `topic` | ***string***| ***(Optional)*** |
## EventgridDomainStatus

Appears on:[EventgridDomain](#eventgriddomain)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[EventgridDomainSpec](#eventgriddomainspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
