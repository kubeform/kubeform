package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type DbSnapshot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DbSnapshotSpec   `json:"spec,omitempty"`
	Status            DbSnapshotStatus `json:"status,omitempty"`
}

type DbSnapshotSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	AllocatedStorage           int    `json:"allocatedStorage" tf:"allocated_storage"`
	AvailabilityZone           string `json:"availabilityZone" tf:"availability_zone"`
	DbInstanceIdentifier       string `json:"dbInstanceIdentifier" tf:"db_instance_identifier"`
	DbSnapshotArn              string `json:"dbSnapshotArn" tf:"db_snapshot_arn"`
	DbSnapshotIdentifier       string `json:"dbSnapshotIdentifier" tf:"db_snapshot_identifier"`
	Encrypted                  bool   `json:"encrypted" tf:"encrypted"`
	Engine                     string `json:"engine" tf:"engine"`
	EngineVersion              string `json:"engineVersion" tf:"engine_version"`
	Iops                       int    `json:"iops" tf:"iops"`
	KmsKeyID                   string `json:"kmsKeyID" tf:"kms_key_id"`
	LicenseModel               string `json:"licenseModel" tf:"license_model"`
	OptionGroupName            string `json:"optionGroupName" tf:"option_group_name"`
	Port                       int    `json:"port" tf:"port"`
	SnapshotType               string `json:"snapshotType" tf:"snapshot_type"`
	SourceDbSnapshotIdentifier string `json:"sourceDbSnapshotIdentifier" tf:"source_db_snapshot_identifier"`
	SourceRegion               string `json:"sourceRegion" tf:"source_region"`
	Status                     string `json:"status" tf:"status"`
	StorageType                string `json:"storageType" tf:"storage_type"`
	// +optional
	Tags  map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	VpcID string            `json:"vpcID" tf:"vpc_id"`
}

type DbSnapshotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DbSnapshotSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DbSnapshotList is a list of DbSnapshots
type DbSnapshotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DbSnapshot CRD objects
	Items []DbSnapshot `json:"items,omitempty"`
}
