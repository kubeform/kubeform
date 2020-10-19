---
title: Instance
menu:
  docs_{{ .version }}:
    identifier: instance-aws.kubeform.com
    name: Instance
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Instance
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Instance` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[InstanceSpec](#instancespec)***||
| `status` | ***[InstanceStatus](#instancestatus)***||
## InstanceSpec

Appears on:[Instance](#instance), [InstanceStatus](#instancestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `ami` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `associatePublicIPAddress` | ***bool***| ***(Optional)*** |
| `availabilityZone` | ***string***| ***(Optional)*** |
| `cpuCoreCount` | ***int64***| ***(Optional)*** |
| `cpuThreadsPerCore` | ***int64***| ***(Optional)*** |
| `creditSpecification` | ***[[]InstanceSpecCreditSpecification](#instancespeccreditspecification)***| ***(Optional)*** |
| `disableAPITermination` | ***bool***| ***(Optional)*** |
| `ebsBlockDevice` | ***[[]InstanceSpecEbsBlockDevice](#instancespecebsblockdevice)***| ***(Optional)*** |
| `ebsOptimized` | ***bool***| ***(Optional)*** |
| `ephemeralBlockDevice` | ***[[]InstanceSpecEphemeralBlockDevice](#instancespecephemeralblockdevice)***| ***(Optional)*** |
| `getPasswordData` | ***bool***| ***(Optional)*** |
| `hostID` | ***string***| ***(Optional)*** |
| `iamInstanceProfile` | ***string***| ***(Optional)*** |
| `instanceInitiatedShutdownBehavior` | ***string***| ***(Optional)*** |
| `instanceState` | ***string***| ***(Optional)*** |
| `instanceType` | ***string***||
| `ipv6AddressCount` | ***int64***| ***(Optional)*** |
| `ipv6Addresses` | ***[]string***| ***(Optional)*** |
| `keyName` | ***string***| ***(Optional)*** |
| `monitoring` | ***bool***| ***(Optional)*** |
| `networkInterface` | ***[[]InstanceSpecNetworkInterface](#instancespecnetworkinterface)***| ***(Optional)*** |
| `passwordData` | ***string***| ***(Optional)*** |
| `placementGroup` | ***string***| ***(Optional)*** |
| `primaryNetworkInterfaceID` | ***string***| ***(Optional)*** |
| `privateDNS` | ***string***| ***(Optional)*** |
| `privateIP` | ***string***| ***(Optional)*** |
| `publicDNS` | ***string***| ***(Optional)*** |
| `publicIP` | ***string***| ***(Optional)*** |
| `rootBlockDevice` | ***[[]InstanceSpecRootBlockDevice](#instancespecrootblockdevice)***| ***(Optional)*** |
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

Appears on:[InstanceSpec](#instancespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cpuCredits` | ***string***| ***(Optional)*** |
## InstanceSpecEbsBlockDevice

Appears on:[InstanceSpec](#instancespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `deleteOnTermination` | ***bool***| ***(Optional)*** |
| `deviceName` | ***string***||
| `encrypted` | ***bool***| ***(Optional)*** |
| `iops` | ***int64***| ***(Optional)*** |
| `kmsKeyID` | ***string***| ***(Optional)*** |
| `snapshotID` | ***string***| ***(Optional)*** |
| `volumeID` | ***string***| ***(Optional)*** |
| `volumeSize` | ***int64***| ***(Optional)*** |
| `volumeType` | ***string***| ***(Optional)*** |
## InstanceSpecEphemeralBlockDevice

Appears on:[InstanceSpec](#instancespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `deviceName` | ***string***||
| `noDevice` | ***bool***| ***(Optional)*** |
| `virtualName` | ***string***| ***(Optional)*** |
## InstanceSpecNetworkInterface

Appears on:[InstanceSpec](#instancespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `deleteOnTermination` | ***bool***| ***(Optional)*** |
| `deviceIndex` | ***int64***||
| `networkInterfaceID` | ***string***||
## InstanceSpecRootBlockDevice

Appears on:[InstanceSpec](#instancespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `deleteOnTermination` | ***bool***| ***(Optional)*** |
| `encrypted` | ***bool***| ***(Optional)*** |
| `iops` | ***int64***| ***(Optional)*** |
| `kmsKeyID` | ***string***| ***(Optional)*** |
| `volumeID` | ***string***| ***(Optional)*** |
| `volumeSize` | ***int64***| ***(Optional)*** |
| `volumeType` | ***string***| ***(Optional)*** |
## InstanceStatus

Appears on:[Instance](#instance)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[InstanceSpec](#instancespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[InstanceStatus](#instancestatus)

---
