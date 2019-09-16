---
title: OpsworksApplication
menu:
  docs_{{ .version }}:
    identifier: opsworksapplication-aws.kubeform.com
    name: OpsworksApplication
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## OpsworksApplication
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `OpsworksApplication` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[OpsworksApplicationSpec](#opsworksapplicationspec)***||
| `status` | ***[OpsworksApplicationStatus](#opsworksapplicationstatus)***||
## OpsworksApplicationSpec

Appears on:[OpsworksApplication](#opsworksapplication), [OpsworksApplicationStatus](#opsworksapplicationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `appSource` | ***[[]OpsworksApplicationSpecAppSource](#opsworksapplicationspecappsource)***| ***(Optional)*** |
| `autoBundleOnDeploy` | ***string***| ***(Optional)*** |
| `awsFlowRubySettings` | ***string***| ***(Optional)*** |
| `dataSourceArn` | ***string***| ***(Optional)*** |
| `dataSourceDatabaseName` | ***string***| ***(Optional)*** |
| `dataSourceType` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `documentRoot` | ***string***| ***(Optional)*** |
| `domains` | ***[]string***| ***(Optional)*** |
| `enableSSL` | ***bool***| ***(Optional)*** |
| `environment` | ***[[]OpsworksApplicationSpecEnvironment](#opsworksapplicationspecenvironment)***| ***(Optional)*** |
| `name` | ***string***||
| `railsEnv` | ***string***| ***(Optional)*** |
| `shortName` | ***string***| ***(Optional)*** |
| `sslConfiguration` | ***[[]OpsworksApplicationSpecSslConfiguration](#opsworksapplicationspecsslconfiguration)***| ***(Optional)*** |
| `stackID` | ***string***||
| `type` | ***string***||
## OpsworksApplicationSpecAppSource

Appears on:[OpsworksApplicationSpec](#opsworksapplicationspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `revision` | ***string***| ***(Optional)*** |
| `sshKey` | ***string***| ***(Optional)*** |
| `type` | ***string***||
| `url` | ***string***| ***(Optional)*** |
| `username` | ***string***| ***(Optional)*** |
## OpsworksApplicationSpecEnvironment

Appears on:[OpsworksApplicationSpec](#opsworksapplicationspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `key` | ***string***||
| `secure` | ***bool***| ***(Optional)*** |
| `value` | ***string***||
## OpsworksApplicationSpecSslConfiguration

Appears on:[OpsworksApplicationSpec](#opsworksapplicationspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `certificate` | ***string***||
| `chain` | ***string***| ***(Optional)*** |
## OpsworksApplicationStatus

Appears on:[OpsworksApplication](#opsworksapplication)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[OpsworksApplicationSpec](#opsworksapplicationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `app_source.<index>.password` | ***string*** ||
| `ssl_configuration.<index>.private_key` | ***string*** ||
