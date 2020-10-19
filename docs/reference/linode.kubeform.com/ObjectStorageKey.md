---
title: ObjectStorageKey
menu:
  docs_{{ .version }}:
    identifier: objectstoragekey-linode.kubeform.com
    name: ObjectStorageKey
    parent: linode.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ObjectStorageKey
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `linode.kubeform.com/v1alpha1` |
|    `kind` | string | `ObjectStorageKey` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ObjectStorageKeySpec](#objectstoragekeyspec)***||
| `status` | ***[ObjectStorageKeyStatus](#objectstoragekeystatus)***||
## ObjectStorageKeySpec

Appears on:[ObjectStorageKey](#objectstoragekey), [ObjectStorageKeyStatus](#objectstoragekeystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `accessKey` | ***string***| ***(Optional)*** This keypair's access key. This is not secret.|
| `label` | ***string***|The label given to this key. For display purposes only.|
## ObjectStorageKeyStatus

Appears on:[ObjectStorageKey](#objectstoragekey)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ObjectStorageKeySpec](#objectstoragekeyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ObjectStorageKeyStatus](#objectstoragekeystatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `secret_key` | ***string*** |This keypair's secret key.|
