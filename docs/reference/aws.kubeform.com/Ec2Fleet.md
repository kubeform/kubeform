## Ec2Fleet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Ec2Fleet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[Ec2FleetSpec](#Ec2FleetSpec)***||
| `status` | ***[Ec2FleetStatus](#Ec2FleetStatus)***||
## Ec2FleetSpec
##### (Appears on:[Ec2Fleet](#Ec2Fleet), [Ec2FleetStatus](#Ec2FleetStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `excessCapacityTerminationPolicy` | ***string***| ***(Optional)*** |
| `launchTemplateConfig` | ***[[]Ec2FleetSpecLaunchTemplateConfig](#Ec2FleetSpecLaunchTemplateConfig)***||
| `onDemandOptions` | ***[[]Ec2FleetSpecOnDemandOptions](#Ec2FleetSpecOnDemandOptions)***| ***(Optional)*** |
| `replaceUnhealthyInstances` | ***bool***| ***(Optional)*** |
| `spotOptions` | ***[[]Ec2FleetSpecSpotOptions](#Ec2FleetSpecSpotOptions)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `targetCapacitySpecification` | ***[[]Ec2FleetSpecTargetCapacitySpecification](#Ec2FleetSpecTargetCapacitySpecification)***||
| `terminateInstances` | ***bool***| ***(Optional)*** |
| `terminateInstancesWithExpiration` | ***bool***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
## Ec2FleetSpecLaunchTemplateConfig
##### (Appears on:[Ec2FleetSpec](#Ec2FleetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `launchTemplateSpecification` | ***[[]Ec2FleetSpecLaunchTemplateConfigLaunchTemplateSpecification](#Ec2FleetSpecLaunchTemplateConfigLaunchTemplateSpecification)***||
| `override` | ***[[]Ec2FleetSpecLaunchTemplateConfigOverride](#Ec2FleetSpecLaunchTemplateConfigOverride)***| ***(Optional)*** |
## Ec2FleetSpecLaunchTemplateConfigLaunchTemplateSpecification
##### (Appears on:[Ec2FleetSpecLaunchTemplateConfig](#Ec2FleetSpecLaunchTemplateConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `launchTemplateID` | ***string***| ***(Optional)*** |
| `launchTemplateName` | ***string***| ***(Optional)*** |
| `version` | ***string***||
## Ec2FleetSpecLaunchTemplateConfigOverride
##### (Appears on:[Ec2FleetSpecLaunchTemplateConfig](#Ec2FleetSpecLaunchTemplateConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `availabilityZone` | ***string***| ***(Optional)*** |
| `instanceType` | ***string***| ***(Optional)*** |
| `maxPrice` | ***string***| ***(Optional)*** |
| `priority` | ***encoding/json.Number***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `weightedCapacity` | ***encoding/json.Number***| ***(Optional)*** |
## Ec2FleetSpecOnDemandOptions
##### (Appears on:[Ec2FleetSpec](#Ec2FleetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `allocationStrategy` | ***string***| ***(Optional)*** |
## Ec2FleetSpecSpotOptions
##### (Appears on:[Ec2FleetSpec](#Ec2FleetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `allocationStrategy` | ***string***| ***(Optional)*** |
| `instanceInterruptionBehavior` | ***string***| ***(Optional)*** |
| `instancePoolsToUseCount` | ***int***| ***(Optional)*** |
## Ec2FleetSpecTargetCapacitySpecification
##### (Appears on:[Ec2FleetSpec](#Ec2FleetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `defaultTargetCapacityType` | ***string***||
| `onDemandTargetCapacity` | ***int***| ***(Optional)*** |
| `spotTargetCapacity` | ***int***| ***(Optional)*** |
| `totalTargetCapacity` | ***int***||
## Ec2FleetStatus
##### (Appears on:[Ec2Fleet](#Ec2Fleet))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[Ec2FleetSpec](#Ec2FleetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
