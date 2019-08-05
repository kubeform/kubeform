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

type AmiCopy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AmiCopySpec   `json:"spec,omitempty"`
	Status            AmiCopyStatus `json:"status,omitempty"`
}

type AmiCopySpecEbsBlockDevice struct {
	DeleteOnTermination bool   `json:"deleteOnTermination" tf:"delete_on_termination"`
	DeviceName          string `json:"deviceName" tf:"device_name"`
	Encrypted           bool   `json:"encrypted" tf:"encrypted"`
	Iops                int    `json:"iops" tf:"iops"`
	SnapshotID          string `json:"snapshotID" tf:"snapshot_id"`
	VolumeSize          int    `json:"volumeSize" tf:"volume_size"`
	VolumeType          string `json:"volumeType" tf:"volume_type"`
}

type AmiCopySpecEphemeralBlockDevice struct {
	DeviceName  string `json:"deviceName" tf:"device_name"`
	VirtualName string `json:"virtualName" tf:"virtual_name"`
}

type AmiCopySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Architecture string `json:"architecture" tf:"architecture"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EbsBlockDevice []AmiCopySpecEbsBlockDevice `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device,omitempty"`
	EnaSupport     bool                        `json:"enaSupport" tf:"ena_support"`
	// +optional
	Encrypted bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EphemeralBlockDevice []AmiCopySpecEphemeralBlockDevice `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device,omitempty"`
	ImageLocation        string                            `json:"imageLocation" tf:"image_location"`
	KernelID             string                            `json:"kernelID" tf:"kernel_id"`
	// +optional
	KmsKeyID           string `json:"kmsKeyID,omitempty" tf:"kms_key_id,omitempty"`
	ManageEbsSnapshots bool   `json:"manageEbsSnapshots" tf:"manage_ebs_snapshots"`
	Name               string `json:"name" tf:"name"`
	RamdiskID          string `json:"ramdiskID" tf:"ramdisk_id"`
	RootDeviceName     string `json:"rootDeviceName" tf:"root_device_name"`
	RootSnapshotID     string `json:"rootSnapshotID" tf:"root_snapshot_id"`
	SourceAmiID        string `json:"sourceAmiID" tf:"source_ami_id"`
	SourceAmiRegion    string `json:"sourceAmiRegion" tf:"source_ami_region"`
	SriovNetSupport    string `json:"sriovNetSupport" tf:"sriov_net_support"`
	// +optional
	Tags               map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	VirtualizationType string            `json:"virtualizationType" tf:"virtualization_type"`
}

type AmiCopyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AmiCopySpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AmiCopyList is a list of AmiCopys
type AmiCopyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AmiCopy CRD objects
	Items []AmiCopy `json:"items,omitempty"`
}
