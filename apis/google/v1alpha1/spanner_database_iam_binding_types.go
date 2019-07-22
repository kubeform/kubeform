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

type SpannerDatabaseIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpannerDatabaseIamBindingSpec   `json:"spec,omitempty"`
	Status            SpannerDatabaseIamBindingStatus `json:"status,omitempty"`
}

type SpannerDatabaseIamBindingSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	Database    string                    `json:"database" tf:"database"`
	Instance    string                    `json:"instance" tf:"instance"`
	// +kubebuilder:validation:UniqueItems=true
	Members []string `json:"members" tf:"members"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	Role    string `json:"role" tf:"role"`
}

type SpannerDatabaseIamBindingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SpannerDatabaseIamBindingList is a list of SpannerDatabaseIamBindings
type SpannerDatabaseIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SpannerDatabaseIamBinding CRD objects
	Items []SpannerDatabaseIamBinding `json:"items,omitempty"`
}
