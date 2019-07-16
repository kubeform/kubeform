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

type NeptuneParameterGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NeptuneParameterGroupSpec   `json:"spec,omitempty"`
	Status            NeptuneParameterGroupStatus `json:"status,omitempty"`
}

type NeptuneParameterGroupSpecParameter struct {
	// +optional
	ApplyMethod string `json:"apply_method,omitempty"`
	Name        string `json:"name"`
	Value       string `json:"value"`
}

type NeptuneParameterGroupSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	Family      string `json:"family"`
	Name        string `json:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Parameter *[]NeptuneParameterGroupSpec `json:"parameter,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type NeptuneParameterGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
