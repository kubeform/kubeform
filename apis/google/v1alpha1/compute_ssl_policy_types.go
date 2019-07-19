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

type ComputeSslPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeSslPolicySpec   `json:"spec,omitempty"`
	Status            ComputeSslPolicyStatus `json:"status,omitempty"`
}

type ComputeSslPolicySpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	CustomFeatures []string `json:"customFeatures,omitempty" tf:"custom_features,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	MinTlsVersion string `json:"minTlsVersion,omitempty" tf:"min_tls_version,omitempty"`
	Name          string `json:"name" tf:"name"`
	// +optional
	Profile string `json:"profile,omitempty" tf:"profile,omitempty"`
	// +optional
	Project     string                    `json:"project,omitempty" tf:"project,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ComputeSslPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeSslPolicyList is a list of ComputeSslPolicys
type ComputeSslPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeSslPolicy CRD objects
	Items []ComputeSslPolicy `json:"items,omitempty"`
}
