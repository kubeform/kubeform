package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type LicensemanagerAssociation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LicensemanagerAssociationSpec   `json:"spec,omitempty"`
	Status            LicensemanagerAssociationStatus `json:"status,omitempty"`
}

type LicensemanagerAssociationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	LicenseConfigurationArn string `json:"licenseConfigurationArn" tf:"license_configuration_arn"`
	ResourceArn             string `json:"resourceArn" tf:"resource_arn"`
}

type LicensemanagerAssociationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *LicensemanagerAssociationSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LicensemanagerAssociationList is a list of LicensemanagerAssociations
type LicensemanagerAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LicensemanagerAssociation CRD objects
	Items []LicensemanagerAssociation `json:"items,omitempty"`
}
