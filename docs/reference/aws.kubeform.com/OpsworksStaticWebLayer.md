---
title: OpsworksStaticWebLayer
menu:
  docs_{{ .version }}:
    identifier: opsworksstaticweblayer-aws.kubeform.com
    name: OpsworksStaticWebLayer
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## OpsworksStaticWebLayer
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `OpsworksStaticWebLayer` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[OpsworksStaticWebLayerSpec](#OpsworksStaticWebLayerSpec)***||
| `status` | ***[OpsworksStaticWebLayerStatus](#OpsworksStaticWebLayerStatus)***||
## OpsworksStaticWebLayerSpec
##### (Appears on:[OpsworksStaticWebLayer](#OpsworksStaticWebLayer), [OpsworksStaticWebLayerStatus](#OpsworksStaticWebLayerStatus))
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
| `ebsVolume` | ***[[]OpsworksStaticWebLayerSpecEbsVolume](#OpsworksStaticWebLayerSpecEbsVolume)***| ***(Optional)*** |
| `elasticLoadBalancer` | ***string***| ***(Optional)*** |
| `installUpdatesOnBoot` | ***bool***| ***(Optional)*** |
| `instanceShutdownTimeout` | ***int***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `stackID` | ***string***||
| `systemPackages` | ***[]string***| ***(Optional)*** |
| `useEbsOptimizedInstances` | ***bool***| ***(Optional)*** |
## OpsworksStaticWebLayerSpecEbsVolume
##### (Appears on:[OpsworksStaticWebLayerSpec](#OpsworksStaticWebLayerSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `iops` | ***int***| ***(Optional)*** |
| `mountPoint` | ***string***||
| `numberOfDisks` | ***int***||
| `raidLevel` | ***string***| ***(Optional)*** |
| `size` | ***int***||
| `type` | ***string***| ***(Optional)*** |
## OpsworksStaticWebLayerStatus
##### (Appears on:[OpsworksStaticWebLayer](#OpsworksStaticWebLayer))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[OpsworksStaticWebLayerSpec](#OpsworksStaticWebLayerSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
