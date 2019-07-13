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

type AwsAmiFromInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAmiFromInstanceSpec   `json:"spec,omitempty"`
	Status            AwsAmiFromInstanceStatus `json:"status,omitempty"`
}

type AwsAmiFromInstanceSpecEbsBlockDevice struct {
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
	Iops                int    `json:"iops"`
}

type AwsAmiFromInstanceSpecEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

type AwsAmiFromInstanceSpec struct {
	Architecture          string                   `json:"architecture"`
	EbsBlockDevice        []AwsAmiFromInstanceSpec `json:"ebs_block_device"`
	EnaSupport            bool                     `json:"ena_support"`
	SourceInstanceId      string                   `json:"source_instance_id"`
	Tags                  map[string]string        `json:"tags"`
	Description           string                   `json:"description"`
	KernelId              string                   `json:"kernel_id"`
	ManageEbsSnapshots    bool                     `json:"manage_ebs_snapshots"`
	SnapshotWithoutReboot bool                     `json:"snapshot_without_reboot"`
	SriovNetSupport       string                   `json:"sriov_net_support"`
	EphemeralBlockDevice  []AwsAmiFromInstanceSpec `json:"ephemeral_block_device"`
	ImageLocation         string                   `json:"image_location"`
	Name                  string                   `json:"name"`
	RamdiskId             string                   `json:"ramdisk_id"`
	RootSnapshotId        string                   `json:"root_snapshot_id"`
	RootDeviceName        string                   `json:"root_device_name"`
	VirtualizationType    string                   `json:"virtualization_type"`
}



type AwsAmiFromInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAmiFromInstanceList is a list of AwsAmiFromInstances
type AwsAmiFromInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAmiFromInstance CRD objects
	Items []AwsAmiFromInstance `json:"items,omitempty"`
}