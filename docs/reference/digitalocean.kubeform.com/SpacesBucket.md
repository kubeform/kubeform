---
title: SpacesBucket
menu:
  docs_{{ .version }}:
    identifier: spacesbucket-digitalocean.kubeform.com
    name: SpacesBucket
    parent: digitalocean.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SpacesBucket
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
|    `kind` | string | `SpacesBucket` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SpacesBucketSpec](#spacesbucketspec)***||
| `status` | ***[SpacesBucketStatus](#spacesbucketstatus)***||
## SpacesBucketSpec

Appears on:[SpacesBucket](#spacesbucket), [SpacesBucketStatus](#spacesbucketstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `acl` | ***string***| ***(Optional)*** Canned ACL applied on bucket creation|
| `bucketDomainName` | ***string***| ***(Optional)*** The FQDN of the bucket|
| `forceDestroy` | ***bool***| ***(Optional)*** Unless true, the bucket will only be destroyed if empty|
| `name` | ***string***|Bucket name|
| `region` | ***string***| ***(Optional)*** Bucket region|
| `urn` | ***string***| ***(Optional)*** the uniform resource name for the bucket|
## SpacesBucketStatus

Appears on:[SpacesBucket](#spacesbucket)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SpacesBucketSpec](#spacesbucketspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
