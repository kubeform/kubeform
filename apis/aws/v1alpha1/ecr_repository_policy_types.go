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

type EcrRepositoryPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EcrRepositoryPolicySpec   `json:"spec,omitempty"`
	Status            EcrRepositoryPolicyStatus `json:"status,omitempty"`
}

type EcrRepositoryPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Policy     string `json:"policy" tf:"policy"`
	Repository string `json:"repository" tf:"repository"`
}

type EcrRepositoryPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *EcrRepositoryPolicySpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EcrRepositoryPolicyList is a list of EcrRepositoryPolicys
type EcrRepositoryPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EcrRepositoryPolicy CRD objects
	Items []EcrRepositoryPolicy `json:"items,omitempty"`
}
