## DbSecurityGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DbSecurityGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DbSecurityGroupSpec](#DbSecurityGroupSpec)***||
| `status` | ***[DbSecurityGroupStatus](#DbSecurityGroupStatus)***||
## DbSecurityGroupSpec
##### (Appears on:[DbSecurityGroup](#DbSecurityGroup), [DbSecurityGroupStatus](#DbSecurityGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `ingress` | ***[[]DbSecurityGroupSpecIngress](#DbSecurityGroupSpecIngress)***||
| `name` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## DbSecurityGroupSpecIngress
##### (Appears on:[DbSecurityGroupSpec](#DbSecurityGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `cidr` | ***string***| ***(Optional)*** |
| `securityGroupID` | ***string***| ***(Optional)*** |
| `securityGroupName` | ***string***| ***(Optional)*** |
| `securityGroupOwnerID` | ***string***| ***(Optional)*** |
## DbSecurityGroupStatus
##### (Appears on:[DbSecurityGroup](#DbSecurityGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DbSecurityGroupSpec](#DbSecurityGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
