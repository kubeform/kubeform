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

type LoggingFolderExclusion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoggingFolderExclusionSpec   `json:"spec,omitempty"`
	Status            LoggingFolderExclusionStatus `json:"status,omitempty"`
}

type LoggingFolderExclusionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Disabled bool   `json:"disabled,omitempty" tf:"disabled,omitempty"`
	Filter   string `json:"filter" tf:"filter"`
	Folder   string `json:"folder" tf:"folder"`
	Name     string `json:"name" tf:"name"`
}

type LoggingFolderExclusionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LoggingFolderExclusionSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LoggingFolderExclusionList is a list of LoggingFolderExclusions
type LoggingFolderExclusionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LoggingFolderExclusion CRD objects
	Items []LoggingFolderExclusion `json:"items,omitempty"`
}
