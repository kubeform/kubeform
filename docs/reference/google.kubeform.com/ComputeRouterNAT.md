---
title: ComputeRouterNAT
menu:
  docs_{{ .version }}:
    identifier: computerouternat-google.kubeform.com
    name: ComputeRouterNAT
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeRouterNAT
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeRouterNAT` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeRouterNATSpec](#computerouternatspec)***||
| `status` | ***[ComputeRouterNATStatus](#computerouternatstatus)***||
## ComputeRouterNATSpec

Appears on:[ComputeRouterNAT](#computerouternat), [ComputeRouterNATStatus](#computerouternatstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `icmpIdleTimeoutSec` | ***int64***| ***(Optional)*** |
| `logConfig` | ***[[]ComputeRouterNATSpecLogConfig](#computerouternatspeclogconfig)***| ***(Optional)*** |
| `minPortsPerVm` | ***int64***| ***(Optional)*** |
| `name` | ***string***||
| `natIPAllocateOption` | ***string***||
| `natIPS` | ***[]string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `router` | ***string***||
| `sourceSubnetworkIPRangesToNAT` | ***string***||
| `subnetwork` | ***[[]ComputeRouterNATSpecSubnetwork](#computerouternatspecsubnetwork)***| ***(Optional)*** |
| `tcpEstablishedIdleTimeoutSec` | ***int64***| ***(Optional)*** |
| `tcpTransitoryIdleTimeoutSec` | ***int64***| ***(Optional)*** |
| `udpIdleTimeoutSec` | ***int64***| ***(Optional)*** |
## ComputeRouterNATSpecLogConfig

Appears on:[ComputeRouterNATSpec](#computerouternatspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enable` | ***bool***||
| `filter` | ***string***||
## ComputeRouterNATSpecSubnetwork

Appears on:[ComputeRouterNATSpec](#computerouternatspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `secondaryIPRangeNames` | ***[]string***| ***(Optional)*** |
| `sourceIPRangesToNAT` | ***[]string***||
## ComputeRouterNATStatus

Appears on:[ComputeRouterNAT](#computerouternat)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeRouterNATSpec](#computerouternatspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeRouterNATStatus](#computerouternatstatus)

---
