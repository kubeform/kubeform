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

type DevTestVirtualNetwork struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DevTestVirtualNetworkSpec   `json:"spec,omitempty"`
	Status            DevTestVirtualNetworkStatus `json:"status,omitempty"`
}

type DevTestVirtualNetworkSpecSubnet struct {
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	UseInVirtualMachineCreation string `json:"useInVirtualMachineCreation,omitempty" tf:"use_in_virtual_machine_creation,omitempty"`
	// +optional
	UsePublicIPAddress string `json:"usePublicIPAddress,omitempty" tf:"use_public_ip_address,omitempty"`
}

type DevTestVirtualNetworkSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Description       string `json:"description,omitempty" tf:"description,omitempty"`
	LabName           string `json:"labName" tf:"lab_name"`
	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Subnet []DevTestVirtualNetworkSpecSubnet `json:"subnet,omitempty" tf:"subnet,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	UniqueIdentifier string `json:"uniqueIdentifier,omitempty" tf:"unique_identifier,omitempty"`
}

type DevTestVirtualNetworkStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DevTestVirtualNetworkSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DevTestVirtualNetworkList is a list of DevTestVirtualNetworks
type DevTestVirtualNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DevTestVirtualNetwork CRD objects
	Items []DevTestVirtualNetwork `json:"items,omitempty"`
}
