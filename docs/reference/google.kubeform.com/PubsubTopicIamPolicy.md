---
title: PubsubTopicIamPolicy
menu:
  docs_{{ .version }}:
    identifier: pubsubtopiciampolicy-google.kubeform.com
    name: PubsubTopicIamPolicy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## PubsubTopicIamPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `PubsubTopicIamPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PubsubTopicIamPolicySpec](#PubsubTopicIamPolicySpec)***||
| `status` | ***[PubsubTopicIamPolicyStatus](#PubsubTopicIamPolicyStatus)***||
## PubsubTopicIamPolicySpec
##### (Appears on:[PubsubTopicIamPolicy](#PubsubTopicIamPolicy), [PubsubTopicIamPolicyStatus](#PubsubTopicIamPolicyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `policyData` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `topic` | ***string***||
## PubsubTopicIamPolicyStatus
##### (Appears on:[PubsubTopicIamPolicy](#PubsubTopicIamPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PubsubTopicIamPolicySpec](#PubsubTopicIamPolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
