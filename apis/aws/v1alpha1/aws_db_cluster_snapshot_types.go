package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbClusterSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDbClusterSnapshotSpec   `json:"spec,omitempty"`
	Status            AwsDbClusterSnapshotStatus `json:"status,omitempty"`
}

type AwsDbClusterSnapshotSpec struct {
	DbClusterIdentifier         string   `json:"db_cluster_identifier"`
	StorageEncrypted            bool     `json:"storage_encrypted"`
	AvailabilityZones           []string `json:"availability_zones"`
	SnapshotType                string   `json:"snapshot_type"`
	Status                      string   `json:"status"`
	DbClusterSnapshotIdentifier string   `json:"db_cluster_snapshot_identifier"`
	DbClusterSnapshotArn        string   `json:"db_cluster_snapshot_arn"`
	Engine                      string   `json:"engine"`
	EngineVersion               string   `json:"engine_version"`
	VpcId                       string   `json:"vpc_id"`
	AllocatedStorage            int      `json:"allocated_storage"`
	KmsKeyId                    string   `json:"kms_key_id"`
	LicenseModel                string   `json:"license_model"`
	Port                        int      `json:"port"`
	SourceDbClusterSnapshotArn  string   `json:"source_db_cluster_snapshot_arn"`
}



type AwsDbClusterSnapshotStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDbClusterSnapshotList is a list of AwsDbClusterSnapshots
type AwsDbClusterSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDbClusterSnapshot CRD objects
	Items []AwsDbClusterSnapshot `json:"items,omitempty"`
}