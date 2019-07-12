package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AzurermNetworkInterface struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermNetworkInterfaceSpec   `json:"spec,omitempty"`
	Status            AzurermNetworkInterfaceStatus `json:"status,omitempty"`
}

type AzurermNetworkInterfaceSpecIpConfiguration struct {
	SubnetId                                 string   `json:"subnet_id"`
	PrivateIpAddressAllocation               string   `json:"private_ip_address_allocation"`
	LoadBalancerBackendAddressPoolsIds       []string `json:"load_balancer_backend_address_pools_ids"`
	ApplicationSecurityGroupIds              []string `json:"application_security_group_ids"`
	Primary                                  bool     `json:"primary"`
	Name                                     string   `json:"name"`
	PrivateIpAddress                         string   `json:"private_ip_address"`
	PrivateIpAddressVersion                  string   `json:"private_ip_address_version"`
	PublicIpAddressId                        string   `json:"public_ip_address_id"`
	ApplicationGatewayBackendAddressPoolsIds []string `json:"application_gateway_backend_address_pools_ids"`
	LoadBalancerInboundNatRulesIds           []string `json:"load_balancer_inbound_nat_rules_ids"`
}

type AzurermNetworkInterfaceSpec struct {
	VirtualMachineId            string                        `json:"virtual_machine_id"`
	EnableIpForwarding          bool                          `json:"enable_ip_forwarding"`
	Tags                        map[string]string             `json:"tags"`
	Name                        string                        `json:"name"`
	Location                    string                        `json:"location"`
	ResourceGroupName           string                        `json:"resource_group_name"`
	EnableAcceleratedNetworking bool                          `json:"enable_accelerated_networking"`
	IpConfiguration             []AzurermNetworkInterfaceSpec `json:"ip_configuration"`
	InternalDnsNameLabel        string                        `json:"internal_dns_name_label"`
	AppliedDnsServers           []string                      `json:"applied_dns_servers"`
	NetworkSecurityGroupId      string                        `json:"network_security_group_id"`
	MacAddress                  string                        `json:"mac_address"`
	PrivateIpAddress            string                        `json:"private_ip_address"`
	PrivateIpAddresses          []string                      `json:"private_ip_addresses"`
	DnsServers                  []string                      `json:"dns_servers"`
	InternalFqdn                string                        `json:"internal_fqdn"`
}

type AzurermNetworkInterfaceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermNetworkInterfaceList is a list of AzurermNetworkInterfaces
type AzurermNetworkInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermNetworkInterface CRD objects
	Items []AzurermNetworkInterface `json:"items,omitempty"`
}
