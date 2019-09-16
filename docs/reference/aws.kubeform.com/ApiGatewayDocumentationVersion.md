---
title: ApiGatewayDocumentationVersion
menu:
  docs_{{ .version }}:
    identifier: apigatewaydocumentationversion-aws.kubeform.com
    name: ApiGatewayDocumentationVersion
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApiGatewayDocumentationVersion
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiGatewayDocumentationVersion` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiGatewayDocumentationVersionSpec](#apigatewaydocumentationversionspec)***||
| `status` | ***[ApiGatewayDocumentationVersionStatus](#apigatewaydocumentationversionstatus)***||
## ApiGatewayDocumentationVersionSpec

Appears on:[ApiGatewayDocumentationVersion](#apigatewaydocumentationversion), [ApiGatewayDocumentationVersionStatus](#apigatewaydocumentationversionstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `restAPIID` | ***string***||
| `version` | ***string***||
## ApiGatewayDocumentationVersionStatus

Appears on:[ApiGatewayDocumentationVersion](#apigatewaydocumentationversion)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiGatewayDocumentationVersionSpec](#apigatewaydocumentationversionspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
