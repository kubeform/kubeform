---
title: VirtualMachineScaleSet
menu:
  docs_{{ .version }}:
    identifier: virtualmachinescaleset-azurerm.kubeform.com
    name: VirtualMachineScaleSet
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## VirtualMachineScaleSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `VirtualMachineScaleSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VirtualMachineScaleSetSpec](#virtualmachinescalesetspec)***||
| `status` | ***[VirtualMachineScaleSetStatus](#virtualmachinescalesetstatus)***||
## Phase(`string` alias)

Appears on:[VirtualMachineScaleSetStatus](#virtualmachinescalesetstatus)

## VirtualMachineScaleSetSpec

Appears on:[VirtualMachineScaleSet](#virtualmachinescaleset), [VirtualMachineScaleSetStatus](#virtualmachinescalesetstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `automaticOsUpgrade` | ***bool***| ***(Optional)*** |
| `bootDiagnostics` | ***[[]VirtualMachineScaleSetSpecBootDiagnostics](#virtualmachinescalesetspecbootdiagnostics)***| ***(Optional)*** |
| `evictionPolicy` | ***string***| ***(Optional)*** |
| `extension` | ***[[]VirtualMachineScaleSetSpecExtension](#virtualmachinescalesetspecextension)***| ***(Optional)*** |
| `healthProbeID` | ***string***| ***(Optional)*** |
| `identity` | ***[[]VirtualMachineScaleSetSpecIdentity](#virtualmachinescalesetspecidentity)***| ***(Optional)*** |
| `licenseType` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `networkProfile` | ***[[]VirtualMachineScaleSetSpecNetworkProfile](#virtualmachinescalesetspecnetworkprofile)***||
| `osProfile` | ***[[]VirtualMachineScaleSetSpecOsProfile](#virtualmachinescalesetspecosprofile)***||
| `osProfileLinuxConfig` | ***[[]VirtualMachineScaleSetSpecOsProfileLinuxConfig](#virtualmachinescalesetspecosprofilelinuxconfig)***| ***(Optional)*** |
| `osProfileSecrets` | ***[[]VirtualMachineScaleSetSpecOsProfileSecrets](#virtualmachinescalesetspecosprofilesecrets)***| ***(Optional)*** |
| `osProfileWindowsConfig` | ***[[]VirtualMachineScaleSetSpecOsProfileWindowsConfig](#virtualmachinescalesetspecosprofilewindowsconfig)***| ***(Optional)*** |
| `overprovision` | ***bool***| ***(Optional)*** |
| `plan` | ***[[]VirtualMachineScaleSetSpecPlan](#virtualmachinescalesetspecplan)***| ***(Optional)*** |
| `priority` | ***string***| ***(Optional)*** |
| `proximityPlacementGroupID` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `rollingUpgradePolicy` | ***[[]VirtualMachineScaleSetSpecRollingUpgradePolicy](#virtualmachinescalesetspecrollingupgradepolicy)***| ***(Optional)*** |
| `singlePlacementGroup` | ***bool***| ***(Optional)*** |
| `sku` | ***[[]VirtualMachineScaleSetSpecSku](#virtualmachinescalesetspecsku)***||
| `storageProfileDataDisk` | ***[[]VirtualMachineScaleSetSpecStorageProfileDataDisk](#virtualmachinescalesetspecstorageprofiledatadisk)***| ***(Optional)*** |
| `storageProfileImageReference` | ***[[]VirtualMachineScaleSetSpecStorageProfileImageReference](#virtualmachinescalesetspecstorageprofileimagereference)***| ***(Optional)*** |
| `storageProfileOsDisk` | ***[[]VirtualMachineScaleSetSpecStorageProfileOsDisk](#virtualmachinescalesetspecstorageprofileosdisk)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `upgradePolicyMode` | ***string***||
| `zones` | ***[]string***| ***(Optional)*** |
## VirtualMachineScaleSetSpecBootDiagnostics

Appears on:[VirtualMachineScaleSetSpec](#virtualmachinescalesetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `storageURI` | ***string***||
## VirtualMachineScaleSetSpecExtension

Appears on:[VirtualMachineScaleSetSpec](#virtualmachinescalesetspec)

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

Appears on:[VirtualMachineScaleSetSpec](#virtualmachinescalesetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `identityIDS` | ***[]string***| ***(Optional)*** |
| `principalID` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## VirtualMachineScaleSetSpecNetworkProfile

Appears on:[VirtualMachineScaleSetSpec](#virtualmachinescalesetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `acceleratedNetworking` | ***bool***| ***(Optional)*** |
| `dnsSettings` | ***[[]VirtualMachineScaleSetSpecNetworkProfileDnsSettings](#virtualmachinescalesetspecnetworkprofilednssettings)***| ***(Optional)*** |
| `ipConfiguration` | ***[[]VirtualMachineScaleSetSpecNetworkProfileIpConfiguration](#virtualmachinescalesetspecnetworkprofileipconfiguration)***||
| `ipForwarding` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `networkSecurityGroupID` | ***string***| ***(Optional)*** |
| `primary` | ***bool***||
## VirtualMachineScaleSetSpecNetworkProfileDnsSettings

Appears on:[VirtualMachineScaleSetSpecNetworkProfile](#virtualmachinescalesetspecnetworkprofile)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `dnsServers` | ***[]string***||
## VirtualMachineScaleSetSpecNetworkProfileIpConfiguration

Appears on:[VirtualMachineScaleSetSpecNetworkProfile](#virtualmachinescalesetspecnetworkprofile)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `applicationGatewayBackendAddressPoolIDS` | ***[]string***| ***(Optional)*** |
| `applicationSecurityGroupIDS` | ***[]string***| ***(Optional)*** |
| `loadBalancerBackendAddressPoolIDS` | ***[]string***| ***(Optional)*** |
| `loadBalancerInboundNATRulesIDS` | ***[]string***| ***(Optional)*** |
| `name` | ***string***||
| `primary` | ***bool***||
| `publicIPAddressConfiguration` | ***[[]VirtualMachineScaleSetSpecNetworkProfileIpConfigurationPublicIPAddressConfiguration](#virtualmachinescalesetspecnetworkprofileipconfigurationpublicipaddressconfiguration)***| ***(Optional)*** |
| `subnetID` | ***string***||
## VirtualMachineScaleSetSpecNetworkProfileIpConfigurationPublicIPAddressConfiguration

Appears on:[VirtualMachineScaleSetSpecNetworkProfileIpConfiguration](#virtualmachinescalesetspecnetworkprofileipconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `domainNameLabel` | ***string***||
| `idleTimeout` | ***int64***||
| `name` | ***string***||
## VirtualMachineScaleSetSpecOsProfile

Appears on:[VirtualMachineScaleSetSpec](#virtualmachinescalesetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `adminUsername` | ***string***||
| `computerNamePrefix` | ***string***||
| `customData` | ***string***| ***(Optional)*** |
## VirtualMachineScaleSetSpecOsProfileLinuxConfig

Appears on:[VirtualMachineScaleSetSpec](#virtualmachinescalesetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `disablePasswordAuthentication` | ***bool***| ***(Optional)*** |
| `sshKeys` | ***[[]VirtualMachineScaleSetSpecOsProfileLinuxConfigSshKeys](#virtualmachinescalesetspecosprofilelinuxconfigsshkeys)***| ***(Optional)*** |
## VirtualMachineScaleSetSpecOsProfileLinuxConfigSshKeys

Appears on:[VirtualMachineScaleSetSpecOsProfileLinuxConfig](#virtualmachinescalesetspecosprofilelinuxconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `keyData` | ***string***| ***(Optional)*** |
| `path` | ***string***||
## VirtualMachineScaleSetSpecOsProfileSecrets

Appears on:[VirtualMachineScaleSetSpec](#virtualmachinescalesetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `sourceVaultID` | ***string***||
| `vaultCertificates` | ***[[]VirtualMachineScaleSetSpecOsProfileSecretsVaultCertificates](#virtualmachinescalesetspecosprofilesecretsvaultcertificates)***| ***(Optional)*** |
## VirtualMachineScaleSetSpecOsProfileSecretsVaultCertificates

Appears on:[VirtualMachineScaleSetSpecOsProfileSecrets](#virtualmachinescalesetspecosprofilesecrets)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `certificateStore` | ***string***| ***(Optional)*** |
| `certificateURL` | ***string***||
## VirtualMachineScaleSetSpecOsProfileWindowsConfig

Appears on:[VirtualMachineScaleSetSpec](#virtualmachinescalesetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `additionalUnattendConfig` | ***[[]VirtualMachineScaleSetSpecOsProfileWindowsConfigAdditionalUnattendConfig](#virtualmachinescalesetspecosprofilewindowsconfigadditionalunattendconfig)***| ***(Optional)*** |
| `enableAutomaticUpgrades` | ***bool***| ***(Optional)*** |
| `provisionVmAgent` | ***bool***| ***(Optional)*** |
| `winrm` | ***[[]VirtualMachineScaleSetSpecOsProfileWindowsConfigWinrm](#virtualmachinescalesetspecosprofilewindowsconfigwinrm)***| ***(Optional)*** |
## VirtualMachineScaleSetSpecOsProfileWindowsConfigAdditionalUnattendConfig

Appears on:[VirtualMachineScaleSetSpecOsProfileWindowsConfig](#virtualmachinescalesetspecosprofilewindowsconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `component` | ***string***||
| `pass` | ***string***||
| `settingName` | ***string***||
## VirtualMachineScaleSetSpecOsProfileWindowsConfigWinrm

Appears on:[VirtualMachineScaleSetSpecOsProfileWindowsConfig](#virtualmachinescalesetspecosprofilewindowsconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `certificateURL` | ***string***| ***(Optional)*** |
| `protocol` | ***string***||
## VirtualMachineScaleSetSpecPlan

Appears on:[VirtualMachineScaleSetSpec](#virtualmachinescalesetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `product` | ***string***||
| `publisher` | ***string***||
## VirtualMachineScaleSetSpecRollingUpgradePolicy

Appears on:[VirtualMachineScaleSetSpec](#virtualmachinescalesetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `maxBatchInstancePercent` | ***int64***| ***(Optional)*** |
| `maxUnhealthyInstancePercent` | ***int64***| ***(Optional)*** |
| `maxUnhealthyUpgradedInstancePercent` | ***int64***| ***(Optional)*** |
| `pauseTimeBetweenBatches` | ***string***| ***(Optional)*** |
## VirtualMachineScaleSetSpecSku

Appears on:[VirtualMachineScaleSetSpec](#virtualmachinescalesetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `capacity` | ***int64***||
| `name` | ***string***||
| `tier` | ***string***| ***(Optional)*** |
## VirtualMachineScaleSetSpecStorageProfileDataDisk

Appears on:[VirtualMachineScaleSetSpec](#virtualmachinescalesetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `caching` | ***string***| ***(Optional)*** |
| `createOption` | ***string***||
| `diskSizeGb` | ***int64***| ***(Optional)*** |
| `lun` | ***int64***||
| `managedDiskType` | ***string***| ***(Optional)*** |
## VirtualMachineScaleSetSpecStorageProfileImageReference

Appears on:[VirtualMachineScaleSetSpec](#virtualmachinescalesetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `offer` | ***string***| ***(Optional)*** |
| `publisher` | ***string***| ***(Optional)*** |
| `sku` | ***string***| ***(Optional)*** |
| `version` | ***string***| ***(Optional)*** |
## VirtualMachineScaleSetSpecStorageProfileOsDisk

Appears on:[VirtualMachineScaleSetSpec](#virtualmachinescalesetspec)

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

Appears on:[VirtualMachineScaleSet](#virtualmachinescaleset)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VirtualMachineScaleSetSpec](#virtualmachinescalesetspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `extension.<index>.protected_settings` | ***string*** ||
| `os_profile.<index>.admin_password` | ***string*** ||
| `os_profile_windows_config.<index>.additional_unattend_config.<index>.content` | ***string*** ||
