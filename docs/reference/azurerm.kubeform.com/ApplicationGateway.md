---
title: ApplicationGateway
menu:
  docs_{{ .version }}:
    identifier: applicationgateway-azurerm.kubeform.com
    name: ApplicationGateway
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApplicationGateway
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ApplicationGateway` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApplicationGatewaySpec](#ApplicationGatewaySpec)***||
| `status` | ***[ApplicationGatewayStatus](#ApplicationGatewayStatus)***||
## ApplicationGatewaySpec
##### (Appears on:[ApplicationGateway](#ApplicationGateway), [ApplicationGatewayStatus](#ApplicationGatewayStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `authenticationCertificate` | ***[[]ApplicationGatewaySpecAuthenticationCertificate](#ApplicationGatewaySpecAuthenticationCertificate)***| ***(Optional)*** |
| `autoscaleConfiguration` | ***[[]ApplicationGatewaySpecAutoscaleConfiguration](#ApplicationGatewaySpecAutoscaleConfiguration)***| ***(Optional)*** |
| `backendAddressPool` | ***[[]ApplicationGatewaySpecBackendAddressPool](#ApplicationGatewaySpecBackendAddressPool)***||
| `backendHTTPSettings` | ***[[]ApplicationGatewaySpecBackendHTTPSettings](#ApplicationGatewaySpecBackendHTTPSettings)***||
| `customErrorConfiguration` | ***[[]ApplicationGatewaySpecCustomErrorConfiguration](#ApplicationGatewaySpecCustomErrorConfiguration)***| ***(Optional)*** |
| `disabledSSLProtocols` | ***[]string***| ***(Optional)*** Deprecated|
| `enableHttp2` | ***bool***| ***(Optional)*** |
| `frontendIPConfiguration` | ***[[]ApplicationGatewaySpecFrontendIPConfiguration](#ApplicationGatewaySpecFrontendIPConfiguration)***||
| `frontendPort` | ***[[]ApplicationGatewaySpecFrontendPort](#ApplicationGatewaySpecFrontendPort)***||
| `gatewayIPConfiguration` | ***[[]ApplicationGatewaySpecGatewayIPConfiguration](#ApplicationGatewaySpecGatewayIPConfiguration)***||
| `httpListener` | ***[[]ApplicationGatewaySpecHttpListener](#ApplicationGatewaySpecHttpListener)***||
| `identity` | ***[[]ApplicationGatewaySpecIdentity](#ApplicationGatewaySpecIdentity)***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `probe` | ***[[]ApplicationGatewaySpecProbe](#ApplicationGatewaySpecProbe)***| ***(Optional)*** |
| `redirectConfiguration` | ***[[]ApplicationGatewaySpecRedirectConfiguration](#ApplicationGatewaySpecRedirectConfiguration)***| ***(Optional)*** |
| `requestRoutingRule` | ***[[]ApplicationGatewaySpecRequestRoutingRule](#ApplicationGatewaySpecRequestRoutingRule)***||
| `resourceGroupName` | ***string***||
| `rewriteRuleSet` | ***[[]ApplicationGatewaySpecRewriteRuleSet](#ApplicationGatewaySpecRewriteRuleSet)***| ***(Optional)*** |
| `sku` | ***[[]ApplicationGatewaySpecSku](#ApplicationGatewaySpecSku)***||
| `sslCertificate` | ***[[]ApplicationGatewaySpecSslCertificate](#ApplicationGatewaySpecSslCertificate)***| ***(Optional)*** |
| `sslPolicy` | ***[[]ApplicationGatewaySpecSslPolicy](#ApplicationGatewaySpecSslPolicy)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `urlPathMap` | ***[[]ApplicationGatewaySpecUrlPathMap](#ApplicationGatewaySpecUrlPathMap)***| ***(Optional)*** |
| `wafConfiguration` | ***[[]ApplicationGatewaySpecWafConfiguration](#ApplicationGatewaySpecWafConfiguration)***| ***(Optional)*** |
| `zones` | ***[]string***| ***(Optional)*** |
## ApplicationGatewaySpecAuthenticationCertificate
##### (Appears on:[ApplicationGatewaySpec](#ApplicationGatewaySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
## ApplicationGatewaySpecAutoscaleConfiguration
##### (Appears on:[ApplicationGatewaySpec](#ApplicationGatewaySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `maxCapacity` | ***int***| ***(Optional)*** |
| `minCapacity` | ***int***||
## ApplicationGatewaySpecBackendAddressPool
##### (Appears on:[ApplicationGatewaySpec](#ApplicationGatewaySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `fqdnList` | ***[]string***| ***(Optional)*** Deprecated|
| `fqdns` | ***[]string***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `ipAddressList` | ***[]string***| ***(Optional)*** Deprecated|
| `ipAddresses` | ***[]string***| ***(Optional)*** |
| `name` | ***string***||
## ApplicationGatewaySpecBackendHTTPSettings
##### (Appears on:[ApplicationGatewaySpec](#ApplicationGatewaySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `affinityCookieName` | ***string***| ***(Optional)*** |
| `authenticationCertificate` | ***[[]ApplicationGatewaySpecBackendHTTPSettingsAuthenticationCertificate](#ApplicationGatewaySpecBackendHTTPSettingsAuthenticationCertificate)***| ***(Optional)*** |
| `connectionDraining` | ***[[]ApplicationGatewaySpecBackendHTTPSettingsConnectionDraining](#ApplicationGatewaySpecBackendHTTPSettingsConnectionDraining)***| ***(Optional)*** |
| `cookieBasedAffinity` | ***string***||
| `hostName` | ***string***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `path` | ***string***| ***(Optional)*** |
| `pickHostNameFromBackendAddress` | ***bool***| ***(Optional)*** |
| `port` | ***int***||
| `probeID` | ***string***| ***(Optional)*** |
| `probeName` | ***string***| ***(Optional)*** |
| `protocol` | ***string***||
| `requestTimeout` | ***int***| ***(Optional)*** |
## ApplicationGatewaySpecBackendHTTPSettingsAuthenticationCertificate
##### (Appears on:[ApplicationGatewaySpecBackendHTTPSettings](#ApplicationGatewaySpecBackendHTTPSettings))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
## ApplicationGatewaySpecBackendHTTPSettingsConnectionDraining
##### (Appears on:[ApplicationGatewaySpecBackendHTTPSettings](#ApplicationGatewaySpecBackendHTTPSettings))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `drainTimeoutSec` | ***int***||
| `enabled` | ***bool***||
## ApplicationGatewaySpecCustomErrorConfiguration
##### (Appears on:[ApplicationGatewaySpec](#ApplicationGatewaySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `customErrorPageURL` | ***string***||
| `ID` | ***string***| ***(Optional)*** |
| `statusCode` | ***string***||
## ApplicationGatewaySpecFrontendIPConfiguration
##### (Appears on:[ApplicationGatewaySpec](#ApplicationGatewaySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `privateIPAddress` | ***string***| ***(Optional)*** |
| `privateIPAddressAllocation` | ***string***| ***(Optional)*** |
| `publicIPAddressID` | ***string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
## ApplicationGatewaySpecFrontendPort
##### (Appears on:[ApplicationGatewaySpec](#ApplicationGatewaySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `port` | ***int***||
## ApplicationGatewaySpecGatewayIPConfiguration
##### (Appears on:[ApplicationGatewaySpec](#ApplicationGatewaySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `subnetID` | ***string***||
## ApplicationGatewaySpecHttpListener
##### (Appears on:[ApplicationGatewaySpec](#ApplicationGatewaySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `customErrorConfiguration` | ***[[]ApplicationGatewaySpecHttpListenerCustomErrorConfiguration](#ApplicationGatewaySpecHttpListenerCustomErrorConfiguration)***| ***(Optional)*** |
| `frontendIPConfigurationID` | ***string***| ***(Optional)*** |
| `frontendIPConfigurationName` | ***string***||
| `frontendPortID` | ***string***| ***(Optional)*** |
| `frontendPortName` | ***string***||
| `hostName` | ***string***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `protocol` | ***string***||
| `requireSni` | ***bool***| ***(Optional)*** |
| `sslCertificateID` | ***string***| ***(Optional)*** |
| `sslCertificateName` | ***string***| ***(Optional)*** |
## ApplicationGatewaySpecHttpListenerCustomErrorConfiguration
##### (Appears on:[ApplicationGatewaySpecHttpListener](#ApplicationGatewaySpecHttpListener))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `customErrorPageURL` | ***string***||
| `ID` | ***string***| ***(Optional)*** |
| `statusCode` | ***string***||
## ApplicationGatewaySpecIdentity
##### (Appears on:[ApplicationGatewaySpec](#ApplicationGatewaySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `identityIDS` | ***[]string***||
| `type` | ***string***| ***(Optional)*** |
## ApplicationGatewaySpecProbe
##### (Appears on:[ApplicationGatewaySpec](#ApplicationGatewaySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `host` | ***string***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `interval` | ***int***||
| `match` | ***[[]ApplicationGatewaySpecProbeMatch](#ApplicationGatewaySpecProbeMatch)***| ***(Optional)*** |
| `minimumServers` | ***int***| ***(Optional)*** |
| `name` | ***string***||
| `path` | ***string***||
| `pickHostNameFromBackendHTTPSettings` | ***bool***| ***(Optional)*** |
| `protocol` | ***string***||
| `timeout` | ***int***||
| `unhealthyThreshold` | ***int***||
## ApplicationGatewaySpecProbeMatch
##### (Appears on:[ApplicationGatewaySpecProbe](#ApplicationGatewaySpecProbe))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `body` | ***string***| ***(Optional)*** |
| `statusCode` | ***[]string***| ***(Optional)*** |
## ApplicationGatewaySpecRedirectConfiguration
##### (Appears on:[ApplicationGatewaySpec](#ApplicationGatewaySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `includePath` | ***bool***| ***(Optional)*** |
| `includeQueryString` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `redirectType` | ***string***||
| `targetListenerID` | ***string***| ***(Optional)*** |
| `targetListenerName` | ***string***| ***(Optional)*** |
| `targetURL` | ***string***| ***(Optional)*** |
## ApplicationGatewaySpecRequestRoutingRule
##### (Appears on:[ApplicationGatewaySpec](#ApplicationGatewaySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `backendAddressPoolID` | ***string***| ***(Optional)*** |
| `backendAddressPoolName` | ***string***| ***(Optional)*** |
| `backendHTTPSettingsID` | ***string***| ***(Optional)*** |
| `backendHTTPSettingsName` | ***string***| ***(Optional)*** |
| `httpListenerID` | ***string***| ***(Optional)*** |
| `httpListenerName` | ***string***||
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `redirectConfigurationID` | ***string***| ***(Optional)*** |
| `redirectConfigurationName` | ***string***| ***(Optional)*** |
| `rewriteRuleSetID` | ***string***| ***(Optional)*** |
| `rewriteRuleSetName` | ***string***| ***(Optional)*** |
| `ruleType` | ***string***||
| `urlPathMapID` | ***string***| ***(Optional)*** |
| `urlPathMapName` | ***string***| ***(Optional)*** |
## ApplicationGatewaySpecRewriteRuleSet
##### (Appears on:[ApplicationGatewaySpec](#ApplicationGatewaySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `rewriteRule` | ***[[]ApplicationGatewaySpecRewriteRuleSetRewriteRule](#ApplicationGatewaySpecRewriteRuleSetRewriteRule)***| ***(Optional)*** |
## ApplicationGatewaySpecRewriteRuleSetRewriteRule
##### (Appears on:[ApplicationGatewaySpecRewriteRuleSet](#ApplicationGatewaySpecRewriteRuleSet))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `condition` | ***[[]ApplicationGatewaySpecRewriteRuleSetRewriteRuleCondition](#ApplicationGatewaySpecRewriteRuleSetRewriteRuleCondition)***| ***(Optional)*** |
| `name` | ***string***||
| `requestHeaderConfiguration` | ***[[]ApplicationGatewaySpecRewriteRuleSetRewriteRuleRequestHeaderConfiguration](#ApplicationGatewaySpecRewriteRuleSetRewriteRuleRequestHeaderConfiguration)***| ***(Optional)*** |
| `responseHeaderConfiguration` | ***[[]ApplicationGatewaySpecRewriteRuleSetRewriteRuleResponseHeaderConfiguration](#ApplicationGatewaySpecRewriteRuleSetRewriteRuleResponseHeaderConfiguration)***| ***(Optional)*** |
| `ruleSequence` | ***int***||
## ApplicationGatewaySpecRewriteRuleSetRewriteRuleCondition
##### (Appears on:[ApplicationGatewaySpecRewriteRuleSetRewriteRule](#ApplicationGatewaySpecRewriteRuleSetRewriteRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ignoreCase` | ***bool***| ***(Optional)*** |
| `negate` | ***bool***| ***(Optional)*** |
| `pattern` | ***string***||
| `variable` | ***string***||
## ApplicationGatewaySpecRewriteRuleSetRewriteRuleRequestHeaderConfiguration
##### (Appears on:[ApplicationGatewaySpecRewriteRuleSetRewriteRule](#ApplicationGatewaySpecRewriteRuleSetRewriteRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `headerName` | ***string***||
| `headerValue` | ***string***||
## ApplicationGatewaySpecRewriteRuleSetRewriteRuleResponseHeaderConfiguration
##### (Appears on:[ApplicationGatewaySpecRewriteRuleSetRewriteRule](#ApplicationGatewaySpecRewriteRuleSetRewriteRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `headerName` | ***string***||
| `headerValue` | ***string***||
## ApplicationGatewaySpecSku
##### (Appears on:[ApplicationGatewaySpec](#ApplicationGatewaySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `capacity` | ***int***| ***(Optional)*** |
| `name` | ***string***||
| `tier` | ***string***||
## ApplicationGatewaySpecSslCertificate
##### (Appears on:[ApplicationGatewaySpec](#ApplicationGatewaySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `publicCertData` | ***string***| ***(Optional)*** |
## ApplicationGatewaySpecSslPolicy
##### (Appears on:[ApplicationGatewaySpec](#ApplicationGatewaySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `cipherSuites` | ***[]string***| ***(Optional)*** |
| `disabledProtocols` | ***[]string***| ***(Optional)*** |
| `minProtocolVersion` | ***string***| ***(Optional)*** |
| `policyName` | ***string***| ***(Optional)*** |
| `policyType` | ***string***| ***(Optional)*** |
## ApplicationGatewaySpecUrlPathMap
##### (Appears on:[ApplicationGatewaySpec](#ApplicationGatewaySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `defaultBackendAddressPoolID` | ***string***| ***(Optional)*** |
| `defaultBackendAddressPoolName` | ***string***| ***(Optional)*** |
| `defaultBackendHTTPSettingsID` | ***string***| ***(Optional)*** |
| `defaultBackendHTTPSettingsName` | ***string***| ***(Optional)*** |
| `defaultRedirectConfigurationID` | ***string***| ***(Optional)*** |
| `defaultRedirectConfigurationName` | ***string***| ***(Optional)*** |
| `defaultRewriteRuleSetID` | ***string***| ***(Optional)*** |
| `defaultRewriteRuleSetName` | ***string***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `pathRule` | ***[[]ApplicationGatewaySpecUrlPathMapPathRule](#ApplicationGatewaySpecUrlPathMapPathRule)***||
## ApplicationGatewaySpecUrlPathMapPathRule
##### (Appears on:[ApplicationGatewaySpecUrlPathMap](#ApplicationGatewaySpecUrlPathMap))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `backendAddressPoolID` | ***string***| ***(Optional)*** |
| `backendAddressPoolName` | ***string***| ***(Optional)*** |
| `backendHTTPSettingsID` | ***string***| ***(Optional)*** |
| `backendHTTPSettingsName` | ***string***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `paths` | ***[]string***||
| `redirectConfigurationID` | ***string***| ***(Optional)*** |
| `redirectConfigurationName` | ***string***| ***(Optional)*** |
| `rewriteRuleSetID` | ***string***| ***(Optional)*** |
| `rewriteRuleSetName` | ***string***| ***(Optional)*** |
## ApplicationGatewaySpecWafConfiguration
##### (Appears on:[ApplicationGatewaySpec](#ApplicationGatewaySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `disabledRuleGroup` | ***[[]ApplicationGatewaySpecWafConfigurationDisabledRuleGroup](#ApplicationGatewaySpecWafConfigurationDisabledRuleGroup)***| ***(Optional)*** |
| `enabled` | ***bool***||
| `exclusion` | ***[[]ApplicationGatewaySpecWafConfigurationExclusion](#ApplicationGatewaySpecWafConfigurationExclusion)***| ***(Optional)*** |
| `fileUploadLimitMb` | ***int***| ***(Optional)*** |
| `firewallMode` | ***string***||
| `maxRequestBodySizeKb` | ***int***| ***(Optional)*** |
| `requestBodyCheck` | ***bool***| ***(Optional)*** |
| `ruleSetType` | ***string***| ***(Optional)*** |
| `ruleSetVersion` | ***string***||
## ApplicationGatewaySpecWafConfigurationDisabledRuleGroup
##### (Appears on:[ApplicationGatewaySpecWafConfiguration](#ApplicationGatewaySpecWafConfiguration))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ruleGroupName` | ***string***||
| `rules` | ***[]int64***| ***(Optional)*** |
## ApplicationGatewaySpecWafConfigurationExclusion
##### (Appears on:[ApplicationGatewaySpecWafConfiguration](#ApplicationGatewaySpecWafConfiguration))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `matchVariable` | ***string***||
| `selector` | ***string***| ***(Optional)*** |
| `selectorMatchOperator` | ***string***| ***(Optional)*** |
## ApplicationGatewayStatus
##### (Appears on:[ApplicationGateway](#ApplicationGateway))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApplicationGatewaySpec](#ApplicationGatewaySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `authentication_certificate.<index>.data` | ***string*** ||
| `ssl_certificate.<index>.data` | ***string*** ||
| `ssl_certificate.<index>.password` | ***string*** ||
