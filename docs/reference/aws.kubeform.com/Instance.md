---
title: Instance
menu:
  docs_v0.0.1:
    identifier: instance-aws.kubeform.com
    name: Instance
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## Instance
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Instance` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[InstanceSpec](#InstanceSpec)***||
| `status` | ***[InstanceStatus](#InstanceStatus)***||
## InstanceSpec
##### (Appears on:[Instance](#Instance), [InstanceStatus](#InstanceStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `ami` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `associatePublicIPAddress` | ***bool***| ***(Optional)*** |
| `availabilityZone` | ***string***| ***(Optional)*** |
| `cpuCoreCount` | ***int***| ***(Optional)*** |
| `cpuThreadsPerCore` | ***int***| ***(Optional)*** |
| `creditSpecification` | ***[[]InstanceSpecCreditSpecification](#InstanceSpecCreditSpecification)***| ***(Optional)*** |
| `disableAPITermination` | ***bool***| ***(Optional)*** |
| `ebsBlockDevice` | ***[[]InstanceSpecEbsBlockDevice](#InstanceSpecEbsBlockDevice)***| ***(Optional)*** |
| `ebsOptimized` | ***bool***| ***(Optional)*** |
| `ephemeralBlockDevice` | ***[[]InstanceSpecEphemeralBlockDevice](#InstanceSpecEphemeralBlockDevice)***| ***(Optional)*** |
| `getPasswordData` | ***bool***| ***(Optional)*** |
| `hostID` | ***string***| ***(Optional)*** |
| `iamInstanceProfile` | ***string***| ***(Optional)*** |
| `instanceInitiatedShutdownBehavior` | ***string***| ***(Optional)*** |
| `instanceState` | ***string***| ***(Optional)*** |
| `instanceType` | ***string***||
| `ipv6AddressCount` | ***int***| ***(Optional)*** |
| `ipv6Addresses` | ***[]string***| ***(Optional)*** |
| `keyName` | ***string***| ***(Optional)*** |
| `monitoring` | ***bool***| ***(Optional)*** |
| `networkInterface` | ***[[]InstanceSpecNetworkInterface](#InstanceSpecNetworkInterface)***| ***(Optional)*** |
| `passwordData` | ***string***| ***(Optional)*** |
| `placementGroup` | ***string***| ***(Optional)*** |
| `primaryNetworkInterfaceID` | ***string***| ***(Optional)*** |
| `privateDNS` | ***string***| ***(Optional)*** |
| `privateIP` | ***string***| ***(Optional)*** |
| `publicDNS` | ***string***| ***(Optional)*** |
| `publicIP` | ***string***| ***(Optional)*** |
| `rootBlockDevice` | ***[[]InstanceSpecRootBlockDevice](#InstanceSpecRootBlockDevice)***| ***(Optional)*** |
| `securityGroups` | ***[]string***| ***(Optional)*** |
| `sourceDestCheck` | ***bool***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `tenancy` | ***string***| ***(Optional)*** |
| `userData` | ***string***| ***(Optional)*** |
| `userDataBase64` | ***string***| ***(Optional)*** |
| `volumeTags` | ***map[string]string***| ***(Optional)*** |
| `vpcSecurityGroupIDS` | ***[]string***| ***(Optional)*** |
## InstanceSpecCreditSpecification
##### (Appears on:[InstanceSpec](#InstanceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `cpuCredits` | ***string***| ***(Optional)*** |
## InstanceSpecEbsBlockDevice
##### (Appears on:[InstanceSpec](#InstanceSpec))
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
## InstanceSpecEphemeralBlockDevice
##### (Appears on:[InstanceSpec](#InstanceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `deviceName` | ***string***||
| `noDevice` | ***bool***| ***(Optional)*** |
| `virtualName` | ***string***| ***(Optional)*** |
## InstanceSpecNetworkInterface
##### (Appears on:[InstanceSpec](#InstanceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `deleteOnTermination` | ***bool***| ***(Optional)*** |
| `deviceIndex` | ***int***||
| `networkInterfaceID` | ***string***||
## InstanceSpecRootBlockDevice
##### (Appears on:[InstanceSpec](#InstanceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `deleteOnTermination` | ***bool***| ***(Optional)*** |
| `iops` | ***int***| ***(Optional)*** |
| `volumeID` | ***string***| ***(Optional)*** |
| `volumeSize` | ***int***| ***(Optional)*** |
| `volumeType` | ***string***| ***(Optional)*** |
## InstanceStatus
##### (Appears on:[Instance](#Instance))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[InstanceSpec](#InstanceSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
