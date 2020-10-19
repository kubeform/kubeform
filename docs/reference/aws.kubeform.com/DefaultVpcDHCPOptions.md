---
title: DefaultVpcDHCPOptions
menu:
  docs_{{ .version }}:
    identifier: defaultvpcdhcpoptions-aws.kubeform.com
    name: DefaultVpcDHCPOptions
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DefaultVpcDHCPOptions
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DefaultVpcDHCPOptions` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DefaultVpcDHCPOptionsSpec](#defaultvpcdhcpoptionsspec)***||
| `status` | ***[DefaultVpcDHCPOptionsStatus](#defaultvpcdhcpoptionsstatus)***||
## DefaultVpcDHCPOptionsSpec

Appears on:[DefaultVpcDHCPOptions](#defaultvpcdhcpoptions), [DefaultVpcDHCPOptionsStatus](#defaultvpcdhcpoptionsstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `domainName` | ***string***| ***(Optional)*** |
| `domainNameServers` | ***string***| ***(Optional)*** |
| `netbiosNameServers` | ***[]string***| ***(Optional)*** |
| `netbiosNodeType` | ***string***| ***(Optional)*** |
| `ntpServers` | ***string***| ***(Optional)*** |
| `ownerID` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## DefaultVpcDHCPOptionsStatus

Appears on:[DefaultVpcDHCPOptions](#defaultvpcdhcpoptions)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DefaultVpcDHCPOptionsSpec](#defaultvpcdhcpoptionsspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DefaultVpcDHCPOptionsStatus](#defaultvpcdhcpoptionsstatus)

---
