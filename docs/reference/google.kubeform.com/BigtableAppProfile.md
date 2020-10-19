---
title: BigtableAppProfile
menu:
  docs_{{ .version }}:
    identifier: bigtableappprofile-google.kubeform.com
    name: BigtableAppProfile
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## BigtableAppProfile
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `BigtableAppProfile` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BigtableAppProfileSpec](#bigtableappprofilespec)***||
| `status` | ***[BigtableAppProfileStatus](#bigtableappprofilestatus)***||
## BigtableAppProfileSpec

Appears on:[BigtableAppProfile](#bigtableappprofile), [BigtableAppProfileStatus](#bigtableappprofilestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `appProfileID` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `ignoreWarnings` | ***bool***| ***(Optional)*** |
| `instance` | ***string***| ***(Optional)*** |
| `multiClusterRoutingUseAny` | ***bool***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `singleClusterRouting` | ***[[]BigtableAppProfileSpecSingleClusterRouting](#bigtableappprofilespecsingleclusterrouting)***| ***(Optional)*** |
## BigtableAppProfileSpecSingleClusterRouting

Appears on:[BigtableAppProfileSpec](#bigtableappprofilespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `allowTransactionalWrites` | ***bool***| ***(Optional)*** |
| `clusterID` | ***string***| ***(Optional)*** |
## BigtableAppProfileStatus

Appears on:[BigtableAppProfile](#bigtableappprofile)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BigtableAppProfileSpec](#bigtableappprofilespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[BigtableAppProfileStatus](#bigtableappprofilestatus)

---
