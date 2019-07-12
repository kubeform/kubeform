package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNeptuneClusterSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsNeptuneClusterSnapshotSpec   `json:"spec,omitempty"`
	Status            AwsNeptuneClusterSnapshotStatus `json:"status,omitempty"`
}

type AwsNeptuneClusterSnapshotSpec struct {
	AllocatedStorage            int      `json:"allocated_storage"`
	Engine                      string   `json:"engine"`
	EngineVersion               string   `json:"engine_version"`
	Port                        int      `json:"port"`
	StorageEncrypted            bool     `json:"storage_encrypted"`
	DbClusterIdentifier         string   `json:"db_cluster_identifier"`
	AvailabilityZones           []string `json:"availability_zones"`
	DbClusterSnapshotArn        string   `json:"db_cluster_snapshot_arn"`
	KmsKeyId                    string   `json:"kms_key_id"`
	SnapshotType                string   `json:"snapshot_type"`
	Status                      string   `json:"status"`
	VpcId                       string   `json:"vpc_id"`
	DbClusterSnapshotIdentifier string   `json:"db_cluster_snapshot_identifier"`
	SourceDbClusterSnapshotArn  string   `json:"source_db_cluster_snapshot_arn"`
	LicenseModel                string   `json:"license_model"`
}

type AwsNeptuneClusterSnapshotStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsNeptuneClusterSnapshotList is a list of AwsNeptuneClusterSnapshots
type AwsNeptuneClusterSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsNeptuneClusterSnapshot CRD objects
	Items []AwsNeptuneClusterSnapshot `json:"items,omitempty"`
}
