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

type SpannerInstanceIamMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpannerInstanceIamMemberSpec   `json:"spec,omitempty"`
	Status            SpannerInstanceIamMemberStatus `json:"status,omitempty"`
}

type SpannerInstanceIamMemberSpec struct {
	Instance string `json:"instance" tf:"instance"`
	Member   string `json:"member" tf:"member"`
	// +optional
	Project     string                    `json:"project,omitempty" tf:"project,omitempty"`
	Role        string                    `json:"role" tf:"role"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type SpannerInstanceIamMemberStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SpannerInstanceIamMemberList is a list of SpannerInstanceIamMembers
type SpannerInstanceIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SpannerInstanceIamMember CRD objects
	Items []SpannerInstanceIamMember `json:"items,omitempty"`
}
