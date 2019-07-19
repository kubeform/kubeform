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

type SesIdentityPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SesIdentityPolicySpec   `json:"spec,omitempty"`
	Status            SesIdentityPolicyStatus `json:"status,omitempty"`
}

type SesIdentityPolicySpec struct {
	Identity    string                    `json:"identity" tf:"identity"`
	Name        string                    `json:"name" tf:"name"`
	Policy      string                    `json:"policy" tf:"policy"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type SesIdentityPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SesIdentityPolicyList is a list of SesIdentityPolicys
type SesIdentityPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SesIdentityPolicy CRD objects
	Items []SesIdentityPolicy `json:"items,omitempty"`
}
