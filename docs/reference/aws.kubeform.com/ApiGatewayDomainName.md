## ApiGatewayDomainName
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiGatewayDomainName` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiGatewayDomainNameSpec](#ApiGatewayDomainNameSpec)***||
| `status` | ***[ApiGatewayDomainNameStatus](#ApiGatewayDomainNameStatus)***||
## ApiGatewayDomainNameSpec
##### (Appears on:[ApiGatewayDomainName](#ApiGatewayDomainName), [ApiGatewayDomainNameStatus](#ApiGatewayDomainNameStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `certificateArn` | ***string***| ***(Optional)*** |
| `certificateBody` | ***string***| ***(Optional)*** |
| `certificateChain` | ***string***| ***(Optional)*** |
| `certificateName` | ***string***| ***(Optional)*** |
| `certificateUploadDate` | ***string***| ***(Optional)*** |
| `cloudfrontDomainName` | ***string***| ***(Optional)*** |
| `cloudfrontZoneID` | ***string***| ***(Optional)*** |
| `domainName` | ***string***||
| `endpointConfiguration` | ***[[]ApiGatewayDomainNameSpecEndpointConfiguration](#ApiGatewayDomainNameSpecEndpointConfiguration)***| ***(Optional)*** |
| `regionalCertificateArn` | ***string***| ***(Optional)*** |
| `regionalCertificateName` | ***string***| ***(Optional)*** |
| `regionalDomainName` | ***string***| ***(Optional)*** |
| `regionalZoneID` | ***string***| ***(Optional)*** |
## ApiGatewayDomainNameSpecEndpointConfiguration
##### (Appears on:[ApiGatewayDomainNameSpec](#ApiGatewayDomainNameSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `types` | ***[]string***||
## ApiGatewayDomainNameStatus
##### (Appears on:[ApiGatewayDomainName](#ApiGatewayDomainName))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiGatewayDomainNameSpec](#ApiGatewayDomainNameSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `certificate_private_key` | ***string*** ||
