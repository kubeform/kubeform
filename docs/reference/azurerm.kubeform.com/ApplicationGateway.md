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
| `spec` | ***[ApplicationGatewaySpec](#applicationgatewayspec)***||
| `status` | ***[ApplicationGatewayStatus](#applicationgatewaystatus)***||
## ApplicationGatewaySpec

Appears on:[ApplicationGateway](#applicationgateway), [ApplicationGatewayStatus](#applicationgatewaystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `authenticationCertificate` | ***[[]ApplicationGatewaySpecAuthenticationCertificate](#applicationgatewayspecauthenticationcertificate)***| ***(Optional)*** |
| `autoscaleConfiguration` | ***[[]ApplicationGatewaySpecAutoscaleConfiguration](#applicationgatewayspecautoscaleconfiguration)***| ***(Optional)*** |
| `backendAddressPool` | ***[[]ApplicationGatewaySpecBackendAddressPool](#applicationgatewayspecbackendaddresspool)***||
| `backendHTTPSettings` | ***[[]ApplicationGatewaySpecBackendHTTPSettings](#applicationgatewayspecbackendhttpsettings)***||
| `customErrorConfiguration` | ***[[]ApplicationGatewaySpecCustomErrorConfiguration](#applicationgatewayspeccustomerrorconfiguration)***| ***(Optional)*** |
| `disabledSSLProtocols` | ***[]string***| ***(Optional)*** Deprecated|
| `enableHttp2` | ***bool***| ***(Optional)*** |
| `frontendIPConfiguration` | ***[[]ApplicationGatewaySpecFrontendIPConfiguration](#applicationgatewayspecfrontendipconfiguration)***||
| `frontendPort` | ***[[]ApplicationGatewaySpecFrontendPort](#applicationgatewayspecfrontendport)***||
| `gatewayIPConfiguration` | ***[[]ApplicationGatewaySpecGatewayIPConfiguration](#applicationgatewayspecgatewayipconfiguration)***||
| `httpListener` | ***[[]ApplicationGatewaySpecHttpListener](#applicationgatewayspechttplistener)***||
| `identity` | ***[[]ApplicationGatewaySpecIdentity](#applicationgatewayspecidentity)***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `probe` | ***[[]ApplicationGatewaySpecProbe](#applicationgatewayspecprobe)***| ***(Optional)*** |
| `redirectConfiguration` | ***[[]ApplicationGatewaySpecRedirectConfiguration](#applicationgatewayspecredirectconfiguration)***| ***(Optional)*** |
| `requestRoutingRule` | ***[[]ApplicationGatewaySpecRequestRoutingRule](#applicationgatewayspecrequestroutingrule)***||
| `resourceGroupName` | ***string***||
| `rewriteRuleSet` | ***[[]ApplicationGatewaySpecRewriteRuleSet](#applicationgatewayspecrewriteruleset)***| ***(Optional)*** |
| `sku` | ***[[]ApplicationGatewaySpecSku](#applicationgatewayspecsku)***||
| `sslCertificate` | ***[[]ApplicationGatewaySpecSslCertificate](#applicationgatewayspecsslcertificate)***| ***(Optional)*** |
| `sslPolicy` | ***[[]ApplicationGatewaySpecSslPolicy](#applicationgatewayspecsslpolicy)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `urlPathMap` | ***[[]ApplicationGatewaySpecUrlPathMap](#applicationgatewayspecurlpathmap)***| ***(Optional)*** |
| `wafConfiguration` | ***[[]ApplicationGatewaySpecWafConfiguration](#applicationgatewayspecwafconfiguration)***| ***(Optional)*** |
| `zones` | ***[]string***| ***(Optional)*** |
## ApplicationGatewaySpecAuthenticationCertificate

Appears on:[ApplicationGatewaySpec](#applicationgatewayspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
## ApplicationGatewaySpecAutoscaleConfiguration

Appears on:[ApplicationGatewaySpec](#applicationgatewayspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `maxCapacity` | ***int64***| ***(Optional)*** |
| `minCapacity` | ***int64***||
## ApplicationGatewaySpecBackendAddressPool

Appears on:[ApplicationGatewaySpec](#applicationgatewayspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `fqdnList` | ***[]string***| ***(Optional)*** Deprecated|
| `fqdns` | ***[]string***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `ipAddressList` | ***[]string***| ***(Optional)*** Deprecated|
| `ipAddresses` | ***[]string***| ***(Optional)*** |
| `name` | ***string***||
## ApplicationGatewaySpecBackendHTTPSettings

Appears on:[ApplicationGatewaySpec](#applicationgatewayspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `affinityCookieName` | ***string***| ***(Optional)*** |
| `authenticationCertificate` | ***[[]ApplicationGatewaySpecBackendHTTPSettingsAuthenticationCertificate](#applicationgatewayspecbackendhttpsettingsauthenticationcertificate)***| ***(Optional)*** |
| `connectionDraining` | ***[[]ApplicationGatewaySpecBackendHTTPSettingsConnectionDraining](#applicationgatewayspecbackendhttpsettingsconnectiondraining)***| ***(Optional)*** |
| `cookieBasedAffinity` | ***string***||
| `hostName` | ***string***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `path` | ***string***| ***(Optional)*** |
| `pickHostNameFromBackendAddress` | ***bool***| ***(Optional)*** |
| `port` | ***int64***||
| `probeID` | ***string***| ***(Optional)*** |
| `probeName` | ***string***| ***(Optional)*** |
| `protocol` | ***string***||
| `requestTimeout` | ***int64***| ***(Optional)*** |
## ApplicationGatewaySpecBackendHTTPSettingsAuthenticationCertificate

Appears on:[ApplicationGatewaySpecBackendHTTPSettings](#applicationgatewayspecbackendhttpsettings)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
## ApplicationGatewaySpecBackendHTTPSettingsConnectionDraining

Appears on:[ApplicationGatewaySpecBackendHTTPSettings](#applicationgatewayspecbackendhttpsettings)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `drainTimeoutSec` | ***int64***||
| `enabled` | ***bool***||
## ApplicationGatewaySpecCustomErrorConfiguration

Appears on:[ApplicationGatewaySpec](#applicationgatewayspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `customErrorPageURL` | ***string***||
| `ID` | ***string***| ***(Optional)*** |
| `statusCode` | ***string***||
## ApplicationGatewaySpecFrontendIPConfiguration

Appears on:[ApplicationGatewaySpec](#applicationgatewayspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `privateIPAddress` | ***string***| ***(Optional)*** |
| `privateIPAddressAllocation` | ***string***| ***(Optional)*** |
| `publicIPAddressID` | ***string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
## ApplicationGatewaySpecFrontendPort

Appears on:[ApplicationGatewaySpec](#applicationgatewayspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `port` | ***int64***||
## ApplicationGatewaySpecGatewayIPConfiguration

Appears on:[ApplicationGatewaySpec](#applicationgatewayspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `subnetID` | ***string***||
## ApplicationGatewaySpecHttpListener

Appears on:[ApplicationGatewaySpec](#applicationgatewayspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `customErrorConfiguration` | ***[[]ApplicationGatewaySpecHttpListenerCustomErrorConfiguration](#applicationgatewayspechttplistenercustomerrorconfiguration)***| ***(Optional)*** |
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

Appears on:[ApplicationGatewaySpecHttpListener](#applicationgatewayspechttplistener)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `customErrorPageURL` | ***string***||
| `ID` | ***string***| ***(Optional)*** |
| `statusCode` | ***string***||
## ApplicationGatewaySpecIdentity

Appears on:[ApplicationGatewaySpec](#applicationgatewayspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `identityIDS` | ***[]string***||
| `type` | ***string***| ***(Optional)*** |
## ApplicationGatewaySpecProbe

Appears on:[ApplicationGatewaySpec](#applicationgatewayspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `host` | ***string***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `interval` | ***int64***||
| `match` | ***[[]ApplicationGatewaySpecProbeMatch](#applicationgatewayspecprobematch)***| ***(Optional)*** |
| `minimumServers` | ***int64***| ***(Optional)*** |
| `name` | ***string***||
| `path` | ***string***||
| `pickHostNameFromBackendHTTPSettings` | ***bool***| ***(Optional)*** |
| `protocol` | ***string***||
| `timeout` | ***int64***||
| `unhealthyThreshold` | ***int64***||
## ApplicationGatewaySpecProbeMatch

Appears on:[ApplicationGatewaySpecProbe](#applicationgatewayspecprobe)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `body` | ***string***| ***(Optional)*** |
| `statusCode` | ***[]string***| ***(Optional)*** |
## ApplicationGatewaySpecRedirectConfiguration

Appears on:[ApplicationGatewaySpec](#applicationgatewayspec)

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

Appears on:[ApplicationGatewaySpec](#applicationgatewayspec)

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

Appears on:[ApplicationGatewaySpec](#applicationgatewayspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `rewriteRule` | ***[[]ApplicationGatewaySpecRewriteRuleSetRewriteRule](#applicationgatewayspecrewriterulesetrewriterule)***| ***(Optional)*** |
## ApplicationGatewaySpecRewriteRuleSetRewriteRule

Appears on:[ApplicationGatewaySpecRewriteRuleSet](#applicationgatewayspecrewriteruleset)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `condition` | ***[[]ApplicationGatewaySpecRewriteRuleSetRewriteRuleCondition](#applicationgatewayspecrewriterulesetrewriterulecondition)***| ***(Optional)*** |
| `name` | ***string***||
| `requestHeaderConfiguration` | ***[[]ApplicationGatewaySpecRewriteRuleSetRewriteRuleRequestHeaderConfiguration](#applicationgatewayspecrewriterulesetrewriterulerequestheaderconfiguration)***| ***(Optional)*** |
| `responseHeaderConfiguration` | ***[[]ApplicationGatewaySpecRewriteRuleSetRewriteRuleResponseHeaderConfiguration](#applicationgatewayspecrewriterulesetrewriteruleresponseheaderconfiguration)***| ***(Optional)*** |
| `ruleSequence` | ***int64***||
## ApplicationGatewaySpecRewriteRuleSetRewriteRuleCondition

Appears on:[ApplicationGatewaySpecRewriteRuleSetRewriteRule](#applicationgatewayspecrewriterulesetrewriterule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ignoreCase` | ***bool***| ***(Optional)*** |
| `negate` | ***bool***| ***(Optional)*** |
| `pattern` | ***string***||
| `variable` | ***string***||
## ApplicationGatewaySpecRewriteRuleSetRewriteRuleRequestHeaderConfiguration

Appears on:[ApplicationGatewaySpecRewriteRuleSetRewriteRule](#applicationgatewayspecrewriterulesetrewriterule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `headerName` | ***string***||
| `headerValue` | ***string***||
## ApplicationGatewaySpecRewriteRuleSetRewriteRuleResponseHeaderConfiguration

Appears on:[ApplicationGatewaySpecRewriteRuleSetRewriteRule](#applicationgatewayspecrewriterulesetrewriterule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `headerName` | ***string***||
| `headerValue` | ***string***||
## ApplicationGatewaySpecSku

Appears on:[ApplicationGatewaySpec](#applicationgatewayspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `capacity` | ***int64***| ***(Optional)*** |
| `name` | ***string***||
| `tier` | ***string***||
## ApplicationGatewaySpecSslCertificate

Appears on:[ApplicationGatewaySpec](#applicationgatewayspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `publicCertData` | ***string***| ***(Optional)*** |
## ApplicationGatewaySpecSslPolicy

Appears on:[ApplicationGatewaySpec](#applicationgatewayspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cipherSuites` | ***[]string***| ***(Optional)*** |
| `disabledProtocols` | ***[]string***| ***(Optional)*** |
| `minProtocolVersion` | ***string***| ***(Optional)*** |
| `policyName` | ***string***| ***(Optional)*** |
| `policyType` | ***string***| ***(Optional)*** |
## ApplicationGatewaySpecUrlPathMap

Appears on:[ApplicationGatewaySpec](#applicationgatewayspec)

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
| `pathRule` | ***[[]ApplicationGatewaySpecUrlPathMapPathRule](#applicationgatewayspecurlpathmappathrule)***||
## ApplicationGatewaySpecUrlPathMapPathRule

Appears on:[ApplicationGatewaySpecUrlPathMap](#applicationgatewayspecurlpathmap)

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

Appears on:[ApplicationGatewaySpec](#applicationgatewayspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `disabledRuleGroup` | ***[[]ApplicationGatewaySpecWafConfigurationDisabledRuleGroup](#applicationgatewayspecwafconfigurationdisabledrulegroup)***| ***(Optional)*** |
| `enabled` | ***bool***||
| `exclusion` | ***[[]ApplicationGatewaySpecWafConfigurationExclusion](#applicationgatewayspecwafconfigurationexclusion)***| ***(Optional)*** |
| `fileUploadLimitMb` | ***int64***| ***(Optional)*** |
| `firewallMode` | ***string***||
| `maxRequestBodySizeKb` | ***int64***| ***(Optional)*** |
| `requestBodyCheck` | ***bool***| ***(Optional)*** |
| `ruleSetType` | ***string***| ***(Optional)*** |
| `ruleSetVersion` | ***string***||
## ApplicationGatewaySpecWafConfigurationDisabledRuleGroup

Appears on:[ApplicationGatewaySpecWafConfiguration](#applicationgatewayspecwafconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ruleGroupName` | ***string***||
| `rules` | ***[]int64***| ***(Optional)*** |
## ApplicationGatewaySpecWafConfigurationExclusion

Appears on:[ApplicationGatewaySpecWafConfiguration](#applicationgatewayspecwafconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `matchVariable` | ***string***||
| `selector` | ***string***| ***(Optional)*** |
| `selectorMatchOperator` | ***string***| ***(Optional)*** |
## ApplicationGatewayStatus

Appears on:[ApplicationGateway](#applicationgateway)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApplicationGatewaySpec](#applicationgatewayspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ApplicationGatewayStatus](#applicationgatewaystatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `authentication_certificate.<index>.data` | ***string*** ||
| `ssl_certificate.<index>.data` | ***string*** ||
| `ssl_certificate.<index>.password` | ***string*** ||
