package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	Day  string `json:"day" tf:"day"`
	Hour string `json:"hour" tf:"hour"`
}

type DatabaseClusterSpec struct {
	Engine string `json:"engine" tf:"engine"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	MaintenanceWindow []DatabaseClusterSpecMaintenanceWindow `json:"maintenanceWindow,omitempty" tf:"maintenance_window,omitempty"`
	Name              string                                 `json:"name" tf:"name"`
	NodeCount         int                                    `json:"nodeCount" tf:"node_count"`
	Region            string                                 `json:"region" tf:"region"`
	Size              string                                 `json:"size" tf:"size"`
	Version           string                                 `json:"version" tf:"version"`
	ProviderRef       core.LocalObjectReference              `json:"providerRef" tf:"-"`
}

type DatabaseClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
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
