---
title: ElasticsearchDomain
menu:
  docs_{{ .version }}:
    identifier: elasticsearchdomain-aws.kubeform.com
    name: ElasticsearchDomain
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ElasticsearchDomain
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ElasticsearchDomain` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ElasticsearchDomainSpec](#elasticsearchdomainspec)***||
| `status` | ***[ElasticsearchDomainStatus](#elasticsearchdomainstatus)***||
## ElasticsearchDomainSpec

Appears on:[ElasticsearchDomain](#elasticsearchdomain), [ElasticsearchDomainStatus](#elasticsearchdomainstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accessPolicies` | ***string***| ***(Optional)*** |
| `advancedOptions` | ***map[string]string***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `clusterConfig` | ***[[]ElasticsearchDomainSpecClusterConfig](#elasticsearchdomainspecclusterconfig)***| ***(Optional)*** |
| `cognitoOptions` | ***[[]ElasticsearchDomainSpecCognitoOptions](#elasticsearchdomainspeccognitooptions)***| ***(Optional)*** |
| `domainID` | ***string***| ***(Optional)*** |
| `domainName` | ***string***||
| `ebsOptions` | ***[[]ElasticsearchDomainSpecEbsOptions](#elasticsearchdomainspecebsoptions)***| ***(Optional)*** |
| `elasticsearchVersion` | ***string***| ***(Optional)*** |
| `encryptAtRest` | ***[[]ElasticsearchDomainSpecEncryptAtRest](#elasticsearchdomainspecencryptatrest)***| ***(Optional)*** |
| `endpoint` | ***string***| ***(Optional)*** |
| `kibanaEndpoint` | ***string***| ***(Optional)*** |
| `logPublishingOptions` | ***[[]ElasticsearchDomainSpecLogPublishingOptions](#elasticsearchdomainspeclogpublishingoptions)***| ***(Optional)*** |
| `nodeToNodeEncryption` | ***[[]ElasticsearchDomainSpecNodeToNodeEncryption](#elasticsearchdomainspecnodetonodeencryption)***| ***(Optional)*** |
| `snapshotOptions` | ***[[]ElasticsearchDomainSpecSnapshotOptions](#elasticsearchdomainspecsnapshotoptions)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpcOptions` | ***[[]ElasticsearchDomainSpecVpcOptions](#elasticsearchdomainspecvpcoptions)***| ***(Optional)*** |
## ElasticsearchDomainSpecClusterConfig

Appears on:[ElasticsearchDomainSpec](#elasticsearchdomainspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `dedicatedMasterCount` | ***int***| ***(Optional)*** |
| `dedicatedMasterEnabled` | ***bool***| ***(Optional)*** |
| `dedicatedMasterType` | ***string***| ***(Optional)*** |
| `instanceCount` | ***int***| ***(Optional)*** |
| `instanceType` | ***string***| ***(Optional)*** |
| `zoneAwarenessEnabled` | ***bool***| ***(Optional)*** |
## ElasticsearchDomainSpecCognitoOptions

Appears on:[ElasticsearchDomainSpec](#elasticsearchdomainspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `identityPoolID` | ***string***||
| `roleArn` | ***string***||
| `userPoolID` | ***string***||
## ElasticsearchDomainSpecEbsOptions

Appears on:[ElasticsearchDomainSpec](#elasticsearchdomainspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ebsEnabled` | ***bool***||
| `iops` | ***int***| ***(Optional)*** |
| `volumeSize` | ***int***| ***(Optional)*** |
| `volumeType` | ***string***| ***(Optional)*** |
## ElasticsearchDomainSpecEncryptAtRest

Appears on:[ElasticsearchDomainSpec](#elasticsearchdomainspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
| `kmsKeyID` | ***string***| ***(Optional)*** |
## ElasticsearchDomainSpecLogPublishingOptions

Appears on:[ElasticsearchDomainSpec](#elasticsearchdomainspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cloudwatchLogGroupArn` | ***string***||
| `enabled` | ***bool***| ***(Optional)*** |
| `logType` | ***string***||
## ElasticsearchDomainSpecNodeToNodeEncryption

Appears on:[ElasticsearchDomainSpec](#elasticsearchdomainspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
## ElasticsearchDomainSpecSnapshotOptions

Appears on:[ElasticsearchDomainSpec](#elasticsearchdomainspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `automatedSnapshotStartHour` | ***int***||
## ElasticsearchDomainSpecVpcOptions

Appears on:[ElasticsearchDomainSpec](#elasticsearchdomainspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `availabilityZones` | ***[]string***| ***(Optional)*** |
| `securityGroupIDS` | ***[]string***| ***(Optional)*** |
| `subnetIDS` | ***[]string***| ***(Optional)*** |
| `vpcID` | ***string***| ***(Optional)*** |
## ElasticsearchDomainStatus

Appears on:[ElasticsearchDomain](#elasticsearchdomain)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ElasticsearchDomainSpec](#elasticsearchdomainspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ElasticsearchDomainStatus](#elasticsearchdomainstatus)

---
