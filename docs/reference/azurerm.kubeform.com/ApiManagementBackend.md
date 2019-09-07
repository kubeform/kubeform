---
title: ApiManagementBackend
menu:
  docs_v0.0.1:
    identifier: apimanagementbackend-azurerm.kubeform.com
    name: ApiManagementBackend
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ApiManagementBackend
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiManagementBackend` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiManagementBackendSpec](#ApiManagementBackendSpec)***||
| `status` | ***[ApiManagementBackendStatus](#ApiManagementBackendStatus)***||
## ApiManagementBackendSpec
##### (Appears on:[ApiManagementBackend](#ApiManagementBackend), [ApiManagementBackendStatus](#ApiManagementBackendStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `apiManagementName` | ***string***||
| `credentials` | ***[[]ApiManagementBackendSpecCredentials](#ApiManagementBackendSpecCredentials)***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `protocol` | ***string***||
| `proxy` | ***[[]ApiManagementBackendSpecProxy](#ApiManagementBackendSpecProxy)***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `resourceID` | ***string***| ***(Optional)*** |
| `serviceFabricCluster` | ***[[]ApiManagementBackendSpecServiceFabricCluster](#ApiManagementBackendSpecServiceFabricCluster)***| ***(Optional)*** |
| `title` | ***string***| ***(Optional)*** |
| `tls` | ***[[]ApiManagementBackendSpecTls](#ApiManagementBackendSpecTls)***| ***(Optional)*** |
| `url` | ***string***||
## ApiManagementBackendSpecCredentials
##### (Appears on:[ApiManagementBackendSpec](#ApiManagementBackendSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `authorization` | ***[[]ApiManagementBackendSpecCredentialsAuthorization](#ApiManagementBackendSpecCredentialsAuthorization)***| ***(Optional)*** |
| `certificate` | ***[]string***| ***(Optional)*** |
| `header` | ***map[string]string***| ***(Optional)*** |
| `query` | ***map[string]string***| ***(Optional)*** |
## ApiManagementBackendSpecCredentialsAuthorization
##### (Appears on:[ApiManagementBackendSpecCredentials](#ApiManagementBackendSpecCredentials))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `parameter` | ***string***| ***(Optional)*** |
| `scheme` | ***string***| ***(Optional)*** |
## ApiManagementBackendSpecProxy
##### (Appears on:[ApiManagementBackendSpec](#ApiManagementBackendSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `url` | ***string***||
| `username` | ***string***||
## ApiManagementBackendSpecServiceFabricCluster
##### (Appears on:[ApiManagementBackendSpec](#ApiManagementBackendSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `clientCertificateThumbprint` | ***string***||
| `managementEndpoints` | ***[]string***||
| `maxPartitionResolutionRetries` | ***int***||
| `serverCertificateThumbprints` | ***[]string***| ***(Optional)*** |
| `serverX509Name` | ***[[]ApiManagementBackendSpecServiceFabricClusterServerX509Name](#ApiManagementBackendSpecServiceFabricClusterServerX509Name)***| ***(Optional)*** |
## ApiManagementBackendSpecServiceFabricClusterServerX509Name
##### (Appears on:[ApiManagementBackendSpecServiceFabricCluster](#ApiManagementBackendSpecServiceFabricCluster))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `issuerCertificateThumbprint` | ***string***||
| `name` | ***string***||
## ApiManagementBackendSpecTls
##### (Appears on:[ApiManagementBackendSpec](#ApiManagementBackendSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `validateCertificateChain` | ***bool***| ***(Optional)*** |
| `validateCertificateName` | ***bool***| ***(Optional)*** |
## ApiManagementBackendStatus
##### (Appears on:[ApiManagementBackend](#ApiManagementBackend))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiManagementBackendSpec](#ApiManagementBackendSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `proxy.<index>.password` | ***string*** ||
