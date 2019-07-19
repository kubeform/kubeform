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

type NeptuneClusterSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NeptuneClusterSnapshotSpec   `json:"spec,omitempty"`
	Status            NeptuneClusterSnapshotStatus `json:"status,omitempty"`
}

type NeptuneClusterSnapshotSpec struct {
	DbClusterIdentifier         string                    `json:"dbClusterIdentifier" tf:"db_cluster_identifier"`
	DbClusterSnapshotIdentifier string                    `json:"dbClusterSnapshotIdentifier" tf:"db_cluster_snapshot_identifier"`
	ProviderRef                 core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type NeptuneClusterSnapshotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NeptuneClusterSnapshotList is a list of NeptuneClusterSnapshots
type NeptuneClusterSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NeptuneClusterSnapshot CRD objects
	Items []NeptuneClusterSnapshot `json:"items,omitempty"`
}
