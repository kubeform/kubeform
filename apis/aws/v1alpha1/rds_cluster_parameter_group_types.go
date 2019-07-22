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

type RdsClusterParameterGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RdsClusterParameterGroupSpec   `json:"spec,omitempty"`
	Status            RdsClusterParameterGroupStatus `json:"status,omitempty"`
}

type RdsClusterParameterGroupSpecParameter struct {
	// +optional
	ApplyMethod string `json:"applyMethod,omitempty" tf:"apply_method,omitempty"`
	Name        string `json:"name" tf:"name"`
	Value       string `json:"value" tf:"value"`
}

type RdsClusterParameterGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Family      string `json:"family" tf:"family"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Parameter []RdsClusterParameterGroupSpecParameter `json:"parameter,omitempty" tf:"parameter,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type RdsClusterParameterGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RdsClusterParameterGroupList is a list of RdsClusterParameterGroups
type RdsClusterParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RdsClusterParameterGroup CRD objects
	Items []RdsClusterParameterGroup `json:"items,omitempty"`
}
