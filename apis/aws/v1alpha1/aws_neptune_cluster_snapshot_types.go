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

type AwsNeptuneClusterSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsNeptuneClusterSnapshotSpec   `json:"spec,omitempty"`
	Status            AwsNeptuneClusterSnapshotStatus `json:"status,omitempty"`
}

type AwsNeptuneClusterSnapshotSpec struct {
	AvailabilityZones           []string `json:"availability_zones"`
	DbClusterSnapshotArn        string   `json:"db_cluster_snapshot_arn"`
	Engine                      string   `json:"engine"`
	LicenseModel                string   `json:"license_model"`
	Port                        int      `json:"port"`
	StorageEncrypted            bool     `json:"storage_encrypted"`
	VpcId                       string   `json:"vpc_id"`
	AllocatedStorage            int      `json:"allocated_storage"`
	DbClusterIdentifier         string   `json:"db_cluster_identifier"`
	DbClusterSnapshotIdentifier string   `json:"db_cluster_snapshot_identifier"`
	SourceDbClusterSnapshotArn  string   `json:"source_db_cluster_snapshot_arn"`
	KmsKeyId                    string   `json:"kms_key_id"`
	SnapshotType                string   `json:"snapshot_type"`
	Status                      string   `json:"status"`
	EngineVersion               string   `json:"engine_version"`
}



type AwsNeptuneClusterSnapshotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsNeptuneClusterSnapshotList is a list of AwsNeptuneClusterSnapshots
type AwsNeptuneClusterSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsNeptuneClusterSnapshot CRD objects
	Items []AwsNeptuneClusterSnapshot `json:"items,omitempty"`
}