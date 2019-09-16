---
title: Certificate
menu:
  docs_{{ .version }}:
    identifier: certificate-digitalocean.kubeform.com
    name: Certificate
    parent: digitalocean.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Certificate
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
|    `kind` | string | `Certificate` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CertificateSpec](#certificatespec)***||
| `status` | ***[CertificateStatus](#certificatestatus)***||
## CertificateSpec

Appears on:[Certificate](#certificate), [CertificateStatus](#certificatestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `certificateChain` | ***string***| ***(Optional)*** |
| `domains` | ***[]string***| ***(Optional)*** |
| `leafCertificate` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `notAfter` | ***string***| ***(Optional)*** |
| `sha1Fingerprint` | ***string***| ***(Optional)*** |
| `state` | ***string***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
## CertificateStatus

Appears on:[Certificate](#certificate)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CertificateSpec](#certificatespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `private_key` | ***string*** ||
