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

type DatabaseCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseClusterSpec   `json:"spec,omitempty"`
	Status            DatabaseClusterStatus `json:"status,omitempty"`
}

type DatabaseClusterSpecMaintenanceWindow struct {
	Day  string `json:"day"`
	Hour string `json:"hour"`
}

type DatabaseClusterSpec struct {
	Engine string `json:"engine"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	MaintenanceWindow *[]DatabaseClusterSpec `json:"maintenance_window,omitempty"`
	Name              string                 `json:"name"`
	NodeCount         int                    `json:"node_count"`
	Region            string                 `json:"region"`
	Size              string                 `json:"size"`
	Version           string                 `json:"version"`
}

type DatabaseClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DatabaseClusterList is a list of DatabaseClusters
type DatabaseClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DatabaseCluster CRD objects
	Items []DatabaseCluster `json:"items,omitempty"`
}
