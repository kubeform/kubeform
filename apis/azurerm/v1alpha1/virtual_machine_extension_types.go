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

type VirtualMachineExtension struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualMachineExtensionSpec   `json:"spec,omitempty"`
	Status            VirtualMachineExtensionStatus `json:"status,omitempty"`
}

type VirtualMachineExtensionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	AutoUpgradeMinorVersion bool   `json:"autoUpgradeMinorVersion,omitempty" tf:"auto_upgrade_minor_version,omitempty"`
	Location                string `json:"location" tf:"location"`
	Name                    string `json:"name" tf:"name"`
	// +optional
	ProtectedSettings string `json:"-" sensitive:"true" tf:"protected_settings,omitempty"`
	Publisher         string `json:"publisher" tf:"publisher"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Settings string `json:"settings,omitempty" tf:"settings,omitempty"`
	// +optional
	Tags               map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	Type               string            `json:"type" tf:"type"`
	TypeHandlerVersion string            `json:"typeHandlerVersion" tf:"type_handler_version"`
	VirtualMachineName string            `json:"virtualMachineName" tf:"virtual_machine_name"`
}

type VirtualMachineExtensionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *VirtualMachineExtensionSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
