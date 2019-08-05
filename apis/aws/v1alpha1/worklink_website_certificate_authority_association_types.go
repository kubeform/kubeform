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

type WorklinkWebsiteCertificateAuthorityAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WorklinkWebsiteCertificateAuthorityAssociationSpec   `json:"spec,omitempty"`
	Status            WorklinkWebsiteCertificateAuthorityAssociationStatus `json:"status,omitempty"`
}

type WorklinkWebsiteCertificateAuthorityAssociationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Certificate string `json:"certificate" tf:"certificate"`
	// +optional
	DisplayName string `json:"displayName,omitempty" tf:"display_name,omitempty"`
	FleetArn    string `json:"fleetArn" tf:"fleet_arn"`
}

type WorklinkWebsiteCertificateAuthorityAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *WorklinkWebsiteCertificateAuthorityAssociationSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// WorklinkWebsiteCertificateAuthorityAssociationList is a list of WorklinkWebsiteCertificateAuthorityAssociations
type WorklinkWebsiteCertificateAuthorityAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of WorklinkWebsiteCertificateAuthorityAssociation CRD objects
	Items []WorklinkWebsiteCertificateAuthorityAssociation `json:"items,omitempty"`
}
