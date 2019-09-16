---
title: KubernetesCluster
menu:
  docs_{{ .version }}:
    identifier: kubernetescluster-azurerm.kubeform.com
    name: KubernetesCluster
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## KubernetesCluster
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `KubernetesCluster` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KubernetesClusterSpec](#KubernetesClusterSpec)***||
| `status` | ***[KubernetesClusterStatus](#KubernetesClusterStatus)***||
## KubernetesClusterSpec

Appears on:[KubernetesCluster](#KubernetesCluster), [KubernetesClusterStatus](#KubernetesClusterStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `addonProfile` | ***[[]KubernetesClusterSpecAddonProfile](#KubernetesClusterSpecAddonProfile)***| ***(Optional)*** |
| `agentPoolProfile` | ***[[]KubernetesClusterSpecAgentPoolProfile](#KubernetesClusterSpecAgentPoolProfile)***||
| `apiServerAuthorizedIPRanges` | ***[]string***| ***(Optional)*** |
| `dnsPrefix` | ***string***||
| `fqdn` | ***string***| ***(Optional)*** |
| `kubeAdminConfig` | ***[[]KubernetesClusterSpecKubeAdminConfig](#KubernetesClusterSpecKubeAdminConfig)***| ***(Optional)*** |
| `kubeConfig` | ***[[]KubernetesClusterSpecKubeConfig](#KubernetesClusterSpecKubeConfig)***| ***(Optional)*** |
| `kubernetesVersion` | ***string***| ***(Optional)*** |
| `linuxProfile` | ***[[]KubernetesClusterSpecLinuxProfile](#KubernetesClusterSpecLinuxProfile)***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `networkProfile` | ***[[]KubernetesClusterSpecNetworkProfile](#KubernetesClusterSpecNetworkProfile)***| ***(Optional)*** |
| `nodeResourceGroup` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `roleBasedAccessControl` | ***[[]KubernetesClusterSpecRoleBasedAccessControl](#KubernetesClusterSpecRoleBasedAccessControl)***| ***(Optional)*** |
| `servicePrincipal` | ***[[]KubernetesClusterSpecServicePrincipal](#KubernetesClusterSpecServicePrincipal)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `windowsProfile` | ***[[]KubernetesClusterSpecWindowsProfile](#KubernetesClusterSpecWindowsProfile)***| ***(Optional)*** |
## KubernetesClusterSpecAddonProfile

Appears on:[KubernetesClusterSpec](#KubernetesClusterSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `aciConnectorLinux` | ***[[]KubernetesClusterSpecAddonProfileAciConnectorLinux](#KubernetesClusterSpecAddonProfileAciConnectorLinux)***| ***(Optional)*** |
| `httpApplicationRouting` | ***[[]KubernetesClusterSpecAddonProfileHttpApplicationRouting](#KubernetesClusterSpecAddonProfileHttpApplicationRouting)***| ***(Optional)*** |
| `omsAgent` | ***[[]KubernetesClusterSpecAddonProfileOmsAgent](#KubernetesClusterSpecAddonProfileOmsAgent)***| ***(Optional)*** |
## KubernetesClusterSpecAddonProfileAciConnectorLinux

Appears on:[KubernetesClusterSpecAddonProfile](#KubernetesClusterSpecAddonProfile)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
| `subnetName` | ***string***||
## KubernetesClusterSpecAddonProfileHttpApplicationRouting

Appears on:[KubernetesClusterSpecAddonProfile](#KubernetesClusterSpecAddonProfile)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
| `httpApplicationRoutingZoneName` | ***string***| ***(Optional)*** |
## KubernetesClusterSpecAddonProfileOmsAgent

Appears on:[KubernetesClusterSpecAddonProfile](#KubernetesClusterSpecAddonProfile)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
| `logAnalyticsWorkspaceID` | ***string***||
## KubernetesClusterSpecAgentPoolProfile

Appears on:[KubernetesClusterSpec](#KubernetesClusterSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `availabilityZones` | ***[]string***| ***(Optional)*** |
| `count` | ***int***| ***(Optional)*** |
| `dnsPrefix` | ***string***| ***(Optional)*** Deprecated|
| `enableAutoScaling` | ***bool***| ***(Optional)*** |
| `fqdn` | ***string***| ***(Optional)*** Deprecated|
| `maxCount` | ***int***| ***(Optional)*** |
| `maxPods` | ***int***| ***(Optional)*** |
| `minCount` | ***int***| ***(Optional)*** |
| `name` | ***string***||
| `nodeTaints` | ***[]string***| ***(Optional)*** |
| `osDiskSizeGb` | ***int***| ***(Optional)*** |
| `osType` | ***string***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
| `vnetSubnetID` | ***string***| ***(Optional)*** |
## KubernetesClusterSpecKubeAdminConfig

Appears on:[KubernetesClusterSpec](#KubernetesClusterSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `clientCertificate` | ***string***| ***(Optional)*** |
| `clusterCaCertificate` | ***string***| ***(Optional)*** |
| `host` | ***string***| ***(Optional)*** |
| `username` | ***string***| ***(Optional)*** |
## KubernetesClusterSpecKubeConfig

Appears on:[KubernetesClusterSpec](#KubernetesClusterSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `clientCertificate` | ***string***| ***(Optional)*** |
| `clusterCaCertificate` | ***string***| ***(Optional)*** |
| `host` | ***string***| ***(Optional)*** |
| `username` | ***string***| ***(Optional)*** |
## KubernetesClusterSpecLinuxProfile

Appears on:[KubernetesClusterSpec](#KubernetesClusterSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `adminUsername` | ***string***||
| `sshKey` | ***[[]KubernetesClusterSpecLinuxProfileSshKey](#KubernetesClusterSpecLinuxProfileSshKey)***||
## KubernetesClusterSpecLinuxProfileSshKey

Appears on:[KubernetesClusterSpecLinuxProfile](#KubernetesClusterSpecLinuxProfile)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `keyData` | ***string***||
## KubernetesClusterSpecNetworkProfile

Appears on:[KubernetesClusterSpec](#KubernetesClusterSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `dnsServiceIP` | ***string***| ***(Optional)*** |
| `dockerBridgeCIDR` | ***string***| ***(Optional)*** |
| `loadBalancerSku` | ***string***| ***(Optional)*** |
| `networkPlugin` | ***string***||
| `networkPolicy` | ***string***| ***(Optional)*** |
| `podCIDR` | ***string***| ***(Optional)*** |
| `serviceCIDR` | ***string***| ***(Optional)*** |
## KubernetesClusterSpecRoleBasedAccessControl

Appears on:[KubernetesClusterSpec](#KubernetesClusterSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `azureActiveDirectory` | ***[[]KubernetesClusterSpecRoleBasedAccessControlAzureActiveDirectory](#KubernetesClusterSpecRoleBasedAccessControlAzureActiveDirectory)***| ***(Optional)*** |
| `enabled` | ***bool***||
## KubernetesClusterSpecRoleBasedAccessControlAzureActiveDirectory

Appears on:[KubernetesClusterSpecRoleBasedAccessControl](#KubernetesClusterSpecRoleBasedAccessControl)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `clientAppID` | ***string***||
| `serverAppID` | ***string***||
| `tenantID` | ***string***| ***(Optional)*** |
## KubernetesClusterSpecServicePrincipal

Appears on:[KubernetesClusterSpec](#KubernetesClusterSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `clientID` | ***string***||
## KubernetesClusterSpecWindowsProfile

Appears on:[KubernetesClusterSpec](#KubernetesClusterSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `adminUsername` | ***string***||
## KubernetesClusterStatus

Appears on:[KubernetesCluster](#KubernetesCluster)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KubernetesClusterSpec](#KubernetesClusterSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `kube_admin_config.<index>.client_key` | ***string*** ||
| `kube_admin_config.<index>.password` | ***string*** ||
| `kube_admin_config_raw` | ***string*** ||
| `kube_config.<index>.client_key` | ***string*** ||
| `kube_config.<index>.password` | ***string*** ||
| `kube_config_raw` | ***string*** ||
| `role_based_access_control.<index>.azure_active_directory.<index>.server_app_secret` | ***string*** ||
| `service_principal.<index>.client_secret` | ***string*** ||
| `windows_profile.<index>.admin_password` | ***string*** ||
