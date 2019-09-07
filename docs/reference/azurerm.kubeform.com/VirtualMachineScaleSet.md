---
title: VirtualMachineScaleSet
menu:
  docs_v0.0.1:
    identifier: virtualmachinescaleset-azurerm.kubeform.com
    name: VirtualMachineScaleSet
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## VirtualMachineScaleSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `VirtualMachineScaleSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VirtualMachineScaleSetSpec](#VirtualMachineScaleSetSpec)***||
| `status` | ***[VirtualMachineScaleSetStatus](#VirtualMachineScaleSetStatus)***||
## VirtualMachineScaleSetSpec
##### (Appears on:[VirtualMachineScaleSet](#VirtualMachineScaleSet), [VirtualMachineScaleSetStatus](#VirtualMachineScaleSetStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `automaticOsUpgrade` | ***bool***| ***(Optional)*** |
| `bootDiagnostics` | ***[[]VirtualMachineScaleSetSpecBootDiagnostics](#VirtualMachineScaleSetSpecBootDiagnostics)***| ***(Optional)*** |
| `evictionPolicy` | ***string***| ***(Optional)*** |
| `extension` | ***[[]VirtualMachineScaleSetSpecExtension](#VirtualMachineScaleSetSpecExtension)***| ***(Optional)*** |
| `healthProbeID` | ***string***| ***(Optional)*** |
| `identity` | ***[[]VirtualMachineScaleSetSpecIdentity](#VirtualMachineScaleSetSpecIdentity)***| ***(Optional)*** |
| `licenseType` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `networkProfile` | ***[[]VirtualMachineScaleSetSpecNetworkProfile](#VirtualMachineScaleSetSpecNetworkProfile)***||
| `osProfile` | ***[[]VirtualMachineScaleSetSpecOsProfile](#VirtualMachineScaleSetSpecOsProfile)***||
| `osProfileLinuxConfig` | ***[[]VirtualMachineScaleSetSpecOsProfileLinuxConfig](#VirtualMachineScaleSetSpecOsProfileLinuxConfig)***| ***(Optional)*** |
| `osProfileSecrets` | ***[[]VirtualMachineScaleSetSpecOsProfileSecrets](#VirtualMachineScaleSetSpecOsProfileSecrets)***| ***(Optional)*** |
| `osProfileWindowsConfig` | ***[[]VirtualMachineScaleSetSpecOsProfileWindowsConfig](#VirtualMachineScaleSetSpecOsProfileWindowsConfig)***| ***(Optional)*** |
| `overprovision` | ***bool***| ***(Optional)*** |
| `plan` | ***[[]VirtualMachineScaleSetSpecPlan](#VirtualMachineScaleSetSpecPlan)***| ***(Optional)*** |
| `priority` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `rollingUpgradePolicy` | ***[[]VirtualMachineScaleSetSpecRollingUpgradePolicy](#VirtualMachineScaleSetSpecRollingUpgradePolicy)***| ***(Optional)*** |
| `singlePlacementGroup` | ***bool***| ***(Optional)*** |
| `sku` | ***[[]VirtualMachineScaleSetSpecSku](#VirtualMachineScaleSetSpecSku)***||
| `storageProfileDataDisk` | ***[[]VirtualMachineScaleSetSpecStorageProfileDataDisk](#VirtualMachineScaleSetSpecStorageProfileDataDisk)***| ***(Optional)*** |
| `storageProfileImageReference` | ***[[]VirtualMachineScaleSetSpecStorageProfileImageReference](#VirtualMachineScaleSetSpecStorageProfileImageReference)***| ***(Optional)*** |
| `storageProfileOsDisk` | ***[[]VirtualMachineScaleSetSpecStorageProfileOsDisk](#VirtualMachineScaleSetSpecStorageProfileOsDisk)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `upgradePolicyMode` | ***string***||
| `zones` | ***[]string***| ***(Optional)*** |
## VirtualMachineScaleSetSpecBootDiagnostics
##### (Appears on:[VirtualMachineScaleSetSpec](#VirtualMachineScaleSetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `storageURI` | ***string***||
## VirtualMachineScaleSetSpecExtension
##### (Appears on:[VirtualMachineScaleSetSpec](#VirtualMachineScaleSetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `autoUpgradeMinorVersion` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `provisionAfterExtensions` | ***[]string***| ***(Optional)*** |
| `publisher` | ***string***||
| `settings` | ***string***| ***(Optional)*** |
| `type` | ***string***||
| `typeHandlerVersion` | ***string***||
## VirtualMachineScaleSetSpecIdentity
##### (Appears on:[VirtualMachineScaleSetSpec](#VirtualMachineScaleSetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `identityIDS` | ***[]string***| ***(Optional)*** |
| `principalID` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## VirtualMachineScaleSetSpecNetworkProfile
##### (Appears on:[VirtualMachineScaleSetSpec](#VirtualMachineScaleSetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `acceleratedNetworking` | ***bool***| ***(Optional)*** |
| `dnsSettings` | ***[[]VirtualMachineScaleSetSpecNetworkProfileDnsSettings](#VirtualMachineScaleSetSpecNetworkProfileDnsSettings)***| ***(Optional)*** |
| `ipConfiguration` | ***[[]VirtualMachineScaleSetSpecNetworkProfileIpConfiguration](#VirtualMachineScaleSetSpecNetworkProfileIpConfiguration)***||
| `ipForwarding` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `networkSecurityGroupID` | ***string***| ***(Optional)*** |
| `primary` | ***bool***||
## VirtualMachineScaleSetSpecNetworkProfileDnsSettings
##### (Appears on:[VirtualMachineScaleSetSpecNetworkProfile](#VirtualMachineScaleSetSpecNetworkProfile))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `dnsServers` | ***[]string***||
## VirtualMachineScaleSetSpecNetworkProfileIpConfiguration
##### (Appears on:[VirtualMachineScaleSetSpecNetworkProfile](#VirtualMachineScaleSetSpecNetworkProfile))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `applicationGatewayBackendAddressPoolIDS` | ***[]string***| ***(Optional)*** |
| `applicationSecurityGroupIDS` | ***[]string***| ***(Optional)*** |
| `loadBalancerBackendAddressPoolIDS` | ***[]string***| ***(Optional)*** |
| `loadBalancerInboundNATRulesIDS` | ***[]string***| ***(Optional)*** |
| `name` | ***string***||
| `primary` | ***bool***||
| `publicIPAddressConfiguration` | ***[[]VirtualMachineScaleSetSpecNetworkProfileIpConfigurationPublicIPAddressConfiguration](#VirtualMachineScaleSetSpecNetworkProfileIpConfigurationPublicIPAddressConfiguration)***| ***(Optional)*** |
| `subnetID` | ***string***||
## VirtualMachineScaleSetSpecNetworkProfileIpConfigurationPublicIPAddressConfiguration
##### (Appears on:[VirtualMachineScaleSetSpecNetworkProfileIpConfiguration](#VirtualMachineScaleSetSpecNetworkProfileIpConfiguration))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `domainNameLabel` | ***string***||
| `idleTimeout` | ***int***||
| `name` | ***string***||
## VirtualMachineScaleSetSpecOsProfile
##### (Appears on:[VirtualMachineScaleSetSpec](#VirtualMachineScaleSetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `adminUsername` | ***string***||
| `computerNamePrefix` | ***string***||
| `customData` | ***string***| ***(Optional)*** |
## VirtualMachineScaleSetSpecOsProfileLinuxConfig
##### (Appears on:[VirtualMachineScaleSetSpec](#VirtualMachineScaleSetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `disablePasswordAuthentication` | ***bool***| ***(Optional)*** |
| `sshKeys` | ***[[]VirtualMachineScaleSetSpecOsProfileLinuxConfigSshKeys](#VirtualMachineScaleSetSpecOsProfileLinuxConfigSshKeys)***| ***(Optional)*** |
## VirtualMachineScaleSetSpecOsProfileLinuxConfigSshKeys
##### (Appears on:[VirtualMachineScaleSetSpecOsProfileLinuxConfig](#VirtualMachineScaleSetSpecOsProfileLinuxConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `keyData` | ***string***| ***(Optional)*** |
| `path` | ***string***||
## VirtualMachineScaleSetSpecOsProfileSecrets
##### (Appears on:[VirtualMachineScaleSetSpec](#VirtualMachineScaleSetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `sourceVaultID` | ***string***||
| `vaultCertificates` | ***[[]VirtualMachineScaleSetSpecOsProfileSecretsVaultCertificates](#VirtualMachineScaleSetSpecOsProfileSecretsVaultCertificates)***| ***(Optional)*** |
## VirtualMachineScaleSetSpecOsProfileSecretsVaultCertificates
##### (Appears on:[VirtualMachineScaleSetSpecOsProfileSecrets](#VirtualMachineScaleSetSpecOsProfileSecrets))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `certificateStore` | ***string***| ***(Optional)*** |
| `certificateURL` | ***string***||
## VirtualMachineScaleSetSpecOsProfileWindowsConfig
##### (Appears on:[VirtualMachineScaleSetSpec](#VirtualMachineScaleSetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `additionalUnattendConfig` | ***[[]VirtualMachineScaleSetSpecOsProfileWindowsConfigAdditionalUnattendConfig](#VirtualMachineScaleSetSpecOsProfileWindowsConfigAdditionalUnattendConfig)***| ***(Optional)*** |
| `enableAutomaticUpgrades` | ***bool***| ***(Optional)*** |
| `provisionVmAgent` | ***bool***| ***(Optional)*** |
| `winrm` | ***[[]VirtualMachineScaleSetSpecOsProfileWindowsConfigWinrm](#VirtualMachineScaleSetSpecOsProfileWindowsConfigWinrm)***| ***(Optional)*** |
## VirtualMachineScaleSetSpecOsProfileWindowsConfigAdditionalUnattendConfig
##### (Appears on:[VirtualMachineScaleSetSpecOsProfileWindowsConfig](#VirtualMachineScaleSetSpecOsProfileWindowsConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `component` | ***string***||
| `pass` | ***string***||
| `settingName` | ***string***||
## VirtualMachineScaleSetSpecOsProfileWindowsConfigWinrm
##### (Appears on:[VirtualMachineScaleSetSpecOsProfileWindowsConfig](#VirtualMachineScaleSetSpecOsProfileWindowsConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `certificateURL` | ***string***| ***(Optional)*** |
| `protocol` | ***string***||
## VirtualMachineScaleSetSpecPlan
##### (Appears on:[VirtualMachineScaleSetSpec](#VirtualMachineScaleSetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `product` | ***string***||
| `publisher` | ***string***||
## VirtualMachineScaleSetSpecRollingUpgradePolicy
##### (Appears on:[VirtualMachineScaleSetSpec](#VirtualMachineScaleSetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `maxBatchInstancePercent` | ***int***| ***(Optional)*** |
| `maxUnhealthyInstancePercent` | ***int***| ***(Optional)*** |
| `maxUnhealthyUpgradedInstancePercent` | ***int***| ***(Optional)*** |
| `pauseTimeBetweenBatches` | ***string***| ***(Optional)*** |
## VirtualMachineScaleSetSpecSku
##### (Appears on:[VirtualMachineScaleSetSpec](#VirtualMachineScaleSetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `capacity` | ***int***||
| `name` | ***string***||
| `tier` | ***string***| ***(Optional)*** |
## VirtualMachineScaleSetSpecStorageProfileDataDisk
##### (Appears on:[VirtualMachineScaleSetSpec](#VirtualMachineScaleSetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `caching` | ***string***| ***(Optional)*** |
| `createOption` | ***string***||
| `diskSizeGb` | ***int***| ***(Optional)*** |
| `lun` | ***int***||
| `managedDiskType` | ***string***| ***(Optional)*** |
## VirtualMachineScaleSetSpecStorageProfileImageReference
##### (Appears on:[VirtualMachineScaleSetSpec](#VirtualMachineScaleSetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `offer` | ***string***| ***(Optional)*** |
| `publisher` | ***string***| ***(Optional)*** |
| `sku` | ***string***| ***(Optional)*** |
| `version` | ***string***| ***(Optional)*** |
## VirtualMachineScaleSetSpecStorageProfileOsDisk
##### (Appears on:[VirtualMachineScaleSetSpec](#VirtualMachineScaleSetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `caching` | ***string***| ***(Optional)*** |
| `createOption` | ***string***||
| `image` | ***string***| ***(Optional)*** |
| `managedDiskType` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `osType` | ***string***| ***(Optional)*** |
| `vhdContainers` | ***[]string***| ***(Optional)*** |
## VirtualMachineScaleSetStatus
##### (Appears on:[VirtualMachineScaleSet](#VirtualMachineScaleSet))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VirtualMachineScaleSetSpec](#VirtualMachineScaleSetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `extension.<index>.protected_settings` | ***string*** ||
| `os_profile.<index>.admin_password` | ***string*** ||
| `os_profile_windows_config.<index>.additional_unattend_config.<index>.content` | ***string*** ||
