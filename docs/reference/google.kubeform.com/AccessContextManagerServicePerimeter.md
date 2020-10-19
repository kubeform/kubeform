---
title: AccessContextManagerServicePerimeter
menu:
  docs_{{ .version }}:
    identifier: accesscontextmanagerserviceperimeter-google.kubeform.com
    name: AccessContextManagerServicePerimeter
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AccessContextManagerServicePerimeter
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `AccessContextManagerServicePerimeter` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AccessContextManagerServicePerimeterSpec](#accesscontextmanagerserviceperimeterspec)***||
| `status` | ***[AccessContextManagerServicePerimeterStatus](#accesscontextmanagerserviceperimeterstatus)***||
## AccessContextManagerServicePerimeterSpec

Appears on:[AccessContextManagerServicePerimeter](#accesscontextmanagerserviceperimeter), [AccessContextManagerServicePerimeterStatus](#accesscontextmanagerserviceperimeterstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `createTime` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `parent` | ***string***||
| `perimeterType` | ***string***| ***(Optional)*** |
| `status` | ***[[]AccessContextManagerServicePerimeterSpecStatus](#accesscontextmanagerserviceperimeterspecstatus)***| ***(Optional)*** |
| `title` | ***string***||
| `updateTime` | ***string***| ***(Optional)*** |
## AccessContextManagerServicePerimeterSpecStatus

Appears on:[AccessContextManagerServicePerimeterSpec](#accesscontextmanagerserviceperimeterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `accessLevels` | ***[]string***| ***(Optional)*** |
| `resources` | ***[]string***| ***(Optional)*** |
| `restrictedServices` | ***[]string***| ***(Optional)*** |
## AccessContextManagerServicePerimeterStatus

Appears on:[AccessContextManagerServicePerimeter](#accesscontextmanagerserviceperimeter)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AccessContextManagerServicePerimeterSpec](#accesscontextmanagerserviceperimeterspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AccessContextManagerServicePerimeterStatus](#accesscontextmanagerserviceperimeterstatus)

---
