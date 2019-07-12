package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermVirtualMachineDataDiskAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermVirtualMachineDataDiskAttachmentSpec   `json:"spec,omitempty"`
	Status            AzurermVirtualMachineDataDiskAttachmentStatus `json:"status,omitempty"`
}

type AzurermVirtualMachineDataDiskAttachmentSpec struct {
	CreateOption            string `json:"create_option"`
	WriteAcceleratorEnabled bool   `json:"write_accelerator_enabled"`
	ManagedDiskId           string `json:"managed_disk_id"`
	VirtualMachineId        string `json:"virtual_machine_id"`
	Lun                     int    `json:"lun"`
	Caching                 string `json:"caching"`
}



type AzurermVirtualMachineDataDiskAttachmentStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermVirtualMachineDataDiskAttachmentList is a list of AzurermVirtualMachineDataDiskAttachments
type AzurermVirtualMachineDataDiskAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermVirtualMachineDataDiskAttachment CRD objects
	Items []AzurermVirtualMachineDataDiskAttachment `json:"items,omitempty"`
}