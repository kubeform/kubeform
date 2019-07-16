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

type DocdbClusterParameterGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DocdbClusterParameterGroupSpec   `json:"spec,omitempty"`
	Status            DocdbClusterParameterGroupStatus `json:"status,omitempty"`
}

type DocdbClusterParameterGroupSpecParameter struct {
	// +optional
	ApplyMethod string `json:"apply_method,omitempty"`
	Name        string `json:"name"`
	Value       string `json:"value"`
}

type DocdbClusterParameterGroupSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	Family      string `json:"family"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Parameter *[]DocdbClusterParameterGroupSpec `json:"parameter,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type DocdbClusterParameterGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DocdbClusterParameterGroupList is a list of DocdbClusterParameterGroups
type DocdbClusterParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DocdbClusterParameterGroup CRD objects
	Items []DocdbClusterParameterGroup `json:"items,omitempty"`
}
