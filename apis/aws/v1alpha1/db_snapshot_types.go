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

	// +optional
	AllocatedStorage int `json:"allocatedStorage,omitempty" tf:"allocated_storage,omitempty"`
	// +optional
	AvailabilityZone     string `json:"availabilityZone,omitempty" tf:"availability_zone,omitempty"`
	DbInstanceIdentifier string `json:"dbInstanceIdentifier" tf:"db_instance_identifier"`
	// +optional
	DbSnapshotArn        string `json:"dbSnapshotArn,omitempty" tf:"db_snapshot_arn,omitempty"`
	DbSnapshotIdentifier string `json:"dbSnapshotIdentifier" tf:"db_snapshot_identifier"`
	// +optional
	Encrypted bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`
	// +optional
	Engine string `json:"engine,omitempty" tf:"engine,omitempty"`
	// +optional
	EngineVersion string `json:"engineVersion,omitempty" tf:"engine_version,omitempty"`
	// +optional
	Iops int `json:"iops,omitempty" tf:"iops,omitempty"`
	// +optional
	KmsKeyID string `json:"kmsKeyID,omitempty" tf:"kms_key_id,omitempty"`
	// +optional
	LicenseModel string `json:"licenseModel,omitempty" tf:"license_model,omitempty"`
	// +optional
	OptionGroupName string `json:"optionGroupName,omitempty" tf:"option_group_name,omitempty"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	SnapshotType string `json:"snapshotType,omitempty" tf:"snapshot_type,omitempty"`
	// +optional
	SourceDbSnapshotIdentifier string `json:"sourceDbSnapshotIdentifier,omitempty" tf:"source_db_snapshot_identifier,omitempty"`
	// +optional
	SourceRegion string `json:"sourceRegion,omitempty" tf:"source_region,omitempty"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// +optional
	StorageType string `json:"storageType,omitempty" tf:"storage_type,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	VpcID string `json:"vpcID,omitempty" tf:"vpc_id,omitempty"`
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
