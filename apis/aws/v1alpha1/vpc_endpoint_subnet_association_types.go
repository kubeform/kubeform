package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type VpcEndpointSubnetAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VpcEndpointSubnetAssociationSpec   `json:"spec,omitempty"`
	Status            VpcEndpointSubnetAssociationStatus `json:"status,omitempty"`
}

type VpcEndpointSubnetAssociationSpec struct {
	SubnetID      string                    `json:"subnetID" tf:"subnet_id"`
	VpcEndpointID string                    `json:"vpcEndpointID" tf:"vpc_endpoint_id"`
	ProviderRef   core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type VpcEndpointSubnetAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VpcEndpointSubnetAssociationList is a list of VpcEndpointSubnetAssociations
type VpcEndpointSubnetAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VpcEndpointSubnetAssociation CRD objects
	Items []VpcEndpointSubnetAssociation `json:"items,omitempty"`
}
