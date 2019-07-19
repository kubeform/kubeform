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

type Token struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TokenSpec   `json:"spec,omitempty"`
	Status            TokenStatus `json:"status,omitempty"`
}

type TokenSpec struct {
	// +optional
	Expiry string `json:"expiry,omitempty" tf:"expiry,omitempty"`
	// +optional
	Label       string                    `json:"label,omitempty" tf:"label,omitempty"`
	Scopes      string                    `json:"scopes" tf:"scopes"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type TokenStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// TokenList is a list of Tokens
type TokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Token CRD objects
	Items []Token `json:"items,omitempty"`
}
