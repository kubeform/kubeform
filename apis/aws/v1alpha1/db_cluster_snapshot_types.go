package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type DbClusterSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbClusterSnapshotSpec   `json:"spec,omitempty"`
	Status            DbClusterSnapshotStatus `json:"status,omitempty"`
}

type DbClusterSnapshotSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AllocatedStorage int `json:"allocatedStorage,omitempty" tf:"allocated_storage,omitempty"`
	// +optional
	AvailabilityZones   []string `json:"availabilityZones,omitempty" tf:"availability_zones,omitempty"`
	DbClusterIdentifier string   `json:"dbClusterIdentifier" tf:"db_cluster_identifier"`
	// +optional
	DbClusterSnapshotArn        string `json:"dbClusterSnapshotArn,omitempty" tf:"db_cluster_snapshot_arn,omitempty"`
	DbClusterSnapshotIdentifier string `json:"dbClusterSnapshotIdentifier" tf:"db_cluster_snapshot_identifier"`
	// +optional
	Engine string `json:"engine,omitempty" tf:"engine,omitempty"`
	// +optional
	EngineVersion string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`
	// +optional
	KmsKeyID string `json:"kmsKeyID,omitempty" tf:"kms_key_id,omitempty"`
	// +optional
	LicenseModel string `json:"licenseModel,omitempty" tf:"license_model,omitempty"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	SnapshotType string `json:"snapshotType,omitempty" tf:"snapshot_type,omitempty"`
	// +optional
	SourceDbClusterSnapshotArn string `json:"sourceDbClusterSnapshotArn,omitempty" tf:"source_db_cluster_snapshot_arn,omitempty"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// +optional
	StorageEncrypted bool `json:"storageEncrypted,omitempty" tf:"storage_encrypted,omitempty"`
	// +optional
	VpcID string `json:"vpcID,omitempty" tf:"vpc_id,omitempty"`
}

type DbClusterSnapshotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DbClusterSnapshotSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
