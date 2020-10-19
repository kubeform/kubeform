---
title: Frontdoor
menu:
  docs_{{ .version }}:
    identifier: frontdoor-azurerm.kubeform.com
    name: Frontdoor
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Frontdoor
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `Frontdoor` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[FrontdoorSpec](#frontdoorspec)***||
| `status` | ***[FrontdoorStatus](#frontdoorstatus)***||
## FrontdoorSpec

Appears on:[Frontdoor](#frontdoor), [FrontdoorStatus](#frontdoorstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `backendPool` | ***[[]FrontdoorSpecBackendPool](#frontdoorspecbackendpool)***||
| `backendPoolHealthProbe` | ***[[]FrontdoorSpecBackendPoolHealthProbe](#frontdoorspecbackendpoolhealthprobe)***||
| `backendPoolLoadBalancing` | ***[[]FrontdoorSpecBackendPoolLoadBalancing](#frontdoorspecbackendpoolloadbalancing)***||
| `cname` | ***string***| ***(Optional)*** |
| `enforceBackendPoolsCertificateNameCheck` | ***bool***||
| `friendlyName` | ***string***| ***(Optional)*** |
| `frontendEndpoint` | ***[[]FrontdoorSpecFrontendEndpoint](#frontdoorspecfrontendendpoint)***||
| `loadBalancerEnabled` | ***bool***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `routingRule` | ***[[]FrontdoorSpecRoutingRule](#frontdoorspecroutingrule)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## FrontdoorSpecBackendPool

Appears on:[FrontdoorSpec](#frontdoorspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `backend` | ***[[]FrontdoorSpecBackendPoolBackend](#frontdoorspecbackendpoolbackend)***||
| `healthProbeName` | ***string***||
| `ID` | ***string***| ***(Optional)*** |
| `loadBalancingName` | ***string***||
| `name` | ***string***||
## FrontdoorSpecBackendPoolBackend

Appears on:[FrontdoorSpecBackendPool](#frontdoorspecbackendpool)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `address` | ***string***||
| `enabled` | ***bool***| ***(Optional)*** |
| `hostHeader` | ***string***||
| `httpPort` | ***int64***||
| `httpsPort` | ***int64***||
| `priority` | ***int64***| ***(Optional)*** |
| `weight` | ***int64***| ***(Optional)*** |
## FrontdoorSpecBackendPoolHealthProbe

Appears on:[FrontdoorSpec](#frontdoorspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `intervalInSeconds` | ***int64***| ***(Optional)*** |
| `name` | ***string***||
| `path` | ***string***| ***(Optional)*** |
| `protocol` | ***string***| ***(Optional)*** |
## FrontdoorSpecBackendPoolLoadBalancing

Appears on:[FrontdoorSpec](#frontdoorspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `additionalLatencyMilliseconds` | ***int64***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `sampleSize` | ***int64***| ***(Optional)*** |
| `successfulSamplesRequired` | ***int64***| ***(Optional)*** |
## FrontdoorSpecFrontendEndpoint

Appears on:[FrontdoorSpec](#frontdoorspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `customHTTPSConfiguration` | ***[[]FrontdoorSpecFrontendEndpointCustomHTTPSConfiguration](#frontdoorspecfrontendendpointcustomhttpsconfiguration)***| ***(Optional)*** |
| `customHTTPSProvisioningEnabled` | ***bool***||
| `hostName` | ***string***||
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `sessionAffinityEnabled` | ***bool***| ***(Optional)*** |
| `sessionAffinityTtlSeconds` | ***int64***| ***(Optional)*** |
| `webApplicationFirewallPolicyLinkID` | ***string***| ***(Optional)*** |
## FrontdoorSpecFrontendEndpointCustomHTTPSConfiguration

Appears on:[FrontdoorSpecFrontendEndpoint](#frontdoorspecfrontendendpoint)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `azureKeyVaultCertificateSecretName` | ***string***| ***(Optional)*** |
| `azureKeyVaultCertificateSecretVersion` | ***string***| ***(Optional)*** |
| `azureKeyVaultCertificateVaultID` | ***string***| ***(Optional)*** |
| `certificateSource` | ***string***| ***(Optional)*** |
| `minimumTLSVersion` | ***string***| ***(Optional)*** |
| `provisioningState` | ***string***| ***(Optional)*** |
| `provisioningSubstate` | ***string***| ***(Optional)*** |
## FrontdoorSpecRoutingRule

Appears on:[FrontdoorSpec](#frontdoorspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `acceptedProtocols` | ***[]string***||
| `enabled` | ***bool***| ***(Optional)*** |
| `forwardingConfiguration` | ***[[]FrontdoorSpecRoutingRuleForwardingConfiguration](#frontdoorspecroutingruleforwardingconfiguration)***| ***(Optional)*** |
| `frontendEndpoints` | ***[]string***||
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `patternsToMatch` | ***[]string***||
| `redirectConfiguration` | ***[[]FrontdoorSpecRoutingRuleRedirectConfiguration](#frontdoorspecroutingruleredirectconfiguration)***| ***(Optional)*** |
## FrontdoorSpecRoutingRuleForwardingConfiguration

Appears on:[FrontdoorSpecRoutingRule](#frontdoorspecroutingrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `backendPoolName` | ***string***||
| `cacheEnabled` | ***bool***| ***(Optional)*** |
| `cacheQueryParameterStripDirective` | ***string***| ***(Optional)*** |
| `cacheUseDynamicCompression` | ***bool***| ***(Optional)*** |
| `customForwardingPath` | ***string***| ***(Optional)*** |
| `forwardingProtocol` | ***string***| ***(Optional)*** |
## FrontdoorSpecRoutingRuleRedirectConfiguration

Appears on:[FrontdoorSpecRoutingRule](#frontdoorspecroutingrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `customFragment` | ***string***| ***(Optional)*** |
| `customHost` | ***string***| ***(Optional)*** |
| `customPath` | ***string***| ***(Optional)*** |
| `customQueryString` | ***string***| ***(Optional)*** |
| `redirectProtocol` | ***string***||
| `redirectType` | ***string***||
## FrontdoorStatus

Appears on:[Frontdoor](#frontdoor)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[FrontdoorSpec](#frontdoorspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[FrontdoorStatus](#frontdoorstatus)

---
