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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ContainerClusterSpec](#containerclusterspec)***||
| `status` | ***[ContainerClusterStatus](#containerclusterstatus)***||
## ContainerClusterSpec

Appears on:[ContainerCluster](#containercluster), [ContainerClusterStatus](#containerclusterstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `additionalZones` | ***[]string***| ***(Optional)*** Deprecated|
| `addonsConfig` | ***[[]ContainerClusterSpecAddonsConfig](#containerclusterspecaddonsconfig)***| ***(Optional)*** |
| `clusterIpv4CIDR` | ***string***| ***(Optional)*** |
| `defaultMaxPodsPerNode` | ***int64***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `enableKubernetesAlpha` | ***bool***| ***(Optional)*** |
| `enableLegacyAbac` | ***bool***| ***(Optional)*** |
| `endpoint` | ***string***| ***(Optional)*** |
| `initialNodeCount` | ***int64***| ***(Optional)*** |
| `instanceGroupUrls` | ***[]string***| ***(Optional)*** |
| `ipAllocationPolicy` | ***[[]ContainerClusterSpecIpAllocationPolicy](#containerclusterspecipallocationpolicy)***| ***(Optional)*** |
| `location` | ***string***| ***(Optional)*** |
| `loggingService` | ***string***| ***(Optional)*** |
| `maintenancePolicy` | ***[[]ContainerClusterSpecMaintenancePolicy](#containerclusterspecmaintenancepolicy)***| ***(Optional)*** |
| `masterAuth` | ***[[]ContainerClusterSpecMasterAuth](#containerclusterspecmasterauth)***| ***(Optional)*** |
| `masterAuthorizedNetworksConfig` | ***[[]ContainerClusterSpecMasterAuthorizedNetworksConfig](#containerclusterspecmasterauthorizednetworksconfig)***| ***(Optional)*** |
| `masterVersion` | ***string***| ***(Optional)*** |
| `minMasterVersion` | ***string***| ***(Optional)*** |
| `monitoringService` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `network` | ***string***| ***(Optional)*** |
| `networkPolicy` | ***[[]ContainerClusterSpecNetworkPolicy](#containerclusterspecnetworkpolicy)***| ***(Optional)*** |
| `nodeConfig` | ***[[]ContainerClusterSpecNodeConfig](#containerclusterspecnodeconfig)***| ***(Optional)*** |
| `nodeLocations` | ***[]string***| ***(Optional)*** |
| `nodePool` | ***[[]ContainerClusterSpecNodePool](#containerclusterspecnodepool)***| ***(Optional)*** |
| `nodeVersion` | ***string***| ***(Optional)*** |
| `privateClusterConfig` | ***[[]ContainerClusterSpecPrivateClusterConfig](#containerclusterspecprivateclusterconfig)***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** Deprecated|
| `removeDefaultNodePool` | ***bool***| ***(Optional)*** |
| `resourceLabels` | ***map[string]string***| ***(Optional)*** |
| `servicesIpv4CIDR` | ***string***| ***(Optional)*** |
| `subnetwork` | ***string***| ***(Optional)*** |
| `zone` | ***string***| ***(Optional)*** Deprecated|
## ContainerClusterSpecAddonsConfig

Appears on:[ContainerClusterSpec](#containerclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `horizontalPodAutoscaling` | ***[[]ContainerClusterSpecAddonsConfigHorizontalPodAutoscaling](#containerclusterspecaddonsconfighorizontalpodautoscaling)***| ***(Optional)*** |
| `httpLoadBalancing` | ***[[]ContainerClusterSpecAddonsConfigHttpLoadBalancing](#containerclusterspecaddonsconfighttploadbalancing)***| ***(Optional)*** |
| `kubernetesDashboard` | ***[[]ContainerClusterSpecAddonsConfigKubernetesDashboard](#containerclusterspecaddonsconfigkubernetesdashboard)***| ***(Optional)*** |
| `networkPolicyConfig` | ***[[]ContainerClusterSpecAddonsConfigNetworkPolicyConfig](#containerclusterspecaddonsconfignetworkpolicyconfig)***| ***(Optional)*** |
## ContainerClusterSpecAddonsConfigHorizontalPodAutoscaling

Appears on:[ContainerClusterSpecAddonsConfig](#containerclusterspecaddonsconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `disabled` | ***bool***| ***(Optional)*** |
## ContainerClusterSpecAddonsConfigHttpLoadBalancing

Appears on:[ContainerClusterSpecAddonsConfig](#containerclusterspecaddonsconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `disabled` | ***bool***| ***(Optional)*** |
## ContainerClusterSpecAddonsConfigKubernetesDashboard

Appears on:[ContainerClusterSpecAddonsConfig](#containerclusterspecaddonsconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `disabled` | ***bool***| ***(Optional)*** |
## ContainerClusterSpecAddonsConfigNetworkPolicyConfig

Appears on:[ContainerClusterSpecAddonsConfig](#containerclusterspecaddonsconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `disabled` | ***bool***| ***(Optional)*** |
## ContainerClusterSpecIpAllocationPolicy

Appears on:[ContainerClusterSpec](#containerclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `clusterIpv4CIDRBlock` | ***string***| ***(Optional)*** |
| `clusterSecondaryRangeName` | ***string***| ***(Optional)*** |
| `createSubnetwork` | ***bool***| ***(Optional)*** |
| `nodeIpv4CIDRBlock` | ***string***| ***(Optional)*** |
| `servicesIpv4CIDRBlock` | ***string***| ***(Optional)*** |
| `servicesSecondaryRangeName` | ***string***| ***(Optional)*** |
| `subnetworkName` | ***string***| ***(Optional)*** |
| `useIPAliases` | ***bool***| ***(Optional)*** |
## ContainerClusterSpecMaintenancePolicy

Appears on:[ContainerClusterSpec](#containerclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `dailyMaintenanceWindow` | ***[[]ContainerClusterSpecMaintenancePolicyDailyMaintenanceWindow](#containerclusterspecmaintenancepolicydailymaintenancewindow)***||
## ContainerClusterSpecMaintenancePolicyDailyMaintenanceWindow

Appears on:[ContainerClusterSpecMaintenancePolicy](#containerclusterspecmaintenancepolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `duration` | ***string***| ***(Optional)*** |
| `startTime` | ***string***||
## ContainerClusterSpecMasterAuth

Appears on:[ContainerClusterSpec](#containerclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `clientCertificate` | ***string***| ***(Optional)*** |
| `clientCertificateConfig` | ***[[]ContainerClusterSpecMasterAuthClientCertificateConfig](#containerclusterspecmasterauthclientcertificateconfig)***| ***(Optional)*** |
| `clusterCaCertificate` | ***string***| ***(Optional)*** |
| `username` | ***string***| ***(Optional)*** |
## ContainerClusterSpecMasterAuthClientCertificateConfig

Appears on:[ContainerClusterSpecMasterAuth](#containerclusterspecmasterauth)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `issueClientCertificate` | ***bool***||
## ContainerClusterSpecMasterAuthorizedNetworksConfig

Appears on:[ContainerClusterSpec](#containerclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cidrBlocks` | ***[[]ContainerClusterSpecMasterAuthorizedNetworksConfigCidrBlocks](#containerclusterspecmasterauthorizednetworksconfigcidrblocks)***| ***(Optional)*** |
## ContainerClusterSpecMasterAuthorizedNetworksConfigCidrBlocks

Appears on:[ContainerClusterSpecMasterAuthorizedNetworksConfig](#containerclusterspecmasterauthorizednetworksconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cidrBlock` | ***string***||
| `displayName` | ***string***| ***(Optional)*** |
## ContainerClusterSpecNetworkPolicy

Appears on:[ContainerClusterSpec](#containerclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `provider` | ***string***| ***(Optional)*** |
## ContainerClusterSpecNodeConfig

Appears on:[ContainerClusterSpec](#containerclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `diskSizeGb` | ***int64***| ***(Optional)*** |
| `diskType` | ***string***| ***(Optional)*** |
| `guestAccelerator` | ***[[]ContainerClusterSpecNodeConfigGuestAccelerator](#containerclusterspecnodeconfigguestaccelerator)***| ***(Optional)*** |
| `imageType` | ***string***| ***(Optional)*** |
| `labels` | ***map[string]string***| ***(Optional)*** |
| `localSsdCount` | ***int64***| ***(Optional)*** |
| `machineType` | ***string***| ***(Optional)*** |
| `metadata` | ***map[string]string***| ***(Optional)*** |
| `minCPUPlatform` | ***string***| ***(Optional)*** |
| `oauthScopes` | ***[]string***| ***(Optional)*** |
| `preemptible` | ***bool***| ***(Optional)*** |
| `serviceAccount` | ***string***| ***(Optional)*** |
| `tags` | ***[]string***| ***(Optional)*** |
## ContainerClusterSpecNodeConfigGuestAccelerator

Appears on:[ContainerClusterSpecNodeConfig](#containerclusterspecnodeconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int64***||
| `type` | ***string***||
## ContainerClusterSpecNodePool

Appears on:[ContainerClusterSpec](#containerclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `autoscaling` | ***[[]ContainerClusterSpecNodePoolAutoscaling](#containerclusterspecnodepoolautoscaling)***| ***(Optional)*** |
| `initialNodeCount` | ***int64***| ***(Optional)*** |
| `instanceGroupUrls` | ***[]string***| ***(Optional)*** |
| `management` | ***[[]ContainerClusterSpecNodePoolManagement](#containerclusterspecnodepoolmanagement)***| ***(Optional)*** |
| `maxPodsPerNode` | ***int64***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `nodeConfig` | ***[[]ContainerClusterSpecNodePoolNodeConfig](#containerclusterspecnodepoolnodeconfig)***| ***(Optional)*** |
| `nodeCount` | ***int64***| ***(Optional)*** |
| `version` | ***string***| ***(Optional)*** |
## ContainerClusterSpecNodePoolAutoscaling

Appears on:[ContainerClusterSpecNodePool](#containerclusterspecnodepool)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `maxNodeCount` | ***int64***||
| `minNodeCount` | ***int64***||
## ContainerClusterSpecNodePoolManagement

Appears on:[ContainerClusterSpecNodePool](#containerclusterspecnodepool)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `autoRepair` | ***bool***| ***(Optional)*** |
| `autoUpgrade` | ***bool***| ***(Optional)*** |
## ContainerClusterSpecNodePoolNodeConfig

Appears on:[ContainerClusterSpecNodePool](#containerclusterspecnodepool)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `diskSizeGb` | ***int64***| ***(Optional)*** |
| `diskType` | ***string***| ***(Optional)*** |
| `guestAccelerator` | ***[[]ContainerClusterSpecNodePoolNodeConfigGuestAccelerator](#containerclusterspecnodepoolnodeconfigguestaccelerator)***| ***(Optional)*** |
| `imageType` | ***string***| ***(Optional)*** |
| `labels` | ***map[string]string***| ***(Optional)*** |
| `localSsdCount` | ***int64***| ***(Optional)*** |
| `machineType` | ***string***| ***(Optional)*** |
| `metadata` | ***map[string]string***| ***(Optional)*** |
| `minCPUPlatform` | ***string***| ***(Optional)*** |
| `oauthScopes` | ***[]string***| ***(Optional)*** |
| `preemptible` | ***bool***| ***(Optional)*** |
| `serviceAccount` | ***string***| ***(Optional)*** |
| `tags` | ***[]string***| ***(Optional)*** |
## ContainerClusterSpecNodePoolNodeConfigGuestAccelerator

Appears on:[ContainerClusterSpecNodePoolNodeConfig](#containerclusterspecnodepoolnodeconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int64***||
| `type` | ***string***||
## ContainerClusterSpecPrivateClusterConfig

Appears on:[ContainerClusterSpec](#containerclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enablePrivateEndpoint` | ***bool***| ***(Optional)*** |
| `enablePrivateNodes` | ***bool***| ***(Optional)*** |
| `masterIpv4CIDRBlock` | ***string***| ***(Optional)*** |
| `privateEndpoint` | ***string***| ***(Optional)*** |
| `publicEndpoint` | ***string***| ***(Optional)*** |
## ContainerClusterStatus

Appears on:[ContainerCluster](#containercluster)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ContainerClusterSpec](#containerclusterspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ContainerClusterStatus](#containerclusterstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `master_auth.<index>.client_key` | ***string*** ||
| `master_auth.<index>.password` | ***string*** ||
