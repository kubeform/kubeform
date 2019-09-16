---
title: CodedeployDeploymentGroup
menu:
  docs_{{ .version }}:
    identifier: codedeploydeploymentgroup-aws.kubeform.com
    name: CodedeployDeploymentGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CodedeployDeploymentGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CodedeployDeploymentGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CodedeployDeploymentGroupSpec](#codedeploydeploymentgroupspec)***||
| `status` | ***[CodedeployDeploymentGroupStatus](#codedeploydeploymentgroupstatus)***||
## CodedeployDeploymentGroupSpec

Appears on:[CodedeployDeploymentGroup](#codedeploydeploymentgroup), [CodedeployDeploymentGroupStatus](#codedeploydeploymentgroupstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `alarmConfiguration` | ***[[]CodedeployDeploymentGroupSpecAlarmConfiguration](#codedeploydeploymentgroupspecalarmconfiguration)***| ***(Optional)*** |
| `appName` | ***string***||
| `autoRollbackConfiguration` | ***[[]CodedeployDeploymentGroupSpecAutoRollbackConfiguration](#codedeploydeploymentgroupspecautorollbackconfiguration)***| ***(Optional)*** |
| `autoscalingGroups` | ***[]string***| ***(Optional)*** |
| `blueGreenDeploymentConfig` | ***[[]CodedeployDeploymentGroupSpecBlueGreenDeploymentConfig](#codedeploydeploymentgroupspecbluegreendeploymentconfig)***| ***(Optional)*** |
| `deploymentConfigName` | ***string***| ***(Optional)*** |
| `deploymentGroupName` | ***string***||
| `deploymentStyle` | ***[[]CodedeployDeploymentGroupSpecDeploymentStyle](#codedeploydeploymentgroupspecdeploymentstyle)***| ***(Optional)*** |
| `ec2TagFilter` | ***[[]CodedeployDeploymentGroupSpecEc2TagFilter](#codedeploydeploymentgroupspecec2tagfilter)***| ***(Optional)*** |
| `ec2TagSet` | ***[[]CodedeployDeploymentGroupSpecEc2TagSet](#codedeploydeploymentgroupspecec2tagset)***| ***(Optional)*** |
| `ecsService` | ***[[]CodedeployDeploymentGroupSpecEcsService](#codedeploydeploymentgroupspececsservice)***| ***(Optional)*** |
| `loadBalancerInfo` | ***[[]CodedeployDeploymentGroupSpecLoadBalancerInfo](#codedeploydeploymentgroupspecloadbalancerinfo)***| ***(Optional)*** |
| `onPremisesInstanceTagFilter` | ***[[]CodedeployDeploymentGroupSpecOnPremisesInstanceTagFilter](#codedeploydeploymentgroupspeconpremisesinstancetagfilter)***| ***(Optional)*** |
| `serviceRoleArn` | ***string***||
| `triggerConfiguration` | ***[[]CodedeployDeploymentGroupSpecTriggerConfiguration](#codedeploydeploymentgroupspectriggerconfiguration)***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecAlarmConfiguration

Appears on:[CodedeployDeploymentGroupSpec](#codedeploydeploymentgroupspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `alarms` | ***[]string***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `ignorePollAlarmFailure` | ***bool***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecAutoRollbackConfiguration

Appears on:[CodedeployDeploymentGroupSpec](#codedeploydeploymentgroupspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `events` | ***[]string***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecBlueGreenDeploymentConfig

Appears on:[CodedeployDeploymentGroupSpec](#codedeploydeploymentgroupspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `deploymentReadyOption` | ***[[]CodedeployDeploymentGroupSpecBlueGreenDeploymentConfigDeploymentReadyOption](#codedeploydeploymentgroupspecbluegreendeploymentconfigdeploymentreadyoption)***| ***(Optional)*** |
| `greenFleetProvisioningOption` | ***[[]CodedeployDeploymentGroupSpecBlueGreenDeploymentConfigGreenFleetProvisioningOption](#codedeploydeploymentgroupspecbluegreendeploymentconfiggreenfleetprovisioningoption)***| ***(Optional)*** |
| `terminateBlueInstancesOnDeploymentSuccess` | ***[[]CodedeployDeploymentGroupSpecBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess](#codedeploydeploymentgroupspecbluegreendeploymentconfigterminateblueinstancesondeploymentsuccess)***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecBlueGreenDeploymentConfigDeploymentReadyOption

Appears on:[CodedeployDeploymentGroupSpecBlueGreenDeploymentConfig](#codedeploydeploymentgroupspecbluegreendeploymentconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `actionOnTimeout` | ***string***| ***(Optional)*** |
| `waitTimeInMinutes` | ***int***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecBlueGreenDeploymentConfigGreenFleetProvisioningOption

Appears on:[CodedeployDeploymentGroupSpecBlueGreenDeploymentConfig](#codedeploydeploymentgroupspecbluegreendeploymentconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***string***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecBlueGreenDeploymentConfigTerminateBlueInstancesOnDeploymentSuccess

Appears on:[CodedeployDeploymentGroupSpecBlueGreenDeploymentConfig](#codedeploydeploymentgroupspecbluegreendeploymentconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***string***| ***(Optional)*** |
| `terminationWaitTimeInMinutes` | ***int***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecDeploymentStyle

Appears on:[CodedeployDeploymentGroupSpec](#codedeploydeploymentgroupspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `deploymentOption` | ***string***| ***(Optional)*** |
| `deploymentType` | ***string***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecEc2TagFilter

Appears on:[CodedeployDeploymentGroupSpec](#codedeploydeploymentgroupspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `key` | ***string***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
| `value` | ***string***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecEc2TagSet

Appears on:[CodedeployDeploymentGroupSpec](#codedeploydeploymentgroupspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ec2TagFilter` | ***[[]CodedeployDeploymentGroupSpecEc2TagSetEc2TagFilter](#codedeploydeploymentgroupspecec2tagsetec2tagfilter)***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecEc2TagSetEc2TagFilter

Appears on:[CodedeployDeploymentGroupSpecEc2TagSet](#codedeploydeploymentgroupspecec2tagset)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `key` | ***string***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
| `value` | ***string***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecEcsService

Appears on:[CodedeployDeploymentGroupSpec](#codedeploydeploymentgroupspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `clusterName` | ***string***||
| `serviceName` | ***string***||
## CodedeployDeploymentGroupSpecLoadBalancerInfo

Appears on:[CodedeployDeploymentGroupSpec](#codedeploydeploymentgroupspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `elbInfo` | ***[[]CodedeployDeploymentGroupSpecLoadBalancerInfoElbInfo](#codedeploydeploymentgroupspecloadbalancerinfoelbinfo)***| ***(Optional)*** |
| `targetGroupInfo` | ***[[]CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupInfo](#codedeploydeploymentgroupspecloadbalancerinfotargetgroupinfo)***| ***(Optional)*** |
| `targetGroupPairInfo` | ***[[]CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo](#codedeploydeploymentgroupspecloadbalancerinfotargetgrouppairinfo)***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecLoadBalancerInfoElbInfo

Appears on:[CodedeployDeploymentGroupSpecLoadBalancerInfo](#codedeploydeploymentgroupspecloadbalancerinfo)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupInfo

Appears on:[CodedeployDeploymentGroupSpecLoadBalancerInfo](#codedeploydeploymentgroupspecloadbalancerinfo)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo

Appears on:[CodedeployDeploymentGroupSpecLoadBalancerInfo](#codedeploydeploymentgroupspecloadbalancerinfo)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `prodTrafficRoute` | ***[[]CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute](#codedeploydeploymentgroupspecloadbalancerinfotargetgrouppairinfoprodtrafficroute)***||
| `targetGroup` | ***[[]CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTargetGroup](#codedeploydeploymentgroupspecloadbalancerinfotargetgrouppairinfotargetgroup)***||
| `testTrafficRoute` | ***[[]CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute](#codedeploydeploymentgroupspecloadbalancerinfotargetgrouppairinfotesttrafficroute)***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoProdTrafficRoute

Appears on:[CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo](#codedeploydeploymentgroupspecloadbalancerinfotargetgrouppairinfo)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `listenerArns` | ***[]string***||
## CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTargetGroup

Appears on:[CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo](#codedeploydeploymentgroupspecloadbalancerinfotargetgrouppairinfo)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
## CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfoTestTrafficRoute

Appears on:[CodedeployDeploymentGroupSpecLoadBalancerInfoTargetGroupPairInfo](#codedeploydeploymentgroupspecloadbalancerinfotargetgrouppairinfo)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `listenerArns` | ***[]string***||
## CodedeployDeploymentGroupSpecOnPremisesInstanceTagFilter

Appears on:[CodedeployDeploymentGroupSpec](#codedeploydeploymentgroupspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `key` | ***string***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
| `value` | ***string***| ***(Optional)*** |
## CodedeployDeploymentGroupSpecTriggerConfiguration

Appears on:[CodedeployDeploymentGroupSpec](#codedeploydeploymentgroupspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `triggerEvents` | ***[]string***||
| `triggerName` | ***string***||
| `triggerTargetArn` | ***string***||
## CodedeployDeploymentGroupStatus

Appears on:[CodedeployDeploymentGroup](#codedeploydeploymentgroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CodedeployDeploymentGroupSpec](#codedeploydeploymentgroupspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
