## EksCluster
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `EksCluster` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[EksClusterSpec](#EksClusterSpec)***||
| `status` | ***[EksClusterStatus](#EksClusterStatus)***||
## EksClusterSpec
##### (Appears on:[EksCluster](#EksCluster), [EksClusterStatus](#EksClusterStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `certificateAuthority` | ***[[]EksClusterSpecCertificateAuthority](#EksClusterSpecCertificateAuthority)***| ***(Optional)*** |
| `createdAt` | ***string***| ***(Optional)*** |
| `enabledClusterLogTypes` | ***[]string***| ***(Optional)*** |
| `endpoint` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `platformVersion` | ***string***| ***(Optional)*** |
| `roleArn` | ***string***||
| `version` | ***string***| ***(Optional)*** |
| `vpcConfig` | ***[[]EksClusterSpecVpcConfig](#EksClusterSpecVpcConfig)***||
## EksClusterSpecCertificateAuthority
##### (Appears on:[EksClusterSpec](#EksClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `data` | ***string***| ***(Optional)*** |
## EksClusterSpecVpcConfig
##### (Appears on:[EksClusterSpec](#EksClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `endpointPrivateAccess` | ***bool***| ***(Optional)*** |
| `endpointPublicAccess` | ***bool***| ***(Optional)*** |
| `securityGroupIDS` | ***[]string***| ***(Optional)*** |
| `subnetIDS` | ***[]string***||
| `vpcID` | ***string***| ***(Optional)*** |
## EksClusterStatus
##### (Appears on:[EksCluster](#EksCluster))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[EksClusterSpec](#EksClusterSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
