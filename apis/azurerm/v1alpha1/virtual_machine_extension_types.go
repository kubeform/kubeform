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

type VirtualMachineExtension struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualMachineExtensionSpec   `json:"spec,omitempty"`
	Status            VirtualMachineExtensionStatus `json:"status,omitempty"`
}

type VirtualMachineExtensionSpec struct {
	// +optional
	AutoUpgradeMinorVersion bool   `json:"auto_upgrade_minor_version,omitempty"`
	Location                string `json:"location"`
	Name                    string `json:"name"`
	// +optional
	ProtectedSettings string `json:"protected_settings,omitempty"`
	Publisher         string `json:"publisher"`
	ResourceGroupName string `json:"resource_group_name"`
	// +optional
	Settings           string `json:"settings,omitempty"`
	Type               string `json:"type"`
	TypeHandlerVersion string `json:"type_handler_version"`
	VirtualMachineName string `json:"virtual_machine_name"`
}

type VirtualMachineExtensionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VirtualMachineExtensionList is a list of VirtualMachineExtensions
type VirtualMachineExtensionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VirtualMachineExtension CRD objects
	Items []VirtualMachineExtension `json:"items,omitempty"`
}
