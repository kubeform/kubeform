package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDocdbClusterSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDocdbClusterSnapshotSpec   `json:"spec,omitempty"`
	Status            AwsDocdbClusterSnapshotStatus `json:"status,omitempty"`
}

type AwsDocdbClusterSnapshotSpec struct {
	SourceDbClusterSnapshotArn  string   `json:"source_db_cluster_snapshot_arn"`
	SnapshotType                string   `json:"snapshot_type"`
	DbClusterSnapshotIdentifier string   `json:"db_cluster_snapshot_identifier"`
	DbClusterSnapshotArn        string   `json:"db_cluster_snapshot_arn"`
	StorageEncrypted            bool     `json:"storage_encrypted"`
	Engine                      string   `json:"engine"`
	EngineVersion               string   `json:"engine_version"`
	Port                        int      `json:"port"`
	Status                      string   `json:"status"`
	DbClusterIdentifier         string   `json:"db_cluster_identifier"`
	AvailabilityZones           []string `json:"availability_zones"`
	KmsKeyId                    string   `json:"kms_key_id"`
	VpcId                       string   `json:"vpc_id"`
}

type AwsDocdbClusterSnapshotStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDocdbClusterSnapshotList is a list of AwsDocdbClusterSnapshots
type AwsDocdbClusterSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDocdbClusterSnapshot CRD objects
	Items []AwsDocdbClusterSnapshot `json:"items,omitempty"`
}
