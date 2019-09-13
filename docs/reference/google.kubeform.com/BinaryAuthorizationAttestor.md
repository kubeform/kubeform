---
title: BinaryAuthorizationAttestor
menu:
  docs_{{ .version }}:
    identifier: binaryauthorizationattestor-google.kubeform.com
    name: BinaryAuthorizationAttestor
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## BinaryAuthorizationAttestor
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `BinaryAuthorizationAttestor` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BinaryAuthorizationAttestorSpec](#BinaryAuthorizationAttestorSpec)***||
| `status` | ***[BinaryAuthorizationAttestorStatus](#BinaryAuthorizationAttestorStatus)***||
## BinaryAuthorizationAttestorSpec
##### (Appears on:[BinaryAuthorizationAttestor](#BinaryAuthorizationAttestor), [BinaryAuthorizationAttestorStatus](#BinaryAuthorizationAttestorStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `attestationAuthorityNote` | ***[[]BinaryAuthorizationAttestorSpecAttestationAuthorityNote](#BinaryAuthorizationAttestorSpecAttestationAuthorityNote)***||
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
## BinaryAuthorizationAttestorSpecAttestationAuthorityNote
##### (Appears on:[BinaryAuthorizationAttestorSpec](#BinaryAuthorizationAttestorSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `delegationServiceAccountEmail` | ***string***| ***(Optional)*** |
| `noteReference` | ***string***||
| `publicKeys` | ***[[]BinaryAuthorizationAttestorSpecAttestationAuthorityNotePublicKeys](#BinaryAuthorizationAttestorSpecAttestationAuthorityNotePublicKeys)***| ***(Optional)*** |
## BinaryAuthorizationAttestorSpecAttestationAuthorityNotePublicKeys
##### (Appears on:[BinaryAuthorizationAttestorSpecAttestationAuthorityNote](#BinaryAuthorizationAttestorSpecAttestationAuthorityNote))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `asciiArmoredPgpPublicKey` | ***string***||
| `comment` | ***string***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
## BinaryAuthorizationAttestorStatus
##### (Appears on:[BinaryAuthorizationAttestor](#BinaryAuthorizationAttestor))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BinaryAuthorizationAttestorSpec](#BinaryAuthorizationAttestorSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
