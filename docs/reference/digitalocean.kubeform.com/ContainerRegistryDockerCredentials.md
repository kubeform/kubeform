---
title: ContainerRegistryDockerCredentials
menu:
  docs_{{ .version }}:
    identifier: containerregistrydockercredentials-digitalocean.kubeform.com
    name: ContainerRegistryDockerCredentials
    parent: digitalocean.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ContainerRegistryDockerCredentials
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
|    `kind` | string | `ContainerRegistryDockerCredentials` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ContainerRegistryDockerCredentialsSpec](#containerregistrydockercredentialsspec)***||
| `status` | ***[ContainerRegistryDockerCredentialsStatus](#containerregistrydockercredentialsstatus)***||
## ContainerRegistryDockerCredentialsSpec

Appears on:[ContainerRegistryDockerCredentials](#containerregistrydockercredentials), [ContainerRegistryDockerCredentialsStatus](#containerregistrydockercredentialsstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `credentialExpirationTime` | ***string***| ***(Optional)*** |
| `expirySeconds` | ***int64***| ***(Optional)*** |
| `registryName` | ***string***||
| `write` | ***bool***| ***(Optional)*** |
## ContainerRegistryDockerCredentialsStatus

Appears on:[ContainerRegistryDockerCredentials](#containerregistrydockercredentials)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ContainerRegistryDockerCredentialsSpec](#containerregistrydockercredentialsspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ContainerRegistryDockerCredentialsStatus](#containerregistrydockercredentialsstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `docker_credentials` | ***string*** ||
