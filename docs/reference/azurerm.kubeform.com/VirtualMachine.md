---
title: VirtualMachine
menu:
  docs_{{ .version }}:
    identifier: virtualmachine-azurerm.kubeform.com
    name: VirtualMachine
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## VirtualMachine
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `VirtualMachine` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VirtualMachineSpec](#virtualmachinespec)***||
| `status` | ***[VirtualMachineStatus](#virtualmachinestatus)***||
## Phase(`string` alias)

Appears on:[VirtualMachineStatus](#virtualmachinestatus)

## VirtualMachineSpec

Appears on:[VirtualMachine](#virtualmachine), [VirtualMachineStatus](#virtualmachinestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `availabilitySetID` | ***string***| ***(Optional)*** |
| `bootDiagnostics` | ***[[]VirtualMachineSpecBootDiagnostics](#virtualmachinespecbootdiagnostics)***| ***(Optional)*** |
| `deleteDataDisksOnTermination` | ***bool***| ***(Optional)*** |
| `deleteOsDiskOnTermination` | ***bool***| ***(Optional)*** |
| `identity` | ***[[]VirtualMachineSpecIdentity](#virtualmachinespecidentity)***| ***(Optional)*** |
| `licenseType` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `networkInterfaceIDS` | ***[]string***||
| `osProfile` | ***[[]VirtualMachineSpecOsProfile](#virtualmachinespecosprofile)***| ***(Optional)*** |
| `osProfileLinuxConfig` | ***[[]VirtualMachineSpecOsProfileLinuxConfig](#virtualmachinespecosprofilelinuxconfig)***| ***(Optional)*** |
| `osProfileSecrets` | ***[[]VirtualMachineSpecOsProfileSecrets](#virtualmachinespecosprofilesecrets)***| ***(Optional)*** |
| `osProfileWindowsConfig` | ***[[]VirtualMachineSpecOsProfileWindowsConfig](#virtualmachinespecosprofilewindowsconfig)***| ***(Optional)*** |
| `plan` | ***[[]VirtualMachineSpecPlan](#virtualmachinespecplan)***| ***(Optional)*** |
| `primaryNetworkInterfaceID` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `storageDataDisk` | ***[[]VirtualMachineSpecStorageDataDisk](#virtualmachinespecstoragedatadisk)***| ***(Optional)*** |
| `storageImageReference` | ***[[]VirtualMachineSpecStorageImageReference](#virtualmachinespecstorageimagereference)***| ***(Optional)*** |
| `storageOsDisk` | ***[[]VirtualMachineSpecStorageOsDisk](#virtualmachinespecstorageosdisk)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vmSize` | ***string***||
| `zones` | ***[]string***| ***(Optional)*** |
## VirtualMachineSpecBootDiagnostics

Appears on:[VirtualMachineSpec](#virtualmachinespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
| `storageURI` | ***string***||
## VirtualMachineSpecIdentity

Appears on:[VirtualMachineSpec](#virtualmachinespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `identityIDS` | ***[]string***| ***(Optional)*** |
| `principalID` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## VirtualMachineSpecOsProfile

Appears on:[VirtualMachineSpec](#virtualmachinespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `adminUsername` | ***string***||
| `computerName` | ***string***||
| `customData` | ***string***| ***(Optional)*** |
## VirtualMachineSpecOsProfileLinuxConfig

Appears on:[VirtualMachineSpec](#virtualmachinespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `disablePasswordAuthentication` | ***bool***||
| `sshKeys` | ***[[]VirtualMachineSpecOsProfileLinuxConfigSshKeys](#virtualmachinespecosprofilelinuxconfigsshkeys)***| ***(Optional)*** |
## VirtualMachineSpecOsProfileLinuxConfigSshKeys

Appears on:[VirtualMachineSpecOsProfileLinuxConfig](#virtualmachinespecosprofilelinuxconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `keyData` | ***string***||
| `path` | ***string***||
## VirtualMachineSpecOsProfileSecrets

Appears on:[VirtualMachineSpec](#virtualmachinespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `sourceVaultID` | ***string***||
| `vaultCertificates` | ***[[]VirtualMachineSpecOsProfileSecretsVaultCertificates](#virtualmachinespecosprofilesecretsvaultcertificates)***| ***(Optional)*** |
## VirtualMachineSpecOsProfileSecretsVaultCertificates

Appears on:[VirtualMachineSpecOsProfileSecrets](#virtualmachinespecosprofilesecrets)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `certificateStore` | ***string***| ***(Optional)*** |
| `certificateURL` | ***string***||
## VirtualMachineSpecOsProfileWindowsConfig

Appears on:[VirtualMachineSpec](#virtualmachinespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `additionalUnattendConfig` | ***[[]VirtualMachineSpecOsProfileWindowsConfigAdditionalUnattendConfig](#virtualmachinespecosprofilewindowsconfigadditionalunattendconfig)***| ***(Optional)*** |
| `enableAutomaticUpgrades` | ***bool***| ***(Optional)*** |
| `provisionVmAgent` | ***bool***| ***(Optional)*** |
| `timezone` | ***string***| ***(Optional)*** |
| `winrm` | ***[[]VirtualMachineSpecOsProfileWindowsConfigWinrm](#virtualmachinespecosprofilewindowsconfigwinrm)***| ***(Optional)*** |
## VirtualMachineSpecOsProfileWindowsConfigAdditionalUnattendConfig

Appears on:[VirtualMachineSpecOsProfileWindowsConfig](#virtualmachinespecosprofilewindowsconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `component` | ***string***||
| `pass` | ***string***||
| `settingName` | ***string***||
## VirtualMachineSpecOsProfileWindowsConfigWinrm

Appears on:[VirtualMachineSpecOsProfileWindowsConfig](#virtualmachinespecosprofilewindowsconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `certificateURL` | ***string***| ***(Optional)*** |
| `protocol` | ***string***||
## VirtualMachineSpecPlan

Appears on:[VirtualMachineSpec](#virtualmachinespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `product` | ***string***||
| `publisher` | ***string***||
## VirtualMachineSpecStorageDataDisk

Appears on:[VirtualMachineSpec](#virtualmachinespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `caching` | ***string***| ***(Optional)*** |
| `createOption` | ***string***||
| `diskSizeGb` | ***int***| ***(Optional)*** |
| `lun` | ***int***||
| `managedDiskID` | ***string***| ***(Optional)*** |
| `managedDiskType` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `vhdURI` | ***string***| ***(Optional)*** |
| `writeAcceleratorEnabled` | ***bool***| ***(Optional)*** |
## VirtualMachineSpecStorageImageReference

Appears on:[VirtualMachineSpec](#virtualmachinespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `offer` | ***string***| ***(Optional)*** |
| `publisher` | ***string***| ***(Optional)*** |
| `sku` | ***string***| ***(Optional)*** |
| `version` | ***string***| ***(Optional)*** |
## VirtualMachineSpecStorageOsDisk

Appears on:[VirtualMachineSpec](#virtualmachinespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `caching` | ***string***| ***(Optional)*** |
| `createOption` | ***string***||
| `diskSizeGb` | ***int***| ***(Optional)*** |
| `imageURI` | ***string***| ***(Optional)*** |
| `managedDiskID` | ***string***| ***(Optional)*** |
| `managedDiskType` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `osType` | ***string***| ***(Optional)*** |
| `vhdURI` | ***string***| ***(Optional)*** |
| `writeAcceleratorEnabled` | ***bool***| ***(Optional)*** |
## VirtualMachineStatus

Appears on:[VirtualMachine](#virtualmachine)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VirtualMachineSpec](#virtualmachinespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `os_profile.<index>.admin_password` | ***string*** ||
| `os_profile_windows_config.<index>.additional_unattend_config.<index>.content` | ***string*** ||
