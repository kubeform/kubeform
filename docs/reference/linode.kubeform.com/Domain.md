---
title: Domain
menu:
  docs_v0.0.1:
    identifier: domain-linode.kubeform.com
    name: Domain
    parent: linode.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## Domain
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `linode.kubeform.com/v1alpha1` |
|    `kind` | string | `Domain` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DomainSpec](#DomainSpec)***||
| `status` | ***[DomainStatus](#DomainStatus)***||
## DomainSpec
##### (Appears on:[Domain](#Domain), [DomainStatus](#DomainStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `axfrIPS` | ***[]string***| ***(Optional)*** The list of IPs that may perform a zone transfer for this Domain. This is potentially dangerous, and should be set to an empty list unless you intend to use it.|
| `description` | ***string***| ***(Optional)*** A description for this Domain. This is for display purposes only.|
| `domain` | ***string***|The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain.|
| `expireSec` | ***int***| ***(Optional)*** The amount of time in seconds that may pass before this Domain is no longer authoritative. Valid values are 0, 00, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.|
| `group` | ***string***| ***(Optional)*** The group this Domain belongs to. This is for display purposes only.|
| `masterIPS` | ***[]string***| ***(Optional)*** The IP addresses representing the master DNS for this Domain.|
| `refreshSec` | ***int***| ***(Optional)*** The amount of time in seconds before this Domain should be refreshed. Valid values are 0, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.|
| `retrySec` | ***int***| ***(Optional)*** The interval, in seconds, at which a failed refresh should be retried. Valid values are 0, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.|
| `soaEmail` | ***string***| ***(Optional)*** Start of Authority email address. This is required for master Domains.|
| `status` | ***string***| ***(Optional)*** Used to control whether this Domain is currently being rendered.|
| `tags` | ***[]string***| ***(Optional)*** An array of tags applied to this object. Tags are for organizational purposes only.|
| `ttlSec` | ***int***| ***(Optional)*** 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 0, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.|
| `type` | ***string***|If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave).|
## DomainStatus
##### (Appears on:[Domain](#Domain))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DomainSpec](#DomainSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
