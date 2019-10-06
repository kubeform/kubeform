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

type ElasticacheParameterGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticacheParameterGroupSpec   `json:"spec,omitempty"`
	Status            ElasticacheParameterGroupStatus `json:"status,omitempty"`
}

type ElasticacheParameterGroupSpecParameter struct {
	Name  string `json:"name" tf:"name"`
	Value string `json:"value" tf:"value"`
}

type ElasticacheParameterGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Family      string `json:"family" tf:"family"`
	Name        string `json:"name" tf:"name"`
	// +optional
	Parameter []ElasticacheParameterGroupSpecParameter `json:"parameter,omitempty" tf:"parameter,omitempty"`
}

type ElasticacheParameterGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ElasticacheParameterGroupSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ElasticacheParameterGroupList is a list of ElasticacheParameterGroups
type ElasticacheParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ElasticacheParameterGroup CRD objects
	Items []ElasticacheParameterGroup `json:"items,omitempty"`
}
