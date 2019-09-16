---
title: CodebuildProject
menu:
  docs_{{ .version }}:
    identifier: codebuildproject-aws.kubeform.com
    name: CodebuildProject
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CodebuildProject
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CodebuildProject` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CodebuildProjectSpec](#codebuildprojectspec)***||
| `status` | ***[CodebuildProjectStatus](#codebuildprojectstatus)***||
## CodebuildProjectSpec

Appears on:[CodebuildProject](#codebuildproject), [CodebuildProjectStatus](#codebuildprojectstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `arn` | ***string***| ***(Optional)*** |
| `artifacts` | ***[[]CodebuildProjectSpecArtifacts](#codebuildprojectspecartifacts)***||
| `badgeEnabled` | ***bool***| ***(Optional)*** |
| `badgeURL` | ***string***| ***(Optional)*** |
| `buildTimeout` | ***int***| ***(Optional)*** |
| `cache` | ***[[]CodebuildProjectSpecCache](#codebuildprojectspeccache)***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `encryptionKey` | ***string***| ***(Optional)*** |
| `environment` | ***[[]CodebuildProjectSpecEnvironment](#codebuildprojectspecenvironment)***||
| `name` | ***string***||
| `secondaryArtifacts` | ***[[]CodebuildProjectSpecSecondaryArtifacts](#codebuildprojectspecsecondaryartifacts)***| ***(Optional)*** |
| `secondarySources` | ***[[]CodebuildProjectSpecSecondarySources](#codebuildprojectspecsecondarysources)***| ***(Optional)*** |
| `serviceRole` | ***string***||
| `source` | ***[[]CodebuildProjectSpecSource](#codebuildprojectspecsource)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpcConfig` | ***[[]CodebuildProjectSpecVpcConfig](#codebuildprojectspecvpcconfig)***| ***(Optional)*** |
## CodebuildProjectSpecArtifacts

Appears on:[CodebuildProjectSpec](#codebuildprojectspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `encryptionDisabled` | ***bool***| ***(Optional)*** |
| `location` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namespaceType` | ***string***| ***(Optional)*** |
| `packaging` | ***string***| ***(Optional)*** |
| `path` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## CodebuildProjectSpecCache

Appears on:[CodebuildProjectSpec](#codebuildprojectspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `location` | ***string***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
## CodebuildProjectSpecEnvironment

Appears on:[CodebuildProjectSpec](#codebuildprojectspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `certificate` | ***string***| ***(Optional)*** |
| `computeType` | ***string***||
| `environmentVariable` | ***[[]CodebuildProjectSpecEnvironmentEnvironmentVariable](#codebuildprojectspecenvironmentenvironmentvariable)***| ***(Optional)*** |
| `image` | ***string***||
| `imagePullCredentialsType` | ***string***| ***(Optional)*** |
| `privilegedMode` | ***bool***| ***(Optional)*** |
| `type` | ***string***||
## CodebuildProjectSpecEnvironmentEnvironmentVariable

Appears on:[CodebuildProjectSpecEnvironment](#codebuildprojectspecenvironment)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `type` | ***string***| ***(Optional)*** |
| `value` | ***string***||
## CodebuildProjectSpecSecondaryArtifacts

Appears on:[CodebuildProjectSpec](#codebuildprojectspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `artifactIdentifier` | ***string***||
| `encryptionDisabled` | ***bool***| ***(Optional)*** |
| `location` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namespaceType` | ***string***| ***(Optional)*** |
| `packaging` | ***string***| ***(Optional)*** |
| `path` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## CodebuildProjectSpecSecondarySources

Appears on:[CodebuildProjectSpec](#codebuildprojectspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `auth` | ***[[]CodebuildProjectSpecSecondarySourcesAuth](#codebuildprojectspecsecondarysourcesauth)***| ***(Optional)*** |
| `buildspec` | ***string***| ***(Optional)*** |
| `gitCloneDepth` | ***int***| ***(Optional)*** |
| `insecureSSL` | ***bool***| ***(Optional)*** |
| `location` | ***string***| ***(Optional)*** |
| `reportBuildStatus` | ***bool***| ***(Optional)*** |
| `sourceIdentifier` | ***string***||
| `type` | ***string***||
## CodebuildProjectSpecSecondarySourcesAuth

Appears on:[CodebuildProjectSpecSecondarySources](#codebuildprojectspecsecondarysources)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
## CodebuildProjectSpecSource

Appears on:[CodebuildProjectSpec](#codebuildprojectspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `auth` | ***[[]CodebuildProjectSpecSourceAuth](#codebuildprojectspecsourceauth)***| ***(Optional)*** |
| `buildspec` | ***string***| ***(Optional)*** |
| `gitCloneDepth` | ***int***| ***(Optional)*** |
| `insecureSSL` | ***bool***| ***(Optional)*** |
| `location` | ***string***| ***(Optional)*** |
| `reportBuildStatus` | ***bool***| ***(Optional)*** |
| `type` | ***string***||
## CodebuildProjectSpecSourceAuth

Appears on:[CodebuildProjectSpecSource](#codebuildprojectspecsource)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
## CodebuildProjectSpecVpcConfig

Appears on:[CodebuildProjectSpec](#codebuildprojectspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `securityGroupIDS` | ***[]string***||
| `subnets` | ***[]string***||
| `vpcID` | ***string***||
## CodebuildProjectStatus

Appears on:[CodebuildProject](#codebuildproject)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CodebuildProjectSpec](#codebuildprojectspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `secondary_sources.<index>.auth.<index>.resource` | ***string*** ||
| `source.<index>.auth.<index>.resource` | ***string*** ||
