---
title: GoogleServiceAccount
menu:
  docs_{{ .version }}:
    identifier: googleserviceaccount-modules.kubeform.com
    name: GoogleServiceAccount
    parent: modules.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## GoogleServiceAccount
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `modules.kubeform.com/v1alpha1` |
|    `kind` | string | `GoogleServiceAccount` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[GoogleServiceAccountSpec](#googleserviceaccountspec)***||
| `status` | ***[GoogleServiceAccountStatus](#googleserviceaccountstatus)***||
## GoogleServiceAccountSpec

Appears on:[GoogleServiceAccount](#googleserviceaccount)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***| ***(Optional)*** |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `source` | ***string***| ***(Optional)*** |
| `billingAccountID` | ***string***| ***(Optional)*** If assigning billing role, specificy a billing account (default is to assign at the organizational level).|
| `description` | ***string***| ***(Optional)*** Descriptions of the created service accounts (defaults to no description)|
| `displayName` | ***string***| ***(Optional)*** Display names of the created service accounts (defaults to 'Terraform-managed service account')|
| `generateKeys` | ***bool***| ***(Optional)*** Generate keys for service accounts.|
| `grantBillingRole` | ***bool***| ***(Optional)*** Grant billing user role.|
| `grantXpnRoles` | ***bool***| ***(Optional)*** Grant roles for shared VPC management.|
| `names` | ***[]string***| ***(Optional)*** Names of the service accounts to create.|
| `orgID` | ***string***| ***(Optional)*** Id of the organization for org-level roles.|
| `prefix` | ***string***| ***(Optional)*** Prefix applied to service account names.|
| `projectID` | ***string***| ***(Optional)*** Project id where service account will be created.|
| `projectRoles` | ***[]string***| ***(Optional)*** Common roles to apply to all service accounts, project=>role as elements.|
## GoogleServiceAccountStatus

Appears on:[GoogleServiceAccount](#googleserviceaccount)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***kubeform.dev/kubeform/apis/modules/v1alpha1.GoogleServiceAccountOutput***| ***(Optional)*** |
| `state` | ***string***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[GoogleServiceAccountStatus](#googleserviceaccountstatus)

---
