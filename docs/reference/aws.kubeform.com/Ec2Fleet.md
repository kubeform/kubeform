---
title: Ec2Fleet
menu:
  docs_{{ .version }}:
    identifier: ec2fleet-aws.kubeform.com
    name: Ec2Fleet
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Ec2Fleet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Ec2Fleet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[Ec2FleetSpec](#ec2fleetspec)***||
| `status` | ***[Ec2FleetStatus](#ec2fleetstatus)***||
## Ec2FleetSpec

Appears on:[Ec2Fleet](#ec2fleet), [Ec2FleetStatus](#ec2fleetstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `excessCapacityTerminationPolicy` | ***string***| ***(Optional)*** |
| `launchTemplateConfig` | ***[[]Ec2FleetSpecLaunchTemplateConfig](#ec2fleetspeclaunchtemplateconfig)***||
| `onDemandOptions` | ***[[]Ec2FleetSpecOnDemandOptions](#ec2fleetspecondemandoptions)***| ***(Optional)*** |
| `replaceUnhealthyInstances` | ***bool***| ***(Optional)*** |
| `spotOptions` | ***[[]Ec2FleetSpecSpotOptions](#ec2fleetspecspotoptions)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `targetCapacitySpecification` | ***[[]Ec2FleetSpecTargetCapacitySpecification](#ec2fleetspectargetcapacityspecification)***||
| `terminateInstances` | ***bool***| ***(Optional)*** |
| `terminateInstancesWithExpiration` | ***bool***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
## Ec2FleetSpecLaunchTemplateConfig

Appears on:[Ec2FleetSpec](#ec2fleetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `launchTemplateSpecification` | ***[[]Ec2FleetSpecLaunchTemplateConfigLaunchTemplateSpecification](#ec2fleetspeclaunchtemplateconfiglaunchtemplatespecification)***||
| `override` | ***[[]Ec2FleetSpecLaunchTemplateConfigOverride](#ec2fleetspeclaunchtemplateconfigoverride)***| ***(Optional)*** |
## Ec2FleetSpecLaunchTemplateConfigLaunchTemplateSpecification

Appears on:[Ec2FleetSpecLaunchTemplateConfig](#ec2fleetspeclaunchtemplateconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `launchTemplateID` | ***string***| ***(Optional)*** |
| `launchTemplateName` | ***string***| ***(Optional)*** |
| `version` | ***string***||
## Ec2FleetSpecLaunchTemplateConfigOverride

Appears on:[Ec2FleetSpecLaunchTemplateConfig](#ec2fleetspeclaunchtemplateconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `availabilityZone` | ***string***| ***(Optional)*** |
| `instanceType` | ***string***| ***(Optional)*** |
| `maxPrice` | ***string***| ***(Optional)*** |
| `priority` | ***encoding/json.Number***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `weightedCapacity` | ***encoding/json.Number***| ***(Optional)*** |
## Ec2FleetSpecOnDemandOptions

Appears on:[Ec2FleetSpec](#ec2fleetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `allocationStrategy` | ***string***| ***(Optional)*** |
## Ec2FleetSpecSpotOptions

Appears on:[Ec2FleetSpec](#ec2fleetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `allocationStrategy` | ***string***| ***(Optional)*** |
| `instanceInterruptionBehavior` | ***string***| ***(Optional)*** |
| `instancePoolsToUseCount` | ***int***| ***(Optional)*** |
## Ec2FleetSpecTargetCapacitySpecification

Appears on:[Ec2FleetSpec](#ec2fleetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `defaultTargetCapacityType` | ***string***||
| `onDemandTargetCapacity` | ***int***| ***(Optional)*** |
| `spotTargetCapacity` | ***int***| ***(Optional)*** |
| `totalTargetCapacity` | ***int***||
## Ec2FleetStatus

Appears on:[Ec2Fleet](#ec2fleet)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[Ec2FleetSpec](#ec2fleetspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[Ec2FleetStatus](#ec2fleetstatus)

---
