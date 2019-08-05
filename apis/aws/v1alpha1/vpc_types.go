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

type Vpc struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcSpec   `json:"spec,omitempty"`
	Status            VpcStatus `json:"status,omitempty"`
}

type VpcSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	AssignGeneratedIpv6CIDRBlock bool   `json:"assignGeneratedIpv6CIDRBlock,omitempty" tf:"assign_generated_ipv6_cidr_block,omitempty"`
	CidrBlock                    string `json:"cidrBlock" tf:"cidr_block"`
	// +optional
	DefaultNetworkACLID string `json:"defaultNetworkACLID,omitempty" tf:"default_network_acl_id,omitempty"`
	// +optional
	DefaultRouteTableID string `json:"defaultRouteTableID,omitempty" tf:"default_route_table_id,omitempty"`
	// +optional
	DefaultSecurityGroupID string `json:"defaultSecurityGroupID,omitempty" tf:"default_security_group_id,omitempty"`
	// +optional
	DhcpOptionsID string `json:"dhcpOptionsID,omitempty" tf:"dhcp_options_id,omitempty"`
	// +optional
	EnableClassiclink bool `json:"enableClassiclink,omitempty" tf:"enable_classiclink,omitempty"`
	// +optional
	EnableClassiclinkDNSSupport bool `json:"enableClassiclinkDNSSupport,omitempty" tf:"enable_classiclink_dns_support,omitempty"`
	// +optional
	EnableDNSHostnames bool `json:"enableDNSHostnames,omitempty" tf:"enable_dns_hostnames,omitempty"`
	// +optional
	EnableDNSSupport bool `json:"enableDNSSupport,omitempty" tf:"enable_dns_support,omitempty"`
	// +optional
	InstanceTenancy string `json:"instanceTenancy,omitempty" tf:"instance_tenancy,omitempty"`
	// +optional
	Ipv6AssociationID string `json:"ipv6AssociationID,omitempty" tf:"ipv6_association_id,omitempty"`
	// +optional
	Ipv6CIDRBlock string `json:"ipv6CIDRBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`
	// +optional
	MainRouteTableID string `json:"mainRouteTableID,omitempty" tf:"main_route_table_id,omitempty"`
	// +optional
	OwnerID string `json:"ownerID,omitempty" tf:"owner_id,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type VpcStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *VpcSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpcList is a list of Vpcs
type VpcList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Vpc CRD objects
	Items []Vpc `json:"items,omitempty"`
}
