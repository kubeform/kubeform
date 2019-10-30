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
| `spec` | ***[ContainerNodePoolSpec](#containernodepoolspec)***||
| `status` | ***[ContainerNodePoolStatus](#containernodepoolstatus)***||
## ContainerNodePoolSpec

Appears on:[ContainerNodePool](#containernodepool), [ContainerNodePoolStatus](#containernodepoolstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `autoscaling` | ***[[]ContainerNodePoolSpecAutoscaling](#containernodepoolspecautoscaling)***| ***(Optional)*** |
| `cluster` | ***string***||
| `initialNodeCount` | ***int***| ***(Optional)*** |
| `instanceGroupUrls` | ***[]string***| ***(Optional)*** |
| `management` | ***[[]ContainerNodePoolSpecManagement](#containernodepoolspecmanagement)***| ***(Optional)*** |
| `maxPodsPerNode` | ***int***| ***(Optional)*** Deprecated|
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** Deprecated|
| `nodeConfig` | ***[[]ContainerNodePoolSpecNodeConfig](#containernodepoolspecnodeconfig)***| ***(Optional)*** |
| `nodeCount` | ***int***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `version` | ***string***| ***(Optional)*** |
| `zone` | ***string***| ***(Optional)*** |
## ContainerNodePoolSpecAutoscaling

Appears on:[ContainerNodePoolSpec](#containernodepoolspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `maxNodeCount` | ***int***||
| `minNodeCount` | ***int***||
## ContainerNodePoolSpecManagement

Appears on:[ContainerNodePoolSpec](#containernodepoolspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `autoRepair` | ***bool***| ***(Optional)*** |
| `autoUpgrade` | ***bool***| ***(Optional)*** |
## ContainerNodePoolSpecNodeConfig

Appears on:[ContainerNodePoolSpec](#containernodepoolspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `diskSizeGb` | ***int***| ***(Optional)*** |
| `diskType` | ***string***| ***(Optional)*** |
| `guestAccelerator` | ***[[]ContainerNodePoolSpecNodeConfigGuestAccelerator](#containernodepoolspecnodeconfigguestaccelerator)***| ***(Optional)*** |
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
| `taint` | ***[[]ContainerNodePoolSpecNodeConfigTaint](#containernodepoolspecnodeconfigtaint)***| ***(Optional)*** Deprecated|
| `workloadMetadataConfig` | ***[[]ContainerNodePoolSpecNodeConfigWorkloadMetadataConfig](#containernodepoolspecnodeconfigworkloadmetadataconfig)***| ***(Optional)*** Deprecated|
## ContainerNodePoolSpecNodeConfigGuestAccelerator

Appears on:[ContainerNodePoolSpecNodeConfig](#containernodepoolspecnodeconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int***||
| `type` | ***string***||
## ContainerNodePoolSpecNodeConfigTaint

Appears on:[ContainerNodePoolSpecNodeConfig](#containernodepoolspecnodeconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `effect` | ***string***||
| `key` | ***string***||
| `value` | ***string***||
## ContainerNodePoolSpecNodeConfigWorkloadMetadataConfig

Appears on:[ContainerNodePoolSpecNodeConfig](#containernodepoolspecnodeconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `nodeMetadata` | ***string***||
## ContainerNodePoolStatus

Appears on:[ContainerNodePool](#containernodepool)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ContainerNodePoolSpec](#containernodepoolspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ContainerNodePoolStatus](#containernodepoolstatus)

---
