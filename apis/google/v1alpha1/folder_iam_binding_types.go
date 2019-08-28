package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type FolderIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FolderIamBindingSpec   `json:"spec,omitempty"`
	Status            FolderIamBindingStatus `json:"status,omitempty"`
}

type FolderIamBindingSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Etag   string `json:"etag,omitempty" tf:"etag,omitempty"`
	Folder string `json:"folder" tf:"folder"`
	// +kubebuilder:validation:UniqueItems=true
	Members []string `json:"members" tf:"members"`
	Role    string   `json:"role" tf:"role"`
}



type FolderIamBindingStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *FolderIamBindingSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// FolderIamBindingList is a list of FolderIamBindings
type FolderIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FolderIamBinding CRD objects
	Items []FolderIamBinding `json:"items,omitempty"`
}