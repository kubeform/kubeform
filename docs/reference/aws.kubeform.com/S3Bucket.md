## S3Bucket
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `S3Bucket` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[S3BucketSpec](#S3BucketSpec)***||
| `status` | ***[S3BucketStatus](#S3BucketStatus)***||
## S3BucketSpec
##### (Appears on:[S3Bucket](#S3Bucket), [S3BucketStatus](#S3BucketStatus))
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
| `corsRule` | ***[[]S3BucketSpecCorsRule](#S3BucketSpecCorsRule)***| ***(Optional)*** |
| `forceDestroy` | ***bool***| ***(Optional)*** |
| `hostedZoneID` | ***string***| ***(Optional)*** |
| `lifecycleRule` | ***[[]S3BucketSpecLifecycleRule](#S3BucketSpecLifecycleRule)***| ***(Optional)*** |
| `logging` | ***[[]S3BucketSpecLogging](#S3BucketSpecLogging)***| ***(Optional)*** |
| `objectLockConfiguration` | ***[[]S3BucketSpecObjectLockConfiguration](#S3BucketSpecObjectLockConfiguration)***| ***(Optional)*** |
| `policy` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `replicationConfiguration` | ***[[]S3BucketSpecReplicationConfiguration](#S3BucketSpecReplicationConfiguration)***| ***(Optional)*** |
| `requestPayer` | ***string***| ***(Optional)*** |
| `serverSideEncryptionConfiguration` | ***[[]S3BucketSpecServerSideEncryptionConfiguration](#S3BucketSpecServerSideEncryptionConfiguration)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `versioning` | ***[[]S3BucketSpecVersioning](#S3BucketSpecVersioning)***| ***(Optional)*** |
| `website` | ***[[]S3BucketSpecWebsite](#S3BucketSpecWebsite)***| ***(Optional)*** |
| `websiteDomain` | ***string***| ***(Optional)*** |
| `websiteEndpoint` | ***string***| ***(Optional)*** |
## S3BucketSpecCorsRule
##### (Appears on:[S3BucketSpec](#S3BucketSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `allowedHeaders` | ***[]string***| ***(Optional)*** |
| `allowedMethods` | ***[]string***||
| `allowedOrigins` | ***[]string***||
| `exposeHeaders` | ***[]string***| ***(Optional)*** |
| `maxAgeSeconds` | ***int***| ***(Optional)*** |
## S3BucketSpecLifecycleRule
##### (Appears on:[S3BucketSpec](#S3BucketSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `abortIncompleteMultipartUploadDays` | ***int***| ***(Optional)*** |
| `enabled` | ***bool***||
| `expiration` | ***[[]S3BucketSpecLifecycleRuleExpiration](#S3BucketSpecLifecycleRuleExpiration)***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `noncurrentVersionExpiration` | ***[[]S3BucketSpecLifecycleRuleNoncurrentVersionExpiration](#S3BucketSpecLifecycleRuleNoncurrentVersionExpiration)***| ***(Optional)*** |
| `noncurrentVersionTransition` | ***[[]S3BucketSpecLifecycleRuleNoncurrentVersionTransition](#S3BucketSpecLifecycleRuleNoncurrentVersionTransition)***| ***(Optional)*** |
| `prefix` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `transition` | ***[[]S3BucketSpecLifecycleRuleTransition](#S3BucketSpecLifecycleRuleTransition)***| ***(Optional)*** |
## S3BucketSpecLifecycleRuleExpiration
##### (Appears on:[S3BucketSpecLifecycleRule](#S3BucketSpecLifecycleRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `date` | ***string***| ***(Optional)*** |
| `days` | ***int***| ***(Optional)*** |
| `expiredObjectDeleteMarker` | ***bool***| ***(Optional)*** |
## S3BucketSpecLifecycleRuleNoncurrentVersionExpiration
##### (Appears on:[S3BucketSpecLifecycleRule](#S3BucketSpecLifecycleRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `days` | ***int***| ***(Optional)*** |
## S3BucketSpecLifecycleRuleNoncurrentVersionTransition
##### (Appears on:[S3BucketSpecLifecycleRule](#S3BucketSpecLifecycleRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `days` | ***int***| ***(Optional)*** |
| `storageClass` | ***string***||
## S3BucketSpecLifecycleRuleTransition
##### (Appears on:[S3BucketSpecLifecycleRule](#S3BucketSpecLifecycleRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `date` | ***string***| ***(Optional)*** |
| `days` | ***int***| ***(Optional)*** |
| `storageClass` | ***string***||
## S3BucketSpecLogging
##### (Appears on:[S3BucketSpec](#S3BucketSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `targetBucket` | ***string***||
| `targetPrefix` | ***string***| ***(Optional)*** |
## S3BucketSpecObjectLockConfiguration
##### (Appears on:[S3BucketSpec](#S3BucketSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `objectLockEnabled` | ***string***||
| `rule` | ***[[]S3BucketSpecObjectLockConfigurationRule](#S3BucketSpecObjectLockConfigurationRule)***| ***(Optional)*** |
## S3BucketSpecObjectLockConfigurationRule
##### (Appears on:[S3BucketSpecObjectLockConfiguration](#S3BucketSpecObjectLockConfiguration))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `defaultRetention` | ***[[]S3BucketSpecObjectLockConfigurationRuleDefaultRetention](#S3BucketSpecObjectLockConfigurationRuleDefaultRetention)***||
## S3BucketSpecObjectLockConfigurationRuleDefaultRetention
##### (Appears on:[S3BucketSpecObjectLockConfigurationRule](#S3BucketSpecObjectLockConfigurationRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `days` | ***int***| ***(Optional)*** |
| `mode` | ***string***||
| `years` | ***int***| ***(Optional)*** |
## S3BucketSpecReplicationConfiguration
##### (Appears on:[S3BucketSpec](#S3BucketSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `role` | ***string***||
| `rules` | ***[[]S3BucketSpecReplicationConfigurationRules](#S3BucketSpecReplicationConfigurationRules)***||
## S3BucketSpecReplicationConfigurationRules
##### (Appears on:[S3BucketSpecReplicationConfiguration](#S3BucketSpecReplicationConfiguration))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `destination` | ***[[]S3BucketSpecReplicationConfigurationRulesDestination](#S3BucketSpecReplicationConfigurationRulesDestination)***||
| `filter` | ***[[]S3BucketSpecReplicationConfigurationRulesFilter](#S3BucketSpecReplicationConfigurationRulesFilter)***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `prefix` | ***string***| ***(Optional)*** |
| `priority` | ***int***| ***(Optional)*** |
| `sourceSelectionCriteria` | ***[[]S3BucketSpecReplicationConfigurationRulesSourceSelectionCriteria](#S3BucketSpecReplicationConfigurationRulesSourceSelectionCriteria)***| ***(Optional)*** |
| `status` | ***string***||
## S3BucketSpecReplicationConfigurationRulesDestination
##### (Appears on:[S3BucketSpecReplicationConfigurationRules](#S3BucketSpecReplicationConfigurationRules))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `accessControlTranslation` | ***[[]S3BucketSpecReplicationConfigurationRulesDestinationAccessControlTranslation](#S3BucketSpecReplicationConfigurationRulesDestinationAccessControlTranslation)***| ***(Optional)*** |
| `accountID` | ***string***| ***(Optional)*** |
| `bucket` | ***string***||
| `replicaKmsKeyID` | ***string***| ***(Optional)*** |
| `storageClass` | ***string***| ***(Optional)*** |
## S3BucketSpecReplicationConfigurationRulesDestinationAccessControlTranslation
##### (Appears on:[S3BucketSpecReplicationConfigurationRulesDestination](#S3BucketSpecReplicationConfigurationRulesDestination))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `owner` | ***string***||
## S3BucketSpecReplicationConfigurationRulesFilter
##### (Appears on:[S3BucketSpecReplicationConfigurationRules](#S3BucketSpecReplicationConfigurationRules))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `prefix` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## S3BucketSpecReplicationConfigurationRulesSourceSelectionCriteria
##### (Appears on:[S3BucketSpecReplicationConfigurationRules](#S3BucketSpecReplicationConfigurationRules))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `sseKmsEncryptedObjects` | ***[[]S3BucketSpecReplicationConfigurationRulesSourceSelectionCriteriaSseKmsEncryptedObjects](#S3BucketSpecReplicationConfigurationRulesSourceSelectionCriteriaSseKmsEncryptedObjects)***| ***(Optional)*** |
## S3BucketSpecReplicationConfigurationRulesSourceSelectionCriteriaSseKmsEncryptedObjects
##### (Appears on:[S3BucketSpecReplicationConfigurationRulesSourceSelectionCriteria](#S3BucketSpecReplicationConfigurationRulesSourceSelectionCriteria))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
## S3BucketSpecServerSideEncryptionConfiguration
##### (Appears on:[S3BucketSpec](#S3BucketSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `rule` | ***[[]S3BucketSpecServerSideEncryptionConfigurationRule](#S3BucketSpecServerSideEncryptionConfigurationRule)***||
## S3BucketSpecServerSideEncryptionConfigurationRule
##### (Appears on:[S3BucketSpecServerSideEncryptionConfiguration](#S3BucketSpecServerSideEncryptionConfiguration))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `applyServerSideEncryptionByDefault` | ***[[]S3BucketSpecServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault](#S3BucketSpecServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault)***||
## S3BucketSpecServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefault
##### (Appears on:[S3BucketSpecServerSideEncryptionConfigurationRule](#S3BucketSpecServerSideEncryptionConfigurationRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `kmsMasterKeyID` | ***string***| ***(Optional)*** |
| `sseAlgorithm` | ***string***||
## S3BucketSpecVersioning
##### (Appears on:[S3BucketSpec](#S3BucketSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `mfaDelete` | ***bool***| ***(Optional)*** |
## S3BucketSpecWebsite
##### (Appears on:[S3BucketSpec](#S3BucketSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `errorDocument` | ***string***| ***(Optional)*** |
| `indexDocument` | ***string***| ***(Optional)*** |
| `redirectAllRequestsTo` | ***string***| ***(Optional)*** |
| `routingRules` | ***string***| ***(Optional)*** |
## S3BucketStatus
##### (Appears on:[S3Bucket](#S3Bucket))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[S3BucketSpec](#S3BucketSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
