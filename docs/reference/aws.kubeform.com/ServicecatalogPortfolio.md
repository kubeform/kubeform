---
title: ServicecatalogPortfolio
menu:
  docs_{{ .version }}:
    identifier: servicecatalogportfolio-aws.kubeform.com
    name: ServicecatalogPortfolio
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ServicecatalogPortfolio
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ServicecatalogPortfolio` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ServicecatalogPortfolioSpec](#servicecatalogportfoliospec)***||
| `status` | ***[ServicecatalogPortfolioStatus](#servicecatalogportfoliostatus)***||
## Phase(`string` alias)

Appears on:[ServicecatalogPortfolioStatus](#servicecatalogportfoliostatus)

## ServicecatalogPortfolioSpec

Appears on:[ServicecatalogPortfolio](#servicecatalogportfolio), [ServicecatalogPortfolioStatus](#servicecatalogportfoliostatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `createdTime` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `providerName` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## ServicecatalogPortfolioStatus

Appears on:[ServicecatalogPortfolio](#servicecatalogportfolio)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ServicecatalogPortfolioSpec](#servicecatalogportfoliospec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
