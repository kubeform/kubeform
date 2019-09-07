---
title: ElasticBeanstalkEnvironment
menu:
  docs_v0.0.1:
    identifier: elasticbeanstalkenvironment-aws.kubeform.com
    name: ElasticBeanstalkEnvironment
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ElasticBeanstalkEnvironment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ElasticBeanstalkEnvironment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ElasticBeanstalkEnvironmentSpec](#ElasticBeanstalkEnvironmentSpec)***||
| `status` | ***[ElasticBeanstalkEnvironmentStatus](#ElasticBeanstalkEnvironmentStatus)***||
## ElasticBeanstalkEnvironmentSpec
##### (Appears on:[ElasticBeanstalkEnvironment](#ElasticBeanstalkEnvironment), [ElasticBeanstalkEnvironmentStatus](#ElasticBeanstalkEnvironmentStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allSettings` | ***[[]ElasticBeanstalkEnvironmentSpecAllSettings](#ElasticBeanstalkEnvironmentSpecAllSettings)***| ***(Optional)*** |
| `application` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `autoscalingGroups` | ***[]string***| ***(Optional)*** |
| `cname` | ***string***| ***(Optional)*** |
| `cnamePrefix` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `instances` | ***[]string***| ***(Optional)*** |
| `launchConfigurations` | ***[]string***| ***(Optional)*** |
| `loadBalancers` | ***[]string***| ***(Optional)*** |
| `name` | ***string***||
| `platformArn` | ***string***| ***(Optional)*** |
| `pollInterval` | ***string***| ***(Optional)*** |
| `queues` | ***[]string***| ***(Optional)*** |
| `setting` | ***[[]ElasticBeanstalkEnvironmentSpecSetting](#ElasticBeanstalkEnvironmentSpecSetting)***| ***(Optional)*** |
| `solutionStackName` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `templateName` | ***string***| ***(Optional)*** |
| `tier` | ***string***| ***(Optional)*** |
| `triggers` | ***[]string***| ***(Optional)*** |
| `versionLabel` | ***string***| ***(Optional)*** |
| `waitForReadyTimeout` | ***string***| ***(Optional)*** |
## ElasticBeanstalkEnvironmentSpecAllSettings
##### (Appears on:[ElasticBeanstalkEnvironmentSpec](#ElasticBeanstalkEnvironmentSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `namespace` | ***string***||
| `resource` | ***string***| ***(Optional)*** |
| `value` | ***string***||
## ElasticBeanstalkEnvironmentSpecSetting
##### (Appears on:[ElasticBeanstalkEnvironmentSpec](#ElasticBeanstalkEnvironmentSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `namespace` | ***string***||
| `resource` | ***string***| ***(Optional)*** |
| `value` | ***string***||
## ElasticBeanstalkEnvironmentStatus
##### (Appears on:[ElasticBeanstalkEnvironment](#ElasticBeanstalkEnvironment))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ElasticBeanstalkEnvironmentSpec](#ElasticBeanstalkEnvironmentSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
