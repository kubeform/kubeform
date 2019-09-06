## AppmeshRoute
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AppmeshRoute` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AppmeshRouteSpec](#AppmeshRouteSpec)***||
| `status` | ***[AppmeshRouteStatus](#AppmeshRouteStatus)***||
## AppmeshRouteSpec
##### (Appears on:[AppmeshRoute](#AppmeshRoute), [AppmeshRouteStatus](#AppmeshRouteStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `createdDate` | ***string***| ***(Optional)*** |
| `lastUpdatedDate` | ***string***| ***(Optional)*** |
| `meshName` | ***string***||
| `name` | ***string***||
| `spec` | ***[[]AppmeshRouteSpecSpec](#AppmeshRouteSpecSpec)***||
| `virtualRouterName` | ***string***||
## AppmeshRouteSpecSpec
##### (Appears on:[AppmeshRouteSpec](#AppmeshRouteSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `httpRoute` | ***[[]AppmeshRouteSpecSpecHttpRoute](#AppmeshRouteSpecSpecHttpRoute)***| ***(Optional)*** |
| `tcpRoute` | ***[[]AppmeshRouteSpecSpecTcpRoute](#AppmeshRouteSpecSpecTcpRoute)***| ***(Optional)*** |
## AppmeshRouteSpecSpecHttpRoute
##### (Appears on:[AppmeshRouteSpecSpec](#AppmeshRouteSpecSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***[[]AppmeshRouteSpecSpecHttpRouteAction](#AppmeshRouteSpecSpecHttpRouteAction)***||
| `match` | ***[[]AppmeshRouteSpecSpecHttpRouteMatch](#AppmeshRouteSpecSpecHttpRouteMatch)***||
## AppmeshRouteSpecSpecHttpRouteAction
##### (Appears on:[AppmeshRouteSpecSpecHttpRoute](#AppmeshRouteSpecSpecHttpRoute))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `weightedTarget` | ***[[]AppmeshRouteSpecSpecHttpRouteActionWeightedTarget](#AppmeshRouteSpecSpecHttpRouteActionWeightedTarget)***||
## AppmeshRouteSpecSpecHttpRouteActionWeightedTarget
##### (Appears on:[AppmeshRouteSpecSpecHttpRouteAction](#AppmeshRouteSpecSpecHttpRouteAction))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `virtualNode` | ***string***||
| `weight` | ***int***||
## AppmeshRouteSpecSpecHttpRouteMatch
##### (Appears on:[AppmeshRouteSpecSpecHttpRoute](#AppmeshRouteSpecSpecHttpRoute))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `prefix` | ***string***||
## AppmeshRouteSpecSpecTcpRoute
##### (Appears on:[AppmeshRouteSpecSpec](#AppmeshRouteSpecSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***[[]AppmeshRouteSpecSpecTcpRouteAction](#AppmeshRouteSpecSpecTcpRouteAction)***||
## AppmeshRouteSpecSpecTcpRouteAction
##### (Appears on:[AppmeshRouteSpecSpecTcpRoute](#AppmeshRouteSpecSpecTcpRoute))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `weightedTarget` | ***[[]AppmeshRouteSpecSpecTcpRouteActionWeightedTarget](#AppmeshRouteSpecSpecTcpRouteActionWeightedTarget)***||
## AppmeshRouteSpecSpecTcpRouteActionWeightedTarget
##### (Appears on:[AppmeshRouteSpecSpecTcpRouteAction](#AppmeshRouteSpecSpecTcpRouteAction))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `virtualNode` | ***string***||
| `weight` | ***int***||
## AppmeshRouteStatus
##### (Appears on:[AppmeshRoute](#AppmeshRoute))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AppmeshRouteSpec](#AppmeshRouteSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
