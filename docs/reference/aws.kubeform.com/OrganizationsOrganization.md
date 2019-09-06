## OrganizationsOrganization
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `OrganizationsOrganization` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[OrganizationsOrganizationSpec](#OrganizationsOrganizationSpec)***||
| `status` | ***[OrganizationsOrganizationStatus](#OrganizationsOrganizationStatus)***||
## OrganizationsOrganizationSpec
##### (Appears on:[OrganizationsOrganization](#OrganizationsOrganization), [OrganizationsOrganizationStatus](#OrganizationsOrganizationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accounts` | ***[[]OrganizationsOrganizationSpecAccounts](#OrganizationsOrganizationSpecAccounts)***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `awsServiceAccessPrincipals` | ***[]string***| ***(Optional)*** |
| `enabledPolicyTypes` | ***[]string***| ***(Optional)*** |
| `featureSet` | ***string***| ***(Optional)*** |
| `masterAccountArn` | ***string***| ***(Optional)*** |
| `masterAccountEmail` | ***string***| ***(Optional)*** |
| `masterAccountID` | ***string***| ***(Optional)*** |
| `roots` | ***[[]OrganizationsOrganizationSpecRoots](#OrganizationsOrganizationSpecRoots)***| ***(Optional)*** |
## OrganizationsOrganizationSpecAccounts
##### (Appears on:[OrganizationsOrganizationSpec](#OrganizationsOrganizationSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `arn` | ***string***| ***(Optional)*** |
| `email` | ***string***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
## OrganizationsOrganizationSpecRoots
##### (Appears on:[OrganizationsOrganizationSpec](#OrganizationsOrganizationSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `arn` | ***string***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `policyTypes` | ***[[]OrganizationsOrganizationSpecRootsPolicyTypes](#OrganizationsOrganizationSpecRootsPolicyTypes)***| ***(Optional)*** |
## OrganizationsOrganizationSpecRootsPolicyTypes
##### (Appears on:[OrganizationsOrganizationSpecRoots](#OrganizationsOrganizationSpecRoots))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `status` | ***string***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
## OrganizationsOrganizationStatus
##### (Appears on:[OrganizationsOrganization](#OrganizationsOrganization))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[OrganizationsOrganizationSpec](#OrganizationsOrganizationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
