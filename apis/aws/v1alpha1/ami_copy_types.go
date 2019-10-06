package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
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
	// +optional
	DeleteOnTermination bool `json:"deleteOnTermination,omitempty" tf:"delete_on_termination,omitempty"`
	// +optional
	DeviceName string `json:"deviceName,omitempty" tf:"device_name,omitempty"`
	// +optional
	Encrypted bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`
	// +optional
	Iops int `json:"iops,omitempty" tf:"iops,omitempty"`
	// +optional
	SnapshotID string `json:"snapshotID,omitempty" tf:"snapshot_id,omitempty"`
	// +optional
	VolumeSize int `json:"volumeSize,omitempty" tf:"volume_size,omitempty"`
	// +optional
	VolumeType string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type AmiCopySpecEphemeralBlockDevice struct {
	// +optional
	DeviceName string `json:"deviceName,omitempty" tf:"device_name,omitempty"`
	// +optional
	VirtualName string `json:"virtualName,omitempty" tf:"virtual_name,omitempty"`
}

type AmiCopySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Architecture string `json:"architecture,omitempty" tf:"architecture,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	EbsBlockDevice []AmiCopySpecEbsBlockDevice `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device,omitempty"`
	// +optional
	EnaSupport bool `json:"enaSupport,omitempty" tf:"ena_support,omitempty"`
	// +optional
	Encrypted bool `json:"encrypted,omitempty" tf:"encrypted,omitempty"`
	// +optional
	EphemeralBlockDevice []AmiCopySpecEphemeralBlockDevice `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device,omitempty"`
	// +optional
	ImageLocation string `json:"imageLocation,omitempty" tf:"image_location,omitempty"`
	// +optional
	KernelID string `json:"kernelID,omitempty" tf:"kernel_id,omitempty"`
	// +optional
	KmsKeyID string `json:"kmsKeyID,omitempty" tf:"kms_key_id,omitempty"`
	// +optional
	ManageEbsSnapshots bool   `json:"manageEbsSnapshots,omitempty" tf:"manage_ebs_snapshots,omitempty"`
	Name               string `json:"name" tf:"name"`
	// +optional
	RamdiskID string `json:"ramdiskID,omitempty" tf:"ramdisk_id,omitempty"`
	// +optional
	RootDeviceName string `json:"rootDeviceName,omitempty" tf:"root_device_name,omitempty"`
	// +optional
	RootSnapshotID  string `json:"rootSnapshotID,omitempty" tf:"root_snapshot_id,omitempty"`
	SourceAmiID     string `json:"sourceAmiID" tf:"source_ami_id"`
	SourceAmiRegion string `json:"sourceAmiRegion" tf:"source_ami_region"`
	// +optional
	SriovNetSupport string `json:"sriovNetSupport,omitempty" tf:"sriov_net_support,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	VirtualizationType string `json:"virtualizationType,omitempty" tf:"virtualization_type,omitempty"`
}

type AmiCopyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AmiCopySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
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
