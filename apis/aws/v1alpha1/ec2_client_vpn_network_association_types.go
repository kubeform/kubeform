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

type Ec2ClientVPNNetworkAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2ClientVPNNetworkAssociationSpec   `json:"spec,omitempty"`
	Status            Ec2ClientVPNNetworkAssociationStatus `json:"status,omitempty"`
}

type Ec2ClientVPNNetworkAssociationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ClientVPNEndpointID string `json:"clientVPNEndpointID" tf:"client_vpn_endpoint_id"`
	// +optional
	SecurityGroups []string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`
	// +optional
	Status   string `json:"status,omitempty" tf:"status,omitempty"`
	SubnetID string `json:"subnetID" tf:"subnet_id"`
	// +optional
	VpcID string `json:"vpcID,omitempty" tf:"vpc_id,omitempty"`
}

type Ec2ClientVPNNetworkAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *Ec2ClientVPNNetworkAssociationSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
