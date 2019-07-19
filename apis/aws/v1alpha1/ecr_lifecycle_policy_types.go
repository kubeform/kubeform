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

type EcrLifecyclePolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EcrLifecyclePolicySpec   `json:"spec,omitempty"`
	Status            EcrLifecyclePolicyStatus `json:"status,omitempty"`
}

type EcrLifecyclePolicySpec struct {
	Policy      string                    `json:"policy" tf:"policy"`
	Repository  string                    `json:"repository" tf:"repository"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type EcrLifecyclePolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EcrLifecyclePolicyList is a list of EcrLifecyclePolicys
type EcrLifecyclePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EcrLifecyclePolicy CRD objects
	Items []EcrLifecyclePolicy `json:"items,omitempty"`
}
