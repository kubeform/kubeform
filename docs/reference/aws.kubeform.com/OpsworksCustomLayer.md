---
title: OpsworksCustomLayer
menu:
  docs_{{ .version }}:
    identifier: opsworkscustomlayer-aws.kubeform.com
    name: OpsworksCustomLayer
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## OpsworksCustomLayer
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `OpsworksCustomLayer` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[OpsworksCustomLayerSpec](#opsworkscustomlayerspec)***||
| `status` | ***[OpsworksCustomLayerStatus](#opsworkscustomlayerstatus)***||
## OpsworksCustomLayerSpec

Appears on:[OpsworksCustomLayer](#opsworkscustomlayer), [OpsworksCustomLayerStatus](#opsworkscustomlayerstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
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
| `ebsVolume` | ***[[]OpsworksCustomLayerSpecEbsVolume](#opsworkscustomlayerspecebsvolume)***| ***(Optional)*** |
| `elasticLoadBalancer` | ***string***| ***(Optional)*** |
| `installUpdatesOnBoot` | ***bool***| ***(Optional)*** |
| `instanceShutdownTimeout` | ***int64***| ***(Optional)*** |
| `name` | ***string***||
| `shortName` | ***string***||
| `stackID` | ***string***||
| `systemPackages` | ***[]string***| ***(Optional)*** |
| `useEbsOptimizedInstances` | ***bool***| ***(Optional)*** |
## OpsworksCustomLayerSpecEbsVolume

Appears on:[OpsworksCustomLayerSpec](#opsworkscustomlayerspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `iops` | ***int64***| ***(Optional)*** |
| `mountPoint` | ***string***||
| `numberOfDisks` | ***int64***||
| `raidLevel` | ***string***| ***(Optional)*** |
| `size` | ***int64***||
| `type` | ***string***| ***(Optional)*** |
## OpsworksCustomLayerStatus

Appears on:[OpsworksCustomLayer](#opsworkscustomlayer)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[OpsworksCustomLayerSpec](#opsworkscustomlayerspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[OpsworksCustomLayerStatus](#opsworkscustomlayerstatus)

---
