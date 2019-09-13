---
title: ContainerGroup
menu:
  docs_{{ .version }}:
    identifier: containergroup-azurerm.kubeform.com
    name: ContainerGroup
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ContainerGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ContainerGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ContainerGroupSpec](#ContainerGroupSpec)***||
| `status` | ***[ContainerGroupStatus](#ContainerGroupStatus)***||
## ContainerGroupSpec
##### (Appears on:[ContainerGroup](#ContainerGroup), [ContainerGroupStatus](#ContainerGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `container` | ***[[]ContainerGroupSpecContainer](#ContainerGroupSpecContainer)***||
| `diagnostics` | ***[[]ContainerGroupSpecDiagnostics](#ContainerGroupSpecDiagnostics)***| ***(Optional)*** |
| `dnsNameLabel` | ***string***| ***(Optional)*** |
| `fqdn` | ***string***| ***(Optional)*** |
| `identity` | ***[[]ContainerGroupSpecIdentity](#ContainerGroupSpecIdentity)***| ***(Optional)*** |
| `imageRegistryCredential` | ***[[]ContainerGroupSpecImageRegistryCredential](#ContainerGroupSpecImageRegistryCredential)***| ***(Optional)*** |
| `ipAddress` | ***string***| ***(Optional)*** |
| `ipAddressType` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `osType` | ***string***||
| `resourceGroupName` | ***string***||
| `restartPolicy` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## ContainerGroupSpecContainer
##### (Appears on:[ContainerGroupSpec](#ContainerGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `command` | ***string***| ***(Optional)*** Deprecated|
| `commands` | ***[]string***| ***(Optional)*** |
| `cpu` | ***encoding/json.Number***||
| `environmentVariables` | ***map[string]string***| ***(Optional)*** |
| `gpu` | ***[[]ContainerGroupSpecContainerGpu](#ContainerGroupSpecContainerGpu)***| ***(Optional)*** |
| `image` | ***string***||
| `livenessProbe` | ***[[]ContainerGroupSpecContainerLivenessProbe](#ContainerGroupSpecContainerLivenessProbe)***| ***(Optional)*** |
| `memory` | ***encoding/json.Number***||
| `name` | ***string***||
| `port` | ***int***| ***(Optional)*** Deprecated|
| `ports` | ***[[]ContainerGroupSpecContainerPorts](#ContainerGroupSpecContainerPorts)***| ***(Optional)*** |
| `protocol` | ***string***| ***(Optional)*** Deprecated|
| `readinessProbe` | ***[[]ContainerGroupSpecContainerReadinessProbe](#ContainerGroupSpecContainerReadinessProbe)***| ***(Optional)*** |
| `volume` | ***[[]ContainerGroupSpecContainerVolume](#ContainerGroupSpecContainerVolume)***| ***(Optional)*** |
## ContainerGroupSpecContainerGpu
##### (Appears on:[ContainerGroupSpecContainer](#ContainerGroupSpecContainer))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int***| ***(Optional)*** |
| `sku` | ***string***| ***(Optional)*** |
## ContainerGroupSpecContainerLivenessProbe
##### (Appears on:[ContainerGroupSpecContainer](#ContainerGroupSpecContainer))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `exec` | ***[]string***| ***(Optional)*** |
| `failureThreshold` | ***int***| ***(Optional)*** |
| `httpGet` | ***[[]ContainerGroupSpecContainerLivenessProbeHttpGet](#ContainerGroupSpecContainerLivenessProbeHttpGet)***| ***(Optional)*** |
| `initialDelaySeconds` | ***int***| ***(Optional)*** |
| `periodSeconds` | ***int***| ***(Optional)*** |
| `successThreshold` | ***int***| ***(Optional)*** |
| `timeoutSeconds` | ***int***| ***(Optional)*** |
## ContainerGroupSpecContainerLivenessProbeHttpGet
##### (Appears on:[ContainerGroupSpecContainerLivenessProbe](#ContainerGroupSpecContainerLivenessProbe))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `path` | ***string***| ***(Optional)*** |
| `port` | ***int***| ***(Optional)*** |
| `scheme` | ***string***| ***(Optional)*** |
## ContainerGroupSpecContainerPorts
##### (Appears on:[ContainerGroupSpecContainer](#ContainerGroupSpecContainer))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `port` | ***int***| ***(Optional)*** |
| `protocol` | ***string***| ***(Optional)*** |
## ContainerGroupSpecContainerReadinessProbe
##### (Appears on:[ContainerGroupSpecContainer](#ContainerGroupSpecContainer))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `exec` | ***[]string***| ***(Optional)*** |
| `failureThreshold` | ***int***| ***(Optional)*** |
| `httpGet` | ***[[]ContainerGroupSpecContainerReadinessProbeHttpGet](#ContainerGroupSpecContainerReadinessProbeHttpGet)***| ***(Optional)*** |
| `initialDelaySeconds` | ***int***| ***(Optional)*** |
| `periodSeconds` | ***int***| ***(Optional)*** |
| `successThreshold` | ***int***| ***(Optional)*** |
| `timeoutSeconds` | ***int***| ***(Optional)*** |
## ContainerGroupSpecContainerReadinessProbeHttpGet
##### (Appears on:[ContainerGroupSpecContainerReadinessProbe](#ContainerGroupSpecContainerReadinessProbe))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `path` | ***string***| ***(Optional)*** |
| `port` | ***int***| ***(Optional)*** |
| `scheme` | ***string***| ***(Optional)*** |
## ContainerGroupSpecContainerVolume
##### (Appears on:[ContainerGroupSpecContainer](#ContainerGroupSpecContainer))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `mountPath` | ***string***||
| `name` | ***string***||
| `readOnly` | ***bool***| ***(Optional)*** |
| `shareName` | ***string***||
| `storageAccountKey` | ***string***||
| `storageAccountName` | ***string***||
## ContainerGroupSpecDiagnostics
##### (Appears on:[ContainerGroupSpec](#ContainerGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `logAnalytics` | ***[[]ContainerGroupSpecDiagnosticsLogAnalytics](#ContainerGroupSpecDiagnosticsLogAnalytics)***||
## ContainerGroupSpecDiagnosticsLogAnalytics
##### (Appears on:[ContainerGroupSpecDiagnostics](#ContainerGroupSpecDiagnostics))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `logType` | ***string***||
| `metadata` | ***map[string]string***| ***(Optional)*** |
| `workspaceID` | ***string***||
## ContainerGroupSpecIdentity
##### (Appears on:[ContainerGroupSpec](#ContainerGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `identityIDS` | ***[]string***| ***(Optional)*** |
| `principalID` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## ContainerGroupSpecImageRegistryCredential
##### (Appears on:[ContainerGroupSpec](#ContainerGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `server` | ***string***||
| `username` | ***string***||
## ContainerGroupStatus
##### (Appears on:[ContainerGroup](#ContainerGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ContainerGroupSpec](#ContainerGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `container.<index>.secure_environment_variables` | ***map[string]string*** ||
| `diagnostics.<index>.log_analytics.<index>.workspace_key` | ***string*** ||
| `image_registry_credential.<index>.password` | ***string*** ||
