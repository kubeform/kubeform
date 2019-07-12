package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDbSnapshotSpec   `json:"spec,omitempty"`
	Status            AwsDbSnapshotStatus `json:"status,omitempty"`
}

type AwsDbSnapshotSpec struct {
	DbInstanceIdentifier       string            `json:"db_instance_identifier"`
	AllocatedStorage           int               `json:"allocated_storage"`
	Iops                       int               `json:"iops"`
	KmsKeyId                   string            `json:"kms_key_id"`
	Port                       int               `json:"port"`
	Status                     string            `json:"status"`
	AvailabilityZone           string            `json:"availability_zone"`
	Encrypted                  bool              `json:"encrypted"`
	EngineVersion              string            `json:"engine_version"`
	VpcId                      string            `json:"vpc_id"`
	DbSnapshotIdentifier       string            `json:"db_snapshot_identifier"`
	OptionGroupName            string            `json:"option_group_name"`
	SourceDbSnapshotIdentifier string            `json:"source_db_snapshot_identifier"`
	SourceRegion               string            `json:"source_region"`
	SnapshotType               string            `json:"snapshot_type"`
	StorageType                string            `json:"storage_type"`
	Tags                       map[string]string `json:"tags"`
	DbSnapshotArn              string            `json:"db_snapshot_arn"`
	Engine                     string            `json:"engine"`
	LicenseModel               string            `json:"license_model"`
}



type AwsDbSnapshotStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDbSnapshotList is a list of AwsDbSnapshots
type AwsDbSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDbSnapshot CRD objects
	Items []AwsDbSnapshot `json:"items,omitempty"`
}