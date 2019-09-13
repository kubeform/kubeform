---
title: NetworkInterface
menu:
  docs_{{ .version }}:
    identifier: networkinterface-azurerm.kubeform.com
    name: NetworkInterface
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## NetworkInterface
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `NetworkInterface` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NetworkInterfaceSpec](#NetworkInterfaceSpec)***||
| `status` | ***[NetworkInterfaceStatus](#NetworkInterfaceStatus)***||
## NetworkInterfaceSpec
##### (Appears on:[NetworkInterface](#NetworkInterface), [NetworkInterfaceStatus](#NetworkInterfaceStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `appliedDNSServers` | ***[]string***| ***(Optional)*** |
| `dnsServers` | ***[]string***| ***(Optional)*** |
| `enableAcceleratedNetworking` | ***bool***| ***(Optional)*** |
| `enableIPForwarding` | ***bool***| ***(Optional)*** |
| `internalDNSNameLabel` | ***string***| ***(Optional)*** |
| `internalFqdn` | ***string***| ***(Optional)*** Deprecated|
| `ipConfiguration` | ***[[]NetworkInterfaceSpecIpConfiguration](#NetworkInterfaceSpecIpConfiguration)***||
| `location` | ***string***||
| `macAddress` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `networkSecurityGroupID` | ***string***| ***(Optional)*** |
| `privateIPAddress` | ***string***| ***(Optional)*** |
| `privateIPAddresses` | ***[]string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `virtualMachineID` | ***string***| ***(Optional)*** |
## NetworkInterfaceSpecIpConfiguration
##### (Appears on:[NetworkInterfaceSpec](#NetworkInterfaceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `applicationGatewayBackendAddressPoolsIDS` | ***[]string***| ***(Optional)*** Deprecated|
| `applicationSecurityGroupIDS` | ***[]string***| ***(Optional)*** Deprecated|
| `loadBalancerBackendAddressPoolsIDS` | ***[]string***| ***(Optional)*** Deprecated|
| `loadBalancerInboundNATRulesIDS` | ***[]string***| ***(Optional)*** Deprecated|
| `name` | ***string***||
| `primary` | ***bool***| ***(Optional)*** |
| `privateIPAddress` | ***string***| ***(Optional)*** |
| `privateIPAddressAllocation` | ***string***||
| `privateIPAddressVersion` | ***string***| ***(Optional)*** |
| `publicIPAddressID` | ***string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
## NetworkInterfaceStatus
##### (Appears on:[NetworkInterface](#NetworkInterface))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NetworkInterfaceSpec](#NetworkInterfaceSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
