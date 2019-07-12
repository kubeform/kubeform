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

type DigitaloceanDatabaseCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DigitaloceanDatabaseClusterSpec   `json:"spec,omitempty"`
	Status            DigitaloceanDatabaseClusterStatus `json:"status,omitempty"`
}

type DigitaloceanDatabaseClusterSpecMaintenanceWindow struct {
	Day  string `json:"day"`
	Hour string `json:"hour"`
}

type DigitaloceanDatabaseClusterSpec struct {
	NodeCount         int                               `json:"node_count"`
	Database          string                            `json:"database"`
	User              string                            `json:"user"`
	Password          string                            `json:"password"`
	Engine            string                            `json:"engine"`
	Version           string                            `json:"version"`
	Size              string                            `json:"size"`
	Host              string                            `json:"host"`
	Port              int                               `json:"port"`
	Uri               string                            `json:"uri"`
	Name              string                            `json:"name"`
	Region            string                            `json:"region"`
	MaintenanceWindow []DigitaloceanDatabaseClusterSpec `json:"maintenance_window"`
}

type DigitaloceanDatabaseClusterStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DigitaloceanDatabaseClusterList is a list of DigitaloceanDatabaseClusters
type DigitaloceanDatabaseClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DigitaloceanDatabaseCluster CRD objects
	Items []DigitaloceanDatabaseCluster `json:"items,omitempty"`
}
