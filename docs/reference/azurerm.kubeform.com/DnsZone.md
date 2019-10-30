---
title: DnsZone
menu:
  docs_{{ .version }}:
    identifier: dnszone-azurerm.kubeform.com
    name: DnsZone
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DnsZone
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DnsZone` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DnsZoneSpec](#dnszonespec)***||
| `status` | ***[DnsZoneStatus](#dnszonestatus)***||
## DnsZoneSpec

Appears on:[DnsZone](#dnszone), [DnsZoneStatus](#dnszonestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `maxNumberOfRecordSets` | ***int***| ***(Optional)*** |
| `name` | ***string***||
| `nameServers` | ***[]string***| ***(Optional)*** |
| `numberOfRecordSets` | ***int***| ***(Optional)*** |
| `registrationVirtualNetworkIDS` | ***[]string***| ***(Optional)*** |
| `resolutionVirtualNetworkIDS` | ***[]string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `zoneType` | ***string***| ***(Optional)*** |
## DnsZoneStatus

Appears on:[DnsZone](#dnszone)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DnsZoneSpec](#dnszonespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DnsZoneStatus](#dnszonestatus)

---
