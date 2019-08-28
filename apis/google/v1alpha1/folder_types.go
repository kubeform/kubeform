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

type Folder struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FolderSpec   `json:"spec,omitempty"`
	Status            FolderStatus `json:"status,omitempty"`
}

type FolderSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CreateTime  string `json:"createTime,omitempty" tf:"create_time,omitempty"`
	DisplayName string `json:"displayName" tf:"display_name"`
	// +optional
	LifecycleState string `json:"lifecycleState,omitempty" tf:"lifecycle_state,omitempty"`
	// +optional
	Name   string `json:"name,omitempty" tf:"name,omitempty"`
	Parent string `json:"parent" tf:"parent"`
}



type FolderStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *FolderSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// FolderList is a list of Folders
type FolderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Folder CRD objects
	Items []Folder `json:"items,omitempty"`
}