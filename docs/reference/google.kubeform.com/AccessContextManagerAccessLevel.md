---
title: AccessContextManagerAccessLevel
menu:
  docs_{{ .version }}:
    identifier: accesscontextmanageraccesslevel-google.kubeform.com
    name: AccessContextManagerAccessLevel
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AccessContextManagerAccessLevel
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `AccessContextManagerAccessLevel` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AccessContextManagerAccessLevelSpec](#accesscontextmanageraccesslevelspec)***||
| `status` | ***[AccessContextManagerAccessLevelStatus](#accesscontextmanageraccesslevelstatus)***||
## AccessContextManagerAccessLevelSpec

Appears on:[AccessContextManagerAccessLevel](#accesscontextmanageraccesslevel), [AccessContextManagerAccessLevelStatus](#accesscontextmanageraccesslevelstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `basic` | ***[[]AccessContextManagerAccessLevelSpecBasic](#accesscontextmanageraccesslevelspecbasic)***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `parent` | ***string***||
| `title` | ***string***||
## AccessContextManagerAccessLevelSpecBasic

Appears on:[AccessContextManagerAccessLevelSpec](#accesscontextmanageraccesslevelspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `combiningFunction` | ***string***| ***(Optional)*** |
| `conditions` | ***[[]AccessContextManagerAccessLevelSpecBasicConditions](#accesscontextmanageraccesslevelspecbasicconditions)***||
## AccessContextManagerAccessLevelSpecBasicConditions

Appears on:[AccessContextManagerAccessLevelSpecBasic](#accesscontextmanageraccesslevelspecbasic)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `devicePolicy` | ***[[]AccessContextManagerAccessLevelSpecBasicConditionsDevicePolicy](#accesscontextmanageraccesslevelspecbasicconditionsdevicepolicy)***| ***(Optional)*** |
| `ipSubnetworks` | ***[]string***| ***(Optional)*** |
| `members` | ***[]string***| ***(Optional)*** |
| `negate` | ***bool***| ***(Optional)*** |
| `requiredAccessLevels` | ***[]string***| ***(Optional)*** |
## AccessContextManagerAccessLevelSpecBasicConditionsDevicePolicy

Appears on:[AccessContextManagerAccessLevelSpecBasicConditions](#accesscontextmanageraccesslevelspecbasicconditions)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `allowedDeviceManagementLevels` | ***[]string***| ***(Optional)*** |
| `allowedEncryptionStatuses` | ***[]string***| ***(Optional)*** |
| `osConstraints` | ***[[]AccessContextManagerAccessLevelSpecBasicConditionsDevicePolicyOsConstraints](#accesscontextmanageraccesslevelspecbasicconditionsdevicepolicyosconstraints)***| ***(Optional)*** |
| `requireScreenLock` | ***bool***| ***(Optional)*** |
## AccessContextManagerAccessLevelSpecBasicConditionsDevicePolicyOsConstraints

Appears on:[AccessContextManagerAccessLevelSpecBasicConditionsDevicePolicy](#accesscontextmanageraccesslevelspecbasicconditionsdevicepolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `minimumVersion` | ***string***| ***(Optional)*** |
| `osType` | ***string***| ***(Optional)*** |
## AccessContextManagerAccessLevelStatus

Appears on:[AccessContextManagerAccessLevel](#accesscontextmanageraccesslevel)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AccessContextManagerAccessLevelSpec](#accesscontextmanageraccesslevelspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AccessContextManagerAccessLevelStatus](#accesscontextmanageraccesslevelstatus)

---
