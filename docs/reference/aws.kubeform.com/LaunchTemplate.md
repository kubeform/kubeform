## LaunchTemplate
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LaunchTemplate` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LaunchTemplateSpec](#LaunchTemplateSpec)***||
| `status` | ***[LaunchTemplateStatus](#LaunchTemplateStatus)***||
## LaunchTemplateSpec
##### (Appears on:[LaunchTemplate](#LaunchTemplate), [LaunchTemplateStatus](#LaunchTemplateStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `blockDeviceMappings` | ***[[]LaunchTemplateSpecBlockDeviceMappings](#LaunchTemplateSpecBlockDeviceMappings)***| ***(Optional)*** |
| `capacityReservationSpecification` | ***[[]LaunchTemplateSpecCapacityReservationSpecification](#LaunchTemplateSpecCapacityReservationSpecification)***| ***(Optional)*** |
| `creditSpecification` | ***[[]LaunchTemplateSpecCreditSpecification](#LaunchTemplateSpecCreditSpecification)***| ***(Optional)*** |
| `defaultVersion` | ***int***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `disableAPITermination` | ***bool***| ***(Optional)*** |
| `ebsOptimized` | ***string***| ***(Optional)*** |
| `elasticGpuSpecifications` | ***[[]LaunchTemplateSpecElasticGpuSpecifications](#LaunchTemplateSpecElasticGpuSpecifications)***| ***(Optional)*** |
| `elasticInferenceAccelerator` | ***[[]LaunchTemplateSpecElasticInferenceAccelerator](#LaunchTemplateSpecElasticInferenceAccelerator)***| ***(Optional)*** |
| `iamInstanceProfile` | ***[[]LaunchTemplateSpecIamInstanceProfile](#LaunchTemplateSpecIamInstanceProfile)***| ***(Optional)*** |
| `imageID` | ***string***| ***(Optional)*** |
| `instanceInitiatedShutdownBehavior` | ***string***| ***(Optional)*** |
| `instanceMarketOptions` | ***[[]LaunchTemplateSpecInstanceMarketOptions](#LaunchTemplateSpecInstanceMarketOptions)***| ***(Optional)*** |
| `instanceType` | ***string***| ***(Optional)*** |
| `kernelID` | ***string***| ***(Optional)*** |
| `keyName` | ***string***| ***(Optional)*** |
| `latestVersion` | ***int***| ***(Optional)*** |
| `licenseSpecification` | ***[[]LaunchTemplateSpecLicenseSpecification](#LaunchTemplateSpecLicenseSpecification)***| ***(Optional)*** |
| `monitoring` | ***[[]LaunchTemplateSpecMonitoring](#LaunchTemplateSpecMonitoring)***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `networkInterfaces` | ***[[]LaunchTemplateSpecNetworkInterfaces](#LaunchTemplateSpecNetworkInterfaces)***| ***(Optional)*** |
| `placement` | ***[[]LaunchTemplateSpecPlacement](#LaunchTemplateSpecPlacement)***| ***(Optional)*** |
| `ramDiskID` | ***string***| ***(Optional)*** |
| `securityGroupNames` | ***[]string***| ***(Optional)*** |
| `tagSpecifications` | ***[[]LaunchTemplateSpecTagSpecifications](#LaunchTemplateSpecTagSpecifications)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `userData` | ***string***| ***(Optional)*** |
| `vpcSecurityGroupIDS` | ***[]string***| ***(Optional)*** |
## LaunchTemplateSpecBlockDeviceMappings
##### (Appears on:[LaunchTemplateSpec](#LaunchTemplateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `deviceName` | ***string***| ***(Optional)*** |
| `ebs` | ***[[]LaunchTemplateSpecBlockDeviceMappingsEbs](#LaunchTemplateSpecBlockDeviceMappingsEbs)***| ***(Optional)*** |
| `noDevice` | ***string***| ***(Optional)*** |
| `virtualName` | ***string***| ***(Optional)*** |
## LaunchTemplateSpecBlockDeviceMappingsEbs
##### (Appears on:[LaunchTemplateSpecBlockDeviceMappings](#LaunchTemplateSpecBlockDeviceMappings))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `deleteOnTermination` | ***string***| ***(Optional)*** |
| `encrypted` | ***string***| ***(Optional)*** |
| `iops` | ***int***| ***(Optional)*** |
| `kmsKeyID` | ***string***| ***(Optional)*** |
| `snapshotID` | ***string***| ***(Optional)*** |
| `volumeSize` | ***int***| ***(Optional)*** |
| `volumeType` | ***string***| ***(Optional)*** |
## LaunchTemplateSpecCapacityReservationSpecification
##### (Appears on:[LaunchTemplateSpec](#LaunchTemplateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `capacityReservationPreference` | ***string***| ***(Optional)*** |
| `capacityReservationTarget` | ***[[]LaunchTemplateSpecCapacityReservationSpecificationCapacityReservationTarget](#LaunchTemplateSpecCapacityReservationSpecificationCapacityReservationTarget)***| ***(Optional)*** |
## LaunchTemplateSpecCapacityReservationSpecificationCapacityReservationTarget
##### (Appears on:[LaunchTemplateSpecCapacityReservationSpecification](#LaunchTemplateSpecCapacityReservationSpecification))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `capacityReservationID` | ***string***| ***(Optional)*** |
## LaunchTemplateSpecCreditSpecification
##### (Appears on:[LaunchTemplateSpec](#LaunchTemplateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `cpuCredits` | ***string***| ***(Optional)*** |
## LaunchTemplateSpecElasticGpuSpecifications
##### (Appears on:[LaunchTemplateSpec](#LaunchTemplateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
## LaunchTemplateSpecElasticInferenceAccelerator
##### (Appears on:[LaunchTemplateSpec](#LaunchTemplateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
## LaunchTemplateSpecIamInstanceProfile
##### (Appears on:[LaunchTemplateSpec](#LaunchTemplateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `arn` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
## LaunchTemplateSpecInstanceMarketOptions
##### (Appears on:[LaunchTemplateSpec](#LaunchTemplateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `marketType` | ***string***| ***(Optional)*** |
| `spotOptions` | ***[[]LaunchTemplateSpecInstanceMarketOptionsSpotOptions](#LaunchTemplateSpecInstanceMarketOptionsSpotOptions)***| ***(Optional)*** |
## LaunchTemplateSpecInstanceMarketOptionsSpotOptions
##### (Appears on:[LaunchTemplateSpecInstanceMarketOptions](#LaunchTemplateSpecInstanceMarketOptions))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `blockDurationMinutes` | ***int***| ***(Optional)*** |
| `instanceInterruptionBehavior` | ***string***| ***(Optional)*** |
| `maxPrice` | ***string***| ***(Optional)*** |
| `spotInstanceType` | ***string***| ***(Optional)*** |
| `validUntil` | ***string***| ***(Optional)*** |
## LaunchTemplateSpecLicenseSpecification
##### (Appears on:[LaunchTemplateSpec](#LaunchTemplateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `licenseConfigurationArn` | ***string***||
## LaunchTemplateSpecMonitoring
##### (Appears on:[LaunchTemplateSpec](#LaunchTemplateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
## LaunchTemplateSpecNetworkInterfaces
##### (Appears on:[LaunchTemplateSpec](#LaunchTemplateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `associatePublicIPAddress` | ***bool***| ***(Optional)*** |
| `deleteOnTermination` | ***bool***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `deviceIndex` | ***int***| ***(Optional)*** |
| `ipv4AddressCount` | ***int***| ***(Optional)*** |
| `ipv4Addresses` | ***[]string***| ***(Optional)*** |
| `ipv6AddressCount` | ***int***| ***(Optional)*** |
| `ipv6Addresses` | ***[]string***| ***(Optional)*** |
| `networkInterfaceID` | ***string***| ***(Optional)*** |
| `privateIPAddress` | ***string***| ***(Optional)*** |
| `securityGroups` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
## LaunchTemplateSpecPlacement
##### (Appears on:[LaunchTemplateSpec](#LaunchTemplateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `affinity` | ***string***| ***(Optional)*** |
| `availabilityZone` | ***string***| ***(Optional)*** |
| `groupName` | ***string***| ***(Optional)*** |
| `hostID` | ***string***| ***(Optional)*** |
| `spreadDomain` | ***string***| ***(Optional)*** |
| `tenancy` | ***string***| ***(Optional)*** |
## LaunchTemplateSpecTagSpecifications
##### (Appears on:[LaunchTemplateSpec](#LaunchTemplateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `resourceType` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## LaunchTemplateStatus
##### (Appears on:[LaunchTemplate](#LaunchTemplate))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LaunchTemplateSpec](#LaunchTemplateSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
