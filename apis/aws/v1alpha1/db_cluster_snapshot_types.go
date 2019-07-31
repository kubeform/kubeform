package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	DbClusterIdentifier         string `json:"dbClusterIdentifier" tf:"db_cluster_identifier"`
	DbClusterSnapshotIdentifier string `json:"dbClusterSnapshotIdentifier" tf:"db_cluster_snapshot_identifier"`
}

type DbClusterSnapshotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
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
