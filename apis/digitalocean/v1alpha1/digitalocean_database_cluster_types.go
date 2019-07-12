package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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
	Version           string                            `json:"version"`
	NodeCount         int                               `json:"node_count"`
	MaintenanceWindow []DigitaloceanDatabaseClusterSpec `json:"maintenance_window"`
	Port              int                               `json:"port"`
	Database          string                            `json:"database"`
	Name              string                            `json:"name"`
	Engine            string                            `json:"engine"`
	Size              string                            `json:"size"`
	Region            string                            `json:"region"`
	Host              string                            `json:"host"`
	Uri               string                            `json:"uri"`
	User              string                            `json:"user"`
	Password          string                            `json:"password"`
}



type DigitaloceanDatabaseClusterStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DigitaloceanDatabaseClusterList is a list of DigitaloceanDatabaseClusters
type DigitaloceanDatabaseClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DigitaloceanDatabaseCluster CRD objects
	Items []DigitaloceanDatabaseCluster `json:"items,omitempty"`
}