package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type NeptuneParameterGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NeptuneParameterGroupSpec   `json:"spec,omitempty"`
	Status            NeptuneParameterGroupStatus `json:"status,omitempty"`
}

type NeptuneParameterGroupSpecParameter struct {
	// +optional
	ApplyMethod string `json:"applyMethod,omitempty" tf:"apply_method,omitempty"`
	Name        string `json:"name" tf:"name"`
	Value       string `json:"value" tf:"value"`
}

type NeptuneParameterGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Family      string `json:"family" tf:"family"`
	Name        string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Parameter []NeptuneParameterGroupSpecParameter `json:"parameter,omitempty" tf:"parameter,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NeptuneParameterGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *NeptuneParameterGroupSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NeptuneParameterGroupList is a list of NeptuneParameterGroups
type NeptuneParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NeptuneParameterGroup CRD objects
	Items []NeptuneParameterGroup `json:"items,omitempty"`
}
