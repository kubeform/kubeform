---
title: ContainerCluster
menu:
  docs_{{ .version }}:
    identifier: containercluster-google.kubeform.com
    name: ContainerCluster
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ContainerCluster
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ContainerCluster` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ContainerClusterSpec](#ContainerClusterSpec)***||
| `status` | ***[ContainerClusterStatus](#ContainerClusterStatus)***||
## ContainerClusterSpec

Appears on:[ContainerCluster](#ContainerCluster), [ContainerClusterStatus](#ContainerClusterStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `additionalZones` | ***[]string***| ***(Optional)*** |
| `addonsConfig` | ***[[]ContainerClusterSpecAddonsConfig](#ContainerClusterSpecAddonsConfig)***| ***(Optional)*** |
| `clusterIpv4CIDR` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `enableBinaryAuthorization` | ***bool***| ***(Optional)*** Deprecated|
| `enableKubernetesAlpha` | ***bool***| ***(Optional)*** |
| `enableLegacyAbac` | ***bool***| ***(Optional)*** |
| `enableTpu` | ***bool***| ***(Optional)*** Deprecated|
| `endpoint` | ***string***| ***(Optional)*** |
| `initialNodeCount` | ***int***| ***(Optional)*** |
| `instanceGroupUrls` | ***[]string***| ***(Optional)*** |
| `ipAllocationPolicy` | ***[[]ContainerClusterSpecIpAllocationPolicy](#ContainerClusterSpecIpAllocationPolicy)***| ***(Optional)*** |
| `loggingService` | ***string***| ***(Optional)*** |
| `maintenancePolicy` | ***[[]ContainerClusterSpecMaintenancePolicy](#ContainerClusterSpecMaintenancePolicy)***| ***(Optional)*** |
| `masterAuth` | ***[[]ContainerClusterSpecMasterAuth](#ContainerClusterSpecMasterAuth)***| ***(Optional)*** |
| `masterAuthorizedNetworksConfig` | ***[[]ContainerClusterSpecMasterAuthorizedNetworksConfig](#ContainerClusterSpecMasterAuthorizedNetworksConfig)***| ***(Optional)*** |
| `masterIpv4CIDRBlock` | ***string***| ***(Optional)*** Deprecated|
| `masterVersion` | ***string***| ***(Optional)*** |
| `minMasterVersion` | ***string***| ***(Optional)*** |
| `monitoringService` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `network` | ***string***| ***(Optional)*** |
| `networkPolicy` | ***[[]ContainerClusterSpecNetworkPolicy](#ContainerClusterSpecNetworkPolicy)***| ***(Optional)*** |
| `nodeConfig` | ***[[]ContainerClusterSpecNodeConfig](#ContainerClusterSpecNodeConfig)***| ***(Optional)*** |
| `nodePool` | ***[[]ContainerClusterSpecNodePool](#ContainerClusterSpecNodePool)***| ***(Optional)*** |
| `nodeVersion` | ***string***| ***(Optional)*** |
| `podSecurityPolicyConfig` | ***[[]ContainerClusterSpecPodSecurityPolicyConfig](#ContainerClusterSpecPodSecurityPolicyConfig)***| ***(Optional)*** Deprecated|
| `privateCluster` | ***bool***| ***(Optional)*** Deprecated|
| `privateClusterConfig` | ***[[]ContainerClusterSpecPrivateClusterConfig](#ContainerClusterSpecPrivateClusterConfig)***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `removeDefaultNodePool` | ***bool***| ***(Optional)*** |
| `resourceLabels` | ***map[string]string***| ***(Optional)*** |
| `subnetwork` | ***string***| ***(Optional)*** |
| `zone` | ***string***| ***(Optional)*** |
## ContainerClusterSpecAddonsConfig

Appears on:[ContainerClusterSpec](#ContainerClusterSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `horizontalPodAutoscaling` | ***[[]ContainerClusterSpecAddonsConfigHorizontalPodAutoscaling](#ContainerClusterSpecAddonsConfigHorizontalPodAutoscaling)***| ***(Optional)*** |
| `httpLoadBalancing` | ***[[]ContainerClusterSpecAddonsConfigHttpLoadBalancing](#ContainerClusterSpecAddonsConfigHttpLoadBalancing)***| ***(Optional)*** |
| `kubernetesDashboard` | ***[[]ContainerClusterSpecAddonsConfigKubernetesDashboard](#ContainerClusterSpecAddonsConfigKubernetesDashboard)***| ***(Optional)*** |
| `networkPolicyConfig` | ***[[]ContainerClusterSpecAddonsConfigNetworkPolicyConfig](#ContainerClusterSpecAddonsConfigNetworkPolicyConfig)***| ***(Optional)*** |
## ContainerClusterSpecAddonsConfigHorizontalPodAutoscaling

Appears on:[ContainerClusterSpecAddonsConfig](#ContainerClusterSpecAddonsConfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `disabled` | ***bool***| ***(Optional)*** |
## ContainerClusterSpecAddonsConfigHttpLoadBalancing

Appears on:[ContainerClusterSpecAddonsConfig](#ContainerClusterSpecAddonsConfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `disabled` | ***bool***| ***(Optional)*** |
## ContainerClusterSpecAddonsConfigKubernetesDashboard

Appears on:[ContainerClusterSpecAddonsConfig](#ContainerClusterSpecAddonsConfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `disabled` | ***bool***| ***(Optional)*** |
## ContainerClusterSpecAddonsConfigNetworkPolicyConfig

Appears on:[ContainerClusterSpecAddonsConfig](#ContainerClusterSpecAddonsConfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `disabled` | ***bool***| ***(Optional)*** |
## ContainerClusterSpecIpAllocationPolicy

Appears on:[ContainerClusterSpec](#ContainerClusterSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `clusterIpv4CIDRBlock` | ***string***| ***(Optional)*** |
| `clusterSecondaryRangeName` | ***string***| ***(Optional)*** |
| `createSubnetwork` | ***bool***| ***(Optional)*** |
| `servicesIpv4CIDRBlock` | ***string***| ***(Optional)*** |
| `servicesSecondaryRangeName` | ***string***| ***(Optional)*** |
| `subnetworkName` | ***string***| ***(Optional)*** |
## ContainerClusterSpecMaintenancePolicy

Appears on:[ContainerClusterSpec](#ContainerClusterSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `dailyMaintenanceWindow` | ***[[]ContainerClusterSpecMaintenancePolicyDailyMaintenanceWindow](#ContainerClusterSpecMaintenancePolicyDailyMaintenanceWindow)***||
## ContainerClusterSpecMaintenancePolicyDailyMaintenanceWindow

Appears on:[ContainerClusterSpecMaintenancePolicy](#ContainerClusterSpecMaintenancePolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `duration` | ***string***| ***(Optional)*** |
| `startTime` | ***string***||
## ContainerClusterSpecMasterAuth

Appears on:[ContainerClusterSpec](#ContainerClusterSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `clientCertificate` | ***string***| ***(Optional)*** |
| `clientCertificateConfig` | ***[[]ContainerClusterSpecMasterAuthClientCertificateConfig](#ContainerClusterSpecMasterAuthClientCertificateConfig)***| ***(Optional)*** |
| `clusterCaCertificate` | ***string***| ***(Optional)*** |
| `username` | ***string***||
## ContainerClusterSpecMasterAuthClientCertificateConfig

Appears on:[ContainerClusterSpecMasterAuth](#ContainerClusterSpecMasterAuth)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `issueClientCertificate` | ***bool***||
## ContainerClusterSpecMasterAuthorizedNetworksConfig

Appears on:[ContainerClusterSpec](#ContainerClusterSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cidrBlocks` | ***[[]ContainerClusterSpecMasterAuthorizedNetworksConfigCidrBlocks](#ContainerClusterSpecMasterAuthorizedNetworksConfigCidrBlocks)***| ***(Optional)*** |
## ContainerClusterSpecMasterAuthorizedNetworksConfigCidrBlocks

Appears on:[ContainerClusterSpecMasterAuthorizedNetworksConfig](#ContainerClusterSpecMasterAuthorizedNetworksConfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cidrBlock` | ***string***||
| `displayName` | ***string***| ***(Optional)*** |
## ContainerClusterSpecNetworkPolicy

Appears on:[ContainerClusterSpec](#ContainerClusterSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `provider` | ***string***| ***(Optional)*** |
## ContainerClusterSpecNodeConfig

Appears on:[ContainerClusterSpec](#ContainerClusterSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `diskSizeGb` | ***int***| ***(Optional)*** |
| `diskType` | ***string***| ***(Optional)*** |
| `guestAccelerator` | ***[[]ContainerClusterSpecNodeConfigGuestAccelerator](#ContainerClusterSpecNodeConfigGuestAccelerator)***| ***(Optional)*** |
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
| `taint` | ***[[]ContainerClusterSpecNodeConfigTaint](#ContainerClusterSpecNodeConfigTaint)***| ***(Optional)*** Deprecated|
| `workloadMetadataConfig` | ***[[]ContainerClusterSpecNodeConfigWorkloadMetadataConfig](#ContainerClusterSpecNodeConfigWorkloadMetadataConfig)***| ***(Optional)*** Deprecated|
## ContainerClusterSpecNodeConfigGuestAccelerator

Appears on:[ContainerClusterSpecNodeConfig](#ContainerClusterSpecNodeConfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int***||
| `type` | ***string***||
## ContainerClusterSpecNodeConfigTaint

Appears on:[ContainerClusterSpecNodeConfig](#ContainerClusterSpecNodeConfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `effect` | ***string***||
| `key` | ***string***||
| `value` | ***string***||
## ContainerClusterSpecNodeConfigWorkloadMetadataConfig

Appears on:[ContainerClusterSpecNodeConfig](#ContainerClusterSpecNodeConfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `nodeMetadata` | ***string***||
## ContainerClusterSpecNodePool

Appears on:[ContainerClusterSpec](#ContainerClusterSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `autoscaling` | ***[[]ContainerClusterSpecNodePoolAutoscaling](#ContainerClusterSpecNodePoolAutoscaling)***| ***(Optional)*** |
| `initialNodeCount` | ***int***| ***(Optional)*** |
| `instanceGroupUrls` | ***[]string***| ***(Optional)*** |
| `management` | ***[[]ContainerClusterSpecNodePoolManagement](#ContainerClusterSpecNodePoolManagement)***| ***(Optional)*** |
| `maxPodsPerNode` | ***int***| ***(Optional)*** Deprecated|
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** Deprecated|
| `nodeConfig` | ***[[]ContainerClusterSpecNodePoolNodeConfig](#ContainerClusterSpecNodePoolNodeConfig)***| ***(Optional)*** |
| `nodeCount` | ***int***| ***(Optional)*** |
| `version` | ***string***| ***(Optional)*** |
## ContainerClusterSpecNodePoolAutoscaling

Appears on:[ContainerClusterSpecNodePool](#ContainerClusterSpecNodePool)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `maxNodeCount` | ***int***||
| `minNodeCount` | ***int***||
## ContainerClusterSpecNodePoolManagement

Appears on:[ContainerClusterSpecNodePool](#ContainerClusterSpecNodePool)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `autoRepair` | ***bool***| ***(Optional)*** |
| `autoUpgrade` | ***bool***| ***(Optional)*** |
## ContainerClusterSpecNodePoolNodeConfig

Appears on:[ContainerClusterSpecNodePool](#ContainerClusterSpecNodePool)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `diskSizeGb` | ***int***| ***(Optional)*** |
| `diskType` | ***string***| ***(Optional)*** |
| `guestAccelerator` | ***[[]ContainerClusterSpecNodePoolNodeConfigGuestAccelerator](#ContainerClusterSpecNodePoolNodeConfigGuestAccelerator)***| ***(Optional)*** |
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
| `taint` | ***[[]ContainerClusterSpecNodePoolNodeConfigTaint](#ContainerClusterSpecNodePoolNodeConfigTaint)***| ***(Optional)*** Deprecated|
| `workloadMetadataConfig` | ***[[]ContainerClusterSpecNodePoolNodeConfigWorkloadMetadataConfig](#ContainerClusterSpecNodePoolNodeConfigWorkloadMetadataConfig)***| ***(Optional)*** Deprecated|
## ContainerClusterSpecNodePoolNodeConfigGuestAccelerator

Appears on:[ContainerClusterSpecNodePoolNodeConfig](#ContainerClusterSpecNodePoolNodeConfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int***||
| `type` | ***string***||
## ContainerClusterSpecNodePoolNodeConfigTaint

Appears on:[ContainerClusterSpecNodePoolNodeConfig](#ContainerClusterSpecNodePoolNodeConfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `effect` | ***string***||
| `key` | ***string***||
| `value` | ***string***||
## ContainerClusterSpecNodePoolNodeConfigWorkloadMetadataConfig

Appears on:[ContainerClusterSpecNodePoolNodeConfig](#ContainerClusterSpecNodePoolNodeConfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `nodeMetadata` | ***string***||
## ContainerClusterSpecPodSecurityPolicyConfig

Appears on:[ContainerClusterSpec](#ContainerClusterSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
## ContainerClusterSpecPrivateClusterConfig

Appears on:[ContainerClusterSpec](#ContainerClusterSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enablePrivateEndpoint` | ***bool***| ***(Optional)*** |
| `enablePrivateNodes` | ***bool***| ***(Optional)*** |
| `masterIpv4CIDRBlock` | ***string***| ***(Optional)*** |
| `privateEndpoint` | ***string***| ***(Optional)*** |
| `publicEndpoint` | ***string***| ***(Optional)*** |
## ContainerClusterStatus

Appears on:[ContainerCluster](#ContainerCluster)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ContainerClusterSpec](#ContainerClusterSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `master_auth.<index>.client_key` | ***string*** ||
| `master_auth.<index>.password` | ***string*** ||
