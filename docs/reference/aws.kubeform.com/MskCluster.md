## MskCluster
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `MskCluster` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MskClusterSpec](#MskClusterSpec)***||
| `status` | ***[MskClusterStatus](#MskClusterStatus)***||
## MskClusterSpec
##### (Appears on:[MskCluster](#MskCluster), [MskClusterStatus](#MskClusterStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `bootstrapBrokers` | ***string***| ***(Optional)*** |
| `bootstrapBrokersTls` | ***string***| ***(Optional)*** |
| `brokerNodeGroupInfo` | ***[[]MskClusterSpecBrokerNodeGroupInfo](#MskClusterSpecBrokerNodeGroupInfo)***||
| `clientAuthentication` | ***[[]MskClusterSpecClientAuthentication](#MskClusterSpecClientAuthentication)***| ***(Optional)*** |
| `clusterName` | ***string***||
| `configurationInfo` | ***[[]MskClusterSpecConfigurationInfo](#MskClusterSpecConfigurationInfo)***| ***(Optional)*** |
| `currentVersion` | ***string***| ***(Optional)*** |
| `encryptionInfo` | ***[[]MskClusterSpecEncryptionInfo](#MskClusterSpecEncryptionInfo)***| ***(Optional)*** |
| `enhancedMonitoring` | ***string***| ***(Optional)*** |
| `kafkaVersion` | ***string***||
| `numberOfBrokerNodes` | ***int***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `zookeeperConnectString` | ***string***| ***(Optional)*** |
## MskClusterSpecBrokerNodeGroupInfo
##### (Appears on:[MskClusterSpec](#MskClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `azDistribution` | ***string***| ***(Optional)*** |
| `clientSubnets` | ***[]string***||
| `ebsVolumeSize` | ***int***||
| `instanceType` | ***string***||
| `securityGroups` | ***[]string***||
## MskClusterSpecClientAuthentication
##### (Appears on:[MskClusterSpec](#MskClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `tls` | ***[[]MskClusterSpecClientAuthenticationTls](#MskClusterSpecClientAuthenticationTls)***| ***(Optional)*** |
## MskClusterSpecClientAuthenticationTls
##### (Appears on:[MskClusterSpecClientAuthentication](#MskClusterSpecClientAuthentication))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `certificateAuthorityArns` | ***[]string***| ***(Optional)*** |
## MskClusterSpecConfigurationInfo
##### (Appears on:[MskClusterSpec](#MskClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `arn` | ***string***||
| `revision` | ***int***||
## MskClusterSpecEncryptionInfo
##### (Appears on:[MskClusterSpec](#MskClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `encryptionAtRestKmsKeyArn` | ***string***| ***(Optional)*** |
| `encryptionInTransit` | ***[[]MskClusterSpecEncryptionInfoEncryptionInTransit](#MskClusterSpecEncryptionInfoEncryptionInTransit)***| ***(Optional)*** |
## MskClusterSpecEncryptionInfoEncryptionInTransit
##### (Appears on:[MskClusterSpecEncryptionInfo](#MskClusterSpecEncryptionInfo))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `clientBroker` | ***string***| ***(Optional)*** |
| `inCluster` | ***bool***| ***(Optional)*** |
## MskClusterStatus
##### (Appears on:[MskCluster](#MskCluster))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MskClusterSpec](#MskClusterSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
