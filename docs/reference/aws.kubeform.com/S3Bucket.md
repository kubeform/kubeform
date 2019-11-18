---
title: S3Bucket
menu:
  docs_{{ .version }}:
    identifier: s3bucket-aws.kubeform.com
    name: S3Bucket
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## S3Bucket
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `S3Bucket` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[S3BucketSpec](#s3bucketspec)***||
| `status` | ***[S3BucketStatus](#s3bucketstatus)***||
## Phase(`string` alias)

Appears on:[S3BucketStatus](#s3bucketstatus)

## S3BucketSpec

Appears on:[S3Bucket](#s3bucket), [S3BucketStatus](#s3bucketstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accelerationStatus` | ***string***| ***(Optional)*** |
| `acl` | ***string***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `bucket` | ***string***| ***(Optional)*** |
| `bucketDomainName` | ***string***| ***(Optional)*** |
| `bucketPrefix` | ***string***| ***(Optional)*** |
| `bucketRegionalDomainName` | ***string***| ***(Optional)*** |
| `corsRule` | ***[[]S3BucketSpecCorsRule](#s3bucketspeccorsrule)***| ***(Optional)*** |
| `forceDestroy` | ***bool***| ***(Optional)*** |
| `hostedZoneID` | ***string***| ***(Optional)*** |
| `lifecycleRule` | ***[[]S3BucketSpecLifecycleRule](#s3bucketspeclifecyclerule)***| ***(Optional)*** |
| `logging` | ***[[]S3BucketSpecLogging](#s3bucketspeclogging)***| ***(Optional)*** |
| `objectLockConfiguration` | ***[[]S3BucketSpecObjectLockConfiguration](#s3bucketspecobjectlockconfiguration)***| ***(Optional)*** |
| `policy` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `replicationConfiguration` | ***[[]S3BucketSpecReplicationConfiguration](#s3bucketspecreplicationconfiguration)***| ***(Optional)*** |
| `requestPayer` | ***string***| ***(Optional)*** |
| `serverSideEncryptionConfiguration` | ***[[]S3BucketSpecServerSideEncryptionConfiguration](#s3bucketspecserversideencryptionconfiguration)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `versioning` | ***[[]S3BucketSpecVersioning](#s3bucketspecversioning)***| ***(Optional)*** |
| `website` | ***[[]S3BucketSpecWebsite](#s3bucketspecwebsite)***| ***(Optional)*** |
| `websiteDomain` | ***string***| ***(Optional)*** |
| `websiteEndpoint` | ***string***| ***(Optional)*** |
## S3BucketSpecCorsRule

Appears on:[S3BucketSpec](#s3bucketspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `allowedHeaders` | ***[]string***| ***(Optional)*** |
| `allowedMethods` | ***[]string***||
| `allowedOrigins` | ***[]string***||
| `exposeHeaders` | ***[]string***| ***(Optional)*** |
| `maxAgeSeconds` | ***int64***| ***(Optional)*** |
## S3BucketSpecLifecycleRule

Appears on:[S3BucketSpec](#s3bucketspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `abortIncompleteMultipartUploadDays` | ***int64***| ***(Optional)*** |
| `enabled` | ***bool***||
| `expiration` | ***[[]S3BucketSpecLifecycleRuleExpiration](#s3bucketspeclifecycleruleexpiration)***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `noncurrentVersionExpiration` | ***[[]S3BucketSpecLifecycleRuleNoncurrentVersionExpiration](#s3bucketspeclifecyclerulenoncurrentversionexpiration)***| ***(Optional)*** |
| `noncurrentVersionTransition` | ***[[]S3BucketSpecLifecycleRuleNoncurrentVersionTransition](#s3bucketspeclifecyclerulenoncurrentversiontransition)***| ***(Optional)*** |
| `prefix` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `transition` | ***[[]S3BucketSpecLifecycleRuleTransition](#s3bucketspeclifecycleruletransition)***| ***(Optional)*** |
## S3BucketSpecLifecycleRuleExpiration

Appears on:[S3BucketSpecLifecycleRule](#s3bucketspeclifecyclerule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `date` | ***string***| ***(Optional)*** |
| `days` | ***int64***| ***(Optional)*** |
| `expiredObjectDeleteMarker` | ***bool***| ***(Optional)*** |
## S3BucketSpecLifecycleRuleNoncurrentVersionExpiration

Appears on:[S3BucketSpecLifecycleRule](#s3bucketspeclifecyclerule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `days` | ***int64***| ***(Optional)*** |
## S3BucketSpecLifecycleRuleNoncurrentVersionTransition

Appears on:[S3BucketSpecLifecycleRule](#s3bucketspeclifecyclerule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `days` | ***int64***| ***(Optional)*** |
| `storageClass` | ***string***||
## S3BucketSpecLifecycleRuleTransition

Appears on:[S3BucketSpecLifecycleRule](#s3bucketspeclifecyclerule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `date` | ***string***| ***(Optional)*** |
| `days` | ***int64***| ***(Optional)*** |
| `storageClass` | ***string***||
## S3BucketSpecLogging

Appears on:[S3BucketSpec](#s3bucketspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `targetBucket` | ***string***||
| `targetPrefix` | ***string***| ***(Optional)*** |
## S3BucketSpecObjectLockConfiguration

Appears on:[S3BucketSpec](#s3bucketspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `objectLockEnabled` | ***string***||
| `rule` | ***[[]S3BucketSpecObjectLockConfigurationRule](#s3bucketspecobjectlockconfigurationrule)***| ***(Optional)*** |
## S3BucketSpecObjectLockConfigurationRule

Appears on:[S3BucketSpecObjectLockConfiguration](#s3bucketspecobjectlockconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `defaultRetention` | ***[[]S3BucketSpecObjectLockConfigurationRuleDefaultRetention](#s3bucketspecobjectlockconfigurationruledefaultretention)***||
## S3BucketSpecObjectLockConfigurationRuleDefaultRetention

Appears on:[S3BucketSpecObjectLockConfigurationRule](#s3bucketspecobjectlockconfigurationrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `days` | ***int64***| ***(Optional)*** |
| `mode` | ***string***||
| `years` | ***int64***| ***(Optional)*** |
## S3BucketSpecReplicationConfiguration

Appears on:[S3BucketSpec](#s3bucketspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `role` | ***string***||
| `rules` | ***[[]S3BucketSpecReplicationConfigurationRules](#s3bucketspecreplicationconfigurationrules)***||
## S3BucketSpecReplicationConfigurationRules

Appears on:[S3BucketSpecReplicationConfiguration](#s3bucketspecreplicationconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `destination` | ***[[]S3BucketSpecReplicationConfigurationRulesDestination](#s3bucketspecreplicationconfigurationrulesdestination)***||
| `filter` | ***[[]S3BucketSpecReplicationConfigurationRulesFilter](#s3bucketspecreplicationconfigurationrulesfilter)***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `prefix` | ***string***| ***(Optional)*** |
| `priority` | ***int64***| ***(Optional)*** |
| `sourceSelectionCriteria` | ***[[]S3BucketSpecReplicationConfigurationRulesSourceSelectionCriteria](#s3bucketspecreplicationconfigurationrulessourceselectioncriteria)***| ***(Optional)*** |
| `status` | ***string***||
## S3BucketSpecReplicationConfigurationRulesDestination

Appears on:[S3BucketSpecReplicationConfigurationRules](#s3bucketspecreplicationconfigurationrules)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `accessControlTranslation` | ***[[]S3BucketSpecReplicationConfigurationRulesDestinationAccessControlTranslation](#s3bucketspecreplicationconfigurationrulesdestinationaccesscontroltranslation)***| ***(Optional)*** |
| `accountID` | ***string***| ***(Optional)*** |
| `bucket` | ***string***||
| `replicaKmsKeyID` | ***string***| ***(Optional)*** |
| `storageClass` | ***string***| ***(Optional)*** |
## S3BucketSpecReplicationConfigurationRulesDestinationAccessControlTranslation

Appears on:[S3BucketSpecReplicationConfigurationRulesDestination](#s3bucketspecreplicationconfigurationrulesdestination)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `owner` | ***string***||
## S3BucketSpecReplicationConfigurationRulesFilter

Appears on:[S3BucketSpecReplicationConfigurationRules](#s3bucketspecreplicationconfigurationrules)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `prefix` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## S3BucketSpecReplicationConfigurationRulesSourceSelectionCriteria

Appears on:[S3BucketSpecReplicationConfigurationRules](#s3bucketspecreplicationconfigurationrules)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `sseKmsEncryptedObjects` | ***[[]S3BucketSpecReplicationConfigurationRulesSourceSelectionCriteriaSseKmsEncryptedObjects](#s3bucketspecreplicationconfigurationrulessourceselectioncriteriassekmsencryptedobjects)***| ***(Optional)*** |
## S3BucketSpecReplicationConfigurationRulesSourceSelectionCriteriaSseKmsEncryptedObjects

Appears on:[S3BucketSpecReplicationConfigurationRulesSourceSelectionCriteria](#s3bucketspecreplicationconfigurationrulessourceselectioncriteria)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
## S3BucketSpecServerSideEncryptionConfiguration

Appears on:[S3BucketSpec](#s3bucketspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `rule` | ***[[]S3BucketSpecServerSideEncryptionConfigurationRule](#s3bucketspecserversideencryptionconfigurationrule)***||
## S3BucketSpecServerSideEncryptionConfigurationRule

Appears on:[S3BucketSpecServerSideEncryptionConfiguration](#s3bucketspecserversideencryptionconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `applyServerSideEncryptionByDefault` | ***[[]S3BucketSpecServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault](#s3bucketspecserversideencryptionconfigurationruleapplyserversideencryptionbydefault)***||
## S3BucketSpecServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault

Appears on:[S3BucketSpecServerSideEncryptionConfigurationRule](#s3bucketspecserversideencryptionconfigurationrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `kmsMasterKeyID` | ***string***| ***(Optional)*** |
| `sseAlgorithm` | ***string***||
## S3BucketSpecVersioning

Appears on:[S3BucketSpec](#s3bucketspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `mfaDelete` | ***bool***| ***(Optional)*** |
## S3BucketSpecWebsite

Appears on:[S3BucketSpec](#s3bucketspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `errorDocument` | ***string***| ***(Optional)*** |
| `indexDocument` | ***string***| ***(Optional)*** |
| `redirectAllRequestsTo` | ***string***| ***(Optional)*** |
| `routingRules` | ***string***| ***(Optional)*** |
## S3BucketStatus

Appears on:[S3Bucket](#s3bucket)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[S3BucketSpec](#s3bucketspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
