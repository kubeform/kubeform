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
	StorageType                string            `json:"storage_type"`
	AvailabilityZone           string            `json:"availability_zone"`
	OptionGroupName            string            `json:"option_group_name"`
	Port                       int               `json:"port"`
	SourceDbSnapshotIdentifier string            `json:"source_db_snapshot_identifier"`
	SourceRegion               string            `json:"source_region"`
	SnapshotType               string            `json:"snapshot_type"`
	AllocatedStorage           int               `json:"allocated_storage"`
	DbSnapshotArn              string            `json:"db_snapshot_arn"`
	EngineVersion              string            `json:"engine_version"`
	KmsKeyId                   string            `json:"kms_key_id"`
	DbInstanceIdentifier       string            `json:"db_instance_identifier"`
	Encrypted                  bool              `json:"encrypted"`
	Iops                       int               `json:"iops"`
	Status                     string            `json:"status"`
	VpcId                      string            `json:"vpc_id"`
	DbSnapshotIdentifier       string            `json:"db_snapshot_identifier"`
	Engine                     string            `json:"engine"`
	LicenseModel               string            `json:"license_model"`
	Tags                       map[string]string `json:"tags"`
}

type AwsDbSnapshotStatus struct {
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
