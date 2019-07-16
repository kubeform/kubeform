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

type DbParameterGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbParameterGroupSpec   `json:"spec,omitempty"`
	Status            DbParameterGroupStatus `json:"status,omitempty"`
}

type DbParameterGroupSpecParameter struct {
	// +optional
	ApplyMethod string `json:"apply_method,omitempty"`
	Name        string `json:"name"`
	Value       string `json:"value"`
}

type DbParameterGroupSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	Family      string `json:"family"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Parameter *[]DbParameterGroupSpec `json:"parameter,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type DbParameterGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DbParameterGroupList is a list of DbParameterGroups
type DbParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DbParameterGroup CRD objects
	Items []DbParameterGroup `json:"items,omitempty"`
}
