---
title: ElasticBeanstalkEnvironment
menu:
  docs_{{ .version }}:
    identifier: elasticbeanstalkenvironment-aws.kubeform.com
    name: ElasticBeanstalkEnvironment
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ElasticBeanstalkEnvironment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ElasticBeanstalkEnvironment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ElasticBeanstalkEnvironmentSpec](#elasticbeanstalkenvironmentspec)***||
| `status` | ***[ElasticBeanstalkEnvironmentStatus](#elasticbeanstalkenvironmentstatus)***||
## ElasticBeanstalkEnvironmentSpec

Appears on:[ElasticBeanstalkEnvironment](#elasticbeanstalkenvironment), [ElasticBeanstalkEnvironmentStatus](#elasticbeanstalkenvironmentstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allSettings` | ***[[]ElasticBeanstalkEnvironmentSpecAllSettings](#elasticbeanstalkenvironmentspecallsettings)***| ***(Optional)*** |
| `application` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `autoscalingGroups` | ***[]string***| ***(Optional)*** |
| `cname` | ***string***| ***(Optional)*** |
| `cnamePrefix` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `endpointURL` | ***string***| ***(Optional)*** |
| `instances` | ***[]string***| ***(Optional)*** |
| `launchConfigurations` | ***[]string***| ***(Optional)*** |
| `loadBalancers` | ***[]string***| ***(Optional)*** |
| `name` | ***string***||
| `platformArn` | ***string***| ***(Optional)*** |
| `pollInterval` | ***string***| ***(Optional)*** |
| `queues` | ***[]string***| ***(Optional)*** |
| `setting` | ***[[]ElasticBeanstalkEnvironmentSpecSetting](#elasticbeanstalkenvironmentspecsetting)***| ***(Optional)*** |
| `solutionStackName` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `templateName` | ***string***| ***(Optional)*** |
| `tier` | ***string***| ***(Optional)*** |
| `triggers` | ***[]string***| ***(Optional)*** |
| `versionLabel` | ***string***| ***(Optional)*** |
| `waitForReadyTimeout` | ***string***| ***(Optional)*** |
## ElasticBeanstalkEnvironmentSpecAllSettings

Appears on:[ElasticBeanstalkEnvironmentSpec](#elasticbeanstalkenvironmentspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `namespace` | ***string***||
| `resource` | ***string***| ***(Optional)*** |
| `value` | ***string***||
## ElasticBeanstalkEnvironmentSpecSetting

Appears on:[ElasticBeanstalkEnvironmentSpec](#elasticbeanstalkenvironmentspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `namespace` | ***string***||
| `resource` | ***string***| ***(Optional)*** |
| `value` | ***string***||
## ElasticBeanstalkEnvironmentStatus

Appears on:[ElasticBeanstalkEnvironment](#elasticbeanstalkenvironment)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ElasticBeanstalkEnvironmentSpec](#elasticbeanstalkenvironmentspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ElasticBeanstalkEnvironmentStatus](#elasticbeanstalkenvironmentstatus)

---
