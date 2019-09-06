## NetworkACL
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `NetworkACL` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NetworkACLSpec](#NetworkACLSpec)***||
| `status` | ***[NetworkACLStatus](#NetworkACLStatus)***||
## NetworkACLSpec
##### (Appears on:[NetworkACL](#NetworkACL), [NetworkACLStatus](#NetworkACLStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `egress` | ***[[]NetworkACLSpecEgress](#NetworkACLSpecEgress)***| ***(Optional)*** |
| `ingress` | ***[[]NetworkACLSpecIngress](#NetworkACLSpecIngress)***| ***(Optional)*** |
| `ownerID` | ***string***| ***(Optional)*** |
| `subnetIDS` | ***[]string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpcID` | ***string***||
## NetworkACLSpecEgress
##### (Appears on:[NetworkACLSpec](#NetworkACLSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***string***||
| `cidrBlock` | ***string***| ***(Optional)*** |
| `fromPort` | ***int***||
| `icmpCode` | ***int***| ***(Optional)*** |
| `icmpType` | ***int***| ***(Optional)*** |
| `ipv6CIDRBlock` | ***string***| ***(Optional)*** |
| `protocol` | ***string***||
| `ruleNo` | ***int***||
| `toPort` | ***int***||
## NetworkACLSpecIngress
##### (Appears on:[NetworkACLSpec](#NetworkACLSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***string***||
| `cidrBlock` | ***string***| ***(Optional)*** |
| `fromPort` | ***int***||
| `icmpCode` | ***int***| ***(Optional)*** |
| `icmpType` | ***int***| ***(Optional)*** |
| `ipv6CIDRBlock` | ***string***| ***(Optional)*** |
| `protocol` | ***string***||
| `ruleNo` | ***int***||
| `toPort` | ***int***||
## NetworkACLStatus
##### (Appears on:[NetworkACL](#NetworkACL))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NetworkACLSpec](#NetworkACLSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
