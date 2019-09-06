## OpsworksStack
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `OpsworksStack` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[OpsworksStackSpec](#OpsworksStackSpec)***||
| `status` | ***[OpsworksStackStatus](#OpsworksStackStatus)***||
## OpsworksStackSpec
##### (Appears on:[OpsworksStack](#OpsworksStack), [OpsworksStackStatus](#OpsworksStackStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `agentVersion` | ***string***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `berkshelfVersion` | ***string***| ***(Optional)*** |
| `color` | ***string***| ***(Optional)*** |
| `configurationManagerName` | ***string***| ***(Optional)*** |
| `configurationManagerVersion` | ***string***| ***(Optional)*** |
| `customCookbooksSource` | ***[[]OpsworksStackSpecCustomCookbooksSource](#OpsworksStackSpecCustomCookbooksSource)***| ***(Optional)*** |
| `customJSON` | ***string***| ***(Optional)*** |
| `defaultAvailabilityZone` | ***string***| ***(Optional)*** |
| `defaultInstanceProfileArn` | ***string***||
| `defaultOs` | ***string***| ***(Optional)*** |
| `defaultRootDeviceType` | ***string***| ***(Optional)*** |
| `defaultSSHKeyName` | ***string***| ***(Optional)*** |
| `defaultSubnetID` | ***string***| ***(Optional)*** |
| `hostnameTheme` | ***string***| ***(Optional)*** |
| `manageBerkshelf` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `region` | ***string***||
| `serviceRoleArn` | ***string***||
| `stackEndpoint` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `useCustomCookbooks` | ***bool***| ***(Optional)*** |
| `useOpsworksSecurityGroups` | ***bool***| ***(Optional)*** |
| `vpcID` | ***string***| ***(Optional)*** |
## OpsworksStackSpecCustomCookbooksSource
##### (Appears on:[OpsworksStackSpec](#OpsworksStackSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `revision` | ***string***| ***(Optional)*** |
| `sshKey` | ***string***| ***(Optional)*** |
| `type` | ***string***||
| `url` | ***string***||
| `username` | ***string***| ***(Optional)*** |
## OpsworksStackStatus
##### (Appears on:[OpsworksStack](#OpsworksStack))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[OpsworksStackSpec](#OpsworksStackSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `custom_cookbooks_source.<index>.password` | ***string*** ||
