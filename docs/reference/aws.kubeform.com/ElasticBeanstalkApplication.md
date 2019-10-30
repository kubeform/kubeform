---
title: ElasticBeanstalkApplication
menu:
  docs_{{ .version }}:
    identifier: elasticbeanstalkapplication-aws.kubeform.com
    name: ElasticBeanstalkApplication
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ElasticBeanstalkApplication
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ElasticBeanstalkApplication` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ElasticBeanstalkApplicationSpec](#elasticbeanstalkapplicationspec)***||
| `status` | ***[ElasticBeanstalkApplicationStatus](#elasticbeanstalkapplicationstatus)***||
## ElasticBeanstalkApplicationSpec

Appears on:[ElasticBeanstalkApplication](#elasticbeanstalkapplication), [ElasticBeanstalkApplicationStatus](#elasticbeanstalkapplicationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `appversionLifecycle` | ***[[]ElasticBeanstalkApplicationSpecAppversionLifecycle](#elasticbeanstalkapplicationspecappversionlifecycle)***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
## ElasticBeanstalkApplicationSpecAppversionLifecycle

Appears on:[ElasticBeanstalkApplicationSpec](#elasticbeanstalkapplicationspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `deleteSourceFromS3` | ***bool***| ***(Optional)*** |
| `maxAgeInDays` | ***int***| ***(Optional)*** |
| `maxCount` | ***int***| ***(Optional)*** |
| `serviceRole` | ***string***||
## ElasticBeanstalkApplicationStatus

Appears on:[ElasticBeanstalkApplication](#elasticbeanstalkapplication)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ElasticBeanstalkApplicationSpec](#elasticbeanstalkapplicationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ElasticBeanstalkApplicationStatus](#elasticbeanstalkapplicationstatus)

---
