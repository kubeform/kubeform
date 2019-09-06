## GlueClassifier
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `GlueClassifier` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[GlueClassifierSpec](#GlueClassifierSpec)***||
| `status` | ***[GlueClassifierStatus](#GlueClassifierStatus)***||
## GlueClassifierSpec
##### (Appears on:[GlueClassifier](#GlueClassifier), [GlueClassifierStatus](#GlueClassifierStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `grokClassifier` | ***[[]GlueClassifierSpecGrokClassifier](#GlueClassifierSpecGrokClassifier)***| ***(Optional)*** |
| `jsonClassifier` | ***[[]GlueClassifierSpecJsonClassifier](#GlueClassifierSpecJsonClassifier)***| ***(Optional)*** |
| `name` | ***string***||
| `xmlClassifier` | ***[[]GlueClassifierSpecXmlClassifier](#GlueClassifierSpecXmlClassifier)***| ***(Optional)*** |
## GlueClassifierSpecGrokClassifier
##### (Appears on:[GlueClassifierSpec](#GlueClassifierSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `classification` | ***string***||
| `customPatterns` | ***string***| ***(Optional)*** |
| `grokPattern` | ***string***||
## GlueClassifierSpecJsonClassifier
##### (Appears on:[GlueClassifierSpec](#GlueClassifierSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `jsonPath` | ***string***||
## GlueClassifierSpecXmlClassifier
##### (Appears on:[GlueClassifierSpec](#GlueClassifierSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `classification` | ***string***||
| `rowTag` | ***string***||
## GlueClassifierStatus
##### (Appears on:[GlueClassifier](#GlueClassifier))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[GlueClassifierSpec](#GlueClassifierSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
