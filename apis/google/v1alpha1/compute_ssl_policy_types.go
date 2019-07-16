package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
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
	CustomFeatures []string `json:"custom_features,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	MinTlsVersion string `json:"min_tls_version,omitempty"`
	Name          string `json:"name"`
	// +optional
	Profile string `json:"profile,omitempty"`
}

type ComputeSslPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
