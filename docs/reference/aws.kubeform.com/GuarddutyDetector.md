---
title: GuarddutyDetector
menu:
  docs_{{ .version }}:
    identifier: guarddutydetector-aws.kubeform.com
    name: GuarddutyDetector
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## GuarddutyDetector
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `GuarddutyDetector` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[GuarddutyDetectorSpec](#guarddutydetectorspec)***||
| `status` | ***[GuarddutyDetectorStatus](#guarddutydetectorstatus)***||
## GuarddutyDetectorSpec

Appears on:[GuarddutyDetector](#guarddutydetector), [GuarddutyDetectorStatus](#guarddutydetectorstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountID` | ***string***| ***(Optional)*** |
| `enable` | ***bool***| ***(Optional)*** |
| `findingPublishingFrequency` | ***string***| ***(Optional)*** |
## GuarddutyDetectorStatus

Appears on:[GuarddutyDetector](#guarddutydetector)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[GuarddutyDetectorSpec](#guarddutydetectorspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[GuarddutyDetectorStatus](#guarddutydetectorstatus)

---
