## SpotInstanceRequest
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SpotInstanceRequest` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SpotInstanceRequestSpec](#SpotInstanceRequestSpec)***||
| `status` | ***[SpotInstanceRequestStatus](#SpotInstanceRequestStatus)***||
## SpotInstanceRequestSpec
##### (Appears on:[SpotInstanceRequest](#SpotInstanceRequest), [SpotInstanceRequestStatus](#SpotInstanceRequestStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `ami` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `associatePublicIPAddress` | ***bool***| ***(Optional)*** |
| `availabilityZone` | ***string***| ***(Optional)*** |
| `blockDurationMinutes` | ***int***| ***(Optional)*** |
| `cpuCoreCount` | ***int***| ***(Optional)*** |
| `cpuThreadsPerCore` | ***int***| ***(Optional)*** |
| `creditSpecification` | ***[[]SpotInstanceRequestSpecCreditSpecification](#SpotInstanceRequestSpecCreditSpecification)***| ***(Optional)*** |
| `disableAPITermination` | ***bool***| ***(Optional)*** |
| `ebsBlockDevice` | ***[[]SpotInstanceRequestSpecEbsBlockDevice](#SpotInstanceRequestSpecEbsBlockDevice)***| ***(Optional)*** |
| `ebsOptimized` | ***bool***| ***(Optional)*** |
| `ephemeralBlockDevice` | ***[[]SpotInstanceRequestSpecEphemeralBlockDevice](#SpotInstanceRequestSpecEphemeralBlockDevice)***| ***(Optional)*** |
| `getPasswordData` | ***bool***| ***(Optional)*** |
| `hostID` | ***string***| ***(Optional)*** |
| `iamInstanceProfile` | ***string***| ***(Optional)*** |
| `instanceInitiatedShutdownBehavior` | ***string***| ***(Optional)*** |
| `instanceInterruptionBehaviour` | ***string***| ***(Optional)*** |
| `instanceState` | ***string***| ***(Optional)*** |
| `instanceType` | ***string***||
| `ipv6AddressCount` | ***int***| ***(Optional)*** |
| `ipv6Addresses` | ***[]string***| ***(Optional)*** |
| `keyName` | ***string***| ***(Optional)*** |
| `launchGroup` | ***string***| ***(Optional)*** |
| `monitoring` | ***bool***| ***(Optional)*** |
| `networkInterface` | ***[[]SpotInstanceRequestSpecNetworkInterface](#SpotInstanceRequestSpecNetworkInterface)***| ***(Optional)*** |
| `passwordData` | ***string***| ***(Optional)*** |
| `placementGroup` | ***string***| ***(Optional)*** |
| `primaryNetworkInterfaceID` | ***string***| ***(Optional)*** |
| `privateDNS` | ***string***| ***(Optional)*** |
| `privateIP` | ***string***| ***(Optional)*** |
| `publicDNS` | ***string***| ***(Optional)*** |
| `publicIP` | ***string***| ***(Optional)*** |
| `rootBlockDevice` | ***[[]SpotInstanceRequestSpecRootBlockDevice](#SpotInstanceRequestSpecRootBlockDevice)***| ***(Optional)*** |
| `securityGroups` | ***[]string***| ***(Optional)*** |
| `sourceDestCheck` | ***bool***| ***(Optional)*** |
| `spotBidStatus` | ***string***| ***(Optional)*** |
| `spotInstanceID` | ***string***| ***(Optional)*** |
| `spotPrice` | ***string***| ***(Optional)*** |
| `spotRequestState` | ***string***| ***(Optional)*** |
| `spotType` | ***string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `tenancy` | ***string***| ***(Optional)*** |
| `userData` | ***string***| ***(Optional)*** |
| `userDataBase64` | ***string***| ***(Optional)*** |
| `validFrom` | ***string***| ***(Optional)*** |
| `validUntil` | ***string***| ***(Optional)*** |
| `volumeTags` | ***map[string]string***| ***(Optional)*** |
| `vpcSecurityGroupIDS` | ***[]string***| ***(Optional)*** |
| `waitForFulfillment` | ***bool***| ***(Optional)*** |
## SpotInstanceRequestSpecCreditSpecification
##### (Appears on:[SpotInstanceRequestSpec](#SpotInstanceRequestSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `cpuCredits` | ***string***| ***(Optional)*** |
## SpotInstanceRequestSpecEbsBlockDevice
##### (Appears on:[SpotInstanceRequestSpec](#SpotInstanceRequestSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `deleteOnTermination` | ***bool***| ***(Optional)*** |
| `deviceName` | ***string***||
| `encrypted` | ***bool***| ***(Optional)*** |
| `iops` | ***int***| ***(Optional)*** |
| `snapshotID` | ***string***| ***(Optional)*** |
| `volumeID` | ***string***| ***(Optional)*** |
| `volumeSize` | ***int***| ***(Optional)*** |
| `volumeType` | ***string***| ***(Optional)*** |
## SpotInstanceRequestSpecEphemeralBlockDevice
##### (Appears on:[SpotInstanceRequestSpec](#SpotInstanceRequestSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `deviceName` | ***string***||
| `noDevice` | ***bool***| ***(Optional)*** |
| `virtualName` | ***string***| ***(Optional)*** |
## SpotInstanceRequestSpecNetworkInterface
##### (Appears on:[SpotInstanceRequestSpec](#SpotInstanceRequestSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `deleteOnTermination` | ***bool***| ***(Optional)*** |
| `deviceIndex` | ***int***||
| `networkInterfaceID` | ***string***||
## SpotInstanceRequestSpecRootBlockDevice
##### (Appears on:[SpotInstanceRequestSpec](#SpotInstanceRequestSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `deleteOnTermination` | ***bool***| ***(Optional)*** |
| `iops` | ***int***| ***(Optional)*** |
| `volumeID` | ***string***| ***(Optional)*** |
| `volumeSize` | ***int***| ***(Optional)*** |
| `volumeType` | ***string***| ***(Optional)*** |
## SpotInstanceRequestStatus
##### (Appears on:[SpotInstanceRequest](#SpotInstanceRequest))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SpotInstanceRequestSpec](#SpotInstanceRequestSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
