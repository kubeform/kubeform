package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
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
	Certificate string `json:"certificate"`
	// +optional
	DisplayName string `json:"display_name,omitempty"`
	FleetArn    string `json:"fleet_arn"`
}

type WorklinkWebsiteCertificateAuthorityAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
