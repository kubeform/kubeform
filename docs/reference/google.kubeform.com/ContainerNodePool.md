---
title: ContainerNodePool
menu:
  docs_{{ .version }}:
    identifier: containernodepool-google.kubeform.com
    name: ContainerNodePool
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ContainerNodePool
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ContainerNodePool` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ContainerNodePoolSpec](#ContainerNodePoolSpec)***||
| `status` | ***[ContainerNodePoolStatus](#ContainerNodePoolStatus)***||
## ContainerNodePoolSpec
##### (Appears on:[ContainerNodePool](#ContainerNodePool), [ContainerNodePoolStatus](#ContainerNodePoolStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `autoscaling` | ***[[]ContainerNodePoolSpecAutoscaling](#ContainerNodePoolSpecAutoscaling)***| ***(Optional)*** |
| `cluster` | ***string***||
| `initialNodeCount` | ***int***| ***(Optional)*** |
| `instanceGroupUrls` | ***[]string***| ***(Optional)*** |
| `management` | ***[[]ContainerNodePoolSpecManagement](#ContainerNodePoolSpecManagement)***| ***(Optional)*** |
| `maxPodsPerNode` | ***int***| ***(Optional)*** Deprecated|
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** Deprecated|
| `nodeConfig` | ***[[]ContainerNodePoolSpecNodeConfig](#ContainerNodePoolSpecNodeConfig)***| ***(Optional)*** |
| `nodeCount` | ***int***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `version` | ***string***| ***(Optional)*** |
| `zone` | ***string***| ***(Optional)*** |
## ContainerNodePoolSpecAutoscaling
##### (Appears on:[ContainerNodePoolSpec](#ContainerNodePoolSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `maxNodeCount` | ***int***||
| `minNodeCount` | ***int***||
## ContainerNodePoolSpecManagement
##### (Appears on:[ContainerNodePoolSpec](#ContainerNodePoolSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `autoRepair` | ***bool***| ***(Optional)*** |
| `autoUpgrade` | ***bool***| ***(Optional)*** |
## ContainerNodePoolSpecNodeConfig
##### (Appears on:[ContainerNodePoolSpec](#ContainerNodePoolSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `diskSizeGb` | ***int***| ***(Optional)*** |
| `diskType` | ***string***| ***(Optional)*** |
| `guestAccelerator` | ***[[]ContainerNodePoolSpecNodeConfigGuestAccelerator](#ContainerNodePoolSpecNodeConfigGuestAccelerator)***| ***(Optional)*** |
| `imageType` | ***string***| ***(Optional)*** |
| `labels` | ***map[string]string***| ***(Optional)*** |
| `localSsdCount` | ***int***| ***(Optional)*** |
| `machineType` | ***string***| ***(Optional)*** |
| `metadata` | ***map[string]string***| ***(Optional)*** |
| `minCPUPlatform` | ***string***| ***(Optional)*** |
| `oauthScopes` | ***[]string***| ***(Optional)*** |
| `preemptible` | ***bool***| ***(Optional)*** |
| `serviceAccount` | ***string***| ***(Optional)*** |
| `tags` | ***[]string***| ***(Optional)*** |
| `taint` | ***[[]ContainerNodePoolSpecNodeConfigTaint](#ContainerNodePoolSpecNodeConfigTaint)***| ***(Optional)*** Deprecated|
| `workloadMetadataConfig` | ***[[]ContainerNodePoolSpecNodeConfigWorkloadMetadataConfig](#ContainerNodePoolSpecNodeConfigWorkloadMetadataConfig)***| ***(Optional)*** Deprecated|
## ContainerNodePoolSpecNodeConfigGuestAccelerator
##### (Appears on:[ContainerNodePoolSpecNodeConfig](#ContainerNodePoolSpecNodeConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int***||
| `type` | ***string***||
## ContainerNodePoolSpecNodeConfigTaint
##### (Appears on:[ContainerNodePoolSpecNodeConfig](#ContainerNodePoolSpecNodeConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `effect` | ***string***||
| `key` | ***string***||
| `value` | ***string***||
## ContainerNodePoolSpecNodeConfigWorkloadMetadataConfig
##### (Appears on:[ContainerNodePoolSpecNodeConfig](#ContainerNodePoolSpecNodeConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `nodeMetadata` | ***string***||
## ContainerNodePoolStatus
##### (Appears on:[ContainerNodePool](#ContainerNodePool))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ContainerNodePoolSpec](#ContainerNodePoolSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
