package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type VpcIpv4CIDRBlockAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcIpv4CIDRBlockAssociationSpec   `json:"spec,omitempty"`
	Status            VpcIpv4CIDRBlockAssociationStatus `json:"status,omitempty"`
}

type VpcIpv4CIDRBlockAssociationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	CidrBlock string `json:"cidrBlock" tf:"cidr_block"`
	VpcID     string `json:"vpcID" tf:"vpc_id"`
}

type VpcIpv4CIDRBlockAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpcIpv4CIDRBlockAssociationList is a list of VpcIpv4CIDRBlockAssociations
type VpcIpv4CIDRBlockAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpcIpv4CIDRBlockAssociation CRD objects
	Items []VpcIpv4CIDRBlockAssociation `json:"items,omitempty"`
}
