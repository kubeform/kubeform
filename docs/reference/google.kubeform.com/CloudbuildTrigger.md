---
title: CloudbuildTrigger
menu:
  docs_{{ .version }}:
    identifier: cloudbuildtrigger-google.kubeform.com
    name: CloudbuildTrigger
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CloudbuildTrigger
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `CloudbuildTrigger` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CloudbuildTriggerSpec](#cloudbuildtriggerspec)***||
| `status` | ***[CloudbuildTriggerStatus](#cloudbuildtriggerstatus)***||
## CloudbuildTriggerSpec

Appears on:[CloudbuildTrigger](#cloudbuildtrigger), [CloudbuildTriggerStatus](#cloudbuildtriggerstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `build` | ***[[]CloudbuildTriggerSpecBuild](#cloudbuildtriggerspecbuild)***| ***(Optional)*** |
| `createTime` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `disabled` | ***bool***| ***(Optional)*** |
| `filename` | ***string***| ***(Optional)*** |
| `ignoredFiles` | ***[]string***| ***(Optional)*** |
| `includedFiles` | ***[]string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `substitutions` | ***map[string]string***| ***(Optional)*** |
| `triggerID` | ***string***| ***(Optional)*** |
| `triggerTemplate` | ***[[]CloudbuildTriggerSpecTriggerTemplate](#cloudbuildtriggerspectriggertemplate)***| ***(Optional)*** |
## CloudbuildTriggerSpecBuild

Appears on:[CloudbuildTriggerSpec](#cloudbuildtriggerspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `images` | ***[]string***| ***(Optional)*** |
| `step` | ***[[]CloudbuildTriggerSpecBuildStep](#cloudbuildtriggerspecbuildstep)***| ***(Optional)*** |
| `tags` | ***[]string***| ***(Optional)*** |
## CloudbuildTriggerSpecBuildStep

Appears on:[CloudbuildTriggerSpecBuild](#cloudbuildtriggerspecbuild)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `args` | ***[]string***| ***(Optional)*** |
| `dir` | ***string***| ***(Optional)*** |
| `entrypoint` | ***string***| ***(Optional)*** |
| `env` | ***[]string***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `secretEnv` | ***[]string***| ***(Optional)*** |
| `timeout` | ***string***| ***(Optional)*** |
| `timing` | ***string***| ***(Optional)*** |
| `volumes` | ***[[]CloudbuildTriggerSpecBuildStepVolumes](#cloudbuildtriggerspecbuildstepvolumes)***| ***(Optional)*** |
| `waitFor` | ***[]string***| ***(Optional)*** |
## CloudbuildTriggerSpecBuildStepVolumes

Appears on:[CloudbuildTriggerSpecBuildStep](#cloudbuildtriggerspecbuildstep)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***| ***(Optional)*** |
| `path` | ***string***| ***(Optional)*** |
## CloudbuildTriggerSpecTriggerTemplate

Appears on:[CloudbuildTriggerSpec](#cloudbuildtriggerspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `branchName` | ***string***| ***(Optional)*** |
| `commitSha` | ***string***| ***(Optional)*** |
| `dir` | ***string***| ***(Optional)*** |
| `projectID` | ***string***| ***(Optional)*** |
| `repoName` | ***string***| ***(Optional)*** |
| `tagName` | ***string***| ***(Optional)*** |
## CloudbuildTriggerStatus

Appears on:[CloudbuildTrigger](#cloudbuildtrigger)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CloudbuildTriggerSpec](#cloudbuildtriggerspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[CloudbuildTriggerStatus](#cloudbuildtriggerstatus)

---
