---
title: LogicAppTriggerHTTPRequest
menu:
  docs_{{ .version }}:
    identifier: logicapptriggerhttprequest-azurerm.kubeform.com
    name: LogicAppTriggerHTTPRequest
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LogicAppTriggerHTTPRequest
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `LogicAppTriggerHTTPRequest` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LogicAppTriggerHTTPRequestSpec](#logicapptriggerhttprequestspec)***||
| `status` | ***[LogicAppTriggerHTTPRequestStatus](#logicapptriggerhttprequeststatus)***||
## LogicAppTriggerHTTPRequestSpec

Appears on:[LogicAppTriggerHTTPRequest](#logicapptriggerhttprequest), [LogicAppTriggerHTTPRequestStatus](#logicapptriggerhttprequeststatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `logicAppID` | ***string***||
| `method` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `relativePath` | ***string***| ***(Optional)*** |
| `schema` | ***string***||
## LogicAppTriggerHTTPRequestStatus

Appears on:[LogicAppTriggerHTTPRequest](#logicapptriggerhttprequest)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LogicAppTriggerHTTPRequestSpec](#logicapptriggerhttprequestspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[LogicAppTriggerHTTPRequestStatus](#logicapptriggerhttprequeststatus)

---