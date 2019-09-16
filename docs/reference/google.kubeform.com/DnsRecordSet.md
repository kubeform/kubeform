---
title: DnsRecordSet
menu:
  docs_{{ .version }}:
    identifier: dnsrecordset-google.kubeform.com
    name: DnsRecordSet
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DnsRecordSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `DnsRecordSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DnsRecordSetSpec](#DnsRecordSetSpec)***||
| `status` | ***[DnsRecordSetStatus](#DnsRecordSetStatus)***||
## DnsRecordSetSpec

Appears on:[DnsRecordSet](#DnsRecordSet), [DnsRecordSetStatus](#DnsRecordSetStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `managedZone` | ***string***||
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `rrdatas` | ***[]string***||
| `ttl` | ***int***||
| `type` | ***string***||
## DnsRecordSetStatus

Appears on:[DnsRecordSet](#DnsRecordSet)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DnsRecordSetSpec](#DnsRecordSetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
