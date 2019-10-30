---
title: ServiceFabricCluster
menu:
  docs_{{ .version }}:
    identifier: servicefabriccluster-azurerm.kubeform.com
    name: ServiceFabricCluster
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ServiceFabricCluster
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ServiceFabricCluster` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ServiceFabricClusterSpec](#servicefabricclusterspec)***||
| `status` | ***[ServiceFabricClusterStatus](#servicefabricclusterstatus)***||
## Phase(`string` alias)

Appears on:[ServiceFabricClusterStatus](#servicefabricclusterstatus)

## ServiceFabricClusterSpec

Appears on:[ServiceFabricCluster](#servicefabriccluster), [ServiceFabricClusterStatus](#servicefabricclusterstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `addOnFeatures` | ***[]string***| ***(Optional)*** |
| `azureActiveDirectory` | ***[[]ServiceFabricClusterSpecAzureActiveDirectory](#servicefabricclusterspecazureactivedirectory)***| ***(Optional)*** |
| `certificate` | ***[[]ServiceFabricClusterSpecCertificate](#servicefabricclusterspeccertificate)***| ***(Optional)*** |
| `certificateCommonNames` | ***[[]ServiceFabricClusterSpecCertificateCommonNames](#servicefabricclusterspeccertificatecommonnames)***| ***(Optional)*** |
| `clientCertificateThumbprint` | ***[[]ServiceFabricClusterSpecClientCertificateThumbprint](#servicefabricclusterspecclientcertificatethumbprint)***| ***(Optional)*** |
| `clusterCodeVersion` | ***string***| ***(Optional)*** |
| `clusterEndpoint` | ***string***| ***(Optional)*** |
| `diagnosticsConfig` | ***[[]ServiceFabricClusterSpecDiagnosticsConfig](#servicefabricclusterspecdiagnosticsconfig)***| ***(Optional)*** |
| `fabricSettings` | ***[[]ServiceFabricClusterSpecFabricSettings](#servicefabricclusterspecfabricsettings)***| ***(Optional)*** |
| `location` | ***string***||
| `managementEndpoint` | ***string***||
| `name` | ***string***||
| `nodeType` | ***[[]ServiceFabricClusterSpecNodeType](#servicefabricclusterspecnodetype)***||
| `reliabilityLevel` | ***string***||
| `resourceGroupName` | ***string***||
| `reverseProxyCertificate` | ***[[]ServiceFabricClusterSpecReverseProxyCertificate](#servicefabricclusterspecreverseproxycertificate)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `upgradeMode` | ***string***||
| `vmImage` | ***string***||
## ServiceFabricClusterSpecAzureActiveDirectory

Appears on:[ServiceFabricClusterSpec](#servicefabricclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `clientApplicationID` | ***string***||
| `clusterApplicationID` | ***string***||
| `tenantID` | ***string***||
## ServiceFabricClusterSpecCertificate

Appears on:[ServiceFabricClusterSpec](#servicefabricclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `thumbprint` | ***string***||
| `thumbprintSecondary` | ***string***| ***(Optional)*** |
| `x509StoreName` | ***string***||
## ServiceFabricClusterSpecCertificateCommonNames

Appears on:[ServiceFabricClusterSpec](#servicefabricclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `commonNames` | ***[[]ServiceFabricClusterSpecCertificateCommonNamesCommonNames](#servicefabricclusterspeccertificatecommonnamescommonnames)***||
| `x509StoreName` | ***string***||
## ServiceFabricClusterSpecCertificateCommonNamesCommonNames

Appears on:[ServiceFabricClusterSpecCertificateCommonNames](#servicefabricclusterspeccertificatecommonnames)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `certificateCommonName` | ***string***||
| `certificateIssuerThumbprint` | ***string***| ***(Optional)*** |
## ServiceFabricClusterSpecClientCertificateThumbprint

Appears on:[ServiceFabricClusterSpec](#servicefabricclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `isAdmin` | ***bool***||
| `thumbprint` | ***string***||
## ServiceFabricClusterSpecDiagnosticsConfig

Appears on:[ServiceFabricClusterSpec](#servicefabricclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `blobEndpoint` | ***string***||
| `protectedAccountKeyName` | ***string***||
| `queueEndpoint` | ***string***||
| `storageAccountName` | ***string***||
| `tableEndpoint` | ***string***||
## ServiceFabricClusterSpecFabricSettings

Appears on:[ServiceFabricClusterSpec](#servicefabricclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `parameters` | ***map[string]string***| ***(Optional)*** |
## ServiceFabricClusterSpecNodeType

Appears on:[ServiceFabricClusterSpec](#servicefabricclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `applicationPorts` | ***[[]ServiceFabricClusterSpecNodeTypeApplicationPorts](#servicefabricclusterspecnodetypeapplicationports)***| ***(Optional)*** |
| `capacities` | ***map[string]string***| ***(Optional)*** |
| `clientEndpointPort` | ***int***||
| `durabilityLevel` | ***string***| ***(Optional)*** |
| `ephemeralPorts` | ***[[]ServiceFabricClusterSpecNodeTypeEphemeralPorts](#servicefabricclusterspecnodetypeephemeralports)***| ***(Optional)*** |
| `httpEndpointPort` | ***int***||
| `instanceCount` | ***int***||
| `isPrimary` | ***bool***||
| `name` | ***string***||
| `placementProperties` | ***map[string]string***| ***(Optional)*** |
| `reverseProxyEndpointPort` | ***int***| ***(Optional)*** |
## ServiceFabricClusterSpecNodeTypeApplicationPorts

Appears on:[ServiceFabricClusterSpecNodeType](#servicefabricclusterspecnodetype)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `endPort` | ***int***||
| `startPort` | ***int***||
## ServiceFabricClusterSpecNodeTypeEphemeralPorts

Appears on:[ServiceFabricClusterSpecNodeType](#servicefabricclusterspecnodetype)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `endPort` | ***int***||
| `startPort` | ***int***||
## ServiceFabricClusterSpecReverseProxyCertificate

Appears on:[ServiceFabricClusterSpec](#servicefabricclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `thumbprint` | ***string***||
| `thumbprintSecondary` | ***string***| ***(Optional)*** |
| `x509StoreName` | ***string***||
## ServiceFabricClusterStatus

Appears on:[ServiceFabricCluster](#servicefabriccluster)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ServiceFabricClusterSpec](#servicefabricclusterspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
