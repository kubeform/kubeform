---
title: LogicAppActionHTTP
menu:
  docs_{{ .version }}:
    identifier: logicappactionhttp-azurerm.kubeform.com
    name: LogicAppActionHTTP
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LogicAppActionHTTP
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `LogicAppActionHTTP` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LogicAppActionHTTPSpec](#logicappactionhttpspec)***||
| `status` | ***[LogicAppActionHTTPStatus](#logicappactionhttpstatus)***||
## LogicAppActionHTTPSpec

Appears on:[LogicAppActionHTTP](#logicappactionhttp), [LogicAppActionHTTPStatus](#logicappactionhttpstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `body` | ***string***| ***(Optional)*** |
| `headers` | ***map[string]string***| ***(Optional)*** |
| `logicAppID` | ***string***||
| `method` | ***string***||
| `name` | ***string***||
| `uri` | ***string***||
## LogicAppActionHTTPStatus

Appears on:[LogicAppActionHTTP](#logicappactionhttp)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LogicAppActionHTTPSpec](#logicappactionhttpspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[LogicAppActionHTTPStatus](#logicappactionhttpstatus)

---
