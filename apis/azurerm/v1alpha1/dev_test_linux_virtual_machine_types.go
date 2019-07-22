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

type DevTestLinuxVirtualMachine struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DevTestLinuxVirtualMachineSpec   `json:"spec,omitempty"`
	Status            DevTestLinuxVirtualMachineStatus `json:"status,omitempty"`
}

type DevTestLinuxVirtualMachineSpecGalleryImageReference struct {
	Offer     string `json:"offer" tf:"offer"`
	Publisher string `json:"publisher" tf:"publisher"`
	Sku       string `json:"sku" tf:"sku"`
	Version   string `json:"version" tf:"version"`
}

type DevTestLinuxVirtualMachineSpecInboundNATRule struct {
	BackendPort int    `json:"backendPort" tf:"backend_port"`
	Protocol    string `json:"protocol" tf:"protocol"`
}

type DevTestLinuxVirtualMachineSpec struct {
	Secret      *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`
	ProviderRef core.LocalObjectReference  `json:"providerRef" tf:"-"`
	// +optional
	AllowClaim bool `json:"allowClaim,omitempty" tf:"allow_claim,omitempty"`
	// +optional
	DisallowPublicIPAddress bool `json:"disallowPublicIPAddress,omitempty" tf:"disallow_public_ip_address,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	GalleryImageReference []DevTestLinuxVirtualMachineSpecGalleryImageReference `json:"galleryImageReference" tf:"gallery_image_reference"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	InboundNATRule      []DevTestLinuxVirtualMachineSpecInboundNATRule `json:"inboundNATRule,omitempty" tf:"inbound_nat_rule,omitempty"`
	LabName             string                                         `json:"labName" tf:"lab_name"`
	LabSubnetName       string                                         `json:"labSubnetName" tf:"lab_subnet_name"`
	LabVirtualNetworkID string                                         `json:"labVirtualNetworkID" tf:"lab_virtual_network_id"`
	Location            string                                         `json:"location" tf:"location"`
	Name                string                                         `json:"name" tf:"name"`
	// +optional
	Notes string `json:"notes,omitempty" tf:"notes,omitempty"`
	// +optional
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	Size              string `json:"size" tf:"size"`
	// +optional
	SshKey      string `json:"sshKey,omitempty" tf:"ssh_key,omitempty"`
	StorageType string `json:"storageType" tf:"storage_type"`
	// +optional
	Tags     map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	Username string            `json:"username" tf:"username"`
}

type DevTestLinuxVirtualMachineStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DevTestLinuxVirtualMachineList is a list of DevTestLinuxVirtualMachines
type DevTestLinuxVirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DevTestLinuxVirtualMachine CRD objects
	Items []DevTestLinuxVirtualMachine `json:"items,omitempty"`
}
