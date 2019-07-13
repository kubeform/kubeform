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

type AwsAmiCopySpecEphemeralBlockDevice struct {
	DeviceName  string `json:"device_name"`
	VirtualName string `json:"virtual_name"`
}

type AwsAmiCopySpecEbsBlockDevice struct {
	VolumeSize          int    `json:"volume_size"`
	VolumeType          string `json:"volume_type"`
	DeleteOnTermination bool   `json:"delete_on_termination"`
	DeviceName          string `json:"device_name"`
	Encrypted           bool   `json:"encrypted"`
	Iops                int    `json:"iops"`
	SnapshotId          string `json:"snapshot_id"`
}

type AwsAmiCopySpec struct {
	EphemeralBlockDevice []AwsAmiCopySpec  `json:"ephemeral_block_device"`
	ImageLocation        string            `json:"image_location"`
	KernelId             string            `json:"kernel_id"`
	KmsKeyId             string            `json:"kms_key_id"`
	SriovNetSupport      string            `json:"sriov_net_support"`
	VirtualizationType   string            `json:"virtualization_type"`
	Description          string            `json:"description"`
	EbsBlockDevice       []AwsAmiCopySpec  `json:"ebs_block_device"`
	EnaSupport           bool              `json:"ena_support"`
	Architecture         string            `json:"architecture"`
	Name                 string            `json:"name"`
	RamdiskId            string            `json:"ramdisk_id"`
	RootSnapshotId       string            `json:"root_snapshot_id"`
	SourceAmiId          string            `json:"source_ami_id"`
	SourceAmiRegion      string            `json:"source_ami_region"`
	ManageEbsSnapshots   bool              `json:"manage_ebs_snapshots"`
	RootDeviceName       string            `json:"root_device_name"`
	Tags                 map[string]string `json:"tags"`
	Encrypted            bool              `json:"encrypted"`
}



type AwsAmiCopyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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