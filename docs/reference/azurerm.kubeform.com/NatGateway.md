---
title: NatGateway
menu:
  docs_{{ .version }}:
    identifier: natgateway-azurerm.kubeform.com
    name: NatGateway
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## NatGateway
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `NatGateway` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NatGatewaySpec](#natgatewayspec)***||
| `status` | ***[NatGatewayStatus](#natgatewaystatus)***||
## NatGatewaySpec

Appears on:[NatGateway](#natgateway), [NatGatewayStatus](#natgatewaystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `idleTimeoutInMinutes` | ***int64***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `publicIPAddressIDS` | ***[]string***| ***(Optional)*** |
| `publicIPPrefixIDS` | ***[]string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `resourceGuid` | ***string***| ***(Optional)*** |
| `skuName` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `zones` | ***[]string***| ***(Optional)*** |
## NatGatewayStatus

Appears on:[NatGateway](#natgateway)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NatGatewaySpec](#natgatewayspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[NatGatewayStatus](#natgatewaystatus)

---
