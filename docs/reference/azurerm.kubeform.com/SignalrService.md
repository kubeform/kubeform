---
title: SignalrService
menu:
  docs_{{ .version }}:
    identifier: signalrservice-azurerm.kubeform.com
    name: SignalrService
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SignalrService
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `SignalrService` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SignalrServiceSpec](#signalrservicespec)***||
| `status` | ***[SignalrServiceStatus](#signalrservicestatus)***||
## Phase(`string` alias)

Appears on:[SignalrServiceStatus](#signalrservicestatus)

## SignalrServiceSpec

Appears on:[SignalrService](#signalrservice), [SignalrServiceStatus](#signalrservicestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `cors` | ***[[]SignalrServiceSpecCors](#signalrservicespeccors)***| ***(Optional)*** |
| `features` | ***[[]SignalrServiceSpecFeatures](#signalrservicespecfeatures)***| ***(Optional)*** |
| `hostname` | ***string***| ***(Optional)*** |
| `ipAddress` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `publicPort` | ***int64***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `serverPort` | ***int64***| ***(Optional)*** |
| `sku` | ***[[]SignalrServiceSpecSku](#signalrservicespecsku)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## SignalrServiceSpecCors

Appears on:[SignalrServiceSpec](#signalrservicespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `allowedOrigins` | ***[]string***||
## SignalrServiceSpecFeatures

Appears on:[SignalrServiceSpec](#signalrservicespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `flag` | ***string***||
| `value` | ***string***||
## SignalrServiceSpecSku

Appears on:[SignalrServiceSpec](#signalrservicespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `capacity` | ***int64***||
| `name` | ***string***||
## SignalrServiceStatus

Appears on:[SignalrService](#signalrservice)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SignalrServiceSpec](#signalrservicespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `primary_access_key` | ***string*** ||
| `primary_connection_string` | ***string*** ||
| `secondary_access_key` | ***string*** ||
| `secondary_connection_string` | ***string*** ||
