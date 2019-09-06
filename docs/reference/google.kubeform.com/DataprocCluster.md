## DataprocCluster
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `DataprocCluster` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DataprocClusterSpec](#DataprocClusterSpec)***||
| `status` | ***[DataprocClusterStatus](#DataprocClusterStatus)***||
## DataprocClusterSpec
##### (Appears on:[DataprocCluster](#DataprocCluster), [DataprocClusterStatus](#DataprocClusterStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `clusterConfig` | ***[[]DataprocClusterSpecClusterConfig](#DataprocClusterSpecClusterConfig)***| ***(Optional)*** |
| `labels` | ***map[string]string***| ***(Optional)*** |
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
## DataprocClusterSpecClusterConfig
##### (Appears on:[DataprocClusterSpec](#DataprocClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucket` | ***string***| ***(Optional)*** |
| `deleteAutogenBucket` | ***bool***| ***(Optional)*** Deprecated|
| `gceClusterConfig` | ***[[]DataprocClusterSpecClusterConfigGceClusterConfig](#DataprocClusterSpecClusterConfigGceClusterConfig)***| ***(Optional)*** |
| `initializationAction` | ***[[]DataprocClusterSpecClusterConfigInitializationAction](#DataprocClusterSpecClusterConfigInitializationAction)***| ***(Optional)*** |
| `masterConfig` | ***[[]DataprocClusterSpecClusterConfigMasterConfig](#DataprocClusterSpecClusterConfigMasterConfig)***| ***(Optional)*** |
| `preemptibleWorkerConfig` | ***[[]DataprocClusterSpecClusterConfigPreemptibleWorkerConfig](#DataprocClusterSpecClusterConfigPreemptibleWorkerConfig)***| ***(Optional)*** |
| `softwareConfig` | ***[[]DataprocClusterSpecClusterConfigSoftwareConfig](#DataprocClusterSpecClusterConfigSoftwareConfig)***| ***(Optional)*** |
| `stagingBucket` | ***string***| ***(Optional)*** |
| `workerConfig` | ***[[]DataprocClusterSpecClusterConfigWorkerConfig](#DataprocClusterSpecClusterConfigWorkerConfig)***| ***(Optional)*** |
## DataprocClusterSpecClusterConfigGceClusterConfig
##### (Appears on:[DataprocClusterSpecClusterConfig](#DataprocClusterSpecClusterConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `internalIPOnly` | ***bool***| ***(Optional)*** |
| `metadata` | ***map[string]string***| ***(Optional)*** |
| `network` | ***string***| ***(Optional)*** |
| `serviceAccount` | ***string***| ***(Optional)*** |
| `serviceAccountScopes` | ***[]string***| ***(Optional)*** |
| `subnetwork` | ***string***| ***(Optional)*** |
| `tags` | ***[]string***| ***(Optional)*** |
| `zone` | ***string***| ***(Optional)*** |
## DataprocClusterSpecClusterConfigInitializationAction
##### (Appears on:[DataprocClusterSpecClusterConfig](#DataprocClusterSpecClusterConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `script` | ***string***||
| `timeoutSec` | ***int***| ***(Optional)*** |
## DataprocClusterSpecClusterConfigMasterConfig
##### (Appears on:[DataprocClusterSpecClusterConfig](#DataprocClusterSpecClusterConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `diskConfig` | ***[[]DataprocClusterSpecClusterConfigMasterConfigDiskConfig](#DataprocClusterSpecClusterConfigMasterConfigDiskConfig)***| ***(Optional)*** |
| `instanceNames` | ***[]string***| ***(Optional)*** |
| `machineType` | ***string***| ***(Optional)*** |
| `numInstances` | ***int***| ***(Optional)*** |
## DataprocClusterSpecClusterConfigMasterConfigDiskConfig
##### (Appears on:[DataprocClusterSpecClusterConfigMasterConfig](#DataprocClusterSpecClusterConfigMasterConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `bootDiskSizeGb` | ***int***| ***(Optional)*** |
| `bootDiskType` | ***string***| ***(Optional)*** |
| `numLocalSsds` | ***int***| ***(Optional)*** |
## DataprocClusterSpecClusterConfigPreemptibleWorkerConfig
##### (Appears on:[DataprocClusterSpecClusterConfig](#DataprocClusterSpecClusterConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `diskConfig` | ***[[]DataprocClusterSpecClusterConfigPreemptibleWorkerConfigDiskConfig](#DataprocClusterSpecClusterConfigPreemptibleWorkerConfigDiskConfig)***| ***(Optional)*** |
| `instanceNames` | ***[]string***| ***(Optional)*** |
| `numInstances` | ***int***| ***(Optional)*** |
## DataprocClusterSpecClusterConfigPreemptibleWorkerConfigDiskConfig
##### (Appears on:[DataprocClusterSpecClusterConfigPreemptibleWorkerConfig](#DataprocClusterSpecClusterConfigPreemptibleWorkerConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `bootDiskSizeGb` | ***int***| ***(Optional)*** |
## DataprocClusterSpecClusterConfigSoftwareConfig
##### (Appears on:[DataprocClusterSpecClusterConfig](#DataprocClusterSpecClusterConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `imageVersion` | ***string***| ***(Optional)*** |
| `overrideProperties` | ***map[string]string***| ***(Optional)*** |
| `properties` | ***map[string]string***| ***(Optional)*** |
## DataprocClusterSpecClusterConfigWorkerConfig
##### (Appears on:[DataprocClusterSpecClusterConfig](#DataprocClusterSpecClusterConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `diskConfig` | ***[[]DataprocClusterSpecClusterConfigWorkerConfigDiskConfig](#DataprocClusterSpecClusterConfigWorkerConfigDiskConfig)***| ***(Optional)*** |
| `instanceNames` | ***[]string***| ***(Optional)*** |
| `machineType` | ***string***| ***(Optional)*** |
| `numInstances` | ***int***| ***(Optional)*** |
## DataprocClusterSpecClusterConfigWorkerConfigDiskConfig
##### (Appears on:[DataprocClusterSpecClusterConfigWorkerConfig](#DataprocClusterSpecClusterConfigWorkerConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `bootDiskSizeGb` | ***int***| ***(Optional)*** |
| `bootDiskType` | ***string***| ***(Optional)*** |
| `numLocalSsds` | ***int***| ***(Optional)*** |
## DataprocClusterStatus
##### (Appears on:[DataprocCluster](#DataprocCluster))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DataprocClusterSpec](#DataprocClusterSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
