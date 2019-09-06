## BatchComputeEnvironment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `BatchComputeEnvironment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BatchComputeEnvironmentSpec](#BatchComputeEnvironmentSpec)***||
| `status` | ***[BatchComputeEnvironmentStatus](#BatchComputeEnvironmentStatus)***||
## BatchComputeEnvironmentSpec
##### (Appears on:[BatchComputeEnvironment](#BatchComputeEnvironment), [BatchComputeEnvironmentStatus](#BatchComputeEnvironmentStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `computeEnvironmentName` | ***string***||
| `computeResources` | ***[[]BatchComputeEnvironmentSpecComputeResources](#BatchComputeEnvironmentSpecComputeResources)***| ***(Optional)*** |
| `ecsClusterArn` | ***string***| ***(Optional)*** |
| `serviceRole` | ***string***||
| `state` | ***string***| ***(Optional)*** |
| `status` | ***string***| ***(Optional)*** |
| `statusReason` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## BatchComputeEnvironmentSpecComputeResources
##### (Appears on:[BatchComputeEnvironmentSpec](#BatchComputeEnvironmentSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `bidPercentage` | ***int***| ***(Optional)*** |
| `desiredVcpus` | ***int***| ***(Optional)*** |
| `ec2KeyPair` | ***string***| ***(Optional)*** |
| `imageID` | ***string***| ***(Optional)*** |
| `instanceRole` | ***string***||
| `instanceType` | ***[]string***||
| `launchTemplate` | ***[[]BatchComputeEnvironmentSpecComputeResourcesLaunchTemplate](#BatchComputeEnvironmentSpecComputeResourcesLaunchTemplate)***| ***(Optional)*** |
| `maxVcpus` | ***int***||
| `minVcpus` | ***int***||
| `securityGroupIDS` | ***[]string***||
| `spotIamFleetRole` | ***string***| ***(Optional)*** |
| `subnets` | ***[]string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `type` | ***string***||
## BatchComputeEnvironmentSpecComputeResourcesLaunchTemplate
##### (Appears on:[BatchComputeEnvironmentSpecComputeResources](#BatchComputeEnvironmentSpecComputeResources))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `launchTemplateID` | ***string***| ***(Optional)*** |
| `launchTemplateName` | ***string***| ***(Optional)*** |
| `version` | ***string***| ***(Optional)*** |
## BatchComputeEnvironmentStatus
##### (Appears on:[BatchComputeEnvironment](#BatchComputeEnvironment))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BatchComputeEnvironmentSpec](#BatchComputeEnvironmentSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
