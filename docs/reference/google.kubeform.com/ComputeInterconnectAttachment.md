---
title: ComputeInterconnectAttachment
menu:
  docs_{{ .version }}:
    identifier: computeinterconnectattachment-google.kubeform.com
    name: ComputeInterconnectAttachment
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeInterconnectAttachment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeInterconnectAttachment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeInterconnectAttachmentSpec](#computeinterconnectattachmentspec)***||
| `status` | ***[ComputeInterconnectAttachmentStatus](#computeinterconnectattachmentstatus)***||
## ComputeInterconnectAttachmentSpec

Appears on:[ComputeInterconnectAttachment](#computeinterconnectattachment), [ComputeInterconnectAttachmentStatus](#computeinterconnectattachmentstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `adminEnabled` | ***bool***| ***(Optional)*** |
| `bandwidth` | ***string***| ***(Optional)*** |
| `candidateSubnets` | ***[]string***| ***(Optional)*** |
| `cloudRouterIPAddress` | ***string***| ***(Optional)*** |
| `creationTimestamp` | ***string***| ***(Optional)*** |
| `customerRouterIPAddress` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `edgeAvailabilityDomain` | ***string***| ***(Optional)*** |
| `googleReferenceID` | ***string***| ***(Optional)*** |
| `interconnect` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `pairingKey` | ***string***| ***(Optional)*** |
| `partnerAsn` | ***string***| ***(Optional)*** |
| `privateInterconnectInfo` | ***[[]ComputeInterconnectAttachmentSpecPrivateInterconnectInfo](#computeinterconnectattachmentspecprivateinterconnectinfo)***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `router` | ***string***||
| `selfLink` | ***string***| ***(Optional)*** |
| `state` | ***string***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
| `vlanTag8021q` | ***int64***| ***(Optional)*** |
## ComputeInterconnectAttachmentSpecPrivateInterconnectInfo

Appears on:[ComputeInterconnectAttachmentSpec](#computeinterconnectattachmentspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `tag8021q` | ***int64***| ***(Optional)*** |
## ComputeInterconnectAttachmentStatus

Appears on:[ComputeInterconnectAttachment](#computeinterconnectattachment)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeInterconnectAttachmentSpec](#computeinterconnectattachmentspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeInterconnectAttachmentStatus](#computeinterconnectattachmentstatus)

---
