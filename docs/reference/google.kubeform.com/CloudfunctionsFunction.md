---
title: CloudfunctionsFunction
menu:
  docs_{{ .version }}:
    identifier: cloudfunctionsfunction-google.kubeform.com
    name: CloudfunctionsFunction
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CloudfunctionsFunction
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `CloudfunctionsFunction` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CloudfunctionsFunctionSpec](#cloudfunctionsfunctionspec)***||
| `status` | ***[CloudfunctionsFunctionStatus](#cloudfunctionsfunctionstatus)***||
## CloudfunctionsFunctionSpec

Appears on:[CloudfunctionsFunction](#cloudfunctionsfunction), [CloudfunctionsFunctionStatus](#cloudfunctionsfunctionstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `availableMemoryMb` | ***int***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `entryPoint` | ***string***| ***(Optional)*** |
| `environmentVariables` | ***map[string]string***| ***(Optional)*** |
| `eventTrigger` | ***[[]CloudfunctionsFunctionSpecEventTrigger](#cloudfunctionsfunctionspeceventtrigger)***| ***(Optional)*** |
| `httpsTriggerURL` | ***string***| ***(Optional)*** |
| `labels` | ***map[string]string***| ***(Optional)*** |
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `retryOnFailure` | ***bool***| ***(Optional)*** Deprecated|
| `runtime` | ***string***| ***(Optional)*** |
| `sourceArchiveBucket` | ***string***||
| `sourceArchiveObject` | ***string***||
| `timeout` | ***int***| ***(Optional)*** |
| `triggerBucket` | ***string***| ***(Optional)*** Deprecated|
| `triggerHTTP` | ***bool***| ***(Optional)*** |
| `triggerTopic` | ***string***| ***(Optional)*** Deprecated|
## CloudfunctionsFunctionSpecEventTrigger

Appears on:[CloudfunctionsFunctionSpec](#cloudfunctionsfunctionspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `eventType` | ***string***||
| `failurePolicy` | ***[[]CloudfunctionsFunctionSpecEventTriggerFailurePolicy](#cloudfunctionsfunctionspeceventtriggerfailurepolicy)***| ***(Optional)*** |
| `resource` | ***string***||
## CloudfunctionsFunctionSpecEventTriggerFailurePolicy

Appears on:[CloudfunctionsFunctionSpecEventTrigger](#cloudfunctionsfunctionspeceventtrigger)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `retry` | ***bool***||
## CloudfunctionsFunctionStatus

Appears on:[CloudfunctionsFunction](#cloudfunctionsfunction)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CloudfunctionsFunctionSpec](#cloudfunctionsfunctionspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[CloudfunctionsFunctionStatus](#cloudfunctionsfunctionstatus)

---
