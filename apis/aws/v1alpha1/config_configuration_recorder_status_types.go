package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type ConfigConfigurationRecorderStatus_ struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigConfigurationRecorderStatus_Spec   `json:"spec,omitempty"`
	Status            ConfigConfigurationRecorderStatus_Status `json:"status,omitempty"`
}

type ConfigConfigurationRecorderStatus_Spec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	IsEnabled bool   `json:"isEnabled" tf:"is_enabled"`
	Name      string `json:"name" tf:"name"`
}

type ConfigConfigurationRecorderStatus_Status struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ConfigConfigurationRecorderStatus_Spec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ConfigConfigurationRecorderStatus_List is a list of ConfigConfigurationRecorderStatus_s
type ConfigConfigurationRecorderStatus_List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ConfigConfigurationRecorderStatus_ CRD objects
	Items []ConfigConfigurationRecorderStatus_ `json:"items,omitempty"`
}
