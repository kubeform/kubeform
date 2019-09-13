---
title: ComputeInstanceGroupManager
menu:
  docs_{{ .version }}:
    identifier: computeinstancegroupmanager-google.kubeform.com
    name: ComputeInstanceGroupManager
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeInstanceGroupManager
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeInstanceGroupManager` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeInstanceGroupManagerSpec](#ComputeInstanceGroupManagerSpec)***||
| `status` | ***[ComputeInstanceGroupManagerStatus](#ComputeInstanceGroupManagerStatus)***||
## ComputeInstanceGroupManagerSpec
##### (Appears on:[ComputeInstanceGroupManager](#ComputeInstanceGroupManager), [ComputeInstanceGroupManagerStatus](#ComputeInstanceGroupManagerStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `autoHealingPolicies` | ***[[]ComputeInstanceGroupManagerSpecAutoHealingPolicies](#ComputeInstanceGroupManagerSpecAutoHealingPolicies)***| ***(Optional)*** Deprecated|
| `baseInstanceName` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `fingerprint` | ***string***| ***(Optional)*** |
| `instanceGroup` | ***string***| ***(Optional)*** |
| `instanceTemplate` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `namedPort` | ***[[]ComputeInstanceGroupManagerSpecNamedPort](#ComputeInstanceGroupManagerSpecNamedPort)***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `rollingUpdatePolicy` | ***[[]ComputeInstanceGroupManagerSpecRollingUpdatePolicy](#ComputeInstanceGroupManagerSpecRollingUpdatePolicy)***| ***(Optional)*** Deprecated|
| `selfLink` | ***string***| ***(Optional)*** |
| `targetPools` | ***[]string***| ***(Optional)*** |
| `targetSize` | ***int***| ***(Optional)*** |
| `updateStrategy` | ***string***| ***(Optional)*** |
| `version` | ***[[]ComputeInstanceGroupManagerSpecVersion](#ComputeInstanceGroupManagerSpecVersion)***| ***(Optional)*** Deprecated|
| `waitForInstances` | ***bool***| ***(Optional)*** |
| `zone` | ***string***| ***(Optional)*** |
## ComputeInstanceGroupManagerSpecAutoHealingPolicies
##### (Appears on:[ComputeInstanceGroupManagerSpec](#ComputeInstanceGroupManagerSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `healthCheck` | ***string***||
| `initialDelaySec` | ***int***||
## ComputeInstanceGroupManagerSpecNamedPort
##### (Appears on:[ComputeInstanceGroupManagerSpec](#ComputeInstanceGroupManagerSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `port` | ***int***||
## ComputeInstanceGroupManagerSpecRollingUpdatePolicy
##### (Appears on:[ComputeInstanceGroupManagerSpec](#ComputeInstanceGroupManagerSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `maxSurgeFixed` | ***int***| ***(Optional)*** |
| `maxSurgePercent` | ***int***| ***(Optional)*** |
| `maxUnavailableFixed` | ***int***| ***(Optional)*** |
| `maxUnavailablePercent` | ***int***| ***(Optional)*** |
| `minReadySec` | ***int***| ***(Optional)*** |
| `minimalAction` | ***string***||
| `type` | ***string***||
## ComputeInstanceGroupManagerSpecVersion
##### (Appears on:[ComputeInstanceGroupManagerSpec](#ComputeInstanceGroupManagerSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `instanceTemplate` | ***string***||
| `name` | ***string***||
| `targetSize` | ***[[]ComputeInstanceGroupManagerSpecVersionTargetSize](#ComputeInstanceGroupManagerSpecVersionTargetSize)***| ***(Optional)*** |
## ComputeInstanceGroupManagerSpecVersionTargetSize
##### (Appears on:[ComputeInstanceGroupManagerSpecVersion](#ComputeInstanceGroupManagerSpecVersion))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `fixed` | ***int***| ***(Optional)*** |
| `percent` | ***int***| ***(Optional)*** |
## ComputeInstanceGroupManagerStatus
##### (Appears on:[ComputeInstanceGroupManager](#ComputeInstanceGroupManager))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeInstanceGroupManagerSpec](#ComputeInstanceGroupManagerSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
