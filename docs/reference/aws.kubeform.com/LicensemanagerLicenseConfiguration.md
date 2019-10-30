---
title: LicensemanagerLicenseConfiguration
menu:
  docs_{{ .version }}:
    identifier: licensemanagerlicenseconfiguration-aws.kubeform.com
    name: LicensemanagerLicenseConfiguration
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LicensemanagerLicenseConfiguration
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LicensemanagerLicenseConfiguration` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LicensemanagerLicenseConfigurationSpec](#licensemanagerlicenseconfigurationspec)***||
| `status` | ***[LicensemanagerLicenseConfigurationStatus](#licensemanagerlicenseconfigurationstatus)***||
## LicensemanagerLicenseConfigurationSpec

Appears on:[LicensemanagerLicenseConfiguration](#licensemanagerlicenseconfiguration), [LicensemanagerLicenseConfigurationStatus](#licensemanagerlicenseconfigurationstatus)

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

Appears on:[LicensemanagerLicenseConfiguration](#licensemanagerlicenseconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LicensemanagerLicenseConfigurationSpec](#licensemanagerlicenseconfigurationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[LicensemanagerLicenseConfigurationStatus](#licensemanagerlicenseconfigurationstatus)

---
