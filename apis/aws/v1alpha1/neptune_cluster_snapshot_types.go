package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	AllocatedStorage            int      `json:"allocatedStorage" tf:"allocated_storage"`
	AvailabilityZones           []string `json:"availabilityZones" tf:"availability_zones"`
	DbClusterIdentifier         string   `json:"dbClusterIdentifier" tf:"db_cluster_identifier"`
	DbClusterSnapshotArn        string   `json:"dbClusterSnapshotArn" tf:"db_cluster_snapshot_arn"`
	DbClusterSnapshotIdentifier string   `json:"dbClusterSnapshotIdentifier" tf:"db_cluster_snapshot_identifier"`
	Engine                      string   `json:"engine" tf:"engine"`
	EngineVersion               string   `json:"engineVersion" tf:"engine_version"`
	KmsKeyID                    string   `json:"kmsKeyID" tf:"kms_key_id"`
	LicenseModel                string   `json:"licenseModel" tf:"license_model"`
	Port                        int      `json:"port" tf:"port"`
	SnapshotType                string   `json:"snapshotType" tf:"snapshot_type"`
	SourceDbClusterSnapshotArn  string   `json:"sourceDbClusterSnapshotArn" tf:"source_db_cluster_snapshot_arn"`
	Status                      string   `json:"status" tf:"status"`
	StorageEncrypted            bool     `json:"storageEncrypted" tf:"storage_encrypted"`
	VpcID                       string   `json:"vpcID" tf:"vpc_id"`
}

type NeptuneClusterSnapshotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *NeptuneClusterSnapshotSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
