---
title: SpotInstanceRequest
menu:
  docs_{{ .version }}:
    identifier: spotinstancerequest-aws.kubeform.com
    name: SpotInstanceRequest
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SpotInstanceRequest
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SpotInstanceRequest` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SpotInstanceRequestSpec](#spotinstancerequestspec)***||
| `status` | ***[SpotInstanceRequestStatus](#spotinstancerequeststatus)***||
## Phase(`string` alias)

Appears on:[SpotInstanceRequestStatus](#spotinstancerequeststatus)

## SpotInstanceRequestSpec

Appears on:[SpotInstanceRequest](#spotinstancerequest), [SpotInstanceRequestStatus](#spotinstancerequeststatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `ami` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `associatePublicIPAddress` | ***bool***| ***(Optional)*** |
| `availabilityZone` | ***string***| ***(Optional)*** |
| `blockDurationMinutes` | ***int64***| ***(Optional)*** |
| `cpuCoreCount` | ***int64***| ***(Optional)*** |
| `cpuThreadsPerCore` | ***int64***| ***(Optional)*** |
| `creditSpecification` | ***[[]SpotInstanceRequestSpecCreditSpecification](#spotinstancerequestspeccreditspecification)***| ***(Optional)*** |
| `disableAPITermination` | ***bool***| ***(Optional)*** |
| `ebsBlockDevice` | ***[[]SpotInstanceRequestSpecEbsBlockDevice](#spotinstancerequestspecebsblockdevice)***| ***(Optional)*** |
| `ebsOptimized` | ***bool***| ***(Optional)*** |
| `ephemeralBlockDevice` | ***[[]SpotInstanceRequestSpecEphemeralBlockDevice](#spotinstancerequestspecephemeralblockdevice)***| ***(Optional)*** |
| `getPasswordData` | ***bool***| ***(Optional)*** |
| `hostID` | ***string***| ***(Optional)*** |
| `iamInstanceProfile` | ***string***| ***(Optional)*** |
| `instanceInitiatedShutdownBehavior` | ***string***| ***(Optional)*** |
| `instanceInterruptionBehaviour` | ***string***| ***(Optional)*** |
| `instanceState` | ***string***| ***(Optional)*** |
| `instanceType` | ***string***||
| `ipv6AddressCount` | ***int64***| ***(Optional)*** |
| `ipv6Addresses` | ***[]string***| ***(Optional)*** |
| `keyName` | ***string***| ***(Optional)*** |
| `launchGroup` | ***string***| ***(Optional)*** |
| `monitoring` | ***bool***| ***(Optional)*** |
| `networkInterface` | ***[[]SpotInstanceRequestSpecNetworkInterface](#spotinstancerequestspecnetworkinterface)***| ***(Optional)*** |
| `passwordData` | ***string***| ***(Optional)*** |
| `placementGroup` | ***string***| ***(Optional)*** |
| `primaryNetworkInterfaceID` | ***string***| ***(Optional)*** |
| `privateDNS` | ***string***| ***(Optional)*** |
| `privateIP` | ***string***| ***(Optional)*** |
| `publicDNS` | ***string***| ***(Optional)*** |
| `publicIP` | ***string***| ***(Optional)*** |
| `rootBlockDevice` | ***[[]SpotInstanceRequestSpecRootBlockDevice](#spotinstancerequestspecrootblockdevice)***| ***(Optional)*** |
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

Appears on:[SpotInstanceRequestSpec](#spotinstancerequestspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cpuCredits` | ***string***| ***(Optional)*** |
## SpotInstanceRequestSpecEbsBlockDevice

Appears on:[SpotInstanceRequestSpec](#spotinstancerequestspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `deleteOnTermination` | ***bool***| ***(Optional)*** |
| `deviceName` | ***string***||
| `encrypted` | ***bool***| ***(Optional)*** |
| `iops` | ***int64***| ***(Optional)*** |
| `snapshotID` | ***string***| ***(Optional)*** |
| `volumeID` | ***string***| ***(Optional)*** |
| `volumeSize` | ***int64***| ***(Optional)*** |
| `volumeType` | ***string***| ***(Optional)*** |
## SpotInstanceRequestSpecEphemeralBlockDevice

Appears on:[SpotInstanceRequestSpec](#spotinstancerequestspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `deviceName` | ***string***||
| `noDevice` | ***bool***| ***(Optional)*** |
| `virtualName` | ***string***| ***(Optional)*** |
## SpotInstanceRequestSpecNetworkInterface

Appears on:[SpotInstanceRequestSpec](#spotinstancerequestspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `deleteOnTermination` | ***bool***| ***(Optional)*** |
| `deviceIndex` | ***int64***||
| `networkInterfaceID` | ***string***||
## SpotInstanceRequestSpecRootBlockDevice

Appears on:[SpotInstanceRequestSpec](#spotinstancerequestspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `deleteOnTermination` | ***bool***| ***(Optional)*** |
| `iops` | ***int64***| ***(Optional)*** |
| `volumeID` | ***string***| ***(Optional)*** |
| `volumeSize` | ***int64***| ***(Optional)*** |
| `volumeType` | ***string***| ***(Optional)*** |
## SpotInstanceRequestStatus

Appears on:[SpotInstanceRequest](#spotinstancerequest)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SpotInstanceRequestSpec](#spotinstancerequestspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
