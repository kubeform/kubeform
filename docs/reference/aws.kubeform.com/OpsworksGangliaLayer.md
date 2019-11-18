---
title: OpsworksGangliaLayer
menu:
  docs_{{ .version }}:
    identifier: opsworksganglialayer-aws.kubeform.com
    name: OpsworksGangliaLayer
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## OpsworksGangliaLayer
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `OpsworksGangliaLayer` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[OpsworksGangliaLayerSpec](#opsworksganglialayerspec)***||
| `status` | ***[OpsworksGangliaLayerStatus](#opsworksganglialayerstatus)***||
## OpsworksGangliaLayerSpec

Appears on:[OpsworksGangliaLayer](#opsworksganglialayer), [OpsworksGangliaLayerStatus](#opsworksganglialayerstatus)

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
| `ebsVolume` | ***[[]OpsworksGangliaLayerSpecEbsVolume](#opsworksganglialayerspecebsvolume)***| ***(Optional)*** |
| `elasticLoadBalancer` | ***string***| ***(Optional)*** |
| `installUpdatesOnBoot` | ***bool***| ***(Optional)*** |
| `instanceShutdownTimeout` | ***int64***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `password` | ***string***||
| `stackID` | ***string***||
| `systemPackages` | ***[]string***| ***(Optional)*** |
| `url` | ***string***| ***(Optional)*** |
| `useEbsOptimizedInstances` | ***bool***| ***(Optional)*** |
| `username` | ***string***| ***(Optional)*** |
## OpsworksGangliaLayerSpecEbsVolume

Appears on:[OpsworksGangliaLayerSpec](#opsworksganglialayerspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `iops` | ***int64***| ***(Optional)*** |
| `mountPoint` | ***string***||
| `numberOfDisks` | ***int64***||
| `raidLevel` | ***string***| ***(Optional)*** |
| `size` | ***int64***||
| `type` | ***string***| ***(Optional)*** |
## OpsworksGangliaLayerStatus

Appears on:[OpsworksGangliaLayer](#opsworksganglialayer)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[OpsworksGangliaLayerSpec](#opsworksganglialayerspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[OpsworksGangliaLayerStatus](#opsworksganglialayerstatus)

---
