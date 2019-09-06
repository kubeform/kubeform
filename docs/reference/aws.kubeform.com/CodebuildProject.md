## CodebuildProject
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CodebuildProject` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CodebuildProjectSpec](#CodebuildProjectSpec)***||
| `status` | ***[CodebuildProjectStatus](#CodebuildProjectStatus)***||
## CodebuildProjectSpec
##### (Appears on:[CodebuildProject](#CodebuildProject), [CodebuildProjectStatus](#CodebuildProjectStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `arn` | ***string***| ***(Optional)*** |
| `artifacts` | ***[[]CodebuildProjectSpecArtifacts](#CodebuildProjectSpecArtifacts)***||
| `badgeEnabled` | ***bool***| ***(Optional)*** |
| `badgeURL` | ***string***| ***(Optional)*** |
| `buildTimeout` | ***int***| ***(Optional)*** |
| `cache` | ***[[]CodebuildProjectSpecCache](#CodebuildProjectSpecCache)***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `encryptionKey` | ***string***| ***(Optional)*** |
| `environment` | ***[[]CodebuildProjectSpecEnvironment](#CodebuildProjectSpecEnvironment)***||
| `name` | ***string***||
| `secondaryArtifacts` | ***[[]CodebuildProjectSpecSecondaryArtifacts](#CodebuildProjectSpecSecondaryArtifacts)***| ***(Optional)*** |
| `secondarySources` | ***[[]CodebuildProjectSpecSecondarySources](#CodebuildProjectSpecSecondarySources)***| ***(Optional)*** |
| `serviceRole` | ***string***||
| `source` | ***[[]CodebuildProjectSpecSource](#CodebuildProjectSpecSource)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpcConfig` | ***[[]CodebuildProjectSpecVpcConfig](#CodebuildProjectSpecVpcConfig)***| ***(Optional)*** |
## CodebuildProjectSpecArtifacts
##### (Appears on:[CodebuildProjectSpec](#CodebuildProjectSpec))
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
##### (Appears on:[CodebuildProjectSpec](#CodebuildProjectSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `location` | ***string***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
## CodebuildProjectSpecEnvironment
##### (Appears on:[CodebuildProjectSpec](#CodebuildProjectSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `certificate` | ***string***| ***(Optional)*** |
| `computeType` | ***string***||
| `environmentVariable` | ***[[]CodebuildProjectSpecEnvironmentEnvironmentVariable](#CodebuildProjectSpecEnvironmentEnvironmentVariable)***| ***(Optional)*** |
| `image` | ***string***||
| `imagePullCredentialsType` | ***string***| ***(Optional)*** |
| `privilegedMode` | ***bool***| ***(Optional)*** |
| `type` | ***string***||
## CodebuildProjectSpecEnvironmentEnvironmentVariable
##### (Appears on:[CodebuildProjectSpecEnvironment](#CodebuildProjectSpecEnvironment))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `type` | ***string***| ***(Optional)*** |
| `value` | ***string***||
## CodebuildProjectSpecSecondaryArtifacts
##### (Appears on:[CodebuildProjectSpec](#CodebuildProjectSpec))
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
##### (Appears on:[CodebuildProjectSpec](#CodebuildProjectSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `auth` | ***[[]CodebuildProjectSpecSecondarySourcesAuth](#CodebuildProjectSpecSecondarySourcesAuth)***| ***(Optional)*** |
| `buildspec` | ***string***| ***(Optional)*** |
| `gitCloneDepth` | ***int***| ***(Optional)*** |
| `insecureSSL` | ***bool***| ***(Optional)*** |
| `location` | ***string***| ***(Optional)*** |
| `reportBuildStatus` | ***bool***| ***(Optional)*** |
| `sourceIdentifier` | ***string***||
| `type` | ***string***||
## CodebuildProjectSpecSecondarySourcesAuth
##### (Appears on:[CodebuildProjectSpecSecondarySources](#CodebuildProjectSpecSecondarySources))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
## CodebuildProjectSpecSource
##### (Appears on:[CodebuildProjectSpec](#CodebuildProjectSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `auth` | ***[[]CodebuildProjectSpecSourceAuth](#CodebuildProjectSpecSourceAuth)***| ***(Optional)*** |
| `buildspec` | ***string***| ***(Optional)*** |
| `gitCloneDepth` | ***int***| ***(Optional)*** |
| `insecureSSL` | ***bool***| ***(Optional)*** |
| `location` | ***string***| ***(Optional)*** |
| `reportBuildStatus` | ***bool***| ***(Optional)*** |
| `type` | ***string***||
## CodebuildProjectSpecSourceAuth
##### (Appears on:[CodebuildProjectSpecSource](#CodebuildProjectSpecSource))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
## CodebuildProjectSpecVpcConfig
##### (Appears on:[CodebuildProjectSpec](#CodebuildProjectSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `securityGroupIDS` | ***[]string***||
| `subnets` | ***[]string***||
| `vpcID` | ***string***||
## CodebuildProjectStatus
##### (Appears on:[CodebuildProject](#CodebuildProject))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CodebuildProjectSpec](#CodebuildProjectSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `secondary_sources.<index>.auth.<index>.resource` | ***string*** ||
| `source.<index>.auth.<index>.resource` | ***string*** ||
