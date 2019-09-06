## Cdn
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
|    `kind` | string | `Cdn` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CdnSpec](#CdnSpec)***||
| `status` | ***[CdnStatus](#CdnStatus)***||
## CdnSpec
##### (Appears on:[Cdn](#Cdn), [CdnStatus](#CdnStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `certificateID` | ***string***| ***(Optional)*** ID of a DigitalOcean managed TLS certificate for use with custom domains|
| `createdAt` | ***string***| ***(Optional)*** The date and time (ISO8601) of when the CDN endpoint was created.|
| `customDomain` | ***string***| ***(Optional)*** fully qualified domain name (FQDN) for custom subdomain, (requires certificate_id)|
| `endpoint` | ***string***| ***(Optional)*** fully qualified domain name (FQDN) to serve the CDN content|
| `origin` | ***string***|fully qualified domain name (FQDN) for the origin server|
| `ttl` | ***int***| ***(Optional)*** The amount of time the content is cached in the CDN|
## CdnStatus
##### (Appears on:[Cdn](#Cdn))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CdnSpec](#CdnSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
