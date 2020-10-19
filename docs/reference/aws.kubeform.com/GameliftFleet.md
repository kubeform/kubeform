---
title: GameliftFleet
menu:
  docs_{{ .version }}:
    identifier: gameliftfleet-aws.kubeform.com
    name: GameliftFleet
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## GameliftFleet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `GameliftFleet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[GameliftFleetSpec](#gameliftfleetspec)***||
| `status` | ***[GameliftFleetStatus](#gameliftfleetstatus)***||
## GameliftFleetSpec

Appears on:[GameliftFleet](#gameliftfleet), [GameliftFleetStatus](#gameliftfleetstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `buildID` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `ec2InboundPermission` | ***[[]GameliftFleetSpecEc2InboundPermission](#gameliftfleetspecec2inboundpermission)***| ***(Optional)*** |
| `ec2InstanceType` | ***string***||
| `logPaths` | ***[]string***| ***(Optional)*** |
| `metricGroups` | ***[]string***| ***(Optional)*** |
| `name` | ***string***||
| `newGameSessionProtectionPolicy` | ***string***| ***(Optional)*** |
| `operatingSystem` | ***string***| ***(Optional)*** |
| `resourceCreationLimitPolicy` | ***[[]GameliftFleetSpecResourceCreationLimitPolicy](#gameliftfleetspecresourcecreationlimitpolicy)***| ***(Optional)*** |
| `runtimeConfiguration` | ***[[]GameliftFleetSpecRuntimeConfiguration](#gameliftfleetspecruntimeconfiguration)***| ***(Optional)*** |
## GameliftFleetSpecEc2InboundPermission

Appears on:[GameliftFleetSpec](#gameliftfleetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `fromPort` | ***int64***||
| `ipRange` | ***string***||
| `protocol` | ***string***||
| `toPort` | ***int64***||
## GameliftFleetSpecResourceCreationLimitPolicy

Appears on:[GameliftFleetSpec](#gameliftfleetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `newGameSessionsPerCreator` | ***int64***| ***(Optional)*** |
| `policyPeriodInMinutes` | ***int64***| ***(Optional)*** |
## GameliftFleetSpecRuntimeConfiguration

Appears on:[GameliftFleetSpec](#gameliftfleetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `gameSessionActivationTimeoutSeconds` | ***int64***| ***(Optional)*** |
| `maxConcurrentGameSessionActivations` | ***int64***| ***(Optional)*** |
| `serverProcess` | ***[[]GameliftFleetSpecRuntimeConfigurationServerProcess](#gameliftfleetspecruntimeconfigurationserverprocess)***| ***(Optional)*** |
## GameliftFleetSpecRuntimeConfigurationServerProcess

Appears on:[GameliftFleetSpecRuntimeConfiguration](#gameliftfleetspecruntimeconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `concurrentExecutions` | ***int64***||
| `launchPath` | ***string***||
| `parameters` | ***string***| ***(Optional)*** |
## GameliftFleetStatus

Appears on:[GameliftFleet](#gameliftfleet)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[GameliftFleetSpec](#gameliftfleetspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[GameliftFleetStatus](#gameliftfleetstatus)

---
