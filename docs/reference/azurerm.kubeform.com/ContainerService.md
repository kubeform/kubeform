---
title: ContainerService
menu:
  docs_v0.0.1:
    identifier: containerservice-azurerm.kubeform.com
    name: ContainerService
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ContainerService
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ContainerService` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ContainerServiceSpec](#ContainerServiceSpec)***||
| `status` | ***[ContainerServiceStatus](#ContainerServiceStatus)***||
## ContainerServiceSpec
##### (Appears on:[ContainerService](#ContainerService), [ContainerServiceStatus](#ContainerServiceStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `agentPoolProfile` | ***[[]ContainerServiceSpecAgentPoolProfile](#ContainerServiceSpecAgentPoolProfile)***||
| `diagnosticsProfile` | ***[[]ContainerServiceSpecDiagnosticsProfile](#ContainerServiceSpecDiagnosticsProfile)***||
| `linuxProfile` | ***[[]ContainerServiceSpecLinuxProfile](#ContainerServiceSpecLinuxProfile)***||
| `location` | ***string***||
| `masterProfile` | ***[[]ContainerServiceSpecMasterProfile](#ContainerServiceSpecMasterProfile)***||
| `name` | ***string***||
| `orchestrationPlatform` | ***string***||
| `resourceGroupName` | ***string***||
| `servicePrincipal` | ***[[]ContainerServiceSpecServicePrincipal](#ContainerServiceSpecServicePrincipal)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## ContainerServiceSpecAgentPoolProfile
##### (Appears on:[ContainerServiceSpec](#ContainerServiceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int***| ***(Optional)*** |
| `dnsPrefix` | ***string***||
| `fqdn` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `vmSize` | ***string***||
## ContainerServiceSpecDiagnosticsProfile
##### (Appears on:[ContainerServiceSpec](#ContainerServiceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
| `storageURI` | ***string***| ***(Optional)*** |
## ContainerServiceSpecLinuxProfile
##### (Appears on:[ContainerServiceSpec](#ContainerServiceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `adminUsername` | ***string***||
| `sshKey` | ***[[]ContainerServiceSpecLinuxProfileSshKey](#ContainerServiceSpecLinuxProfileSshKey)***||
## ContainerServiceSpecLinuxProfileSshKey
##### (Appears on:[ContainerServiceSpecLinuxProfile](#ContainerServiceSpecLinuxProfile))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `keyData` | ***string***||
## ContainerServiceSpecMasterProfile
##### (Appears on:[ContainerServiceSpec](#ContainerServiceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int***| ***(Optional)*** |
| `dnsPrefix` | ***string***||
| `fqdn` | ***string***| ***(Optional)*** |
## ContainerServiceSpecServicePrincipal
##### (Appears on:[ContainerServiceSpec](#ContainerServiceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `clientID` | ***string***||
## ContainerServiceStatus
##### (Appears on:[ContainerService](#ContainerService))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ContainerServiceSpec](#ContainerServiceSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `service_principal.<index>.client_secret` | ***string*** ||
