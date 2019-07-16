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

type BigtableTable struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BigtableTableSpec   `json:"spec,omitempty"`
	Status            BigtableTableStatus `json:"status,omitempty"`
}

type BigtableTableSpec struct {
	InstanceName string `json:"instance_name"`
	Name         string `json:"name"`
	// +optional
	SplitKeys []string `json:"split_keys,omitempty"`
}

type BigtableTableStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BigtableTableList is a list of BigtableTables
type BigtableTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BigtableTable CRD objects
	Items []BigtableTable `json:"items,omitempty"`
}
