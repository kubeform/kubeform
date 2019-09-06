## LoggingBillingAccountSink
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `LoggingBillingAccountSink` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LoggingBillingAccountSinkSpec](#LoggingBillingAccountSinkSpec)***||
| `status` | ***[LoggingBillingAccountSinkStatus](#LoggingBillingAccountSinkStatus)***||
## LoggingBillingAccountSinkSpec
##### (Appears on:[LoggingBillingAccountSink](#LoggingBillingAccountSink), [LoggingBillingAccountSinkStatus](#LoggingBillingAccountSinkStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `billingAccount` | ***string***||
| `destination` | ***string***||
| `filter` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `writerIdentity` | ***string***| ***(Optional)*** |
## LoggingBillingAccountSinkStatus
##### (Appears on:[LoggingBillingAccountSink](#LoggingBillingAccountSink))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LoggingBillingAccountSinkSpec](#LoggingBillingAccountSinkSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
