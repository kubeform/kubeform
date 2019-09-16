---
title: LoggingBillingAccountExclusion
menu:
  docs_{{ .version }}:
    identifier: loggingbillingaccountexclusion-google.kubeform.com
    name: LoggingBillingAccountExclusion
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LoggingBillingAccountExclusion
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `LoggingBillingAccountExclusion` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LoggingBillingAccountExclusionSpec](#loggingbillingaccountexclusionspec)***||
| `status` | ***[LoggingBillingAccountExclusionStatus](#loggingbillingaccountexclusionstatus)***||
## LoggingBillingAccountExclusionSpec

Appears on:[LoggingBillingAccountExclusion](#loggingbillingaccountexclusion), [LoggingBillingAccountExclusionStatus](#loggingbillingaccountexclusionstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `billingAccount` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `disabled` | ***bool***| ***(Optional)*** |
| `filter` | ***string***||
| `name` | ***string***||
## LoggingBillingAccountExclusionStatus

Appears on:[LoggingBillingAccountExclusion](#loggingbillingaccountexclusion)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LoggingBillingAccountExclusionSpec](#loggingbillingaccountexclusionspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
