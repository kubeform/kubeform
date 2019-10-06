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

type VpcDHCPOptionsAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcDHCPOptionsAssociationSpec   `json:"spec,omitempty"`
	Status            VpcDHCPOptionsAssociationStatus `json:"status,omitempty"`
}

type VpcDHCPOptionsAssociationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	DhcpOptionsID string `json:"dhcpOptionsID" tf:"dhcp_options_id"`
	VpcID         string `json:"vpcID" tf:"vpc_id"`
}

type VpcDHCPOptionsAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *VpcDHCPOptionsAssociationSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpcDHCPOptionsAssociationList is a list of VpcDHCPOptionsAssociations
type VpcDHCPOptionsAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpcDHCPOptionsAssociation CRD objects
	Items []VpcDHCPOptionsAssociation `json:"items,omitempty"`
}
