---
title: AutoscalingGroup
menu:
  docs_v0.0.1:
    identifier: autoscalinggroup-aws.kubeform.com
    name: AutoscalingGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## AutoscalingGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AutoscalingGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AutoscalingGroupSpec](#AutoscalingGroupSpec)***||
| `status` | ***[AutoscalingGroupStatus](#AutoscalingGroupStatus)***||
## AutoscalingGroupSpec
##### (Appears on:[AutoscalingGroup](#AutoscalingGroup), [AutoscalingGroupStatus](#AutoscalingGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `availabilityZones` | ***[]string***| ***(Optional)*** |
| `defaultCooldown` | ***int***| ***(Optional)*** |
| `desiredCapacity` | ***int***| ***(Optional)*** |
| `enabledMetrics` | ***[]string***| ***(Optional)*** |
| `forceDelete` | ***bool***| ***(Optional)*** |
| `healthCheckGracePeriod` | ***int***| ***(Optional)*** |
| `healthCheckType` | ***string***| ***(Optional)*** |
| `initialLifecycleHook` | ***[[]AutoscalingGroupSpecInitialLifecycleHook](#AutoscalingGroupSpecInitialLifecycleHook)***| ***(Optional)*** |
| `launchConfiguration` | ***string***| ***(Optional)*** |
| `launchTemplate` | ***[[]AutoscalingGroupSpecLaunchTemplate](#AutoscalingGroupSpecLaunchTemplate)***| ***(Optional)*** |
| `loadBalancers` | ***[]string***| ***(Optional)*** |
| `maxSize` | ***int***||
| `metricsGranularity` | ***string***| ***(Optional)*** |
| `minElbCapacity` | ***int***| ***(Optional)*** |
| `minSize` | ***int***||
| `mixedInstancesPolicy` | ***[[]AutoscalingGroupSpecMixedInstancesPolicy](#AutoscalingGroupSpecMixedInstancesPolicy)***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `placementGroup` | ***string***| ***(Optional)*** |
| `protectFromScaleIn` | ***bool***| ***(Optional)*** |
| `serviceLinkedRoleArn` | ***string***| ***(Optional)*** |
| `suspendedProcesses` | ***[]string***| ***(Optional)*** |
| `tag` | ***[[]AutoscalingGroupSpecTag](#AutoscalingGroupSpecTag)***| ***(Optional)*** |
| `targetGroupArns` | ***[]string***| ***(Optional)*** |
| `terminationPolicies` | ***[]string***| ***(Optional)*** |
| `vpcZoneIdentifier` | ***[]string***| ***(Optional)*** |
| `waitForCapacityTimeout` | ***string***| ***(Optional)*** |
| `waitForElbCapacity` | ***int***| ***(Optional)*** |
## AutoscalingGroupSpecInitialLifecycleHook
##### (Appears on:[AutoscalingGroupSpec](#AutoscalingGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `defaultResult` | ***string***| ***(Optional)*** |
| `heartbeatTimeout` | ***int***| ***(Optional)*** |
| `lifecycleTransition` | ***string***||
| `name` | ***string***||
| `notificationMetadata` | ***string***| ***(Optional)*** |
| `notificationTargetArn` | ***string***| ***(Optional)*** |
| `roleArn` | ***string***| ***(Optional)*** |
## AutoscalingGroupSpecLaunchTemplate
##### (Appears on:[AutoscalingGroupSpec](#AutoscalingGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `version` | ***string***| ***(Optional)*** |
## AutoscalingGroupSpecMixedInstancesPolicy
##### (Appears on:[AutoscalingGroupSpec](#AutoscalingGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `instancesDistribution` | ***[[]AutoscalingGroupSpecMixedInstancesPolicyInstancesDistribution](#AutoscalingGroupSpecMixedInstancesPolicyInstancesDistribution)***| ***(Optional)*** |
| `launchTemplate` | ***[[]AutoscalingGroupSpecMixedInstancesPolicyLaunchTemplate](#AutoscalingGroupSpecMixedInstancesPolicyLaunchTemplate)***||
## AutoscalingGroupSpecMixedInstancesPolicyInstancesDistribution
##### (Appears on:[AutoscalingGroupSpecMixedInstancesPolicy](#AutoscalingGroupSpecMixedInstancesPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `onDemandAllocationStrategy` | ***string***| ***(Optional)*** |
| `onDemandBaseCapacity` | ***int***| ***(Optional)*** |
| `onDemandPercentageAboveBaseCapacity` | ***int***| ***(Optional)*** |
| `spotAllocationStrategy` | ***string***| ***(Optional)*** |
| `spotInstancePools` | ***int***| ***(Optional)*** |
| `spotMaxPrice` | ***string***| ***(Optional)*** |
## AutoscalingGroupSpecMixedInstancesPolicyLaunchTemplate
##### (Appears on:[AutoscalingGroupSpecMixedInstancesPolicy](#AutoscalingGroupSpecMixedInstancesPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `launchTemplateSpecification` | ***[[]AutoscalingGroupSpecMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecification](#AutoscalingGroupSpecMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecification)***||
| `override` | ***[[]AutoscalingGroupSpecMixedInstancesPolicyLaunchTemplateOverride](#AutoscalingGroupSpecMixedInstancesPolicyLaunchTemplateOverride)***| ***(Optional)*** |
## AutoscalingGroupSpecMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecification
##### (Appears on:[AutoscalingGroupSpecMixedInstancesPolicyLaunchTemplate](#AutoscalingGroupSpecMixedInstancesPolicyLaunchTemplate))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `launchTemplateID` | ***string***| ***(Optional)*** |
| `launchTemplateName` | ***string***| ***(Optional)*** |
| `version` | ***string***| ***(Optional)*** |
## AutoscalingGroupSpecMixedInstancesPolicyLaunchTemplateOverride
##### (Appears on:[AutoscalingGroupSpecMixedInstancesPolicyLaunchTemplate](#AutoscalingGroupSpecMixedInstancesPolicyLaunchTemplate))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `instanceType` | ***string***| ***(Optional)*** |
## AutoscalingGroupSpecTag
##### (Appears on:[AutoscalingGroupSpec](#AutoscalingGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `key` | ***string***||
| `propagateAtLaunch` | ***bool***||
| `value` | ***string***||
## AutoscalingGroupStatus
##### (Appears on:[AutoscalingGroup](#AutoscalingGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AutoscalingGroupSpec](#AutoscalingGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
