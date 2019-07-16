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

type DbClusterSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbClusterSnapshotSpec   `json:"spec,omitempty"`
	Status            DbClusterSnapshotStatus `json:"status,omitempty"`
}

type DbClusterSnapshotSpec struct {
	DbClusterIdentifier         string `json:"db_cluster_identifier"`
	DbClusterSnapshotIdentifier string `json:"db_cluster_snapshot_identifier"`
}

type DbClusterSnapshotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DbClusterSnapshotList is a list of DbClusterSnapshots
type DbClusterSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DbClusterSnapshot CRD objects
	Items []DbClusterSnapshot `json:"items,omitempty"`
}
