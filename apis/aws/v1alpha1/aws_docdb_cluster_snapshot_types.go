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

type AwsDocdbClusterSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDocdbClusterSnapshotSpec   `json:"spec,omitempty"`
	Status            AwsDocdbClusterSnapshotStatus `json:"status,omitempty"`
}

type AwsDocdbClusterSnapshotSpec struct {
	DbClusterSnapshotIdentifier string   `json:"db_cluster_snapshot_identifier"`
	SnapshotType                string   `json:"snapshot_type"`
	VpcId                       string   `json:"vpc_id"`
	DbClusterIdentifier         string   `json:"db_cluster_identifier"`
	AvailabilityZones           []string `json:"availability_zones"`
	DbClusterSnapshotArn        string   `json:"db_cluster_snapshot_arn"`
	StorageEncrypted            bool     `json:"storage_encrypted"`
	Engine                      string   `json:"engine"`
	EngineVersion               string   `json:"engine_version"`
	KmsKeyId                    string   `json:"kms_key_id"`
	Port                        int      `json:"port"`
	SourceDbClusterSnapshotArn  string   `json:"source_db_cluster_snapshot_arn"`
	Status                      string   `json:"status"`
}



type AwsDocdbClusterSnapshotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDocdbClusterSnapshotList is a list of AwsDocdbClusterSnapshots
type AwsDocdbClusterSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDocdbClusterSnapshot CRD objects
	Items []AwsDocdbClusterSnapshot `json:"items,omitempty"`
}