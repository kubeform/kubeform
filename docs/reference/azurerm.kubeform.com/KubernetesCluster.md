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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KubernetesClusterSpec](#kubernetesclusterspec)***||
| `status` | ***[KubernetesClusterStatus](#kubernetesclusterstatus)***||
## KubernetesClusterSpec

Appears on:[KubernetesCluster](#kubernetescluster), [KubernetesClusterStatus](#kubernetesclusterstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `addonProfile` | ***[[]KubernetesClusterSpecAddonProfile](#kubernetesclusterspecaddonprofile)***| ***(Optional)*** |
| `agentPoolProfile` | ***[[]KubernetesClusterSpecAgentPoolProfile](#kubernetesclusterspecagentpoolprofile)***| ***(Optional)*** Deprecated|
| `apiServerAuthorizedIPRanges` | ***[]string***| ***(Optional)*** |
| `defaultNodePool` | ***[[]KubernetesClusterSpecDefaultNodePool](#kubernetesclusterspecdefaultnodepool)***| ***(Optional)*** |
| `dnsPrefix` | ***string***||
| `enablePodSecurityPolicy` | ***bool***| ***(Optional)*** |
| `fqdn` | ***string***| ***(Optional)*** |
| `identity` | ***[[]KubernetesClusterSpecIdentity](#kubernetesclusterspecidentity)***| ***(Optional)*** |
| `kubeAdminConfig` | ***[[]KubernetesClusterSpecKubeAdminConfig](#kubernetesclusterspeckubeadminconfig)***| ***(Optional)*** |
| `kubeConfig` | ***[[]KubernetesClusterSpecKubeConfig](#kubernetesclusterspeckubeconfig)***| ***(Optional)*** |
| `kubernetesVersion` | ***string***| ***(Optional)*** |
| `linuxProfile` | ***[[]KubernetesClusterSpecLinuxProfile](#kubernetesclusterspeclinuxprofile)***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `networkProfile` | ***[[]KubernetesClusterSpecNetworkProfile](#kubernetesclusterspecnetworkprofile)***| ***(Optional)*** |
| `nodeResourceGroup` | ***string***| ***(Optional)*** |
| `privateFqdn` | ***string***| ***(Optional)*** |
| `privateLinkEnabled` | ***bool***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `roleBasedAccessControl` | ***[[]KubernetesClusterSpecRoleBasedAccessControl](#kubernetesclusterspecrolebasedaccesscontrol)***| ***(Optional)*** |
| `servicePrincipal` | ***[[]KubernetesClusterSpecServicePrincipal](#kubernetesclusterspecserviceprincipal)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `windowsProfile` | ***[[]KubernetesClusterSpecWindowsProfile](#kubernetesclusterspecwindowsprofile)***| ***(Optional)*** |
## KubernetesClusterSpecAddonProfile

Appears on:[KubernetesClusterSpec](#kubernetesclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `aciConnectorLinux` | ***[[]KubernetesClusterSpecAddonProfileAciConnectorLinux](#kubernetesclusterspecaddonprofileaciconnectorlinux)***| ***(Optional)*** |
| `azurePolicy` | ***[[]KubernetesClusterSpecAddonProfileAzurePolicy](#kubernetesclusterspecaddonprofileazurepolicy)***| ***(Optional)*** |
| `httpApplicationRouting` | ***[[]KubernetesClusterSpecAddonProfileHttpApplicationRouting](#kubernetesclusterspecaddonprofilehttpapplicationrouting)***| ***(Optional)*** |
| `kubeDashboard` | ***[[]KubernetesClusterSpecAddonProfileKubeDashboard](#kubernetesclusterspecaddonprofilekubedashboard)***| ***(Optional)*** |
| `omsAgent` | ***[[]KubernetesClusterSpecAddonProfileOmsAgent](#kubernetesclusterspecaddonprofileomsagent)***| ***(Optional)*** |
## KubernetesClusterSpecAddonProfileAciConnectorLinux

Appears on:[KubernetesClusterSpecAddonProfile](#kubernetesclusterspecaddonprofile)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
| `subnetName` | ***string***| ***(Optional)*** |
## KubernetesClusterSpecAddonProfileAzurePolicy

Appears on:[KubernetesClusterSpecAddonProfile](#kubernetesclusterspecaddonprofile)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
## KubernetesClusterSpecAddonProfileHttpApplicationRouting

Appears on:[KubernetesClusterSpecAddonProfile](#kubernetesclusterspecaddonprofile)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
| `httpApplicationRoutingZoneName` | ***string***| ***(Optional)*** |
## KubernetesClusterSpecAddonProfileKubeDashboard

Appears on:[KubernetesClusterSpecAddonProfile](#kubernetesclusterspecaddonprofile)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
## KubernetesClusterSpecAddonProfileOmsAgent

Appears on:[KubernetesClusterSpecAddonProfile](#kubernetesclusterspecaddonprofile)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
| `logAnalyticsWorkspaceID` | ***string***| ***(Optional)*** |
## KubernetesClusterSpecAgentPoolProfile

Appears on:[KubernetesClusterSpec](#kubernetesclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `availabilityZones` | ***[]string***| ***(Optional)*** |
| `count` | ***int64***| ***(Optional)*** |
| `dnsPrefix` | ***string***| ***(Optional)*** Deprecated|
| `enableAutoScaling` | ***bool***| ***(Optional)*** |
| `enableNodePublicIP` | ***bool***| ***(Optional)*** |
| `fqdn` | ***string***| ***(Optional)*** Deprecated|
| `maxCount` | ***int64***| ***(Optional)*** |
| `maxPods` | ***int64***| ***(Optional)*** |
| `minCount` | ***int64***| ***(Optional)*** |
| `name` | ***string***||
| `nodeTaints` | ***[]string***| ***(Optional)*** |
| `osDiskSizeGb` | ***int64***| ***(Optional)*** |
| `osType` | ***string***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
| `vnetSubnetID` | ***string***| ***(Optional)*** |
## KubernetesClusterSpecDefaultNodePool

Appears on:[KubernetesClusterSpec](#kubernetesclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `availabilityZones` | ***[]string***| ***(Optional)*** |
| `enableAutoScaling` | ***bool***| ***(Optional)*** |
| `enableNodePublicIP` | ***bool***| ***(Optional)*** |
| `maxCount` | ***int64***| ***(Optional)*** |
| `maxPods` | ***int64***| ***(Optional)*** |
| `minCount` | ***int64***| ***(Optional)*** |
| `name` | ***string***||
| `nodeCount` | ***int64***| ***(Optional)*** |
| `nodeTaints` | ***[]string***| ***(Optional)*** |
| `osDiskSizeGb` | ***int64***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
| `vnetSubnetID` | ***string***| ***(Optional)*** |
## KubernetesClusterSpecIdentity

Appears on:[KubernetesClusterSpec](#kubernetesclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `principalID` | ***string***| ***(Optional)*** |
| `tenantID` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## KubernetesClusterSpecKubeAdminConfig

Appears on:[KubernetesClusterSpec](#kubernetesclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `clientCertificate` | ***string***| ***(Optional)*** |
| `clusterCaCertificate` | ***string***| ***(Optional)*** |
| `host` | ***string***| ***(Optional)*** |
| `username` | ***string***| ***(Optional)*** |
## KubernetesClusterSpecKubeConfig

Appears on:[KubernetesClusterSpec](#kubernetesclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `clientCertificate` | ***string***| ***(Optional)*** |
| `clusterCaCertificate` | ***string***| ***(Optional)*** |
| `host` | ***string***| ***(Optional)*** |
| `username` | ***string***| ***(Optional)*** |
## KubernetesClusterSpecLinuxProfile

Appears on:[KubernetesClusterSpec](#kubernetesclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `adminUsername` | ***string***||
| `sshKey` | ***[[]KubernetesClusterSpecLinuxProfileSshKey](#kubernetesclusterspeclinuxprofilesshkey)***||
## KubernetesClusterSpecLinuxProfileSshKey

Appears on:[KubernetesClusterSpecLinuxProfile](#kubernetesclusterspeclinuxprofile)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `keyData` | ***string***||
## KubernetesClusterSpecNetworkProfile

Appears on:[KubernetesClusterSpec](#kubernetesclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `dnsServiceIP` | ***string***| ***(Optional)*** |
| `dockerBridgeCIDR` | ***string***| ***(Optional)*** |
| `loadBalancerProfile` | ***[[]KubernetesClusterSpecNetworkProfileLoadBalancerProfile](#kubernetesclusterspecnetworkprofileloadbalancerprofile)***| ***(Optional)*** |
| `loadBalancerSku` | ***string***| ***(Optional)*** |
| `networkPlugin` | ***string***||
| `networkPolicy` | ***string***| ***(Optional)*** |
| `podCIDR` | ***string***| ***(Optional)*** |
| `serviceCIDR` | ***string***| ***(Optional)*** |
## KubernetesClusterSpecNetworkProfileLoadBalancerProfile

Appears on:[KubernetesClusterSpecNetworkProfile](#kubernetesclusterspecnetworkprofile)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `effectiveOutboundIPS` | ***[]string***| ***(Optional)*** |
| `managedOutboundIPCount` | ***int64***| ***(Optional)*** |
| `outboundIPAddressIDS` | ***[]string***| ***(Optional)*** |
| `outboundIPPrefixIDS` | ***[]string***| ***(Optional)*** |
## KubernetesClusterSpecRoleBasedAccessControl

Appears on:[KubernetesClusterSpec](#kubernetesclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `azureActiveDirectory` | ***[[]KubernetesClusterSpecRoleBasedAccessControlAzureActiveDirectory](#kubernetesclusterspecrolebasedaccesscontrolazureactivedirectory)***| ***(Optional)*** |
| `enabled` | ***bool***||
## KubernetesClusterSpecRoleBasedAccessControlAzureActiveDirectory

Appears on:[KubernetesClusterSpecRoleBasedAccessControl](#kubernetesclusterspecrolebasedaccesscontrol)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `clientAppID` | ***string***||
| `serverAppID` | ***string***||
| `tenantID` | ***string***| ***(Optional)*** |
## KubernetesClusterSpecServicePrincipal

Appears on:[KubernetesClusterSpec](#kubernetesclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `clientID` | ***string***||
## KubernetesClusterSpecWindowsProfile

Appears on:[KubernetesClusterSpec](#kubernetesclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `adminUsername` | ***string***||
## KubernetesClusterStatus

Appears on:[KubernetesCluster](#kubernetescluster)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KubernetesClusterSpec](#kubernetesclusterspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[KubernetesClusterStatus](#kubernetesclusterstatus)

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
