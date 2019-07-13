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

type AwsDbSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDbSnapshotSpec   `json:"spec,omitempty"`
	Status            AwsDbSnapshotStatus `json:"status,omitempty"`
}

type AwsDbSnapshotSpec struct {
	AvailabilityZone           string            `json:"availability_zone"`
	StorageType                string            `json:"storage_type"`
	Status                     string            `json:"status"`
	Tags                       map[string]string `json:"tags"`
	Engine                     string            `json:"engine"`
	KmsKeyId                   string            `json:"kms_key_id"`
	OptionGroupName            string            `json:"option_group_name"`
	Port                       int               `json:"port"`
	SnapshotType               string            `json:"snapshot_type"`
	DbSnapshotIdentifier       string            `json:"db_snapshot_identifier"`
	AllocatedStorage           int               `json:"allocated_storage"`
	SourceRegion               string            `json:"source_region"`
	VpcId                      string            `json:"vpc_id"`
	LicenseModel               string            `json:"license_model"`
	SourceDbSnapshotIdentifier string            `json:"source_db_snapshot_identifier"`
	DbInstanceIdentifier       string            `json:"db_instance_identifier"`
	DbSnapshotArn              string            `json:"db_snapshot_arn"`
	Encrypted                  bool              `json:"encrypted"`
	EngineVersion              string            `json:"engine_version"`
	Iops                       int               `json:"iops"`
}



type AwsDbSnapshotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDbSnapshotList is a list of AwsDbSnapshots
type AwsDbSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDbSnapshot CRD objects
	Items []AwsDbSnapshot `json:"items,omitempty"`
}