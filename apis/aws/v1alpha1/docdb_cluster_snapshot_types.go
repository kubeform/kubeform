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

type DocdbClusterSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DocdbClusterSnapshotSpec   `json:"spec,omitempty"`
	Status            DocdbClusterSnapshotStatus `json:"status,omitempty"`
}

type DocdbClusterSnapshotSpec struct {
	DbClusterIdentifier         string `json:"db_cluster_identifier"`
	DbClusterSnapshotIdentifier string `json:"db_cluster_snapshot_identifier"`
}

type DocdbClusterSnapshotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DocdbClusterSnapshotList is a list of DocdbClusterSnapshots
type DocdbClusterSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DocdbClusterSnapshot CRD objects
	Items []DocdbClusterSnapshot `json:"items,omitempty"`
}