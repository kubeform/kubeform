---
title: ApiGatewayVpcLink
menu:
  docs_{{ .version }}:
    identifier: apigatewayvpclink-aws.kubeform.com
    name: ApiGatewayVpcLink
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApiGatewayVpcLink
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiGatewayVpcLink` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiGatewayVpcLinkSpec](#apigatewayvpclinkspec)***||
| `status` | ***[ApiGatewayVpcLinkStatus](#apigatewayvpclinkstatus)***||
## ApiGatewayVpcLinkSpec

Appears on:[ApiGatewayVpcLink](#apigatewayvpclink), [ApiGatewayVpcLinkStatus](#apigatewayvpclinkstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `targetArns` | ***[]string***||
## ApiGatewayVpcLinkStatus

Appears on:[ApiGatewayVpcLink](#apigatewayvpclink)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiGatewayVpcLinkSpec](#apigatewayvpclinkspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ApiGatewayVpcLinkStatus](#apigatewayvpclinkstatus)

---
