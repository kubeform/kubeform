---
title: SpotFleetRequest
menu:
  docs_{{ .version }}:
    identifier: spotfleetrequest-aws.kubeform.com
    name: SpotFleetRequest
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SpotFleetRequest
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SpotFleetRequest` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SpotFleetRequestSpec](#SpotFleetRequestSpec)***||
| `status` | ***[SpotFleetRequestStatus](#SpotFleetRequestStatus)***||
## SpotFleetRequestSpec
##### (Appears on:[SpotFleetRequest](#SpotFleetRequest), [SpotFleetRequestStatus](#SpotFleetRequestStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allocationStrategy` | ***string***| ***(Optional)*** |
| `clientToken` | ***string***| ***(Optional)*** |
| `excessCapacityTerminationPolicy` | ***string***| ***(Optional)*** |
| `fleetType` | ***string***| ***(Optional)*** |
| `iamFleetRole` | ***string***||
| `instanceInterruptionBehaviour` | ***string***| ***(Optional)*** |
| `instancePoolsToUseCount` | ***int***| ***(Optional)*** |
| `launchSpecification` | ***[[]SpotFleetRequestSpecLaunchSpecification](#SpotFleetRequestSpecLaunchSpecification)***||
| `loadBalancers` | ***[]string***| ***(Optional)*** |
| `replaceUnhealthyInstances` | ***bool***| ***(Optional)*** |
| `spotPrice` | ***string***| ***(Optional)*** |
| `spotRequestState` | ***string***| ***(Optional)*** |
| `targetCapacity` | ***int***||
| `targetGroupArns` | ***[]string***| ***(Optional)*** |
| `terminateInstancesWithExpiration` | ***bool***| ***(Optional)*** |
| `validFrom` | ***string***| ***(Optional)*** |
| `validUntil` | ***string***| ***(Optional)*** |
| `waitForFulfillment` | ***bool***| ***(Optional)*** |
## SpotFleetRequestSpecLaunchSpecification
##### (Appears on:[SpotFleetRequestSpec](#SpotFleetRequestSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ami` | ***string***||
| `associatePublicIPAddress` | ***bool***| ***(Optional)*** |
| `availabilityZone` | ***string***| ***(Optional)*** |
| `ebsBlockDevice` | ***[[]SpotFleetRequestSpecLaunchSpecificationEbsBlockDevice](#SpotFleetRequestSpecLaunchSpecificationEbsBlockDevice)***| ***(Optional)*** |
| `ebsOptimized` | ***bool***| ***(Optional)*** |
| `ephemeralBlockDevice` | ***[[]SpotFleetRequestSpecLaunchSpecificationEphemeralBlockDevice](#SpotFleetRequestSpecLaunchSpecificationEphemeralBlockDevice)***| ***(Optional)*** |
| `iamInstanceProfile` | ***string***| ***(Optional)*** |
| `iamInstanceProfileArn` | ***string***| ***(Optional)*** |
| `instanceType` | ***string***||
| `keyName` | ***string***| ***(Optional)*** |
| `monitoring` | ***bool***| ***(Optional)*** |
| `placementGroup` | ***string***| ***(Optional)*** |
| `placementTenancy` | ***string***| ***(Optional)*** |
| `rootBlockDevice` | ***[[]SpotFleetRequestSpecLaunchSpecificationRootBlockDevice](#SpotFleetRequestSpecLaunchSpecificationRootBlockDevice)***| ***(Optional)*** |
| `spotPrice` | ***string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `userData` | ***string***| ***(Optional)*** |
| `vpcSecurityGroupIDS` | ***[]string***| ***(Optional)*** |
| `weightedCapacity` | ***string***| ***(Optional)*** |
## SpotFleetRequestSpecLaunchSpecificationEbsBlockDevice
##### (Appears on:[SpotFleetRequestSpecLaunchSpecification](#SpotFleetRequestSpecLaunchSpecification))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `deleteOnTermination` | ***bool***| ***(Optional)*** |
| `deviceName` | ***string***||
| `encrypted` | ***bool***| ***(Optional)*** |
| `iops` | ***int***| ***(Optional)*** |
| `snapshotID` | ***string***| ***(Optional)*** |
| `volumeSize` | ***int***| ***(Optional)*** |
| `volumeType` | ***string***| ***(Optional)*** |
## SpotFleetRequestSpecLaunchSpecificationEphemeralBlockDevice
##### (Appears on:[SpotFleetRequestSpecLaunchSpecification](#SpotFleetRequestSpecLaunchSpecification))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `deviceName` | ***string***||
| `virtualName` | ***string***||
## SpotFleetRequestSpecLaunchSpecificationRootBlockDevice
##### (Appears on:[SpotFleetRequestSpecLaunchSpecification](#SpotFleetRequestSpecLaunchSpecification))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `deleteOnTermination` | ***bool***| ***(Optional)*** |
| `iops` | ***int***| ***(Optional)*** |
| `volumeSize` | ***int***| ***(Optional)*** |
| `volumeType` | ***string***| ***(Optional)*** |
## SpotFleetRequestStatus
##### (Appears on:[SpotFleetRequest](#SpotFleetRequest))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SpotFleetRequestSpec](#SpotFleetRequestSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
