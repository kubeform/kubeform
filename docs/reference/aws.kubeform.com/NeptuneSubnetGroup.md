---
title: NeptuneSubnetGroup
menu:
  docs_{{ .version }}:
    identifier: neptunesubnetgroup-aws.kubeform.com
    name: NeptuneSubnetGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## NeptuneSubnetGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `NeptuneSubnetGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NeptuneSubnetGroupSpec](#neptunesubnetgroupspec)***||
| `status` | ***[NeptuneSubnetGroupStatus](#neptunesubnetgroupstatus)***||
## NeptuneSubnetGroupSpec

Appears on:[NeptuneSubnetGroup](#neptunesubnetgroup), [NeptuneSubnetGroupStatus](#neptunesubnetgroupstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `subnetIDS` | ***[]string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## NeptuneSubnetGroupStatus

Appears on:[NeptuneSubnetGroup](#neptunesubnetgroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NeptuneSubnetGroupSpec](#neptunesubnetgroupspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[NeptuneSubnetGroupStatus](#neptunesubnetgroupstatus)

---
