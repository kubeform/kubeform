## PublicIP
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `PublicIP` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PublicIPSpec](#PublicIPSpec)***||
| `status` | ***[PublicIPStatus](#PublicIPStatus)***||
## PublicIPSpec
##### (Appears on:[PublicIP](#PublicIP), [PublicIPStatus](#PublicIPStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allocationMethod` | ***string***| ***(Optional)*** |
| `domainNameLabel` | ***string***| ***(Optional)*** |
| `fqdn` | ***string***| ***(Optional)*** |
| `idleTimeoutInMinutes` | ***int***| ***(Optional)*** |
| `ipAddress` | ***string***| ***(Optional)*** |
| `ipVersion` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `publicIPAddressAllocation` | ***string***| ***(Optional)*** Deprecated|
| `publicIPPrefixID` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `reverseFqdn` | ***string***| ***(Optional)*** |
| `sku` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `zones` | ***[]string***| ***(Optional)*** |
## PublicIPStatus
##### (Appears on:[PublicIP](#PublicIP))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PublicIPSpec](#PublicIPSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
