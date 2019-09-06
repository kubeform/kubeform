## DefaultVpcDHCPOptions
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DefaultVpcDHCPOptions` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DefaultVpcDHCPOptionsSpec](#DefaultVpcDHCPOptionsSpec)***||
| `status` | ***[DefaultVpcDHCPOptionsStatus](#DefaultVpcDHCPOptionsStatus)***||
## DefaultVpcDHCPOptionsSpec
##### (Appears on:[DefaultVpcDHCPOptions](#DefaultVpcDHCPOptions), [DefaultVpcDHCPOptionsStatus](#DefaultVpcDHCPOptionsStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `domainName` | ***string***| ***(Optional)*** |
| `domainNameServers` | ***string***| ***(Optional)*** |
| `netbiosNameServers` | ***[]string***| ***(Optional)*** |
| `netbiosNodeType` | ***string***| ***(Optional)*** |
| `ntpServers` | ***string***| ***(Optional)*** |
| `ownerID` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## DefaultVpcDHCPOptionsStatus
##### (Appears on:[DefaultVpcDHCPOptions](#DefaultVpcDHCPOptions))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DefaultVpcDHCPOptionsSpec](#DefaultVpcDHCPOptionsSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
