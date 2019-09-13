---
title: OpsworksHaproxyLayer
menu:
  docs_{{ .version }}:
    identifier: opsworkshaproxylayer-aws.kubeform.com
    name: OpsworksHaproxyLayer
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## OpsworksHaproxyLayer
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `OpsworksHaproxyLayer` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[OpsworksHaproxyLayerSpec](#OpsworksHaproxyLayerSpec)***||
| `status` | ***[OpsworksHaproxyLayerStatus](#OpsworksHaproxyLayerStatus)***||
## OpsworksHaproxyLayerSpec
##### (Appears on:[OpsworksHaproxyLayer](#OpsworksHaproxyLayer), [OpsworksHaproxyLayerStatus](#OpsworksHaproxyLayerStatus))
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
| `ebsVolume` | ***[[]OpsworksHaproxyLayerSpecEbsVolume](#OpsworksHaproxyLayerSpecEbsVolume)***| ***(Optional)*** |
| `elasticLoadBalancer` | ***string***| ***(Optional)*** |
| `healthcheckMethod` | ***string***| ***(Optional)*** |
| `healthcheckURL` | ***string***| ***(Optional)*** |
| `installUpdatesOnBoot` | ***bool***| ***(Optional)*** |
| `instanceShutdownTimeout` | ***int***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `stackID` | ***string***||
| `statsEnabled` | ***bool***| ***(Optional)*** |
| `statsPassword` | ***string***||
| `statsURL` | ***string***| ***(Optional)*** |
| `statsUser` | ***string***| ***(Optional)*** |
| `systemPackages` | ***[]string***| ***(Optional)*** |
| `useEbsOptimizedInstances` | ***bool***| ***(Optional)*** |
## OpsworksHaproxyLayerSpecEbsVolume
##### (Appears on:[OpsworksHaproxyLayerSpec](#OpsworksHaproxyLayerSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `iops` | ***int***| ***(Optional)*** |
| `mountPoint` | ***string***||
| `numberOfDisks` | ***int***||
| `raidLevel` | ***string***| ***(Optional)*** |
| `size` | ***int***||
| `type` | ***string***| ***(Optional)*** |
## OpsworksHaproxyLayerStatus
##### (Appears on:[OpsworksHaproxyLayer](#OpsworksHaproxyLayer))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[OpsworksHaproxyLayerSpec](#OpsworksHaproxyLayerSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
