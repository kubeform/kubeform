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

type RdsGlobalCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RdsGlobalClusterSpec   `json:"spec,omitempty"`
	Status            RdsGlobalClusterStatus `json:"status,omitempty"`
}

type RdsGlobalClusterSpec struct {
	// +optional
	DatabaseName string `json:"database_name,omitempty"`
	// +optional
	DeletionProtection bool `json:"deletion_protection,omitempty"`
	// +optional
	Engine                  string `json:"engine,omitempty"`
	GlobalClusterIdentifier string `json:"global_cluster_identifier"`
	// +optional
	StorageEncrypted bool `json:"storage_encrypted,omitempty"`
}

type RdsGlobalClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RdsGlobalClusterList is a list of RdsGlobalClusters
type RdsGlobalClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RdsGlobalCluster CRD objects
	Items []RdsGlobalCluster `json:"items,omitempty"`
}
