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
	DbClusterIdentifier         string   `json:"db_cluster_identifier"`
	LicenseModel                string   `json:"license_model"`
	Engine                      string   `json:"engine"`
	SourceDbClusterSnapshotArn  string   `json:"source_db_cluster_snapshot_arn"`
	SnapshotType                string   `json:"snapshot_type"`
	Status                      string   `json:"status"`
	VpcId                       string   `json:"vpc_id"`
	DbClusterSnapshotIdentifier string   `json:"db_cluster_snapshot_identifier"`
	AllocatedStorage            int      `json:"allocated_storage"`
	DbClusterSnapshotArn        string   `json:"db_cluster_snapshot_arn"`
	EngineVersion               string   `json:"engine_version"`
	Port                        int      `json:"port"`
	AvailabilityZones           []string `json:"availability_zones"`
	StorageEncrypted            bool     `json:"storage_encrypted"`
	KmsKeyId                    string   `json:"kms_key_id"`
}



type AwsDbClusterSnapshotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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