---
title: RedisInstance
menu:
  docs_{{ .version }}:
    identifier: redisinstance-google.kubeform.com
    name: RedisInstance
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## RedisInstance
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `RedisInstance` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RedisInstanceSpec](#redisinstancespec)***||
| `status` | ***[RedisInstanceStatus](#redisinstancestatus)***||
## Phase(`string` alias)

Appears on:[RedisInstanceStatus](#redisinstancestatus)

## RedisInstanceSpec

Appears on:[RedisInstance](#redisinstance), [RedisInstanceStatus](#redisinstancestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `alternativeLocationID` | ***string***| ***(Optional)*** |
| `authorizedNetwork` | ***string***| ***(Optional)*** |
| `createTime` | ***string***| ***(Optional)*** |
| `currentLocationID` | ***string***| ***(Optional)*** |
| `displayName` | ***string***| ***(Optional)*** |
| `host` | ***string***| ***(Optional)*** |
| `labels` | ***map[string]string***| ***(Optional)*** |
| `locationID` | ***string***| ***(Optional)*** |
| `memorySizeGb` | ***int***||
| `name` | ***string***||
| `port` | ***int***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `redisConfigs` | ***map[string]string***| ***(Optional)*** |
| `redisVersion` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `reservedIPRange` | ***string***| ***(Optional)*** |
| `tier` | ***string***| ***(Optional)*** |
## RedisInstanceStatus

Appears on:[RedisInstance](#redisinstance)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RedisInstanceSpec](#redisinstancespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
