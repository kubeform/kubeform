---
title: ElasticBeanstalkConfigurationTemplate
menu:
  docs_v0.0.1:
    identifier: elasticbeanstalkconfigurationtemplate-aws.kubeform.com
    name: ElasticBeanstalkConfigurationTemplate
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ElasticBeanstalkConfigurationTemplate
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ElasticBeanstalkConfigurationTemplate` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ElasticBeanstalkConfigurationTemplateSpec](#ElasticBeanstalkConfigurationTemplateSpec)***||
| `status` | ***[ElasticBeanstalkConfigurationTemplateStatus](#ElasticBeanstalkConfigurationTemplateStatus)***||
## ElasticBeanstalkConfigurationTemplateSpec
##### (Appears on:[ElasticBeanstalkConfigurationTemplate](#ElasticBeanstalkConfigurationTemplate), [ElasticBeanstalkConfigurationTemplateStatus](#ElasticBeanstalkConfigurationTemplateStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `application` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `environmentID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `setting` | ***[[]ElasticBeanstalkConfigurationTemplateSpecSetting](#ElasticBeanstalkConfigurationTemplateSpecSetting)***| ***(Optional)*** |
| `solutionStackName` | ***string***| ***(Optional)*** |
## ElasticBeanstalkConfigurationTemplateSpecSetting
##### (Appears on:[ElasticBeanstalkConfigurationTemplateSpec](#ElasticBeanstalkConfigurationTemplateSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `namespace` | ***string***||
| `resource` | ***string***| ***(Optional)*** |
| `value` | ***string***||
## ElasticBeanstalkConfigurationTemplateStatus
##### (Appears on:[ElasticBeanstalkConfigurationTemplate](#ElasticBeanstalkConfigurationTemplate))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ElasticBeanstalkConfigurationTemplateSpec](#ElasticBeanstalkConfigurationTemplateSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
