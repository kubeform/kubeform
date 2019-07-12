package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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
	Name         string                       `json:"name"`
	Zone         string                       `json:"zone"`
	NumNodes     int                          `json:"num_nodes"`
	InstanceType string                       `json:"instance_type"`
	Project      string                       `json:"project"`
	ClusterId    string                       `json:"cluster_id"`
	Cluster      []GoogleBigtableInstanceSpec `json:"cluster"`
	DisplayName  string                       `json:"display_name"`
	StorageType  string                       `json:"storage_type"`
}

type GoogleBigtableInstanceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleBigtableInstanceList is a list of GoogleBigtableInstances
type GoogleBigtableInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleBigtableInstance CRD objects
	Items []GoogleBigtableInstance `json:"items,omitempty"`
}
