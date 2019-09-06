## CodedeployDeploymentGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CodedeployDeploymentGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CodedeployDeploymentGroupSpec](#CodedeployDeploymentGroupSpec)***||
| `status` | ***[CodedeployDeploymentGroupStatus](#CodedeployDeploymentGroupStatus)***||
## CodedeployDeploymentGroupSpec
##### (Appears on:[CodedeployDeploymentGroup](#CodedeployDeploymentGroup), [CodedeployDeploymentGroupStatus](#CodedeployDeploymentGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `alarmConfiguration` | ***[[]CodedeployDeploymentGroupSpecAlarmConfiguration](#CodedeployDeploymentGroupSpecAlarmConfiguration)***| ***(Optional)*** |
| `appName` | ***string***||
| `autoRollbackConfiguration` | ***[[]CodedeployDeploymentGroupSpecAutoRollbackConfiguration](#CodedeployDeploymentGroupSpecAutoRollbackConfiguration)***| ***(Optional)*** |
| `autoscalingGroups` | ***[]string***| ***(Optional)*** |
| `blueGreenDeploymentConfig` | ***[[]CodedeployDeploymentGroupSpecBlueGreenDeploymentConfig](#CodedeployDeploymentGroupSpecBlueGreenDeploymentConfig)***| ***(Optional)*** |
| `deploymentConfigName` | ***string***| ***(Optional)*** |
| `deploymentGroupName` | ***string***||
| `deploymentStyle` | ***[[]CodedeployDeploymentGroupSpecDeploymentStyle](#CodedeployDeploymentGroupSpecDeploymentStyle)***| ***(Optional)*** |
| `ec2TagFilter` | ***[[]CodedeployDeploymentGroupSpecEc2TagFilter](#CodedeployDeploymentGroupSpecEc2TagFilter)***| ***(Optional)*** |
| `ec2TagSet` | ***[[]CodedeployDeploymentGroupSpecEc2TagSet](#CodedeployDeploymentGroupSpecEc2TagSet)***| ***(Optional)*** |
| `ecsService` | ***[[]CodedeployDeploymentGroupSpecEcsService](#CodedeployDeploymentGroupSpecEcsService)***| ***(Optional)*** |
| `loadBalancerInfo` | ***[[]CodedeployDeploymentGroupSpecLoadBalancerInfo](#CodedeployDeploymentGroupSpecLoadBalancerInfo)***| ***(Optional)*** |
| `onPremisesInstanceTagFilter` | ***[[]CodedeployDeploymentGroupSpecOnPremisesInstanceTagFilter](#CodedeployDeploymentGroupSpecOnPremisesInstanceTagFilter)***| ***(Optional)*** |
| `serviceRoleArn` | ***string***||
| `triggerConfiguration` | ***[[]CodedeployDeploymentGroupSpecTriggerConfiguration](#CodedeployDeploymentGroupSpecTriggerConfiguration)***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecAlarmConfiguration
##### (Appears on:[CodedeployDeploymentGroupSpec](#CodedeployDeploymentGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `alarms` | ***[]string***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `ignorePollAlarmFailure` | ***bool***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecAutoRollbackConfiguration
##### (Appears on:[CodedeployDeploymentGroupSpec](#CodedeployDeploymentGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `events` | ***[]string***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecBlueGreenDeploymentConfig
##### (Appears on:[CodedeployDeploymentGroupSpec](#CodedeployDeploymentGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `deploymentReadyOption` | ***[[]CodedeployDeploymentGroupSpecBlueGreenDeploymentConfigDeploymentReadyOption](#CodedeployDeploymentGroupSpecBlueGreenDeploymentConfigDeploymentReadyOption)***| ***(Optional)*** |
| `greenFleetProvisioningOption` | ***[[]CodedeployDeploymentGroupSpecBlueGreenDeploymentConfigGreenFleetProvisioningOption](#CodedeployDeploymentGroupSpecBlueGreenDeploymentConfigGreenFleetProvisioningOption)***| ***(Optional)*** |
| `terminateBlueInstancesOnDeploymentSuccess` | ***[[]CodedeployDeploymentGroupSpecBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess](#CodedeployDeploymentGroupSpecBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess)***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecBlueGreenDeploymentConfigDeploymentReadyOption
##### (Appears on:[CodedeployDeploymentGroupSpecBlueGreenDeploymentConfig](#CodedeployDeploymentGroupSpecBlueGreenDeploymentConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `actionOnTimeout` | ***string***| ***(Optional)*** |
| `waitTimeInMinutes` | ***int***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecBlueGreenDeploymentConfigGreenFleetProvisioningOption
##### (Appears on:[CodedeployDeploymentGroupSpecBlueGreenDeploymentConfig](#CodedeployDeploymentGroupSpecBlueGreenDeploymentConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***string***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess
##### (Appears on:[CodedeployDeploymentGroupSpecBlueGreenDeploymentConfig](#CodedeployDeploymentGroupSpecBlueGreenDeploymentConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***string***| ***(Optional)*** |
| `terminationWaitTimeInMinutes` | ***int***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecDeploymentStyle
##### (Appears on:[CodedeployDeploymentGroupSpec](#CodedeployDeploymentGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `deploymentOption` | ***string***| ***(Optional)*** |
| `deploymentType` | ***string***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecEc2TagFilter
##### (Appears on:[CodedeployDeploymentGroupSpec](#CodedeployDeploymentGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `key` | ***string***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
| `value` | ***string***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecEc2TagSet
##### (Appears on:[CodedeployDeploymentGroupSpec](#CodedeployDeploymentGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ec2TagFilter` | ***[[]CodedeployDeploymentGroupSpecEc2TagSetEc2TagFilter](#CodedeployDeploymentGroupSpecEc2TagSetEc2TagFilter)***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecEc2TagSetEc2TagFilter
##### (Appears on:[CodedeployDeploymentGroupSpecEc2TagSet](#CodedeployDeploymentGroupSpecEc2TagSet))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `key` | ***string***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
| `value` | ***string***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecEcsService
##### (Appears on:[CodedeployDeploymentGroupSpec](#CodedeployDeploymentGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `clusterName` | ***string***||
| `serviceName` | ***string***||
## CodedeployDeploymentGroupSpecLoadBalancerInfo
##### (Appears on:[CodedeployDeploymentGroupSpec](#CodedeployDeploymentGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `elbInfo` | ***[[]CodedeployDeploymentGroupSpecLoadBalancerInfoElbInfo](#CodedeployDeploymentGroupSpecLoadBalancerInfoElbInfo)***| ***(Optional)*** |
| `targetGroupInfo` | ***[[]CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupInfo](#CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupInfo)***| ***(Optional)*** |
| `targetGroupPairInfo` | ***[[]CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo](#CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo)***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecLoadBalancerInfoElbInfo
##### (Appears on:[CodedeployDeploymentGroupSpecLoadBalancerInfo](#CodedeployDeploymentGroupSpecLoadBalancerInfo))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupInfo
##### (Appears on:[CodedeployDeploymentGroupSpecLoadBalancerInfo](#CodedeployDeploymentGroupSpecLoadBalancerInfo))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo
##### (Appears on:[CodedeployDeploymentGroupSpecLoadBalancerInfo](#CodedeployDeploymentGroupSpecLoadBalancerInfo))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `prodTrafficRoute` | ***[[]CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute](#CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute)***||
| `targetGroup` | ***[[]CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTargetGroup](#CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTargetGroup)***||
| `testTrafficRoute` | ***[[]CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute](#CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute)***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute
##### (Appears on:[CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo](#CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `listenerArns` | ***[]string***||
## CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTargetGroup
##### (Appears on:[CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo](#CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
## CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute
##### (Appears on:[CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo](#CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `listenerArns` | ***[]string***||
## CodedeployDeploymentGroupSpecOnPremisesInstanceTagFilter
##### (Appears on:[CodedeployDeploymentGroupSpec](#CodedeployDeploymentGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `key` | ***string***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
| `value` | ***string***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecTriggerConfiguration
##### (Appears on:[CodedeployDeploymentGroupSpec](#CodedeployDeploymentGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `triggerEvents` | ***[]string***||
| `triggerName` | ***string***||
| `triggerTargetArn` | ***string***||
## CodedeployDeploymentGroupStatus
##### (Appears on:[CodedeployDeploymentGroup](#CodedeployDeploymentGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CodedeployDeploymentGroupSpec](#CodedeployDeploymentGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
