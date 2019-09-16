---
title: CloudiotRegistry
menu:
  docs_{{ .version }}:
    identifier: cloudiotregistry-google.kubeform.com
    name: CloudiotRegistry
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CloudiotRegistry
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `CloudiotRegistry` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CloudiotRegistrySpec](#cloudiotregistryspec)***||
| `status` | ***[CloudiotRegistryStatus](#cloudiotregistrystatus)***||
## CloudiotRegistrySpec

Appears on:[CloudiotRegistry](#cloudiotregistry), [CloudiotRegistryStatus](#cloudiotregistrystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `credentials` | ***[[]CloudiotRegistrySpecCredentials](#cloudiotregistryspeccredentials)***| ***(Optional)*** |
| `eventNotificationConfig` | ***map[string]kubeform.dev/kubeform/apis/google/v1alpha1.CloudiotRegistrySpecEventNotificationConfig***| ***(Optional)*** |
| `httpConfig` | ***map[string]kubeform.dev/kubeform/apis/google/v1alpha1.CloudiotRegistrySpecHttpConfig***| ***(Optional)*** |
| `mqttConfig` | ***map[string]kubeform.dev/kubeform/apis/google/v1alpha1.CloudiotRegistrySpecMqttConfig***| ***(Optional)*** |
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `stateNotificationConfig` | ***map[string]kubeform.dev/kubeform/apis/google/v1alpha1.CloudiotRegistrySpecStateNotificationConfig***| ***(Optional)*** |
## CloudiotRegistrySpecCredentials

Appears on:[CloudiotRegistrySpec](#cloudiotregistryspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `publicKeyCertificate` | ***map[string]kubeform.dev/kubeform/apis/google/v1alpha1.CloudiotRegistrySpecCredentialsPublicKeyCertificate***| ***(Optional)*** |
## CloudiotRegistryStatus

Appears on:[CloudiotRegistry](#cloudiotregistry)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CloudiotRegistrySpec](#cloudiotregistryspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
