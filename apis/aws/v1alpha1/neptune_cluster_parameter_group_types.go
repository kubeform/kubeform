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

type NeptuneClusterParameterGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NeptuneClusterParameterGroupSpec   `json:"spec,omitempty"`
	Status            NeptuneClusterParameterGroupStatus `json:"status,omitempty"`
}

type NeptuneClusterParameterGroupSpecParameter struct {
	// +optional
	ApplyMethod string `json:"applyMethod,omitempty" tf:"apply_method,omitempty"`
	Name        string `json:"name" tf:"name"`
	Value       string `json:"value" tf:"value"`
}

type NeptuneClusterParameterGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Family      string `json:"family" tf:"family"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	NamePrefix string `json:"namePrefix,omitempty" tf:"name_prefix,omitempty"`
	// +optional
	Parameter []NeptuneClusterParameterGroupSpecParameter `json:"parameter,omitempty" tf:"parameter,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NeptuneClusterParameterGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *NeptuneClusterParameterGroupSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NeptuneClusterParameterGroupList is a list of NeptuneClusterParameterGroups
type NeptuneClusterParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NeptuneClusterParameterGroup CRD objects
	Items []NeptuneClusterParameterGroup `json:"items,omitempty"`
}
