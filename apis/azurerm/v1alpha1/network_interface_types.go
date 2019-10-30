/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type NetworkInterface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkInterfaceSpec   `json:"spec,omitempty"`
	Status            NetworkInterfaceStatus `json:"status,omitempty"`
}

type NetworkInterfaceSpecIpConfiguration struct {
	// +optional
	// Deprecated
	ApplicationGatewayBackendAddressPoolsIDS []string `json:"applicationGatewayBackendAddressPoolsIDS,omitempty" tf:"application_gateway_backend_address_pools_ids,omitempty"`
	// +optional
	// Deprecated
	ApplicationSecurityGroupIDS []string `json:"applicationSecurityGroupIDS,omitempty" tf:"application_security_group_ids,omitempty"`
	// +optional
	// Deprecated
	LoadBalancerBackendAddressPoolsIDS []string `json:"loadBalancerBackendAddressPoolsIDS,omitempty" tf:"load_balancer_backend_address_pools_ids,omitempty"`
	// +optional
	// Deprecated
	LoadBalancerInboundNATRulesIDS []string `json:"loadBalancerInboundNATRulesIDS,omitempty" tf:"load_balancer_inbound_nat_rules_ids,omitempty"`
	Name                           string   `json:"name" tf:"name"`
	// +optional
	Primary bool `json:"primary,omitempty" tf:"primary,omitempty"`
	// +optional
	PrivateIPAddress           string `json:"privateIPAddress,omitempty" tf:"private_ip_address,omitempty"`
	PrivateIPAddressAllocation string `json:"privateIPAddressAllocation" tf:"private_ip_address_allocation"`
	// +optional
	PrivateIPAddressVersion string `json:"privateIPAddressVersion,omitempty" tf:"private_ip_address_version,omitempty"`
	// +optional
	PublicIPAddressID string `json:"publicIPAddressID,omitempty" tf:"public_ip_address_id,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
}

type NetworkInterfaceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AppliedDNSServers []string `json:"appliedDNSServers,omitempty" tf:"applied_dns_servers,omitempty"`
	// +optional
	DnsServers []string `json:"dnsServers,omitempty" tf:"dns_servers,omitempty"`
	// +optional
	EnableAcceleratedNetworking bool `json:"enableAcceleratedNetworking,omitempty" tf:"enable_accelerated_networking,omitempty"`
	// +optional
	EnableIPForwarding bool `json:"enableIPForwarding,omitempty" tf:"enable_ip_forwarding,omitempty"`
	// +optional
	InternalDNSNameLabel string `json:"internalDNSNameLabel,omitempty" tf:"internal_dns_name_label,omitempty"`
	// +optional
	// Deprecated
	InternalFqdn    string                                `json:"internalFqdn,omitempty" tf:"internal_fqdn,omitempty"`
	IpConfiguration []NetworkInterfaceSpecIpConfiguration `json:"ipConfiguration" tf:"ip_configuration"`
	Location        string                                `json:"location" tf:"location"`
	// +optional
	MacAddress string `json:"macAddress,omitempty" tf:"mac_address,omitempty"`
	Name       string `json:"name" tf:"name"`
	// +optional
	NetworkSecurityGroupID string `json:"networkSecurityGroupID,omitempty" tf:"network_security_group_id,omitempty"`
	// +optional
	PrivateIPAddress string `json:"privateIPAddress,omitempty" tf:"private_ip_address,omitempty"`
	// +optional
	PrivateIPAddresses []string `json:"privateIPAddresses,omitempty" tf:"private_ip_addresses,omitempty"`
	ResourceGroupName  string   `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	VirtualMachineID string `json:"virtualMachineID,omitempty" tf:"virtual_machine_id,omitempty"`
}

type NetworkInterfaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *NetworkInterfaceSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NetworkInterfaceList is a list of NetworkInterfaces
type NetworkInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NetworkInterface CRD objects
	Items []NetworkInterface `json:"items,omitempty"`
}
