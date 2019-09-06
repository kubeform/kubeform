## SsmMaintenanceWindowTarget
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SsmMaintenanceWindowTarget` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SsmMaintenanceWindowTargetSpec](#SsmMaintenanceWindowTargetSpec)***||
| `status` | ***[SsmMaintenanceWindowTargetStatus](#SsmMaintenanceWindowTargetStatus)***||
## SsmMaintenanceWindowTargetSpec
##### (Appears on:[SsmMaintenanceWindowTarget](#SsmMaintenanceWindowTarget), [SsmMaintenanceWindowTargetStatus](#SsmMaintenanceWindowTargetStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `ownerInformation` | ***string***| ***(Optional)*** |
| `resourceType` | ***string***||
| `targets` | ***[[]SsmMaintenanceWindowTargetSpecTargets](#SsmMaintenanceWindowTargetSpecTargets)***||
| `windowID` | ***string***||
## SsmMaintenanceWindowTargetSpecTargets
##### (Appears on:[SsmMaintenanceWindowTargetSpec](#SsmMaintenanceWindowTargetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `key` | ***string***||
| `values` | ***[]string***||
## SsmMaintenanceWindowTargetStatus
##### (Appears on:[SsmMaintenanceWindowTarget](#SsmMaintenanceWindowTarget))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SsmMaintenanceWindowTargetSpec](#SsmMaintenanceWindowTargetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
