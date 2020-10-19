---
title: Droplet
menu:
  docs_{{ .version }}:
    identifier: droplet-digitalocean.kubeform.com
    name: Droplet
    parent: digitalocean.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Droplet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
|    `kind` | string | `Droplet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DropletSpec](#dropletspec)***||
| `status` | ***[DropletStatus](#dropletstatus)***||
## DropletSpec

Appears on:[Droplet](#droplet), [DropletStatus](#dropletstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `backups` | ***bool***| ***(Optional)*** |
| `createdAt` | ***string***| ***(Optional)*** |
| `disk` | ***int64***| ***(Optional)*** |
| `image` | ***string***||
| `ipv4Address` | ***string***| ***(Optional)*** |
| `ipv4AddressPrivate` | ***string***| ***(Optional)*** |
| `ipv6` | ***bool***| ***(Optional)*** |
| `ipv6Address` | ***string***| ***(Optional)*** |
| `locked` | ***bool***| ***(Optional)*** |
| `memory` | ***int64***| ***(Optional)*** |
| `monitoring` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `priceHourly` | ***float64***| ***(Optional)*** |
| `priceMonthly` | ***float64***| ***(Optional)*** |
| `privateNetworking` | ***bool***| ***(Optional)*** |
| `region` | ***string***||
| `resizeDisk` | ***bool***| ***(Optional)*** |
| `size` | ***string***||
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `status` | ***string***| ***(Optional)*** |
| `tags` | ***[]string***| ***(Optional)*** |
| `urn` | ***string***| ***(Optional)*** |
| `userData` | ***string***| ***(Optional)*** |
| `vcpus` | ***int64***| ***(Optional)*** |
| `volumeIDS` | ***[]string***| ***(Optional)*** |
| `vpcUUID` | ***string***| ***(Optional)*** |
## DropletStatus

Appears on:[Droplet](#droplet)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DropletSpec](#dropletspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DropletStatus](#dropletstatus)

---
