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

type DevTestLinuxVirtualMachine struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DevTestLinuxVirtualMachineSpec   `json:"spec,omitempty"`
	Status            DevTestLinuxVirtualMachineStatus `json:"status,omitempty"`
}

type DevTestLinuxVirtualMachineSpecGalleryImageReference struct {
	Offer     string `json:"offer"`
	Publisher string `json:"publisher"`
	Sku       string `json:"sku"`
	Version   string `json:"version"`
}

type DevTestLinuxVirtualMachineSpecInboundNatRule struct {
	BackendPort int    `json:"backend_port"`
	Protocol    string `json:"protocol"`
}

type DevTestLinuxVirtualMachineSpec struct {
	// +optional
	AllowClaim bool `json:"allow_claim,omitempty"`
	// +optional
	DisallowPublicIpAddress bool `json:"disallow_public_ip_address,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	GalleryImageReference []DevTestLinuxVirtualMachineSpec `json:"gallery_image_reference"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	InboundNatRule      *[]DevTestLinuxVirtualMachineSpec `json:"inbound_nat_rule,omitempty"`
	LabName             string                            `json:"lab_name"`
	LabSubnetName       string                            `json:"lab_subnet_name"`
	LabVirtualNetworkId string                            `json:"lab_virtual_network_id"`
	Location            string                            `json:"location"`
	Name                string                            `json:"name"`
	// +optional
	Notes string `json:"notes,omitempty"`
	// +optional
	Password          string `json:"password,omitempty"`
	ResourceGroupName string `json:"resource_group_name"`
	Size              string `json:"size"`
	// +optional
	SshKey      string `json:"ssh_key,omitempty"`
	StorageType string `json:"storage_type"`
	Username    string `json:"username"`
}

type DevTestLinuxVirtualMachineStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
