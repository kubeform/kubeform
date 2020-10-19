---
title: HealthcareService
menu:
  docs_{{ .version }}:
    identifier: healthcareservice-azurerm.kubeform.com
    name: HealthcareService
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## HealthcareService
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `HealthcareService` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[HealthcareServiceSpec](#healthcareservicespec)***||
| `status` | ***[HealthcareServiceStatus](#healthcareservicestatus)***||
## HealthcareServiceSpec

Appears on:[HealthcareService](#healthcareservice), [HealthcareServiceStatus](#healthcareservicestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accessPolicyObjectIDS` | ***[]string***||
| `authenticationConfiguration` | ***[[]HealthcareServiceSpecAuthenticationConfiguration](#healthcareservicespecauthenticationconfiguration)***| ***(Optional)*** |
| `corsConfiguration` | ***[[]HealthcareServiceSpecCorsConfiguration](#healthcareservicespeccorsconfiguration)***| ***(Optional)*** |
| `cosmosdbThroughput` | ***int64***| ***(Optional)*** |
| `kind` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## HealthcareServiceSpecAuthenticationConfiguration

Appears on:[HealthcareServiceSpec](#healthcareservicespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `audience` | ***string***| ***(Optional)*** |
| `authority` | ***string***| ***(Optional)*** |
| `smartProxyEnabled` | ***bool***| ***(Optional)*** |
## HealthcareServiceSpecCorsConfiguration

Appears on:[HealthcareServiceSpec](#healthcareservicespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `allowCredentials` | ***bool***| ***(Optional)*** |
| `allowedHeaders` | ***[]string***| ***(Optional)*** |
| `allowedMethods` | ***[]string***| ***(Optional)*** |
| `allowedOrigins` | ***[]string***| ***(Optional)*** |
| `maxAgeInSeconds` | ***int64***| ***(Optional)*** |
## HealthcareServiceStatus

Appears on:[HealthcareService](#healthcareservice)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[HealthcareServiceSpec](#healthcareservicespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[HealthcareServiceStatus](#healthcareservicestatus)

---
