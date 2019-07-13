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

type GoogleLoggingOrganizationExclusion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleLoggingOrganizationExclusionSpec   `json:"spec,omitempty"`
	Status            GoogleLoggingOrganizationExclusionStatus `json:"status,omitempty"`
}

type GoogleLoggingOrganizationExclusionSpec struct {
	Filter      string `json:"filter"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Disabled    bool   `json:"disabled"`
	OrgId       string `json:"org_id"`
}



type GoogleLoggingOrganizationExclusionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleLoggingOrganizationExclusionList is a list of GoogleLoggingOrganizationExclusions
type GoogleLoggingOrganizationExclusionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleLoggingOrganizationExclusion CRD objects
	Items []GoogleLoggingOrganizationExclusion `json:"items,omitempty"`
}