## OrganizationPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `OrganizationPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[OrganizationPolicySpec](#OrganizationPolicySpec)***||
| `status` | ***[OrganizationPolicyStatus](#OrganizationPolicyStatus)***||
## OrganizationPolicySpec
##### (Appears on:[OrganizationPolicy](#OrganizationPolicy), [OrganizationPolicyStatus](#OrganizationPolicyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `booleanPolicy` | ***[[]OrganizationPolicySpecBooleanPolicy](#OrganizationPolicySpecBooleanPolicy)***| ***(Optional)*** |
| `constraint` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `listPolicy` | ***[[]OrganizationPolicySpecListPolicy](#OrganizationPolicySpecListPolicy)***| ***(Optional)*** |
| `orgID` | ***string***||
| `restorePolicy` | ***[[]OrganizationPolicySpecRestorePolicy](#OrganizationPolicySpecRestorePolicy)***| ***(Optional)*** |
| `updateTime` | ***string***| ***(Optional)*** |
| `version` | ***int***| ***(Optional)*** |
## OrganizationPolicySpecBooleanPolicy
##### (Appears on:[OrganizationPolicySpec](#OrganizationPolicySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `enforced` | ***bool***||
## OrganizationPolicySpecListPolicy
##### (Appears on:[OrganizationPolicySpec](#OrganizationPolicySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `allow` | ***[[]OrganizationPolicySpecListPolicyAllow](#OrganizationPolicySpecListPolicyAllow)***| ***(Optional)*** |
| `deny` | ***[[]OrganizationPolicySpecListPolicyDeny](#OrganizationPolicySpecListPolicyDeny)***| ***(Optional)*** |
| `suggestedValue` | ***string***| ***(Optional)*** |
## OrganizationPolicySpecListPolicyAllow
##### (Appears on:[OrganizationPolicySpecListPolicy](#OrganizationPolicySpecListPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `all` | ***bool***| ***(Optional)*** |
| `values` | ***[]string***| ***(Optional)*** |
## OrganizationPolicySpecListPolicyDeny
##### (Appears on:[OrganizationPolicySpecListPolicy](#OrganizationPolicySpecListPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `all` | ***bool***| ***(Optional)*** |
| `values` | ***[]string***| ***(Optional)*** |
## OrganizationPolicySpecRestorePolicy
##### (Appears on:[OrganizationPolicySpec](#OrganizationPolicySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `default` | ***bool***||
## OrganizationPolicyStatus
##### (Appears on:[OrganizationPolicy](#OrganizationPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[OrganizationPolicySpec](#OrganizationPolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
