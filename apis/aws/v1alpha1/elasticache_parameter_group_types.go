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

type ElasticacheParameterGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElasticacheParameterGroupSpec   `json:"spec,omitempty"`
	Status            ElasticacheParameterGroupStatus `json:"status,omitempty"`
}

type ElasticacheParameterGroupSpecParameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type ElasticacheParameterGroupSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	Family      string `json:"family"`
	Name        string `json:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Parameter *[]ElasticacheParameterGroupSpec `json:"parameter,omitempty"`
}

type ElasticacheParameterGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
