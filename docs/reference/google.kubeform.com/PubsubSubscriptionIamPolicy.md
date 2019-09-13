---
title: PubsubSubscriptionIamPolicy
menu:
  docs_{{ .version }}:
    identifier: pubsubsubscriptioniampolicy-google.kubeform.com
    name: PubsubSubscriptionIamPolicy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## PubsubSubscriptionIamPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `PubsubSubscriptionIamPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PubsubSubscriptionIamPolicySpec](#PubsubSubscriptionIamPolicySpec)***||
| `status` | ***[PubsubSubscriptionIamPolicyStatus](#PubsubSubscriptionIamPolicyStatus)***||
## PubsubSubscriptionIamPolicySpec
##### (Appears on:[PubsubSubscriptionIamPolicy](#PubsubSubscriptionIamPolicy), [PubsubSubscriptionIamPolicyStatus](#PubsubSubscriptionIamPolicyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `policyData` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `subscription` | ***string***||
## PubsubSubscriptionIamPolicyStatus
##### (Appears on:[PubsubSubscriptionIamPolicy](#PubsubSubscriptionIamPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PubsubSubscriptionIamPolicySpec](#PubsubSubscriptionIamPolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
