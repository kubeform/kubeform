---
title: ProxyProtocolPolicy
menu:
  docs_{{ .version }}:
    identifier: proxyprotocolpolicy-aws.kubeform.com
    name: ProxyProtocolPolicy
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ProxyProtocolPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ProxyProtocolPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ProxyProtocolPolicySpec](#ProxyProtocolPolicySpec)***||
| `status` | ***[ProxyProtocolPolicyStatus](#ProxyProtocolPolicyStatus)***||
## ProxyProtocolPolicySpec

Appears on:[ProxyProtocolPolicy](#ProxyProtocolPolicy), [ProxyProtocolPolicyStatus](#ProxyProtocolPolicyStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `instancePorts` | ***[]string***||
| `loadBalancer` | ***string***||
## ProxyProtocolPolicyStatus

Appears on:[ProxyProtocolPolicy](#ProxyProtocolPolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ProxyProtocolPolicySpec](#ProxyProtocolPolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
