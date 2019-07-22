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

type SourcerepoRepository struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SourcerepoRepositorySpec   `json:"spec,omitempty"`
	Status            SourcerepoRepositoryStatus `json:"status,omitempty"`
}

type SourcerepoRepositorySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	Name        string                    `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
}

type SourcerepoRepositoryStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SourcerepoRepositoryList is a list of SourcerepoRepositorys
type SourcerepoRepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SourcerepoRepository CRD objects
	Items []SourcerepoRepository `json:"items,omitempty"`
}
