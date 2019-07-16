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

type BigtableInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BigtableInstanceSpec   `json:"spec,omitempty"`
	Status            BigtableInstanceStatus `json:"status,omitempty"`
}

type BigtableInstanceSpecCluster struct {
	// +optional
	ClusterId string `json:"cluster_id,omitempty"`
	// +optional
	NumNodes int `json:"num_nodes,omitempty"`
	// +optional
	StorageType string `json:"storage_type,omitempty"`
}

type BigtableInstanceSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	Cluster *[]BigtableInstanceSpec `json:"cluster,omitempty"`
	// +optional
	// Deprecated
	ClusterId string `json:"cluster_id,omitempty"`
	// +optional
	InstanceType string `json:"instance_type,omitempty"`
	Name         string `json:"name"`
	// +optional
	// Deprecated
	NumNodes int `json:"num_nodes,omitempty"`
	// +optional
	// Deprecated
	StorageType string `json:"storage_type,omitempty"`
}

type BigtableInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BigtableInstanceList is a list of BigtableInstances
type BigtableInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BigtableInstance CRD objects
	Items []BigtableInstance `json:"items,omitempty"`
}
