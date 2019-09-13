---
title: CloudwatchEventTarget
menu:
  docs_{{ .version }}:
    identifier: cloudwatcheventtarget-aws.kubeform.com
    name: CloudwatchEventTarget
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CloudwatchEventTarget
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CloudwatchEventTarget` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CloudwatchEventTargetSpec](#CloudwatchEventTargetSpec)***||
| `status` | ***[CloudwatchEventTargetStatus](#CloudwatchEventTargetStatus)***||
## CloudwatchEventTargetSpec
##### (Appears on:[CloudwatchEventTarget](#CloudwatchEventTarget), [CloudwatchEventTargetStatus](#CloudwatchEventTargetStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***||
| `batchTarget` | ***[[]CloudwatchEventTargetSpecBatchTarget](#CloudwatchEventTargetSpecBatchTarget)***| ***(Optional)*** |
| `ecsTarget` | ***[[]CloudwatchEventTargetSpecEcsTarget](#CloudwatchEventTargetSpecEcsTarget)***| ***(Optional)*** |
| `input` | ***string***| ***(Optional)*** |
| `inputPath` | ***string***| ***(Optional)*** |
| `inputTransformer` | ***[[]CloudwatchEventTargetSpecInputTransformer](#CloudwatchEventTargetSpecInputTransformer)***| ***(Optional)*** |
| `kinesisTarget` | ***[[]CloudwatchEventTargetSpecKinesisTarget](#CloudwatchEventTargetSpecKinesisTarget)***| ***(Optional)*** |
| `roleArn` | ***string***| ***(Optional)*** |
| `rule` | ***string***||
| `runCommandTargets` | ***[[]CloudwatchEventTargetSpecRunCommandTargets](#CloudwatchEventTargetSpecRunCommandTargets)***| ***(Optional)*** |
| `sqsTarget` | ***[[]CloudwatchEventTargetSpecSqsTarget](#CloudwatchEventTargetSpecSqsTarget)***| ***(Optional)*** |
| `targetID` | ***string***| ***(Optional)*** |
## CloudwatchEventTargetSpecBatchTarget
##### (Appears on:[CloudwatchEventTargetSpec](#CloudwatchEventTargetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `arraySize` | ***int***| ***(Optional)*** |
| `jobAttempts` | ***int***| ***(Optional)*** |
| `jobDefinition` | ***string***||
| `jobName` | ***string***||
## CloudwatchEventTargetSpecEcsTarget
##### (Appears on:[CloudwatchEventTargetSpec](#CloudwatchEventTargetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `group` | ***string***| ***(Optional)*** |
| `launchType` | ***string***| ***(Optional)*** |
| `networkConfiguration` | ***[[]CloudwatchEventTargetSpecEcsTargetNetworkConfiguration](#CloudwatchEventTargetSpecEcsTargetNetworkConfiguration)***| ***(Optional)*** |
| `platformVersion` | ***string***| ***(Optional)*** |
| `taskCount` | ***int***| ***(Optional)*** |
| `taskDefinitionArn` | ***string***||
## CloudwatchEventTargetSpecEcsTargetNetworkConfiguration
##### (Appears on:[CloudwatchEventTargetSpecEcsTarget](#CloudwatchEventTargetSpecEcsTarget))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `assignPublicIP` | ***bool***| ***(Optional)*** |
| `securityGroups` | ***[]string***| ***(Optional)*** |
| `subnets` | ***[]string***||
## CloudwatchEventTargetSpecInputTransformer
##### (Appears on:[CloudwatchEventTargetSpec](#CloudwatchEventTargetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `inputPaths` | ***map[string]string***| ***(Optional)*** |
| `inputTemplate` | ***string***||
## CloudwatchEventTargetSpecKinesisTarget
##### (Appears on:[CloudwatchEventTargetSpec](#CloudwatchEventTargetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `partitionKeyPath` | ***string***| ***(Optional)*** |
## CloudwatchEventTargetSpecRunCommandTargets
##### (Appears on:[CloudwatchEventTargetSpec](#CloudwatchEventTargetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `key` | ***string***||
| `values` | ***[]string***||
## CloudwatchEventTargetSpecSqsTarget
##### (Appears on:[CloudwatchEventTargetSpec](#CloudwatchEventTargetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `messageGroupID` | ***string***| ***(Optional)*** |
## CloudwatchEventTargetStatus
##### (Appears on:[CloudwatchEventTarget](#CloudwatchEventTarget))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CloudwatchEventTargetSpec](#CloudwatchEventTargetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
