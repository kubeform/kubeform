---
title: Vpc
menu:
  docs_{{ .version }}:
    identifier: vpc-digitalocean.kubeform.com
    name: Vpc
    parent: digitalocean.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Vpc
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
|    `kind` | string | `Vpc` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VpcSpec](#vpcspec)***||
| `status` | ***[VpcStatus](#vpcstatus)***||
## Phase(`string` alias)

Appears on:[VpcStatus](#vpcstatus)

## VpcSpec

Appears on:[Vpc](#vpc), [VpcStatus](#vpcstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `createdAt` | ***string***| ***(Optional)*** The date and time of when the VPC was created|
| `default` | ***bool***| ***(Optional)*** Whether or not the VPC is the default one for the region|
| `description` | ***string***| ***(Optional)*** A free-form description for the VPC|
| `ipRange` | ***string***| ***(Optional)*** The range of IP addresses for the VPC in CIDR notation|
| `name` | ***string***|The name of the VPC|
| `region` | ***string***|DigitalOcean region slug for the VPC's location|
| `urn` | ***string***| ***(Optional)*** The uniform resource name (URN) for the VPC|
## VpcStatus

Appears on:[Vpc](#vpc)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VpcSpec](#vpcspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
