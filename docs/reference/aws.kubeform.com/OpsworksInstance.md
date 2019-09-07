---
title: OpsworksInstance
menu:
  docs_v0.0.1:
    identifier: opsworksinstance-aws.kubeform.com
    name: OpsworksInstance
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## OpsworksInstance
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `OpsworksInstance` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[OpsworksInstanceSpec](#OpsworksInstanceSpec)***||
| `status` | ***[OpsworksInstanceStatus](#OpsworksInstanceStatus)***||
## OpsworksInstanceSpec
##### (Appears on:[OpsworksInstance](#OpsworksInstance), [OpsworksInstanceStatus](#OpsworksInstanceStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `agentVersion` | ***string***| ***(Optional)*** |
| `amiID` | ***string***| ***(Optional)*** |
| `architecture` | ***string***| ***(Optional)*** |
| `autoScalingType` | ***string***| ***(Optional)*** |
| `availabilityZone` | ***string***| ***(Optional)*** |
| `createdAt` | ***string***| ***(Optional)*** |
| `deleteEbs` | ***bool***| ***(Optional)*** |
| `deleteEip` | ***bool***| ***(Optional)*** |
| `ebsBlockDevice` | ***[[]OpsworksInstanceSpecEbsBlockDevice](#OpsworksInstanceSpecEbsBlockDevice)***| ***(Optional)*** |
| `ebsOptimized` | ***bool***| ***(Optional)*** |
| `ec2InstanceID` | ***string***| ***(Optional)*** |
| `ecsClusterArn` | ***string***| ***(Optional)*** |
| `elasticIP` | ***string***| ***(Optional)*** |
| `ephemeralBlockDevice` | ***[[]OpsworksInstanceSpecEphemeralBlockDevice](#OpsworksInstanceSpecEphemeralBlockDevice)***| ***(Optional)*** |
| `hostname` | ***string***| ***(Optional)*** |
| `infrastructureClass` | ***string***| ***(Optional)*** |
| `installUpdatesOnBoot` | ***bool***| ***(Optional)*** |
| `instanceProfileArn` | ***string***| ***(Optional)*** |
| `instanceType` | ***string***| ***(Optional)*** |
| `lastServiceErrorID` | ***string***| ***(Optional)*** |
| `layerIDS` | ***[]string***||
| `os` | ***string***| ***(Optional)*** |
| `platform` | ***string***| ***(Optional)*** |
| `privateDNS` | ***string***| ***(Optional)*** |
| `privateIP` | ***string***| ***(Optional)*** |
| `publicDNS` | ***string***| ***(Optional)*** |
| `publicIP` | ***string***| ***(Optional)*** |
| `registeredBy` | ***string***| ***(Optional)*** |
| `reportedAgentVersion` | ***string***| ***(Optional)*** |
| `reportedOsFamily` | ***string***| ***(Optional)*** |
| `reportedOsName` | ***string***| ***(Optional)*** |
| `reportedOsVersion` | ***string***| ***(Optional)*** |
| `rootBlockDevice` | ***[[]OpsworksInstanceSpecRootBlockDevice](#OpsworksInstanceSpecRootBlockDevice)***| ***(Optional)*** |
| `rootDeviceType` | ***string***| ***(Optional)*** |
| `rootDeviceVolumeID` | ***string***| ***(Optional)*** |
| `securityGroupIDS` | ***[]string***| ***(Optional)*** |
| `sshHostDsaKeyFingerprint` | ***string***| ***(Optional)*** |
| `sshHostRsaKeyFingerprint` | ***string***| ***(Optional)*** |
| `sshKeyName` | ***string***| ***(Optional)*** |
| `stackID` | ***string***||
| `state` | ***string***| ***(Optional)*** |
| `status` | ***string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `tenancy` | ***string***| ***(Optional)*** |
| `virtualizationType` | ***string***| ***(Optional)*** |
## OpsworksInstanceSpecEbsBlockDevice
##### (Appears on:[OpsworksInstanceSpec](#OpsworksInstanceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `deleteOnTermination` | ***bool***| ***(Optional)*** |
| `deviceName` | ***string***||
| `iops` | ***int***| ***(Optional)*** |
| `snapshotID` | ***string***| ***(Optional)*** |
| `volumeSize` | ***int***| ***(Optional)*** |
| `volumeType` | ***string***| ***(Optional)*** |
## OpsworksInstanceSpecEphemeralBlockDevice
##### (Appears on:[OpsworksInstanceSpec](#OpsworksInstanceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `deviceName` | ***string***||
| `virtualName` | ***string***||
## OpsworksInstanceSpecRootBlockDevice
##### (Appears on:[OpsworksInstanceSpec](#OpsworksInstanceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `deleteOnTermination` | ***bool***| ***(Optional)*** |
| `iops` | ***int***| ***(Optional)*** |
| `volumeSize` | ***int***| ***(Optional)*** |
| `volumeType` | ***string***| ***(Optional)*** |
## OpsworksInstanceStatus
##### (Appears on:[OpsworksInstance](#OpsworksInstance))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[OpsworksInstanceSpec](#OpsworksInstanceSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
