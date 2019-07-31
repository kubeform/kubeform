package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ApiGatewayAPIKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayAPIKeySpec   `json:"spec,omitempty"`
	Status            ApiGatewayAPIKeyStatus `json:"status,omitempty"`
}

type ApiGatewayAPIKeySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Enabled bool   `json:"enabled,omitempty" tf:"enabled,omitempty"`
	Name    string `json:"name" tf:"name"`
	// +optional
	Value string `json:"-" sensitive:"true" tf:"value,omitempty"`
}

type ApiGatewayAPIKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiGatewayAPIKeyList is a list of ApiGatewayAPIKeys
type ApiGatewayAPIKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiGatewayAPIKey CRD objects
	Items []ApiGatewayAPIKey `json:"items,omitempty"`
}
