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

type DefaultSubnet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultSubnetSpec   `json:"spec,omitempty"`
	Status            DefaultSubnetStatus `json:"status,omitempty"`
}

type DefaultSubnetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Arn                         string `json:"arn" tf:"arn"`
	AssignIpv6AddressOnCreation bool   `json:"assignIpv6AddressOnCreation" tf:"assign_ipv6_address_on_creation"`
	AvailabilityZone            string `json:"availabilityZone" tf:"availability_zone"`
	AvailabilityZoneID          string `json:"availabilityZoneID" tf:"availability_zone_id"`
	CidrBlock                   string `json:"cidrBlock" tf:"cidr_block"`
	Ipv6CIDRBlock               string `json:"ipv6CIDRBlock" tf:"ipv6_cidr_block"`
	Ipv6CIDRBlockAssociationID  string `json:"ipv6CIDRBlockAssociationID" tf:"ipv6_cidr_block_association_id"`
	// +optional
	MapPublicIPOnLaunch bool   `json:"mapPublicIPOnLaunch,omitempty" tf:"map_public_ip_on_launch,omitempty"`
	OwnerID             string `json:"ownerID" tf:"owner_id"`
	// +optional
	Tags  map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	VpcID string            `json:"vpcID" tf:"vpc_id"`
}

type DefaultSubnetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DefaultSubnetSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DefaultSubnetList is a list of DefaultSubnets
type DefaultSubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DefaultSubnet CRD objects
	Items []DefaultSubnet `json:"items,omitempty"`
}
