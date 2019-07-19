package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	Caching string `json:"caching" tf:"caching"`
	// +optional
	CreateOption     string `json:"createOption,omitempty" tf:"create_option,omitempty"`
	Lun              int    `json:"lun" tf:"lun"`
	ManagedDiskID    string `json:"managedDiskID" tf:"managed_disk_id"`
	VirtualMachineID string `json:"virtualMachineID" tf:"virtual_machine_id"`
	// +optional
	WriteAcceleratorEnabled bool                      `json:"writeAcceleratorEnabled,omitempty" tf:"write_accelerator_enabled,omitempty"`
	ProviderRef             core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type VirtualMachineDataDiskAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
