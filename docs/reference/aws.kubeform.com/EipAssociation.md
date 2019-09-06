## EipAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `EipAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[EipAssociationSpec](#EipAssociationSpec)***||
| `status` | ***[EipAssociationStatus](#EipAssociationStatus)***||
## EipAssociationSpec
##### (Appears on:[EipAssociation](#EipAssociation), [EipAssociationStatus](#EipAssociationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allocationID` | ***string***| ***(Optional)*** |
| `allowReassociation` | ***bool***| ***(Optional)*** |
| `instanceID` | ***string***| ***(Optional)*** |
| `networkInterfaceID` | ***string***| ***(Optional)*** |
| `privateIPAddress` | ***string***| ***(Optional)*** |
| `publicIP` | ***string***| ***(Optional)*** |
## EipAssociationStatus
##### (Appears on:[EipAssociation](#EipAssociation))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[EipAssociationSpec](#EipAssociationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
