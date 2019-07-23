package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Disabled bool   `json:"disabled,omitempty" tf:"disabled,omitempty"`
	Filter   string `json:"filter" tf:"filter"`
	Name     string `json:"name" tf:"name"`
	OrgID    string `json:"orgID" tf:"org_id"`
}

type LoggingOrganizationExclusionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
