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

type SpannerInstanceIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpannerInstanceIamBindingSpec   `json:"spec,omitempty"`
	Status            SpannerInstanceIamBindingStatus `json:"status,omitempty"`
}

type SpannerInstanceIamBindingSpec struct {
	Instance string `json:"instance" tf:"instance"`
	// +kubebuilder:validation:UniqueItems=true
	Members []string `json:"members" tf:"members"`
	// +optional
	Project     string                    `json:"project,omitempty" tf:"project,omitempty"`
	Role        string                    `json:"role" tf:"role"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type SpannerInstanceIamBindingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SpannerInstanceIamBindingList is a list of SpannerInstanceIamBindings
type SpannerInstanceIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SpannerInstanceIamBinding CRD objects
	Items []SpannerInstanceIamBinding `json:"items,omitempty"`
}
