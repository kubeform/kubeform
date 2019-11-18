---
title: MskCluster
menu:
  docs_{{ .version }}:
    identifier: mskcluster-aws.kubeform.com
    name: MskCluster
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## MskCluster
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `MskCluster` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MskClusterSpec](#mskclusterspec)***||
| `status` | ***[MskClusterStatus](#mskclusterstatus)***||
## MskClusterSpec



Appears on:[MskCluster](#mskcluster), [MskClusterStatus](#mskclusterstatus)



| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `bootstrapBrokers` | ***string***| ***(Optional)*** |
| `bootstrapBrokersTls` | ***string***| ***(Optional)*** |
| `brokerNodeGroupInfo` | ***[[]MskClusterSpecBrokerNodeGroupInfo](#mskclusterspecbrokernodegroupinfo)***||
| `clientAuthentication` | ***[[]MskClusterSpecClientAuthentication](#mskclusterspecclientauthentication)***| ***(Optional)*** |
| `clusterName` | ***string***||
| `configurationInfo` | ***[[]MskClusterSpecConfigurationInfo](#mskclusterspecconfigurationinfo)***| ***(Optional)*** |
| `currentVersion` | ***string***| ***(Optional)*** |
| `encryptionInfo` | ***[[]MskClusterSpecEncryptionInfo](#mskclusterspecencryptioninfo)***| ***(Optional)*** |
| `enhancedMonitoring` | ***string***| ***(Optional)*** |
| `kafkaVersion` | ***string***||
| `numberOfBrokerNodes` | ***int***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `zookeeperConnectString` | ***string***| ***(Optional)*** |
## MskClusterSpecBrokerNodeGroupInfo



Appears on:[MskClusterSpec](#mskclusterspec)



| Field | Type | Description |
| ------ | ----- | ----------- |
| `azDistribution` | ***string***| ***(Optional)*** |
| `clientSubnets` | ***[]string***||
| `ebsVolumeSize` | ***int***||
| `instanceType` | ***string***||
| `securityGroups` | ***[]string***||
## MskClusterSpecClientAuthentication



Appears on:[MskClusterSpec](#mskclusterspec)



| Field | Type | Description |
| ------ | ----- | ----------- |
| `tls` | ***[[]MskClusterSpecClientAuthenticationTls](#mskclusterspecclientauthenticationtls)***| ***(Optional)*** |
## MskClusterSpecClientAuthenticationTls



Appears on:[MskClusterSpecClientAuthentication](#mskclusterspecclientauthentication)



| Field | Type | Description |
| ------ | ----- | ----------- |
| `certificateAuthorityArns` | ***[]string***| ***(Optional)*** |
## MskClusterSpecConfigurationInfo



Appears on:[MskClusterSpec](#mskclusterspec)



| Field | Type | Description |
| ------ | ----- | ----------- |
| `arn` | ***string***||
| `revision` | ***int***||
## MskClusterSpecEncryptionInfo



Appears on:[MskClusterSpec](#mskclusterspec)



| Field | Type | Description |
| ------ | ----- | ----------- |
| `encryptionAtRestKmsKeyArn` | ***string***| ***(Optional)*** |
| `encryptionInTransit` | ***[[]MskClusterSpecEncryptionInfoEncryptionInTransit](#mskclusterspecencryptioninfoencryptionintransit)***| ***(Optional)*** |
## MskClusterSpecEncryptionInfoEncryptionInTransit



Appears on:[MskClusterSpecEncryptionInfo](#mskclusterspecencryptioninfo)



| Field | Type | Description |
| ------ | ----- | ----------- |
| `clientBroker` | ***string***| ***(Optional)*** |
| `inCluster` | ***bool***| ***(Optional)*** |
## MskClusterStatus



Appears on:[MskCluster](#mskcluster)



| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MskClusterSpec](#mskclusterspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
