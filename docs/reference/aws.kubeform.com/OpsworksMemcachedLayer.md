---
title: OpsworksMemcachedLayer
menu:
  docs_v0.0.1:
    identifier: opsworksmemcachedlayer-aws.kubeform.com
    name: OpsworksMemcachedLayer
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## OpsworksMemcachedLayer
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `OpsworksMemcachedLayer` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[OpsworksMemcachedLayerSpec](#OpsworksMemcachedLayerSpec)***||
| `status` | ***[OpsworksMemcachedLayerStatus](#OpsworksMemcachedLayerStatus)***||
## OpsworksMemcachedLayerSpec
##### (Appears on:[OpsworksMemcachedLayer](#OpsworksMemcachedLayer), [OpsworksMemcachedLayerStatus](#OpsworksMemcachedLayerStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allocatedMemory` | ***int***| ***(Optional)*** |
| `autoAssignElasticIPS` | ***bool***| ***(Optional)*** |
| `autoAssignPublicIPS` | ***bool***| ***(Optional)*** |
| `autoHealing` | ***bool***| ***(Optional)*** |
| `customConfigureRecipes` | ***[]string***| ***(Optional)*** |
| `customDeployRecipes` | ***[]string***| ***(Optional)*** |
| `customInstanceProfileArn` | ***string***| ***(Optional)*** |
| `customJSON` | ***string***| ***(Optional)*** |
| `customSecurityGroupIDS` | ***[]string***| ***(Optional)*** |
| `customSetupRecipes` | ***[]string***| ***(Optional)*** |
| `customShutdownRecipes` | ***[]string***| ***(Optional)*** |
| `customUndeployRecipes` | ***[]string***| ***(Optional)*** |
| `drainElbOnShutdown` | ***bool***| ***(Optional)*** |
| `ebsVolume` | ***[[]OpsworksMemcachedLayerSpecEbsVolume](#OpsworksMemcachedLayerSpecEbsVolume)***| ***(Optional)*** |
| `elasticLoadBalancer` | ***string***| ***(Optional)*** |
| `installUpdatesOnBoot` | ***bool***| ***(Optional)*** |
| `instanceShutdownTimeout` | ***int***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `stackID` | ***string***||
| `systemPackages` | ***[]string***| ***(Optional)*** |
| `useEbsOptimizedInstances` | ***bool***| ***(Optional)*** |
## OpsworksMemcachedLayerSpecEbsVolume
##### (Appears on:[OpsworksMemcachedLayerSpec](#OpsworksMemcachedLayerSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `iops` | ***int***| ***(Optional)*** |
| `mountPoint` | ***string***||
| `numberOfDisks` | ***int***||
| `raidLevel` | ***string***| ***(Optional)*** |
| `size` | ***int***||
| `type` | ***string***| ***(Optional)*** |
## OpsworksMemcachedLayerStatus
##### (Appears on:[OpsworksMemcachedLayer](#OpsworksMemcachedLayer))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[OpsworksMemcachedLayerSpec](#OpsworksMemcachedLayerSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
