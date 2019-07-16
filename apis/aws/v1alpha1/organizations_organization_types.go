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

type OrganizationsOrganization struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationsOrganizationSpec   `json:"spec,omitempty"`
	Status            OrganizationsOrganizationStatus `json:"status,omitempty"`
}

type OrganizationsOrganizationSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AwsServiceAccessPrincipals []string `json:"aws_service_access_principals,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EnabledPolicyTypes []string `json:"enabled_policy_types,omitempty"`
	// +optional
	FeatureSet string `json:"feature_set,omitempty"`
}

type OrganizationsOrganizationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// OrganizationsOrganizationList is a list of OrganizationsOrganizations
type OrganizationsOrganizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OrganizationsOrganization CRD objects
	Items []OrganizationsOrganization `json:"items,omitempty"`
}
