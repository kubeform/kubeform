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

type Ec2ClientVPNNetworkAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2ClientVPNNetworkAssociationSpec   `json:"spec,omitempty"`
	Status            Ec2ClientVPNNetworkAssociationStatus `json:"status,omitempty"`
}

type Ec2ClientVPNNetworkAssociationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ClientVPNEndpointID string `json:"clientVPNEndpointID" tf:"client_vpn_endpoint_id"`
	SubnetID            string `json:"subnetID" tf:"subnet_id"`
}

type Ec2ClientVPNNetworkAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// Ec2ClientVPNNetworkAssociationList is a list of Ec2ClientVPNNetworkAssociations
type Ec2ClientVPNNetworkAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Ec2ClientVPNNetworkAssociation CRD objects
	Items []Ec2ClientVPNNetworkAssociation `json:"items,omitempty"`
}
