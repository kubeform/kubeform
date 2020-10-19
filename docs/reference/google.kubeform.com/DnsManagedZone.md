---
title: DnsManagedZone
menu:
  docs_{{ .version }}:
    identifier: dnsmanagedzone-google.kubeform.com
    name: DnsManagedZone
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DnsManagedZone
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `DnsManagedZone` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DnsManagedZoneSpec](#dnsmanagedzonespec)***||
| `status` | ***[DnsManagedZoneStatus](#dnsmanagedzonestatus)***||
## DnsManagedZoneSpec

Appears on:[DnsManagedZone](#dnsmanagedzone), [DnsManagedZoneStatus](#dnsmanagedzonestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `dnsName` | ***string***||
| `dnssecConfig` | ***[[]DnsManagedZoneSpecDnssecConfig](#dnsmanagedzonespecdnssecconfig)***| ***(Optional)*** |
| `labels` | ***map[string]string***| ***(Optional)*** |
| `name` | ***string***||
| `nameServers` | ***[]string***| ***(Optional)*** |
| `privateVisibilityConfig` | ***[[]DnsManagedZoneSpecPrivateVisibilityConfig](#dnsmanagedzonespecprivatevisibilityconfig)***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `visibility` | ***string***| ***(Optional)*** |
## DnsManagedZoneSpecDnssecConfig

Appears on:[DnsManagedZoneSpec](#dnsmanagedzonespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `defaultKeySpecs` | ***[[]DnsManagedZoneSpecDnssecConfigDefaultKeySpecs](#dnsmanagedzonespecdnssecconfigdefaultkeyspecs)***| ***(Optional)*** |
| `kind` | ***string***| ***(Optional)*** |
| `nonExistence` | ***string***| ***(Optional)*** |
| `state` | ***string***| ***(Optional)*** |
## DnsManagedZoneSpecDnssecConfigDefaultKeySpecs

Appears on:[DnsManagedZoneSpecDnssecConfig](#dnsmanagedzonespecdnssecconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `algorithm` | ***string***| ***(Optional)*** |
| `keyLength` | ***int64***| ***(Optional)*** |
| `keyType` | ***string***| ***(Optional)*** |
| `kind` | ***string***| ***(Optional)*** |
## DnsManagedZoneSpecPrivateVisibilityConfig

Appears on:[DnsManagedZoneSpec](#dnsmanagedzonespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `networks` | ***[[]DnsManagedZoneSpecPrivateVisibilityConfigNetworks](#dnsmanagedzonespecprivatevisibilityconfignetworks)***| ***(Optional)*** |
## DnsManagedZoneSpecPrivateVisibilityConfigNetworks

Appears on:[DnsManagedZoneSpecPrivateVisibilityConfig](#dnsmanagedzonespecprivatevisibilityconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `networkURL` | ***string***| ***(Optional)*** |
## DnsManagedZoneStatus

Appears on:[DnsManagedZone](#dnsmanagedzone)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DnsManagedZoneSpec](#dnsmanagedzonespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DnsManagedZoneStatus](#dnsmanagedzonestatus)

---
