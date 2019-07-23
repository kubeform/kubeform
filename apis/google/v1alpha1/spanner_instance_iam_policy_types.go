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

type SpannerInstanceIamPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SpannerInstanceIamPolicySpec   `json:"spec,omitempty"`
	Status            SpannerInstanceIamPolicyStatus `json:"status,omitempty"`
}

type SpannerInstanceIamPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Instance   string `json:"instance" tf:"instance"`
	PolicyData string `json:"policyData" tf:"policy_data"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
}

type SpannerInstanceIamPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SpannerInstanceIamPolicyList is a list of SpannerInstanceIamPolicys
type SpannerInstanceIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SpannerInstanceIamPolicy CRD objects
	Items []SpannerInstanceIamPolicy `json:"items,omitempty"`
}
