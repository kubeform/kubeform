---
title: Ec2CapacityReservation
menu:
  docs_{{ .version }}:
    identifier: ec2capacityreservation-aws.kubeform.com
    name: Ec2CapacityReservation
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Ec2CapacityReservation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Ec2CapacityReservation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[Ec2CapacityReservationSpec](#ec2capacityreservationspec)***||
| `status` | ***[Ec2CapacityReservationStatus](#ec2capacityreservationstatus)***||
## Ec2CapacityReservationSpec

Appears on:[Ec2CapacityReservation](#ec2capacityreservation), [Ec2CapacityReservationStatus](#ec2capacityreservationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `availabilityZone` | ***string***||
| `ebsOptimized` | ***bool***| ***(Optional)*** |
| `endDate` | ***string***| ***(Optional)*** |
| `endDateType` | ***string***| ***(Optional)*** |
| `ephemeralStorage` | ***bool***| ***(Optional)*** |
| `instanceCount` | ***int***||
| `instanceMatchCriteria` | ***string***| ***(Optional)*** |
| `instancePlatform` | ***string***||
| `instanceType` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `tenancy` | ***string***| ***(Optional)*** |
## Ec2CapacityReservationStatus

Appears on:[Ec2CapacityReservation](#ec2capacityreservation)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[Ec2CapacityReservationSpec](#ec2capacityreservationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
