package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermVirtualMachineExtension struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermVirtualMachineExtensionSpec   `json:"spec,omitempty"`
	Status            AzurermVirtualMachineExtensionStatus `json:"status,omitempty"`
}

type AzurermVirtualMachineExtensionSpec struct {
	Publisher               string            `json:"publisher"`
	Type                    string            `json:"type"`
	TypeHandlerVersion      string            `json:"type_handler_version"`
	ProtectedSettings       string            `json:"protected_settings"`
	Name                    string            `json:"name"`
	ResourceGroupName       string            `json:"resource_group_name"`
	AutoUpgradeMinorVersion bool              `json:"auto_upgrade_minor_version"`
	Settings                string            `json:"settings"`
	Tags                    map[string]string `json:"tags"`
	Location                string            `json:"location"`
	VirtualMachineName      string            `json:"virtual_machine_name"`
}



type AzurermVirtualMachineExtensionStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermVirtualMachineExtensionList is a list of AzurermVirtualMachineExtensions
type AzurermVirtualMachineExtensionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermVirtualMachineExtension CRD objects
	Items []AzurermVirtualMachineExtension `json:"items,omitempty"`
}