## NetworkSecurityGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `NetworkSecurityGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NetworkSecurityGroupSpec](#NetworkSecurityGroupSpec)***||
| `status` | ***[NetworkSecurityGroupStatus](#NetworkSecurityGroupStatus)***||
## NetworkSecurityGroupSpec
##### (Appears on:[NetworkSecurityGroup](#NetworkSecurityGroup), [NetworkSecurityGroupStatus](#NetworkSecurityGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `securityRule` | ***[[]NetworkSecurityGroupSpecSecurityRule](#NetworkSecurityGroupSpecSecurityRule)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## NetworkSecurityGroupSpecSecurityRule
##### (Appears on:[NetworkSecurityGroupSpec](#NetworkSecurityGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `access` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `destinationAddressPrefix` | ***string***| ***(Optional)*** |
| `destinationAddressPrefixes` | ***[]string***| ***(Optional)*** |
| `destinationApplicationSecurityGroupIDS` | ***[]string***| ***(Optional)*** |
| `destinationPortRange` | ***string***| ***(Optional)*** |
| `destinationPortRanges` | ***[]string***| ***(Optional)*** |
| `direction` | ***string***||
| `name` | ***string***||
| `priority` | ***int***||
| `protocol` | ***string***||
| `sourceAddressPrefix` | ***string***| ***(Optional)*** |
| `sourceAddressPrefixes` | ***[]string***| ***(Optional)*** |
| `sourceApplicationSecurityGroupIDS` | ***[]string***| ***(Optional)*** |
| `sourcePortRange` | ***string***| ***(Optional)*** |
| `sourcePortRanges` | ***[]string***| ***(Optional)*** |
## NetworkSecurityGroupStatus
##### (Appears on:[NetworkSecurityGroup](#NetworkSecurityGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NetworkSecurityGroupSpec](#NetworkSecurityGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
