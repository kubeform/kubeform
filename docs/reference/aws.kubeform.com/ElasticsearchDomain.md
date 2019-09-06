## ElasticsearchDomain
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ElasticsearchDomain` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ElasticsearchDomainSpec](#ElasticsearchDomainSpec)***||
| `status` | ***[ElasticsearchDomainStatus](#ElasticsearchDomainStatus)***||
## ElasticsearchDomainSpec
##### (Appears on:[ElasticsearchDomain](#ElasticsearchDomain), [ElasticsearchDomainStatus](#ElasticsearchDomainStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accessPolicies` | ***string***| ***(Optional)*** |
| `advancedOptions` | ***map[string]string***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `clusterConfig` | ***[[]ElasticsearchDomainSpecClusterConfig](#ElasticsearchDomainSpecClusterConfig)***| ***(Optional)*** |
| `cognitoOptions` | ***[[]ElasticsearchDomainSpecCognitoOptions](#ElasticsearchDomainSpecCognitoOptions)***| ***(Optional)*** |
| `domainID` | ***string***| ***(Optional)*** |
| `domainName` | ***string***||
| `ebsOptions` | ***[[]ElasticsearchDomainSpecEbsOptions](#ElasticsearchDomainSpecEbsOptions)***| ***(Optional)*** |
| `elasticsearchVersion` | ***string***| ***(Optional)*** |
| `encryptAtRest` | ***[[]ElasticsearchDomainSpecEncryptAtRest](#ElasticsearchDomainSpecEncryptAtRest)***| ***(Optional)*** |
| `endpoint` | ***string***| ***(Optional)*** |
| `kibanaEndpoint` | ***string***| ***(Optional)*** |
| `logPublishingOptions` | ***[[]ElasticsearchDomainSpecLogPublishingOptions](#ElasticsearchDomainSpecLogPublishingOptions)***| ***(Optional)*** |
| `nodeToNodeEncryption` | ***[[]ElasticsearchDomainSpecNodeToNodeEncryption](#ElasticsearchDomainSpecNodeToNodeEncryption)***| ***(Optional)*** |
| `snapshotOptions` | ***[[]ElasticsearchDomainSpecSnapshotOptions](#ElasticsearchDomainSpecSnapshotOptions)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpcOptions` | ***[[]ElasticsearchDomainSpecVpcOptions](#ElasticsearchDomainSpecVpcOptions)***| ***(Optional)*** |
## ElasticsearchDomainSpecClusterConfig
##### (Appears on:[ElasticsearchDomainSpec](#ElasticsearchDomainSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `dedicatedMasterCount` | ***int***| ***(Optional)*** |
| `dedicatedMasterEnabled` | ***bool***| ***(Optional)*** |
| `dedicatedMasterType` | ***string***| ***(Optional)*** |
| `instanceCount` | ***int***| ***(Optional)*** |
| `instanceType` | ***string***| ***(Optional)*** |
| `zoneAwarenessEnabled` | ***bool***| ***(Optional)*** |
## ElasticsearchDomainSpecCognitoOptions
##### (Appears on:[ElasticsearchDomainSpec](#ElasticsearchDomainSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `identityPoolID` | ***string***||
| `roleArn` | ***string***||
| `userPoolID` | ***string***||
## ElasticsearchDomainSpecEbsOptions
##### (Appears on:[ElasticsearchDomainSpec](#ElasticsearchDomainSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ebsEnabled` | ***bool***||
| `iops` | ***int***| ***(Optional)*** |
| `volumeSize` | ***int***| ***(Optional)*** |
| `volumeType` | ***string***| ***(Optional)*** |
## ElasticsearchDomainSpecEncryptAtRest
##### (Appears on:[ElasticsearchDomainSpec](#ElasticsearchDomainSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
| `kmsKeyID` | ***string***| ***(Optional)*** |
## ElasticsearchDomainSpecLogPublishingOptions
##### (Appears on:[ElasticsearchDomainSpec](#ElasticsearchDomainSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `cloudwatchLogGroupArn` | ***string***||
| `enabled` | ***bool***| ***(Optional)*** |
| `logType` | ***string***||
## ElasticsearchDomainSpecNodeToNodeEncryption
##### (Appears on:[ElasticsearchDomainSpec](#ElasticsearchDomainSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
## ElasticsearchDomainSpecSnapshotOptions
##### (Appears on:[ElasticsearchDomainSpec](#ElasticsearchDomainSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `automatedSnapshotStartHour` | ***int***||
## ElasticsearchDomainSpecVpcOptions
##### (Appears on:[ElasticsearchDomainSpec](#ElasticsearchDomainSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `availabilityZones` | ***[]string***| ***(Optional)*** |
| `securityGroupIDS` | ***[]string***| ***(Optional)*** |
| `subnetIDS` | ***[]string***| ***(Optional)*** |
| `vpcID` | ***string***| ***(Optional)*** |
## ElasticsearchDomainStatus
##### (Appears on:[ElasticsearchDomain](#ElasticsearchDomain))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ElasticsearchDomainSpec](#ElasticsearchDomainSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
