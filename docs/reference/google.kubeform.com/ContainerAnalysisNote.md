## ContainerAnalysisNote
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ContainerAnalysisNote` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ContainerAnalysisNoteSpec](#ContainerAnalysisNoteSpec)***||
| `status` | ***[ContainerAnalysisNoteStatus](#ContainerAnalysisNoteStatus)***||
## ContainerAnalysisNoteSpec
##### (Appears on:[ContainerAnalysisNote](#ContainerAnalysisNote), [ContainerAnalysisNoteStatus](#ContainerAnalysisNoteStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `attestationAuthority` | ***[[]ContainerAnalysisNoteSpecAttestationAuthority](#ContainerAnalysisNoteSpecAttestationAuthority)***||
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
## ContainerAnalysisNoteSpecAttestationAuthority
##### (Appears on:[ContainerAnalysisNoteSpec](#ContainerAnalysisNoteSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `hint` | ***[[]ContainerAnalysisNoteSpecAttestationAuthorityHint](#ContainerAnalysisNoteSpecAttestationAuthorityHint)***||
## ContainerAnalysisNoteSpecAttestationAuthorityHint
##### (Appears on:[ContainerAnalysisNoteSpecAttestationAuthority](#ContainerAnalysisNoteSpecAttestationAuthority))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `humanReadableName` | ***string***||
## ContainerAnalysisNoteStatus
##### (Appears on:[ContainerAnalysisNote](#ContainerAnalysisNote))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ContainerAnalysisNoteSpec](#ContainerAnalysisNoteSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
