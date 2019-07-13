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

type GoogleBigtableInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleBigtableInstanceSpec   `json:"spec,omitempty"`
	Status            GoogleBigtableInstanceStatus `json:"status,omitempty"`
}

type GoogleBigtableInstanceSpecCluster struct {
	ClusterId   string `json:"cluster_id"`
	Zone        string `json:"zone"`
	NumNodes    int    `json:"num_nodes"`
	StorageType string `json:"storage_type"`
}

type GoogleBigtableInstanceSpec struct {
	StorageType  string                       `json:"storage_type"`
	Name         string                       `json:"name"`
	ClusterId    string                       `json:"cluster_id"`
	NumNodes     int                          `json:"num_nodes"`
	InstanceType string                       `json:"instance_type"`
	Cluster      []GoogleBigtableInstanceSpec `json:"cluster"`
	Zone         string                       `json:"zone"`
	DisplayName  string                       `json:"display_name"`
	Project      string                       `json:"project"`
}



type GoogleBigtableInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleBigtableInstanceList is a list of GoogleBigtableInstances
type GoogleBigtableInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleBigtableInstance CRD objects
	Items []GoogleBigtableInstance `json:"items,omitempty"`
}