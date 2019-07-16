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

type LoggingOrganizationExclusion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoggingOrganizationExclusionSpec   `json:"spec,omitempty"`
	Status            LoggingOrganizationExclusionStatus `json:"status,omitempty"`
}

type LoggingOrganizationExclusionSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Disabled bool   `json:"disabled,omitempty"`
	Filter   string `json:"filter"`
	Name     string `json:"name"`
	OrgId    string `json:"org_id"`
}

type LoggingOrganizationExclusionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LoggingOrganizationExclusionList is a list of LoggingOrganizationExclusions
type LoggingOrganizationExclusionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LoggingOrganizationExclusion CRD objects
	Items []LoggingOrganizationExclusion `json:"items,omitempty"`
}
