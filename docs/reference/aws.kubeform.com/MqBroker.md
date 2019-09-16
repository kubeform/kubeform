---
title: MqBroker
menu:
  docs_{{ .version }}:
    identifier: mqbroker-aws.kubeform.com
    name: MqBroker
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## MqBroker
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `MqBroker` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MqBrokerSpec](#MqBrokerSpec)***||
| `status` | ***[MqBrokerStatus](#MqBrokerStatus)***||
## MqBrokerSpec

Appears on:[MqBroker](#MqBroker), [MqBrokerStatus](#MqBrokerStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `applyImmediately` | ***bool***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `autoMinorVersionUpgrade` | ***bool***| ***(Optional)*** |
| `brokerName` | ***string***||
| `configuration` | ***[[]MqBrokerSpecConfiguration](#MqBrokerSpecConfiguration)***| ***(Optional)*** |
| `deploymentMode` | ***string***| ***(Optional)*** |
| `engineType` | ***string***||
| `engineVersion` | ***string***||
| `hostInstanceType` | ***string***||
| `instances` | ***[[]MqBrokerSpecInstances](#MqBrokerSpecInstances)***| ***(Optional)*** |
| `logs` | ***[[]MqBrokerSpecLogs](#MqBrokerSpecLogs)***| ***(Optional)*** |
| `maintenanceWindowStartTime` | ***[[]MqBrokerSpecMaintenanceWindowStartTime](#MqBrokerSpecMaintenanceWindowStartTime)***| ***(Optional)*** |
| `publiclyAccessible` | ***bool***| ***(Optional)*** |
| `securityGroups` | ***[]string***||
| `subnetIDS` | ***[]string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `user` | ***[[]MqBrokerSpecUser](#MqBrokerSpecUser)***||
## MqBrokerSpecConfiguration

Appears on:[MqBrokerSpec](#MqBrokerSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `revision` | ***int***| ***(Optional)*** |
## MqBrokerSpecInstances

Appears on:[MqBrokerSpec](#MqBrokerSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `consoleURL` | ***string***| ***(Optional)*** |
| `endpoints` | ***[]string***| ***(Optional)*** |
| `ipAddress` | ***string***| ***(Optional)*** |
## MqBrokerSpecLogs

Appears on:[MqBrokerSpec](#MqBrokerSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `audit` | ***bool***| ***(Optional)*** |
| `general` | ***bool***| ***(Optional)*** |
## MqBrokerSpecMaintenanceWindowStartTime

Appears on:[MqBrokerSpec](#MqBrokerSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `dayOfWeek` | ***string***||
| `timeOfDay` | ***string***||
| `timeZone` | ***string***||
## MqBrokerSpecUser

Appears on:[MqBrokerSpec](#MqBrokerSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `consoleAccess` | ***bool***| ***(Optional)*** |
| `groups` | ***[]string***| ***(Optional)*** |
| `username` | ***string***||
## MqBrokerStatus

Appears on:[MqBroker](#MqBroker)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MqBrokerSpec](#MqBrokerSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `user.<index>.password` | ***string*** ||
