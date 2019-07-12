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

type AwsWorklinkWebsiteCertificateAuthorityAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsWorklinkWebsiteCertificateAuthorityAssociationSpec   `json:"spec,omitempty"`
	Status            AwsWorklinkWebsiteCertificateAuthorityAssociationStatus `json:"status,omitempty"`
}

type AwsWorklinkWebsiteCertificateAuthorityAssociationSpec struct {
	FleetArn    string `json:"fleet_arn"`
	Certificate string `json:"certificate"`
	DisplayName string `json:"display_name"`
	WebsiteCaId string `json:"website_ca_id"`
}

type AwsWorklinkWebsiteCertificateAuthorityAssociationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsWorklinkWebsiteCertificateAuthorityAssociationList is a list of AwsWorklinkWebsiteCertificateAuthorityAssociations
type AwsWorklinkWebsiteCertificateAuthorityAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsWorklinkWebsiteCertificateAuthorityAssociation CRD objects
	Items []AwsWorklinkWebsiteCertificateAuthorityAssociation `json:"items,omitempty"`
}
