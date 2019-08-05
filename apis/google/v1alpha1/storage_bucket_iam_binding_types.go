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

type StorageBucketIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageBucketIamBindingSpec   `json:"spec,omitempty"`
	Status            StorageBucketIamBindingStatus `json:"status,omitempty"`
}

type StorageBucketIamBindingSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Bucket string `json:"bucket" tf:"bucket"`
	// +kubebuilder:validation:UniqueItems=true
	Members []string `json:"members" tf:"members"`
	Role    string   `json:"role" tf:"role"`
}

type StorageBucketIamBindingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *StorageBucketIamBindingSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageBucketIamBindingList is a list of StorageBucketIamBindings
type StorageBucketIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageBucketIamBinding CRD objects
	Items []StorageBucketIamBinding `json:"items,omitempty"`
}
