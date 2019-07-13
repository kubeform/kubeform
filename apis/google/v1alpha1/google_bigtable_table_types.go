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

type GoogleBigtableTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleBigtableTableSpec   `json:"spec,omitempty"`
	Status            GoogleBigtableTableStatus `json:"status,omitempty"`
}

type GoogleBigtableTableSpec struct {
	Project      string   `json:"project"`
	Name         string   `json:"name"`
	InstanceName string   `json:"instance_name"`
	SplitKeys    []string `json:"split_keys"`
}



type GoogleBigtableTableStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleBigtableTableList is a list of GoogleBigtableTables
type GoogleBigtableTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleBigtableTable CRD objects
	Items []GoogleBigtableTable `json:"items,omitempty"`
}