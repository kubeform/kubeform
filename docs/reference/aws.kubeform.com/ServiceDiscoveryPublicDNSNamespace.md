---
title: ServiceDiscoveryPublicDNSNamespace
menu:
  docs_{{ .version }}:
    identifier: servicediscoverypublicdnsnamespace-aws.kubeform.com
    name: ServiceDiscoveryPublicDNSNamespace
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ServiceDiscoveryPublicDNSNamespace
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ServiceDiscoveryPublicDNSNamespace` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ServiceDiscoveryPublicDNSNamespaceSpec](#servicediscoverypublicdnsnamespacespec)***||
| `status` | ***[ServiceDiscoveryPublicDNSNamespaceStatus](#servicediscoverypublicdnsnamespacestatus)***||
## Phase(`string` alias)

Appears on:[ServiceDiscoveryPublicDNSNamespaceStatus](#servicediscoverypublicdnsnamespacestatus)

## ServiceDiscoveryPublicDNSNamespaceSpec

Appears on:[ServiceDiscoveryPublicDNSNamespace](#servicediscoverypublicdnsnamespace), [ServiceDiscoveryPublicDNSNamespaceStatus](#servicediscoverypublicdnsnamespacestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `hostedZone` | ***string***| ***(Optional)*** |
| `name` | ***string***||
## ServiceDiscoveryPublicDNSNamespaceStatus

Appears on:[ServiceDiscoveryPublicDNSNamespace](#servicediscoverypublicdnsnamespace)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ServiceDiscoveryPublicDNSNamespaceSpec](#servicediscoverypublicdnsnamespacespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
