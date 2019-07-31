package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type SubnetNetworkSecurityGroupAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubnetNetworkSecurityGroupAssociationSpec   `json:"spec,omitempty"`
	Status            SubnetNetworkSecurityGroupAssociationStatus `json:"status,omitempty"`
}

type SubnetNetworkSecurityGroupAssociationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	NetworkSecurityGroupID string `json:"networkSecurityGroupID" tf:"network_security_group_id"`
	SubnetID               string `json:"subnetID" tf:"subnet_id"`
}

type SubnetNetworkSecurityGroupAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SubnetNetworkSecurityGroupAssociationList is a list of SubnetNetworkSecurityGroupAssociations
type SubnetNetworkSecurityGroupAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SubnetNetworkSecurityGroupAssociation CRD objects
	Items []SubnetNetworkSecurityGroupAssociation `json:"items,omitempty"`
}
