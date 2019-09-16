---
title: ComputeRegionInstanceGroupManager
menu:
  docs_{{ .version }}:
    identifier: computeregioninstancegroupmanager-google.kubeform.com
    name: ComputeRegionInstanceGroupManager
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeRegionInstanceGroupManager
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeRegionInstanceGroupManager` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeRegionInstanceGroupManagerSpec](#ComputeRegionInstanceGroupManagerSpec)***||
| `status` | ***[ComputeRegionInstanceGroupManagerStatus](#ComputeRegionInstanceGroupManagerStatus)***||
## ComputeRegionInstanceGroupManagerSpec

Appears on:[ComputeRegionInstanceGroupManager](#ComputeRegionInstanceGroupManager), [ComputeRegionInstanceGroupManagerStatus](#ComputeRegionInstanceGroupManagerStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `autoHealingPolicies` | ***[[]ComputeRegionInstanceGroupManagerSpecAutoHealingPolicies](#ComputeRegionInstanceGroupManagerSpecAutoHealingPolicies)***| ***(Optional)*** Deprecated|
| `baseInstanceName` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `distributionPolicyZones` | ***[]string***| ***(Optional)*** |
| `fingerprint` | ***string***| ***(Optional)*** |
| `instanceGroup` | ***string***| ***(Optional)*** |
| `instanceTemplate` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `namedPort` | ***[[]ComputeRegionInstanceGroupManagerSpecNamedPort](#ComputeRegionInstanceGroupManagerSpecNamedPort)***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***||
| `rollingUpdatePolicy` | ***[[]ComputeRegionInstanceGroupManagerSpecRollingUpdatePolicy](#ComputeRegionInstanceGroupManagerSpecRollingUpdatePolicy)***| ***(Optional)*** Deprecated|
| `selfLink` | ***string***| ***(Optional)*** |
| `targetPools` | ***[]string***| ***(Optional)*** |
| `targetSize` | ***int***| ***(Optional)*** |
| `updateStrategy` | ***string***| ***(Optional)*** Deprecated|
| `version` | ***[[]ComputeRegionInstanceGroupManagerSpecVersion](#ComputeRegionInstanceGroupManagerSpecVersion)***| ***(Optional)*** Deprecated|
| `waitForInstances` | ***bool***| ***(Optional)*** |
## ComputeRegionInstanceGroupManagerSpecAutoHealingPolicies

Appears on:[ComputeRegionInstanceGroupManagerSpec](#ComputeRegionInstanceGroupManagerSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `healthCheck` | ***string***||
| `initialDelaySec` | ***int***||
## ComputeRegionInstanceGroupManagerSpecNamedPort

Appears on:[ComputeRegionInstanceGroupManagerSpec](#ComputeRegionInstanceGroupManagerSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `port` | ***int***||
## ComputeRegionInstanceGroupManagerSpecRollingUpdatePolicy

Appears on:[ComputeRegionInstanceGroupManagerSpec](#ComputeRegionInstanceGroupManagerSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `maxSurgeFixed` | ***int***| ***(Optional)*** |
| `maxSurgePercent` | ***int***| ***(Optional)*** |
| `maxUnavailableFixed` | ***int***| ***(Optional)*** |
| `maxUnavailablePercent` | ***int***| ***(Optional)*** |
| `minReadySec` | ***int***| ***(Optional)*** |
| `minimalAction` | ***string***||
| `type` | ***string***||
## ComputeRegionInstanceGroupManagerSpecVersion

Appears on:[ComputeRegionInstanceGroupManagerSpec](#ComputeRegionInstanceGroupManagerSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `instanceTemplate` | ***string***||
| `name` | ***string***||
| `targetSize` | ***[[]ComputeRegionInstanceGroupManagerSpecVersionTargetSize](#ComputeRegionInstanceGroupManagerSpecVersionTargetSize)***| ***(Optional)*** |
## ComputeRegionInstanceGroupManagerSpecVersionTargetSize

Appears on:[ComputeRegionInstanceGroupManagerSpecVersion](#ComputeRegionInstanceGroupManagerSpecVersion)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `fixed` | ***int***| ***(Optional)*** |
| `percent` | ***int***| ***(Optional)*** |
## ComputeRegionInstanceGroupManagerStatus

Appears on:[ComputeRegionInstanceGroupManager](#ComputeRegionInstanceGroupManager)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeRegionInstanceGroupManagerSpec](#ComputeRegionInstanceGroupManagerSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
