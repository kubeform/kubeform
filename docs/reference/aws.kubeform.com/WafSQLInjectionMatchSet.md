---
title: WafSQLInjectionMatchSet
menu:
  docs_{{ .version }}:
    identifier: wafsqlinjectionmatchset-aws.kubeform.com
    name: WafSQLInjectionMatchSet
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## WafSQLInjectionMatchSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafSQLInjectionMatchSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafSQLInjectionMatchSetSpec](#wafsqlinjectionmatchsetspec)***||
| `status` | ***[WafSQLInjectionMatchSetStatus](#wafsqlinjectionmatchsetstatus)***||
## Phase(`string` alias)

Appears on:[WafSQLInjectionMatchSetStatus](#wafsqlinjectionmatchsetstatus)

## WafSQLInjectionMatchSetSpec

Appears on:[WafSQLInjectionMatchSet](#wafsqlinjectionmatchset), [WafSQLInjectionMatchSetStatus](#wafsqlinjectionmatchsetstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `sqlInjectionMatchTuples` | ***[[]WafSQLInjectionMatchSetSpecSqlInjectionMatchTuples](#wafsqlinjectionmatchsetspecsqlinjectionmatchtuples)***| ***(Optional)*** |
## WafSQLInjectionMatchSetSpecSqlInjectionMatchTuples

Appears on:[WafSQLInjectionMatchSetSpec](#wafsqlinjectionmatchsetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `fieldToMatch` | ***[[]WafSQLInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch](#wafsqlinjectionmatchsetspecsqlinjectionmatchtuplesfieldtomatch)***||
| `textTransformation` | ***string***||
## WafSQLInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch

Appears on:[WafSQLInjectionMatchSetSpecSqlInjectionMatchTuples](#wafsqlinjectionmatchsetspecsqlinjectionmatchtuples)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `data` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## WafSQLInjectionMatchSetStatus

Appears on:[WafSQLInjectionMatchSet](#wafsqlinjectionmatchset)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafSQLInjectionMatchSetSpec](#wafsqlinjectionmatchsetspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
