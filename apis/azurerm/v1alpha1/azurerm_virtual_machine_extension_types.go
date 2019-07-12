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

type AzurermVirtualMachineExtension struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermVirtualMachineExtensionSpec   `json:"spec,omitempty"`
	Status            AzurermVirtualMachineExtensionStatus `json:"status,omitempty"`
}

type AzurermVirtualMachineExtensionSpec struct {
	ResourceGroupName       string            `json:"resource_group_name"`
	VirtualMachineName      string            `json:"virtual_machine_name"`
	Type                    string            `json:"type"`
	AutoUpgradeMinorVersion bool              `json:"auto_upgrade_minor_version"`
	Settings                string            `json:"settings"`
	ProtectedSettings       string            `json:"protected_settings"`
	Tags                    map[string]string `json:"tags"`
	Name                    string            `json:"name"`
	Location                string            `json:"location"`
	Publisher               string            `json:"publisher"`
	TypeHandlerVersion      string            `json:"type_handler_version"`
}

type AzurermVirtualMachineExtensionStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermVirtualMachineExtensionList is a list of AzurermVirtualMachineExtensions
type AzurermVirtualMachineExtensionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermVirtualMachineExtension CRD objects
	Items []AzurermVirtualMachineExtension `json:"items,omitempty"`
}
