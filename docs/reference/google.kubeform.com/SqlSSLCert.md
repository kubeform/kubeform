---
title: SqlSSLCert
menu:
  docs_{{ .version }}:
    identifier: sqlsslcert-google.kubeform.com
    name: SqlSSLCert
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SqlSSLCert
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `SqlSSLCert` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SqlSSLCertSpec](#sqlsslcertspec)***||
| `status` | ***[SqlSSLCertStatus](#sqlsslcertstatus)***||
## Phase(`string` alias)

Appears on:[SqlSSLCertStatus](#sqlsslcertstatus)

## SqlSSLCertSpec

Appears on:[SqlSSLCert](#sqlsslcert), [SqlSSLCertStatus](#sqlsslcertstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `cert` | ***string***| ***(Optional)*** |
| `certSerialNumber` | ***string***| ***(Optional)*** |
| `commonName` | ***string***||
| `createTime` | ***string***| ***(Optional)*** |
| `expirationTime` | ***string***| ***(Optional)*** |
| `instance` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `serverCaCert` | ***string***| ***(Optional)*** |
| `sha1Fingerprint` | ***string***| ***(Optional)*** |
## SqlSSLCertStatus

Appears on:[SqlSSLCert](#sqlsslcert)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SqlSSLCertSpec](#sqlsslcertspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `private_key` | ***string*** ||
