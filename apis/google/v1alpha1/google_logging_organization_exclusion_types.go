package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleLoggingOrganizationExclusion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleLoggingOrganizationExclusionSpec   `json:"spec,omitempty"`
	Status            GoogleLoggingOrganizationExclusionStatus `json:"status,omitempty"`
}

type GoogleLoggingOrganizationExclusionSpec struct {
	OrgId       string `json:"org_id"`
	Description string `json:"description"`
	Disabled    bool   `json:"disabled"`
	Filter      string `json:"filter"`
	Name        string `json:"name"`
}

type GoogleLoggingOrganizationExclusionStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleLoggingOrganizationExclusionList is a list of GoogleLoggingOrganizationExclusions
type GoogleLoggingOrganizationExclusionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleLoggingOrganizationExclusion CRD objects
	Items []GoogleLoggingOrganizationExclusion `json:"items,omitempty"`
}
