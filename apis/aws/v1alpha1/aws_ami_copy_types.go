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

type AwsAmiCopy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAmiCopySpec   `json:"spec,omitempty"`
	Status            AwsAmiCopyStatus `json:"status,omitempty"`
}

type AwsAmiCopySpecEbsBlockDevice struct {
	Iops                int    `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
}

type AwsAmiCopySpecEphemeralBlockDevice struct {
	VirtualName string `json:"virtual_name"`
	DeviceName  string `json:"device_name"`
}

type AwsAmiCopySpec struct {
	Description          string            `json:"description"`
	KernelId             string            `json:"kernel_id"`
	KmsKeyId             string            `json:"kms_key_id"`
	Tags                 map[string]string `json:"tags"`
	EbsBlockDevice       []AwsAmiCopySpec  `json:"ebs_block_device"`
	EnaSupport           bool              `json:"ena_support"`
	ManageEbsSnapshots   bool              `json:"manage_ebs_snapshots"`
	Name                 string            `json:"name"`
	RootDeviceName       string            `json:"root_device_name"`
	SourceAmiRegion      string            `json:"source_ami_region"`
	VirtualizationType   string            `json:"virtualization_type"`
	EphemeralBlockDevice []AwsAmiCopySpec  `json:"ephemeral_block_device"`
	ImageLocation        string            `json:"image_location"`
	RootSnapshotId       string            `json:"root_snapshot_id"`
	SourceAmiId          string            `json:"source_ami_id"`
	SriovNetSupport      string            `json:"sriov_net_support"`
	Architecture         string            `json:"architecture"`
	Encrypted            bool              `json:"encrypted"`
	RamdiskId            string            `json:"ramdisk_id"`
}

type AwsAmiCopyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAmiCopyList is a list of AwsAmiCopys
type AwsAmiCopyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAmiCopy CRD objects
	Items []AwsAmiCopy `json:"items,omitempty"`
}
