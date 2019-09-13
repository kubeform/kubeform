---
title: OpsworksPhpAppLayer
menu:
  docs_{{ .version }}:
    identifier: opsworksphpapplayer-aws.kubeform.com
    name: OpsworksPhpAppLayer
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## OpsworksPhpAppLayer
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `OpsworksPhpAppLayer` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[OpsworksPhpAppLayerSpec](#OpsworksPhpAppLayerSpec)***||
| `status` | ***[OpsworksPhpAppLayerStatus](#OpsworksPhpAppLayerStatus)***||
## OpsworksPhpAppLayerSpec
##### (Appears on:[OpsworksPhpAppLayer](#OpsworksPhpAppLayer), [OpsworksPhpAppLayerStatus](#OpsworksPhpAppLayerStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
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
| `ebsVolume` | ***[[]OpsworksPhpAppLayerSpecEbsVolume](#OpsworksPhpAppLayerSpecEbsVolume)***| ***(Optional)*** |
| `elasticLoadBalancer` | ***string***| ***(Optional)*** |
| `installUpdatesOnBoot` | ***bool***| ***(Optional)*** |
| `instanceShutdownTimeout` | ***int***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `stackID` | ***string***||
| `systemPackages` | ***[]string***| ***(Optional)*** |
| `useEbsOptimizedInstances` | ***bool***| ***(Optional)*** |
## OpsworksPhpAppLayerSpecEbsVolume
##### (Appears on:[OpsworksPhpAppLayerSpec](#OpsworksPhpAppLayerSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `iops` | ***int***| ***(Optional)*** |
| `mountPoint` | ***string***||
| `numberOfDisks` | ***int***||
| `raidLevel` | ***string***| ***(Optional)*** |
| `size` | ***int***||
| `type` | ***string***| ***(Optional)*** |
## OpsworksPhpAppLayerStatus
##### (Appears on:[OpsworksPhpAppLayer](#OpsworksPhpAppLayer))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[OpsworksPhpAppLayerSpec](#OpsworksPhpAppLayerSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
