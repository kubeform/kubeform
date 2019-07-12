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

type AwsDbClusterSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDbClusterSnapshotSpec   `json:"spec,omitempty"`
	Status            AwsDbClusterSnapshotStatus `json:"status,omitempty"`
}

type AwsDbClusterSnapshotSpec struct {
	AllocatedStorage            int      `json:"allocated_storage"`
	AvailabilityZones           []string `json:"availability_zones"`
	DbClusterSnapshotArn        string   `json:"db_cluster_snapshot_arn"`
	KmsKeyId                    string   `json:"kms_key_id"`
	SnapshotType                string   `json:"snapshot_type"`
	DbClusterSnapshotIdentifier string   `json:"db_cluster_snapshot_identifier"`
	DbClusterIdentifier         string   `json:"db_cluster_identifier"`
	Engine                      string   `json:"engine"`
	Port                        int      `json:"port"`
	StorageEncrypted            bool     `json:"storage_encrypted"`
	EngineVersion               string   `json:"engine_version"`
	LicenseModel                string   `json:"license_model"`
	SourceDbClusterSnapshotArn  string   `json:"source_db_cluster_snapshot_arn"`
	Status                      string   `json:"status"`
	VpcId                       string   `json:"vpc_id"`
}

type AwsDbClusterSnapshotStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDbClusterSnapshotList is a list of AwsDbClusterSnapshots
type AwsDbClusterSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDbClusterSnapshot CRD objects
	Items []AwsDbClusterSnapshot `json:"items,omitempty"`
}
