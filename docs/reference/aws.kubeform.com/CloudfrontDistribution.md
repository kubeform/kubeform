---
title: CloudfrontDistribution
menu:
  docs_{{ .version }}:
    identifier: cloudfrontdistribution-aws.kubeform.com
    name: CloudfrontDistribution
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CloudfrontDistribution
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CloudfrontDistribution` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CloudfrontDistributionSpec](#CloudfrontDistributionSpec)***||
| `status` | ***[CloudfrontDistributionStatus](#CloudfrontDistributionStatus)***||
## CloudfrontDistributionSpec

Appears on:[CloudfrontDistribution](#CloudfrontDistribution), [CloudfrontDistributionStatus](#CloudfrontDistributionStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `activeTrustedSigners` | ***map[string]string***| ***(Optional)*** |
| `aliases` | ***[]string***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `callerReference` | ***string***| ***(Optional)*** |
| `comment` | ***string***| ***(Optional)*** |
| `customErrorResponse` | ***[[]CloudfrontDistributionSpecCustomErrorResponse](#CloudfrontDistributionSpecCustomErrorResponse)***| ***(Optional)*** |
| `defaultCacheBehavior` | ***[[]CloudfrontDistributionSpecDefaultCacheBehavior](#CloudfrontDistributionSpecDefaultCacheBehavior)***||
| `defaultRootObject` | ***string***| ***(Optional)*** |
| `domainName` | ***string***| ***(Optional)*** |
| `enabled` | ***bool***||
| `etag` | ***string***| ***(Optional)*** |
| `hostedZoneID` | ***string***| ***(Optional)*** |
| `httpVersion` | ***string***| ***(Optional)*** |
| `inProgressValidationBatches` | ***int***| ***(Optional)*** |
| `isIpv6Enabled` | ***bool***| ***(Optional)*** |
| `lastModifiedTime` | ***string***| ***(Optional)*** |
| `loggingConfig` | ***[[]CloudfrontDistributionSpecLoggingConfig](#CloudfrontDistributionSpecLoggingConfig)***| ***(Optional)*** |
| `orderedCacheBehavior` | ***[[]CloudfrontDistributionSpecOrderedCacheBehavior](#CloudfrontDistributionSpecOrderedCacheBehavior)***| ***(Optional)*** |
| `origin` | ***[[]CloudfrontDistributionSpecOrigin](#CloudfrontDistributionSpecOrigin)***||
| `originGroup` | ***[[]CloudfrontDistributionSpecOriginGroup](#CloudfrontDistributionSpecOriginGroup)***| ***(Optional)*** |
| `priceClass` | ***string***| ***(Optional)*** |
| `restrictions` | ***[[]CloudfrontDistributionSpecRestrictions](#CloudfrontDistributionSpecRestrictions)***||
| `retainOnDelete` | ***bool***| ***(Optional)*** |
| `status` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `viewerCertificate` | ***[[]CloudfrontDistributionSpecViewerCertificate](#CloudfrontDistributionSpecViewerCertificate)***||
| `waitForDeployment` | ***bool***| ***(Optional)*** |
| `webACLID` | ***string***| ***(Optional)*** |
## CloudfrontDistributionSpecCustomErrorResponse

Appears on:[CloudfrontDistributionSpec](#CloudfrontDistributionSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `errorCachingMinTtl` | ***int***| ***(Optional)*** |
| `errorCode` | ***int***||
| `responseCode` | ***int***| ***(Optional)*** |
| `responsePagePath` | ***string***| ***(Optional)*** |
## CloudfrontDistributionSpecDefaultCacheBehavior

Appears on:[CloudfrontDistributionSpec](#CloudfrontDistributionSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `allowedMethods` | ***[]string***||
| `cachedMethods` | ***[]string***||
| `compress` | ***bool***| ***(Optional)*** |
| `defaultTtl` | ***int***| ***(Optional)*** |
| `fieldLevelEncryptionID` | ***string***| ***(Optional)*** |
| `forwardedValues` | ***[[]CloudfrontDistributionSpecDefaultCacheBehaviorForwardedValues](#CloudfrontDistributionSpecDefaultCacheBehaviorForwardedValues)***||
| `lambdaFunctionAssociation` | ***[[]CloudfrontDistributionSpecDefaultCacheBehaviorLambdaFunctionAssociation](#CloudfrontDistributionSpecDefaultCacheBehaviorLambdaFunctionAssociation)***| ***(Optional)*** |
| `maxTtl` | ***int***| ***(Optional)*** |
| `minTtl` | ***int***| ***(Optional)*** |
| `smoothStreaming` | ***bool***| ***(Optional)*** |
| `targetOriginID` | ***string***||
| `trustedSigners` | ***[]string***| ***(Optional)*** |
| `viewerProtocolPolicy` | ***string***||
## CloudfrontDistributionSpecDefaultCacheBehaviorForwardedValues

Appears on:[CloudfrontDistributionSpecDefaultCacheBehavior](#CloudfrontDistributionSpecDefaultCacheBehavior)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cookies` | ***[[]CloudfrontDistributionSpecDefaultCacheBehaviorForwardedValuesCookies](#CloudfrontDistributionSpecDefaultCacheBehaviorForwardedValuesCookies)***||
| `headers` | ***[]string***| ***(Optional)*** |
| `queryString` | ***bool***||
| `queryStringCacheKeys` | ***[]string***| ***(Optional)*** |
## CloudfrontDistributionSpecDefaultCacheBehaviorForwardedValuesCookies

Appears on:[CloudfrontDistributionSpecDefaultCacheBehaviorForwardedValues](#CloudfrontDistributionSpecDefaultCacheBehaviorForwardedValues)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `forward` | ***string***||
| `whitelistedNames` | ***[]string***| ***(Optional)*** |
## CloudfrontDistributionSpecDefaultCacheBehaviorLambdaFunctionAssociation

Appears on:[CloudfrontDistributionSpecDefaultCacheBehavior](#CloudfrontDistributionSpecDefaultCacheBehavior)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `eventType` | ***string***||
| `includeBody` | ***bool***| ***(Optional)*** |
| `lambdaArn` | ***string***||
## CloudfrontDistributionSpecLoggingConfig

Appears on:[CloudfrontDistributionSpec](#CloudfrontDistributionSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucket` | ***string***||
| `includeCookies` | ***bool***| ***(Optional)*** |
| `prefix` | ***string***| ***(Optional)*** |
## CloudfrontDistributionSpecOrderedCacheBehavior

Appears on:[CloudfrontDistributionSpec](#CloudfrontDistributionSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `allowedMethods` | ***[]string***||
| `cachedMethods` | ***[]string***||
| `compress` | ***bool***| ***(Optional)*** |
| `defaultTtl` | ***int***| ***(Optional)*** |
| `fieldLevelEncryptionID` | ***string***| ***(Optional)*** |
| `forwardedValues` | ***[[]CloudfrontDistributionSpecOrderedCacheBehaviorForwardedValues](#CloudfrontDistributionSpecOrderedCacheBehaviorForwardedValues)***||
| `lambdaFunctionAssociation` | ***[[]CloudfrontDistributionSpecOrderedCacheBehaviorLambdaFunctionAssociation](#CloudfrontDistributionSpecOrderedCacheBehaviorLambdaFunctionAssociation)***| ***(Optional)*** |
| `maxTtl` | ***int***| ***(Optional)*** |
| `minTtl` | ***int***| ***(Optional)*** |
| `pathPattern` | ***string***||
| `smoothStreaming` | ***bool***| ***(Optional)*** |
| `targetOriginID` | ***string***||
| `trustedSigners` | ***[]string***| ***(Optional)*** |
| `viewerProtocolPolicy` | ***string***||
## CloudfrontDistributionSpecOrderedCacheBehaviorForwardedValues

Appears on:[CloudfrontDistributionSpecOrderedCacheBehavior](#CloudfrontDistributionSpecOrderedCacheBehavior)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cookies` | ***[[]CloudfrontDistributionSpecOrderedCacheBehaviorForwardedValuesCookies](#CloudfrontDistributionSpecOrderedCacheBehaviorForwardedValuesCookies)***||
| `headers` | ***[]string***| ***(Optional)*** |
| `queryString` | ***bool***||
| `queryStringCacheKeys` | ***[]string***| ***(Optional)*** |
## CloudfrontDistributionSpecOrderedCacheBehaviorForwardedValuesCookies

Appears on:[CloudfrontDistributionSpecOrderedCacheBehaviorForwardedValues](#CloudfrontDistributionSpecOrderedCacheBehaviorForwardedValues)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `forward` | ***string***||
| `whitelistedNames` | ***[]string***| ***(Optional)*** |
## CloudfrontDistributionSpecOrderedCacheBehaviorLambdaFunctionAssociation

Appears on:[CloudfrontDistributionSpecOrderedCacheBehavior](#CloudfrontDistributionSpecOrderedCacheBehavior)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `eventType` | ***string***||
| `includeBody` | ***bool***| ***(Optional)*** |
| `lambdaArn` | ***string***||
## CloudfrontDistributionSpecOrigin

Appears on:[CloudfrontDistributionSpec](#CloudfrontDistributionSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `customHeader` | ***[[]CloudfrontDistributionSpecOriginCustomHeader](#CloudfrontDistributionSpecOriginCustomHeader)***| ***(Optional)*** |
| `customOriginConfig` | ***[[]CloudfrontDistributionSpecOriginCustomOriginConfig](#CloudfrontDistributionSpecOriginCustomOriginConfig)***| ***(Optional)*** |
| `domainName` | ***string***||
| `originID` | ***string***||
| `originPath` | ***string***| ***(Optional)*** |
| `s3OriginConfig` | ***[[]CloudfrontDistributionSpecOriginS3OriginConfig](#CloudfrontDistributionSpecOriginS3OriginConfig)***| ***(Optional)*** |
## CloudfrontDistributionSpecOriginCustomHeader

Appears on:[CloudfrontDistributionSpecOrigin](#CloudfrontDistributionSpecOrigin)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `value` | ***string***||
## CloudfrontDistributionSpecOriginCustomOriginConfig

Appears on:[CloudfrontDistributionSpecOrigin](#CloudfrontDistributionSpecOrigin)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `httpPort` | ***int***||
| `httpsPort` | ***int***||
| `originKeepaliveTimeout` | ***int***| ***(Optional)*** |
| `originProtocolPolicy` | ***string***||
| `originReadTimeout` | ***int***| ***(Optional)*** |
| `originSSLProtocols` | ***[]string***||
## CloudfrontDistributionSpecOriginGroup

Appears on:[CloudfrontDistributionSpec](#CloudfrontDistributionSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `failoverCriteria` | ***[[]CloudfrontDistributionSpecOriginGroupFailoverCriteria](#CloudfrontDistributionSpecOriginGroupFailoverCriteria)***||
| `member` | ***[[]CloudfrontDistributionSpecOriginGroupMember](#CloudfrontDistributionSpecOriginGroupMember)***||
| `originID` | ***string***||
## CloudfrontDistributionSpecOriginGroupFailoverCriteria

Appears on:[CloudfrontDistributionSpecOriginGroup](#CloudfrontDistributionSpecOriginGroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `statusCodes` | ***[]int64***||
## CloudfrontDistributionSpecOriginGroupMember

Appears on:[CloudfrontDistributionSpecOriginGroup](#CloudfrontDistributionSpecOriginGroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `originID` | ***string***||
## CloudfrontDistributionSpecOriginS3OriginConfig

Appears on:[CloudfrontDistributionSpecOrigin](#CloudfrontDistributionSpecOrigin)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `originAccessIdentity` | ***string***||
## CloudfrontDistributionSpecRestrictions

Appears on:[CloudfrontDistributionSpec](#CloudfrontDistributionSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `geoRestriction` | ***[[]CloudfrontDistributionSpecRestrictionsGeoRestriction](#CloudfrontDistributionSpecRestrictionsGeoRestriction)***||
## CloudfrontDistributionSpecRestrictionsGeoRestriction

Appears on:[CloudfrontDistributionSpecRestrictions](#CloudfrontDistributionSpecRestrictions)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `locations` | ***[]string***| ***(Optional)*** |
| `restrictionType` | ***string***||
## CloudfrontDistributionSpecViewerCertificate

Appears on:[CloudfrontDistributionSpec](#CloudfrontDistributionSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `acmCertificateArn` | ***string***| ***(Optional)*** |
| `cloudfrontDefaultCertificate` | ***bool***| ***(Optional)*** |
| `iamCertificateID` | ***string***| ***(Optional)*** |
| `minimumProtocolVersion` | ***string***| ***(Optional)*** |
| `sslSupportMethod` | ***string***| ***(Optional)*** |
## CloudfrontDistributionStatus

Appears on:[CloudfrontDistribution](#CloudfrontDistribution)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CloudfrontDistributionSpec](#CloudfrontDistributionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
