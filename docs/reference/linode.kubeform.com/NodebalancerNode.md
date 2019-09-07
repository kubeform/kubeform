---
title: NodebalancerNode
menu:
  docs_v0.0.1:
    identifier: nodebalancernode-linode.kubeform.com
    name: NodebalancerNode
    parent: linode.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## NodebalancerNode
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `linode.kubeform.com/v1alpha1` |
|    `kind` | string | `NodebalancerNode` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NodebalancerNodeSpec](#NodebalancerNodeSpec)***||
| `status` | ***[NodebalancerNodeStatus](#NodebalancerNodeStatus)***||
## NodebalancerNodeSpec
##### (Appears on:[NodebalancerNode](#NodebalancerNode), [NodebalancerNodeStatus](#NodebalancerNodeStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `address` | ***string***|The private IP Address and port (IP:PORT) where this backend can be reached. This must be a private IP address.|
| `configID` | ***int***|The ID of the NodeBalancerConfig to access.|
| `label` | ***string***|The label for this node. This is for display purposes only.|
| `mode` | ***string***| ***(Optional)*** The mode this NodeBalancer should use when sending traffic to this backend. If set to `accept` this backend is accepting traffic. If set to `reject` this backend will not receive traffic. If set to `drain` this backend will not receive new traffic, but connections already pinned to it will continue to be routed to it.|
| `nodebalancerID` | ***int***|The ID of the NodeBalancer to access.|
| `status` | ***string***| ***(Optional)*** The current status of this node, based on the configured checks of its NodeBalancer Config. (unknown, UP, DOWN)|
| `weight` | ***int***| ***(Optional)*** Used when picking a backend to serve a request and is not pinned to a single backend yet. Nodes with a higher weight will receive more traffic. (1-255)|
## NodebalancerNodeStatus
##### (Appears on:[NodebalancerNode](#NodebalancerNode))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NodebalancerNodeSpec](#NodebalancerNodeSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
