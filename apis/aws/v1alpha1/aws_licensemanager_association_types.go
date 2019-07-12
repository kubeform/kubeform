package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLicensemanagerAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLicensemanagerAssociationSpec   `json:"spec,omitempty"`
	Status            AwsLicensemanagerAssociationStatus `json:"status,omitempty"`
}

type AwsLicensemanagerAssociationSpec struct {
	ResourceArn             string `json:"resource_arn"`
	LicenseConfigurationArn string `json:"license_configuration_arn"`
}



type AwsLicensemanagerAssociationStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLicensemanagerAssociationList is a list of AwsLicensemanagerAssociations
type AwsLicensemanagerAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLicensemanagerAssociation CRD objects
	Items []AwsLicensemanagerAssociation `json:"items,omitempty"`
}