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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BinaryAuthorizationAttestorSpec](#binaryauthorizationattestorspec)***||
| `status` | ***[BinaryAuthorizationAttestorStatus](#binaryauthorizationattestorstatus)***||
## BinaryAuthorizationAttestorSpec

Appears on:[BinaryAuthorizationAttestor](#binaryauthorizationattestor), [BinaryAuthorizationAttestorStatus](#binaryauthorizationattestorstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `attestationAuthorityNote` | ***[[]BinaryAuthorizationAttestorSpecAttestationAuthorityNote](#binaryauthorizationattestorspecattestationauthoritynote)***||
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
## BinaryAuthorizationAttestorSpecAttestationAuthorityNote

Appears on:[BinaryAuthorizationAttestorSpec](#binaryauthorizationattestorspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `delegationServiceAccountEmail` | ***string***| ***(Optional)*** |
| `noteReference` | ***string***||
| `publicKeys` | ***[[]BinaryAuthorizationAttestorSpecAttestationAuthorityNotePublicKeys](#binaryauthorizationattestorspecattestationauthoritynotepublickeys)***| ***(Optional)*** |
## BinaryAuthorizationAttestorSpecAttestationAuthorityNotePublicKeys

Appears on:[BinaryAuthorizationAttestorSpecAttestationAuthorityNote](#binaryauthorizationattestorspecattestationauthoritynote)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `asciiArmoredPgpPublicKey` | ***string***| ***(Optional)*** |
| `comment` | ***string***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `pkixPublicKey` | ***[[]BinaryAuthorizationAttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKey](#binaryauthorizationattestorspecattestationauthoritynotepublickeyspkixpublickey)***| ***(Optional)*** |
## BinaryAuthorizationAttestorSpecAttestationAuthorityNotePublicKeysPkixPublicKey

Appears on:[BinaryAuthorizationAttestorSpecAttestationAuthorityNotePublicKeys](#binaryauthorizationattestorspecattestationauthoritynotepublickeys)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `publicKeyPem` | ***string***| ***(Optional)*** |
| `signatureAlgorithm` | ***string***| ***(Optional)*** |
## BinaryAuthorizationAttestorStatus

Appears on:[BinaryAuthorizationAttestor](#binaryauthorizationattestor)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BinaryAuthorizationAttestorSpec](#binaryauthorizationattestorspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[BinaryAuthorizationAttestorStatus](#binaryauthorizationattestorstatus)

---
