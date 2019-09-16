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
| `spec` | ***[ComputeRegionInstanceGroupManagerSpec](#computeregioninstancegroupmanagerspec)***||
| `status` | ***[ComputeRegionInstanceGroupManagerStatus](#computeregioninstancegroupmanagerstatus)***||
## ComputeRegionInstanceGroupManagerSpec

Appears on:[ComputeRegionInstanceGroupManager](#computeregioninstancegroupmanager), [ComputeRegionInstanceGroupManagerStatus](#computeregioninstancegroupmanagerstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `autoHealingPolicies` | ***[[]ComputeRegionInstanceGroupManagerSpecAutoHealingPolicies](#computeregioninstancegroupmanagerspecautohealingpolicies)***| ***(Optional)*** Deprecated|
| `baseInstanceName` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `distributionPolicyZones` | ***[]string***| ***(Optional)*** |
| `fingerprint` | ***string***| ***(Optional)*** |
| `instanceGroup` | ***string***| ***(Optional)*** |
| `instanceTemplate` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `namedPort` | ***[[]ComputeRegionInstanceGroupManagerSpecNamedPort](#computeregioninstancegroupmanagerspecnamedport)***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***||
| `rollingUpdatePolicy` | ***[[]ComputeRegionInstanceGroupManagerSpecRollingUpdatePolicy](#computeregioninstancegroupmanagerspecrollingupdatepolicy)***| ***(Optional)*** Deprecated|
| `selfLink` | ***string***| ***(Optional)*** |
| `targetPools` | ***[]string***| ***(Optional)*** |
| `targetSize` | ***int***| ***(Optional)*** |
| `updateStrategy` | ***string***| ***(Optional)*** Deprecated|
| `version` | ***[[]ComputeRegionInstanceGroupManagerSpecVersion](#computeregioninstancegroupmanagerspecversion)***| ***(Optional)*** Deprecated|
| `waitForInstances` | ***bool***| ***(Optional)*** |
## ComputeRegionInstanceGroupManagerSpecAutoHealingPolicies

Appears on:[ComputeRegionInstanceGroupManagerSpec](#computeregioninstancegroupmanagerspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `healthCheck` | ***string***||
| `initialDelaySec` | ***int***||
## ComputeRegionInstanceGroupManagerSpecNamedPort

Appears on:[ComputeRegionInstanceGroupManagerSpec](#computeregioninstancegroupmanagerspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `port` | ***int***||
## ComputeRegionInstanceGroupManagerSpecRollingUpdatePolicy

Appears on:[ComputeRegionInstanceGroupManagerSpec](#computeregioninstancegroupmanagerspec)

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

Appears on:[ComputeRegionInstanceGroupManagerSpec](#computeregioninstancegroupmanagerspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `instanceTemplate` | ***string***||
| `name` | ***string***||
| `targetSize` | ***[[]ComputeRegionInstanceGroupManagerSpecVersionTargetSize](#computeregioninstancegroupmanagerspecversiontargetsize)***| ***(Optional)*** |
## ComputeRegionInstanceGroupManagerSpecVersionTargetSize

Appears on:[ComputeRegionInstanceGroupManagerSpecVersion](#computeregioninstancegroupmanagerspecversion)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `fixed` | ***int***| ***(Optional)*** |
| `percent` | ***int***| ***(Optional)*** |
## ComputeRegionInstanceGroupManagerStatus

Appears on:[ComputeRegionInstanceGroupManager](#computeregioninstancegroupmanager)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeRegionInstanceGroupManagerSpec](#computeregioninstancegroupmanagerspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
