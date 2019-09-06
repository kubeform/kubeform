## GameliftFleet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `GameliftFleet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[GameliftFleetSpec](#GameliftFleetSpec)***||
| `status` | ***[GameliftFleetStatus](#GameliftFleetStatus)***||
## GameliftFleetSpec
##### (Appears on:[GameliftFleet](#GameliftFleet), [GameliftFleetStatus](#GameliftFleetStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `buildID` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `ec2InboundPermission` | ***[[]GameliftFleetSpecEc2InboundPermission](#GameliftFleetSpecEc2InboundPermission)***| ***(Optional)*** |
| `ec2InstanceType` | ***string***||
| `logPaths` | ***[]string***| ***(Optional)*** |
| `metricGroups` | ***[]string***| ***(Optional)*** |
| `name` | ***string***||
| `newGameSessionProtectionPolicy` | ***string***| ***(Optional)*** |
| `operatingSystem` | ***string***| ***(Optional)*** |
| `resourceCreationLimitPolicy` | ***[[]GameliftFleetSpecResourceCreationLimitPolicy](#GameliftFleetSpecResourceCreationLimitPolicy)***| ***(Optional)*** |
| `runtimeConfiguration` | ***[[]GameliftFleetSpecRuntimeConfiguration](#GameliftFleetSpecRuntimeConfiguration)***| ***(Optional)*** |
## GameliftFleetSpecEc2InboundPermission
##### (Appears on:[GameliftFleetSpec](#GameliftFleetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `fromPort` | ***int***||
| `ipRange` | ***string***||
| `protocol` | ***string***||
| `toPort` | ***int***||
## GameliftFleetSpecResourceCreationLimitPolicy
##### (Appears on:[GameliftFleetSpec](#GameliftFleetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `newGameSessionsPerCreator` | ***int***| ***(Optional)*** |
| `policyPeriodInMinutes` | ***int***| ***(Optional)*** |
## GameliftFleetSpecRuntimeConfiguration
##### (Appears on:[GameliftFleetSpec](#GameliftFleetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `gameSessionActivationTimeoutSeconds` | ***int***| ***(Optional)*** |
| `maxConcurrentGameSessionActivations` | ***int***| ***(Optional)*** |
| `serverProcess` | ***[[]GameliftFleetSpecRuntimeConfigurationServerProcess](#GameliftFleetSpecRuntimeConfigurationServerProcess)***| ***(Optional)*** |
## GameliftFleetSpecRuntimeConfigurationServerProcess
##### (Appears on:[GameliftFleetSpecRuntimeConfiguration](#GameliftFleetSpecRuntimeConfiguration))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `concurrentExecutions` | ***int***||
| `launchPath` | ***string***||
| `parameters` | ***string***| ***(Optional)*** |
## GameliftFleetStatus
##### (Appears on:[GameliftFleet](#GameliftFleet))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[GameliftFleetSpec](#GameliftFleetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
