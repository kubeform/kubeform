## LicensemanagerLicenseConfiguration
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LicensemanagerLicenseConfiguration` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LicensemanagerLicenseConfigurationSpec](#LicensemanagerLicenseConfigurationSpec)***||
| `status` | ***[LicensemanagerLicenseConfigurationStatus](#LicensemanagerLicenseConfigurationStatus)***||
## LicensemanagerLicenseConfigurationSpec
##### (Appears on:[LicensemanagerLicenseConfiguration](#LicensemanagerLicenseConfiguration), [LicensemanagerLicenseConfigurationStatus](#LicensemanagerLicenseConfigurationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `licenseCount` | ***int***| ***(Optional)*** |
| `licenseCountHardLimit` | ***bool***| ***(Optional)*** |
| `licenseCountingType` | ***string***||
| `licenseRules` | ***[]string***| ***(Optional)*** |
| `name` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## LicensemanagerLicenseConfigurationStatus
##### (Appears on:[LicensemanagerLicenseConfiguration](#LicensemanagerLicenseConfiguration))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LicensemanagerLicenseConfigurationSpec](#LicensemanagerLicenseConfigurationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
