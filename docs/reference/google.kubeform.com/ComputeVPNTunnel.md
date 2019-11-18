---
title: ComputeVPNTunnel
menu:
  docs_{{ .version }}:
    identifier: computevpntunnel-google.kubeform.com
    name: ComputeVPNTunnel
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeVPNTunnel
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeVPNTunnel` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeVPNTunnelSpec](#computevpntunnelspec)***||
| `status` | ***[ComputeVPNTunnelStatus](#computevpntunnelstatus)***||
## ComputeVPNTunnelSpec

Appears on:[ComputeVPNTunnel](#computevpntunnel), [ComputeVPNTunnelStatus](#computevpntunnelstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `creationTimestamp` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `detailedStatus` | ***string***| ***(Optional)*** |
| `ikeVersion` | ***int64***| ***(Optional)*** |
| `labelFingerprint` | ***string***| ***(Optional)*** |
| `labels` | ***map[string]string***| ***(Optional)*** |
| `localTrafficSelector` | ***[]string***| ***(Optional)*** |
| `name` | ***string***||
| `peerIP` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `remoteTrafficSelector` | ***[]string***| ***(Optional)*** |
| `router` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `sharedSecretHash` | ***string***| ***(Optional)*** |
| `targetVPNGateway` | ***string***||
## ComputeVPNTunnelStatus

Appears on:[ComputeVPNTunnel](#computevpntunnel)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeVPNTunnelSpec](#computevpntunnelspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeVPNTunnelStatus](#computevpntunnelstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `shared_secret` | ***string*** ||
