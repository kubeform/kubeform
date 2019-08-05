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

type AmiFromInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AmiFromInstanceSpec   `json:"spec,omitempty"`
	Status            AmiFromInstanceStatus `json:"status,omitempty"`
}

type AmiFromInstanceSpecEbsBlockDevice struct {
	DeleteOnTermination bool   `json:"deleteOnTermination" tf:"delete_on_termination"`
	DeviceName          string `json:"deviceName" tf:"device_name"`
	Encrypted           bool   `json:"encrypted" tf:"encrypted"`
	Iops                int    `json:"iops" tf:"iops"`
	SnapshotID          string `json:"snapshotID" tf:"snapshot_id"`
	VolumeSize          int    `json:"volumeSize" tf:"volume_size"`
	VolumeType          string `json:"volumeType" tf:"volume_type"`
}

type AmiFromInstanceSpecEphemeralBlockDevice struct {
	DeviceName  string `json:"deviceName" tf:"device_name"`
	VirtualName string `json:"virtualName" tf:"virtual_name"`
}

type AmiFromInstanceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Architecture string `json:"architecture" tf:"architecture"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EbsBlockDevice []AmiFromInstanceSpecEbsBlockDevice `json:"ebsBlockDevice,omitempty" tf:"ebs_block_device,omitempty"`
	EnaSupport     bool                                `json:"enaSupport" tf:"ena_support"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	EphemeralBlockDevice []AmiFromInstanceSpecEphemeralBlockDevice `json:"ephemeralBlockDevice,omitempty" tf:"ephemeral_block_device,omitempty"`
	ImageLocation        string                                    `json:"imageLocation" tf:"image_location"`
	KernelID             string                                    `json:"kernelID" tf:"kernel_id"`
	ManageEbsSnapshots   bool                                      `json:"manageEbsSnapshots" tf:"manage_ebs_snapshots"`
	Name                 string                                    `json:"name" tf:"name"`
	RamdiskID            string                                    `json:"ramdiskID" tf:"ramdisk_id"`
	RootDeviceName       string                                    `json:"rootDeviceName" tf:"root_device_name"`
	RootSnapshotID       string                                    `json:"rootSnapshotID" tf:"root_snapshot_id"`
	// +optional
	SnapshotWithoutReboot bool   `json:"snapshotWithoutReboot,omitempty" tf:"snapshot_without_reboot,omitempty"`
	SourceInstanceID      string `json:"sourceInstanceID" tf:"source_instance_id"`
	SriovNetSupport       string `json:"sriovNetSupport" tf:"sriov_net_support"`
	// +optional
	Tags               map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	VirtualizationType string            `json:"virtualizationType" tf:"virtualization_type"`
}

type AmiFromInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AmiFromInstanceSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AmiFromInstanceList is a list of AmiFromInstances
type AmiFromInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AmiFromInstance CRD objects
	Items []AmiFromInstance `json:"items,omitempty"`
}
