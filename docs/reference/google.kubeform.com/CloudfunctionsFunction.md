---
title: CloudfunctionsFunction
menu:
  docs_v0.0.1:
    identifier: cloudfunctionsfunction-google.kubeform.com
    name: CloudfunctionsFunction
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## CloudfunctionsFunction
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `CloudfunctionsFunction` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CloudfunctionsFunctionSpec](#CloudfunctionsFunctionSpec)***||
| `status` | ***[CloudfunctionsFunctionStatus](#CloudfunctionsFunctionStatus)***||
## CloudfunctionsFunctionSpec
##### (Appears on:[CloudfunctionsFunction](#CloudfunctionsFunction), [CloudfunctionsFunctionStatus](#CloudfunctionsFunctionStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `availableMemoryMb` | ***int***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `entryPoint` | ***string***| ***(Optional)*** |
| `environmentVariables` | ***map[string]string***| ***(Optional)*** |
| `eventTrigger` | ***[[]CloudfunctionsFunctionSpecEventTrigger](#CloudfunctionsFunctionSpecEventTrigger)***| ***(Optional)*** |
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
##### (Appears on:[CloudfunctionsFunctionSpec](#CloudfunctionsFunctionSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `eventType` | ***string***||
| `failurePolicy` | ***[[]CloudfunctionsFunctionSpecEventTriggerFailurePolicy](#CloudfunctionsFunctionSpecEventTriggerFailurePolicy)***| ***(Optional)*** |
| `resource` | ***string***||
## CloudfunctionsFunctionSpecEventTriggerFailurePolicy
##### (Appears on:[CloudfunctionsFunctionSpecEventTrigger](#CloudfunctionsFunctionSpecEventTrigger))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `retry` | ***bool***||
## CloudfunctionsFunctionStatus
##### (Appears on:[CloudfunctionsFunction](#CloudfunctionsFunction))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CloudfunctionsFunctionSpec](#CloudfunctionsFunctionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
