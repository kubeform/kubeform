### SqlSslCert
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `SqlSslCert` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SqlSslCertSpec](#SqlSslCertSpec)***||
| `status` | ***[SqlSslCertStatus](#SqlSslCertStatus)***||
### SqlSslCertSpec
(<em>Appears on:</em>
[SqlSslCert](#SqlSslCert), 
[SqlSslCertStatus](#SqlSslCertStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `cert` | ***string***| ***(Optional)*** |
| `certSerialNumber` | ***string***| ***(Optional)*** |
| `commonName` | ***string***||
| `createTime` | ***string***| ***(Optional)*** |
| `expirationTime` | ***string***| ***(Optional)*** |
| `instance` | ***string***||
| `serverCaCert` | ***string***| ***(Optional)*** |
| `sha1Fingerprint` | ***string***| ***(Optional)*** |
### SqlSslCertStatus
(<em>Appears on:</em>
[SqlSslCert](#SqlSslCert))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SqlSslCertSpec](#SqlSslCertSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `private_key` |string||
