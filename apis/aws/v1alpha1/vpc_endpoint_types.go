package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type VpcEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcEndpointSpec   `json:"spec,omitempty"`
	Status            VpcEndpointStatus `json:"status,omitempty"`
}

type VpcEndpointSpecDnsEntry struct {
	// +optional
	DnsName string `json:"dnsName,omitempty" tf:"dns_name,omitempty"`
	// +optional
	HostedZoneID string `json:"hostedZoneID,omitempty" tf:"hosted_zone_id,omitempty"`
}

type VpcEndpointSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AutoAccept bool `json:"autoAccept,omitempty" tf:"auto_accept,omitempty"`
	// +optional
	CidrBlocks []string `json:"cidrBlocks,omitempty" tf:"cidr_blocks,omitempty"`
	// +optional
	DnsEntry []VpcEndpointSpecDnsEntry `json:"dnsEntry,omitempty" tf:"dns_entry,omitempty"`
	// +optional
	NetworkInterfaceIDS []string `json:"networkInterfaceIDS,omitempty" tf:"network_interface_ids,omitempty"`
	// +optional
	Policy string `json:"policy,omitempty" tf:"policy,omitempty"`
	// +optional
	PrefixListID string `json:"prefixListID,omitempty" tf:"prefix_list_id,omitempty"`
	// +optional
	PrivateDNSEnabled bool `json:"privateDNSEnabled,omitempty" tf:"private_dns_enabled,omitempty"`
	// +optional
	RouteTableIDS []string `json:"routeTableIDS,omitempty" tf:"route_table_ids,omitempty"`
	// +optional
	SecurityGroupIDS []string `json:"securityGroupIDS,omitempty" tf:"security_group_ids,omitempty"`
	ServiceName      string   `json:"serviceName" tf:"service_name"`
	// +optional
	State string `json:"state,omitempty" tf:"state,omitempty"`
	// +optional
	SubnetIDS []string `json:"subnetIDS,omitempty" tf:"subnet_ids,omitempty"`
	// +optional
	VpcEndpointType string `json:"vpcEndpointType,omitempty" tf:"vpc_endpoint_type,omitempty"`
	VpcID           string `json:"vpcID" tf:"vpc_id"`
}

type VpcEndpointStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *VpcEndpointSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpcEndpointList is a list of VpcEndpoints
type VpcEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpcEndpoint CRD objects
	Items []VpcEndpoint `json:"items,omitempty"`
}
