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

type VirtualMachineDataDiskAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualMachineDataDiskAttachmentSpec   `json:"spec,omitempty"`
	Status            VirtualMachineDataDiskAttachmentStatus `json:"status,omitempty"`
}

type VirtualMachineDataDiskAttachmentSpec struct {
	Caching string `json:"caching"`
	// +optional
	CreateOption     string `json:"create_option,omitempty"`
	Lun              int    `json:"lun"`
	ManagedDiskId    string `json:"managed_disk_id"`
	VirtualMachineId string `json:"virtual_machine_id"`
	// +optional
	WriteAcceleratorEnabled bool `json:"write_accelerator_enabled,omitempty"`
}

type VirtualMachineDataDiskAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VirtualMachineDataDiskAttachmentList is a list of VirtualMachineDataDiskAttachments
type VirtualMachineDataDiskAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VirtualMachineDataDiskAttachment CRD objects
	Items []VirtualMachineDataDiskAttachment `json:"items,omitempty"`
}
