## ServiceFabricCluster
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ServiceFabricCluster` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ServiceFabricClusterSpec](#ServiceFabricClusterSpec)***||
| `status` | ***[ServiceFabricClusterStatus](#ServiceFabricClusterStatus)***||
## ServiceFabricClusterSpec
##### (Appears on:[ServiceFabricCluster](#ServiceFabricCluster), [ServiceFabricClusterStatus](#ServiceFabricClusterStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `addOnFeatures` | ***[]string***| ***(Optional)*** |
| `azureActiveDirectory` | ***[[]ServiceFabricClusterSpecAzureActiveDirectory](#ServiceFabricClusterSpecAzureActiveDirectory)***| ***(Optional)*** |
| `certificate` | ***[[]ServiceFabricClusterSpecCertificate](#ServiceFabricClusterSpecCertificate)***| ***(Optional)*** |
| `certificateCommonNames` | ***[[]ServiceFabricClusterSpecCertificateCommonNames](#ServiceFabricClusterSpecCertificateCommonNames)***| ***(Optional)*** |
| `clientCertificateThumbprint` | ***[[]ServiceFabricClusterSpecClientCertificateThumbprint](#ServiceFabricClusterSpecClientCertificateThumbprint)***| ***(Optional)*** |
| `clusterCodeVersion` | ***string***| ***(Optional)*** |
| `clusterEndpoint` | ***string***| ***(Optional)*** |
| `diagnosticsConfig` | ***[[]ServiceFabricClusterSpecDiagnosticsConfig](#ServiceFabricClusterSpecDiagnosticsConfig)***| ***(Optional)*** |
| `fabricSettings` | ***[[]ServiceFabricClusterSpecFabricSettings](#ServiceFabricClusterSpecFabricSettings)***| ***(Optional)*** |
| `location` | ***string***||
| `managementEndpoint` | ***string***||
| `name` | ***string***||
| `nodeType` | ***[[]ServiceFabricClusterSpecNodeType](#ServiceFabricClusterSpecNodeType)***||
| `reliabilityLevel` | ***string***||
| `resourceGroupName` | ***string***||
| `reverseProxyCertificate` | ***[[]ServiceFabricClusterSpecReverseProxyCertificate](#ServiceFabricClusterSpecReverseProxyCertificate)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `upgradeMode` | ***string***||
| `vmImage` | ***string***||
## ServiceFabricClusterSpecAzureActiveDirectory
##### (Appears on:[ServiceFabricClusterSpec](#ServiceFabricClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `clientApplicationID` | ***string***||
| `clusterApplicationID` | ***string***||
| `tenantID` | ***string***||
## ServiceFabricClusterSpecCertificate
##### (Appears on:[ServiceFabricClusterSpec](#ServiceFabricClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `thumbprint` | ***string***||
| `thumbprintSecondary` | ***string***| ***(Optional)*** |
| `x509StoreName` | ***string***||
## ServiceFabricClusterSpecCertificateCommonNames
##### (Appears on:[ServiceFabricClusterSpec](#ServiceFabricClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `commonNames` | ***[[]ServiceFabricClusterSpecCertificateCommonNamesCommonNames](#ServiceFabricClusterSpecCertificateCommonNamesCommonNames)***||
| `x509StoreName` | ***string***||
## ServiceFabricClusterSpecCertificateCommonNamesCommonNames
##### (Appears on:[ServiceFabricClusterSpecCertificateCommonNames](#ServiceFabricClusterSpecCertificateCommonNames))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `certificateCommonName` | ***string***||
| `certificateIssuerThumbprint` | ***string***| ***(Optional)*** |
## ServiceFabricClusterSpecClientCertificateThumbprint
##### (Appears on:[ServiceFabricClusterSpec](#ServiceFabricClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `isAdmin` | ***bool***||
| `thumbprint` | ***string***||
## ServiceFabricClusterSpecDiagnosticsConfig
##### (Appears on:[ServiceFabricClusterSpec](#ServiceFabricClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `blobEndpoint` | ***string***||
| `protectedAccountKeyName` | ***string***||
| `queueEndpoint` | ***string***||
| `storageAccountName` | ***string***||
| `tableEndpoint` | ***string***||
## ServiceFabricClusterSpecFabricSettings
##### (Appears on:[ServiceFabricClusterSpec](#ServiceFabricClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `parameters` | ***map[string]string***| ***(Optional)*** |
## ServiceFabricClusterSpecNodeType
##### (Appears on:[ServiceFabricClusterSpec](#ServiceFabricClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `applicationPorts` | ***[[]ServiceFabricClusterSpecNodeTypeApplicationPorts](#ServiceFabricClusterSpecNodeTypeApplicationPorts)***| ***(Optional)*** |
| `capacities` | ***map[string]string***| ***(Optional)*** |
| `clientEndpointPort` | ***int***||
| `durabilityLevel` | ***string***| ***(Optional)*** |
| `ephemeralPorts` | ***[[]ServiceFabricClusterSpecNodeTypeEphemeralPorts](#ServiceFabricClusterSpecNodeTypeEphemeralPorts)***| ***(Optional)*** |
| `httpEndpointPort` | ***int***||
| `instanceCount` | ***int***||
| `isPrimary` | ***bool***||
| `name` | ***string***||
| `placementProperties` | ***map[string]string***| ***(Optional)*** |
| `reverseProxyEndpointPort` | ***int***| ***(Optional)*** |
## ServiceFabricClusterSpecNodeTypeApplicationPorts
##### (Appears on:[ServiceFabricClusterSpecNodeType](#ServiceFabricClusterSpecNodeType))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `endPort` | ***int***||
| `startPort` | ***int***||
## ServiceFabricClusterSpecNodeTypeEphemeralPorts
##### (Appears on:[ServiceFabricClusterSpecNodeType](#ServiceFabricClusterSpecNodeType))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `endPort` | ***int***||
| `startPort` | ***int***||
## ServiceFabricClusterSpecReverseProxyCertificate
##### (Appears on:[ServiceFabricClusterSpec](#ServiceFabricClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `thumbprint` | ***string***||
| `thumbprintSecondary` | ***string***| ***(Optional)*** |
| `x509StoreName` | ***string***||
## ServiceFabricClusterStatus
##### (Appears on:[ServiceFabricCluster](#ServiceFabricCluster))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ServiceFabricClusterSpec](#ServiceFabricClusterSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
