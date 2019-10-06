package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ComputeProjectMetadataItem struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeProjectMetadataItemSpec   `json:"spec,omitempty"`
	Status            ComputeProjectMetadataItemStatus `json:"status,omitempty"`
}

type ComputeProjectMetadataItemSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	Key string `json:"key" tf:"key"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	Value   string `json:"value" tf:"value"`
}

type ComputeProjectMetadataItemStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeProjectMetadataItemSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeProjectMetadataItemList is a list of ComputeProjectMetadataItems
type ComputeProjectMetadataItemList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeProjectMetadataItem CRD objects
	Items []ComputeProjectMetadataItem `json:"items,omitempty"`
}
