---
title: CloudwatchDashboard
menu:
  docs_{{ .version }}:
    identifier: cloudwatchdashboard-aws.kubeform.com
    name: CloudwatchDashboard
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CloudwatchDashboard
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CloudwatchDashboard` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CloudwatchDashboardSpec](#cloudwatchdashboardspec)***||
| `status` | ***[CloudwatchDashboardStatus](#cloudwatchdashboardstatus)***||
## CloudwatchDashboardSpec

Appears on:[CloudwatchDashboard](#cloudwatchdashboard), [CloudwatchDashboardStatus](#cloudwatchdashboardstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `dashboardArn` | ***string***| ***(Optional)*** |
| `dashboardBody` | ***string***||
| `dashboardName` | ***string***||
## CloudwatchDashboardStatus

Appears on:[CloudwatchDashboard](#cloudwatchdashboard)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CloudwatchDashboardSpec](#cloudwatchdashboardspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[CloudwatchDashboardStatus](#cloudwatchdashboardstatus)

---
