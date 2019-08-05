package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type VpcEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcEndpointSpec   `json:"spec,omitempty"`
	Status            VpcEndpointStatus `json:"status,omitempty"`
}

type VpcEndpointSpecDnsEntry struct {
	DnsName      string `json:"dnsName" tf:"dns_name"`
	HostedZoneID string `json:"hostedZoneID" tf:"hosted_zone_id"`
}

type VpcEndpointSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	AutoAccept bool                      `json:"autoAccept,omitempty" tf:"auto_accept,omitempty"`
	CidrBlocks []string                  `json:"cidrBlocks" tf:"cidr_blocks"`
	DnsEntry   []VpcEndpointSpecDnsEntry `json:"dnsEntry" tf:"dns_entry"`
	// +kubebuilder:validation:UniqueItems=true
	NetworkInterfaceIDS []string `json:"networkInterfaceIDS" tf:"network_interface_ids"`
	OwnerID             string   `json:"ownerID" tf:"owner_id"`
	// +optional
	Policy       string `json:"policy,omitempty" tf:"policy,omitempty"`
	PrefixListID string `json:"prefixListID" tf:"prefix_list_id"`
	// +optional
	PrivateDNSEnabled bool `json:"privateDNSEnabled,omitempty" tf:"private_dns_enabled,omitempty"`
	RequesterManaged  bool `json:"requesterManaged" tf:"requester_managed"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	RouteTableIDS []string `json:"routeTableIDS,omitempty" tf:"route_table_ids,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SecurityGroupIDS []string `json:"securityGroupIDS,omitempty" tf:"security_group_ids,omitempty"`
	ServiceName      string   `json:"serviceName" tf:"service_name"`
	State            string   `json:"state" tf:"state"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SubnetIDS []string `json:"subnetIDS,omitempty" tf:"subnet_ids,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
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
	State *apis.State `json:"state,omitempty"`
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
