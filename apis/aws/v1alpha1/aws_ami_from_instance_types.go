package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAmiFromInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAmiFromInstanceSpec   `json:"spec,omitempty"`
	Status            AwsAmiFromInstanceStatus `json:"status,omitempty"`
}

type AwsAmiFromInstanceSpecEbsBlockDevice struct {
	Iops                int    `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
}

type AwsAmiFromInstanceSpecEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

type AwsAmiFromInstanceSpec struct {
	ManageEbsSnapshots    bool                     `json:"manage_ebs_snapshots"`
	SnapshotWithoutReboot bool                     `json:"snapshot_without_reboot"`
	SriovNetSupport       string                   `json:"sriov_net_support"`
	Tags                  map[string]string        `json:"tags"`
	SourceInstanceId      string                   `json:"source_instance_id"`
	Architecture          string                   `json:"architecture"`
	EbsBlockDevice        []AwsAmiFromInstanceSpec `json:"ebs_block_device"`
	EphemeralBlockDevice  []AwsAmiFromInstanceSpec `json:"ephemeral_block_device"`
	ImageLocation         string                   `json:"image_location"`
	Name                  string                   `json:"name"`
	RootDeviceName        string                   `json:"root_device_name"`
	RootSnapshotId        string                   `json:"root_snapshot_id"`
	Description           string                   `json:"description"`
	EnaSupport            bool                     `json:"ena_support"`
	KernelId              string                   `json:"kernel_id"`
	RamdiskId             string                   `json:"ramdisk_id"`
	VirtualizationType    string                   `json:"virtualization_type"`
}



type AwsAmiFromInstanceStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAmiFromInstanceList is a list of AwsAmiFromInstances
type AwsAmiFromInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAmiFromInstance CRD objects
	Items []AwsAmiFromInstance `json:"items,omitempty"`
}