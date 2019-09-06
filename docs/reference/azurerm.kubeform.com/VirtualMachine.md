## VirtualMachine
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `VirtualMachine` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VirtualMachineSpec](#VirtualMachineSpec)***||
| `status` | ***[VirtualMachineStatus](#VirtualMachineStatus)***||
## VirtualMachineSpec
##### (Appears on:[VirtualMachine](#VirtualMachine), [VirtualMachineStatus](#VirtualMachineStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `availabilitySetID` | ***string***| ***(Optional)*** |
| `bootDiagnostics` | ***[[]VirtualMachineSpecBootDiagnostics](#VirtualMachineSpecBootDiagnostics)***| ***(Optional)*** |
| `deleteDataDisksOnTermination` | ***bool***| ***(Optional)*** |
| `deleteOsDiskOnTermination` | ***bool***| ***(Optional)*** |
| `identity` | ***[[]VirtualMachineSpecIdentity](#VirtualMachineSpecIdentity)***| ***(Optional)*** |
| `licenseType` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `networkInterfaceIDS` | ***[]string***||
| `osProfile` | ***[[]VirtualMachineSpecOsProfile](#VirtualMachineSpecOsProfile)***| ***(Optional)*** |
| `osProfileLinuxConfig` | ***[[]VirtualMachineSpecOsProfileLinuxConfig](#VirtualMachineSpecOsProfileLinuxConfig)***| ***(Optional)*** |
| `osProfileSecrets` | ***[[]VirtualMachineSpecOsProfileSecrets](#VirtualMachineSpecOsProfileSecrets)***| ***(Optional)*** |
| `osProfileWindowsConfig` | ***[[]VirtualMachineSpecOsProfileWindowsConfig](#VirtualMachineSpecOsProfileWindowsConfig)***| ***(Optional)*** |
| `plan` | ***[[]VirtualMachineSpecPlan](#VirtualMachineSpecPlan)***| ***(Optional)*** |
| `primaryNetworkInterfaceID` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `storageDataDisk` | ***[[]VirtualMachineSpecStorageDataDisk](#VirtualMachineSpecStorageDataDisk)***| ***(Optional)*** |
| `storageImageReference` | ***[[]VirtualMachineSpecStorageImageReference](#VirtualMachineSpecStorageImageReference)***| ***(Optional)*** |
| `storageOsDisk` | ***[[]VirtualMachineSpecStorageOsDisk](#VirtualMachineSpecStorageOsDisk)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vmSize` | ***string***||
| `zones` | ***[]string***| ***(Optional)*** |
## VirtualMachineSpecBootDiagnostics
##### (Appears on:[VirtualMachineSpec](#VirtualMachineSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
| `storageURI` | ***string***||
## VirtualMachineSpecIdentity
##### (Appears on:[VirtualMachineSpec](#VirtualMachineSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `identityIDS` | ***[]string***| ***(Optional)*** |
| `principalID` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## VirtualMachineSpecOsProfile
##### (Appears on:[VirtualMachineSpec](#VirtualMachineSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `adminUsername` | ***string***||
| `computerName` | ***string***||
| `customData` | ***string***| ***(Optional)*** |
## VirtualMachineSpecOsProfileLinuxConfig
##### (Appears on:[VirtualMachineSpec](#VirtualMachineSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `disablePasswordAuthentication` | ***bool***||
| `sshKeys` | ***[[]VirtualMachineSpecOsProfileLinuxConfigSshKeys](#VirtualMachineSpecOsProfileLinuxConfigSshKeys)***| ***(Optional)*** |
## VirtualMachineSpecOsProfileLinuxConfigSshKeys
##### (Appears on:[VirtualMachineSpecOsProfileLinuxConfig](#VirtualMachineSpecOsProfileLinuxConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `keyData` | ***string***||
| `path` | ***string***||
## VirtualMachineSpecOsProfileSecrets
##### (Appears on:[VirtualMachineSpec](#VirtualMachineSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `sourceVaultID` | ***string***||
| `vaultCertificates` | ***[[]VirtualMachineSpecOsProfileSecretsVaultCertificates](#VirtualMachineSpecOsProfileSecretsVaultCertificates)***| ***(Optional)*** |
## VirtualMachineSpecOsProfileSecretsVaultCertificates
##### (Appears on:[VirtualMachineSpecOsProfileSecrets](#VirtualMachineSpecOsProfileSecrets))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `certificateStore` | ***string***| ***(Optional)*** |
| `certificateURL` | ***string***||
## VirtualMachineSpecOsProfileWindowsConfig
##### (Appears on:[VirtualMachineSpec](#VirtualMachineSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `additionalUnattendConfig` | ***[[]VirtualMachineSpecOsProfileWindowsConfigAdditionalUnattendConfig](#VirtualMachineSpecOsProfileWindowsConfigAdditionalUnattendConfig)***| ***(Optional)*** |
| `enableAutomaticUpgrades` | ***bool***| ***(Optional)*** |
| `provisionVmAgent` | ***bool***| ***(Optional)*** |
| `timezone` | ***string***| ***(Optional)*** |
| `winrm` | ***[[]VirtualMachineSpecOsProfileWindowsConfigWinrm](#VirtualMachineSpecOsProfileWindowsConfigWinrm)***| ***(Optional)*** |
## VirtualMachineSpecOsProfileWindowsConfigAdditionalUnattendConfig
##### (Appears on:[VirtualMachineSpecOsProfileWindowsConfig](#VirtualMachineSpecOsProfileWindowsConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `component` | ***string***||
| `pass` | ***string***||
| `settingName` | ***string***||
## VirtualMachineSpecOsProfileWindowsConfigWinrm
##### (Appears on:[VirtualMachineSpecOsProfileWindowsConfig](#VirtualMachineSpecOsProfileWindowsConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `certificateURL` | ***string***| ***(Optional)*** |
| `protocol` | ***string***||
## VirtualMachineSpecPlan
##### (Appears on:[VirtualMachineSpec](#VirtualMachineSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `product` | ***string***||
| `publisher` | ***string***||
## VirtualMachineSpecStorageDataDisk
##### (Appears on:[VirtualMachineSpec](#VirtualMachineSpec))
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
##### (Appears on:[VirtualMachineSpec](#VirtualMachineSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `offer` | ***string***| ***(Optional)*** |
| `publisher` | ***string***| ***(Optional)*** |
| `sku` | ***string***| ***(Optional)*** |
| `version` | ***string***| ***(Optional)*** |
## VirtualMachineSpecStorageOsDisk
##### (Appears on:[VirtualMachineSpec](#VirtualMachineSpec))
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
##### (Appears on:[VirtualMachine](#VirtualMachine))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VirtualMachineSpec](#VirtualMachineSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `os_profile.<index>.admin_password` | ***string*** ||
| `os_profile_windows_config.<index>.additional_unattend_config.<index>.content` | ***string*** ||
