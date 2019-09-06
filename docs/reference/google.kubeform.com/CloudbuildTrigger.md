## CloudbuildTrigger
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `CloudbuildTrigger` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CloudbuildTriggerSpec](#CloudbuildTriggerSpec)***||
| `status` | ***[CloudbuildTriggerStatus](#CloudbuildTriggerStatus)***||
## CloudbuildTriggerSpec
##### (Appears on:[CloudbuildTrigger](#CloudbuildTrigger), [CloudbuildTriggerStatus](#CloudbuildTriggerStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `build` | ***[[]CloudbuildTriggerSpecBuild](#CloudbuildTriggerSpecBuild)***| ***(Optional)*** Contents of the build template.|
| `description` | ***string***| ***(Optional)*** |
| `filename` | ***string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `substitutions` | ***map[string]string***| ***(Optional)*** |
| `triggerTemplate` | ***[[]CloudbuildTriggerSpecTriggerTemplate](#CloudbuildTriggerSpecTriggerTemplate)***| ***(Optional)*** |
## CloudbuildTriggerSpecBuild
##### (Appears on:[CloudbuildTriggerSpec](#CloudbuildTriggerSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `images` | ***[]string***| ***(Optional)*** |
| `step` | ***[[]CloudbuildTriggerSpecBuildStep](#CloudbuildTriggerSpecBuildStep)***| ***(Optional)*** |
| `tags` | ***[]string***| ***(Optional)*** |
## CloudbuildTriggerSpecBuildStep
##### (Appears on:[CloudbuildTriggerSpecBuild](#CloudbuildTriggerSpecBuild))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `args` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
## CloudbuildTriggerSpecTriggerTemplate
##### (Appears on:[CloudbuildTriggerSpec](#CloudbuildTriggerSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `branchName` | ***string***| ***(Optional)*** |
| `commitSha` | ***string***| ***(Optional)*** |
| `dir` | ***string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `repoName` | ***string***| ***(Optional)*** |
| `tagName` | ***string***| ***(Optional)*** |
## CloudbuildTriggerStatus
##### (Appears on:[CloudbuildTrigger](#CloudbuildTrigger))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CloudbuildTriggerSpec](#CloudbuildTriggerSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
