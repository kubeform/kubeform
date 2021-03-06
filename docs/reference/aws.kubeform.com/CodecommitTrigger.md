---
title: CodecommitTrigger
menu:
  docs_{{ .version }}:
    identifier: codecommittrigger-aws.kubeform.com
    name: CodecommitTrigger
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CodecommitTrigger
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CodecommitTrigger` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CodecommitTriggerSpec](#codecommittriggerspec)***||
| `status` | ***[CodecommitTriggerStatus](#codecommittriggerstatus)***||
## CodecommitTriggerSpec

Appears on:[CodecommitTrigger](#codecommittrigger), [CodecommitTriggerStatus](#codecommittriggerstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `configurationID` | ***string***| ***(Optional)*** |
| `repositoryName` | ***string***||
| `trigger` | ***[[]CodecommitTriggerSpecTrigger](#codecommittriggerspectrigger)***||
## CodecommitTriggerSpecTrigger

Appears on:[CodecommitTriggerSpec](#codecommittriggerspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `branches` | ***[]string***| ***(Optional)*** |
| `customData` | ***string***| ***(Optional)*** |
| `destinationArn` | ***string***||
| `events` | ***[]string***||
| `name` | ***string***||
## CodecommitTriggerStatus

Appears on:[CodecommitTrigger](#codecommittrigger)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CodecommitTriggerSpec](#codecommittriggerspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[CodecommitTriggerStatus](#codecommittriggerstatus)

---
