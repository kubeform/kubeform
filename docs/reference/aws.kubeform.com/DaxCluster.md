## DaxCluster
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DaxCluster` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DaxClusterSpec](#DaxClusterSpec)***||
| `status` | ***[DaxClusterStatus](#DaxClusterStatus)***||
## DaxClusterSpec
##### (Appears on:[DaxCluster](#DaxCluster), [DaxClusterStatus](#DaxClusterStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `availabilityZones` | ***[]string***| ***(Optional)*** |
| `clusterAddress` | ***string***| ***(Optional)*** |
| `clusterName` | ***string***||
| `configurationEndpoint` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `iamRoleArn` | ***string***||
| `maintenanceWindow` | ***string***| ***(Optional)*** |
| `nodeType` | ***string***||
| `nodes` | ***[[]DaxClusterSpecNodes](#DaxClusterSpecNodes)***| ***(Optional)*** |
| `notificationTopicArn` | ***string***| ***(Optional)*** |
| `parameterGroupName` | ***string***| ***(Optional)*** |
| `port` | ***int***| ***(Optional)*** |
| `replicationFactor` | ***int***||
| `securityGroupIDS` | ***[]string***| ***(Optional)*** |
| `serverSideEncryption` | ***[[]DaxClusterSpecServerSideEncryption](#DaxClusterSpecServerSideEncryption)***| ***(Optional)*** |
| `subnetGroupName` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## DaxClusterSpecNodes
##### (Appears on:[DaxClusterSpec](#DaxClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `address` | ***string***| ***(Optional)*** |
| `availabilityZone` | ***string***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `port` | ***int***| ***(Optional)*** |
## DaxClusterSpecServerSideEncryption
##### (Appears on:[DaxClusterSpec](#DaxClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
## DaxClusterStatus
##### (Appears on:[DaxCluster](#DaxCluster))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DaxClusterSpec](#DaxClusterSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
