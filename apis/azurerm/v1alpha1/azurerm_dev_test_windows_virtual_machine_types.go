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

type AzurermDevTestWindowsVirtualMachine struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDevTestWindowsVirtualMachineSpec   `json:"spec,omitempty"`
	Status            AzurermDevTestWindowsVirtualMachineStatus `json:"status,omitempty"`
}

type AzurermDevTestWindowsVirtualMachineSpecGalleryImageReference struct {
	Offer     string `json:"offer"`
	Publisher string `json:"publisher"`
	Sku       string `json:"sku"`
	Version   string `json:"version"`
}

type AzurermDevTestWindowsVirtualMachineSpecInboundNatRule struct {
	BackendPort  int    `json:"backend_port"`
	FrontendPort int    `json:"frontend_port"`
	Protocol     string `json:"protocol"`
}

type AzurermDevTestWindowsVirtualMachineSpec struct {
	Location                string                                    `json:"location"`
	StorageType             string                                    `json:"storage_type"`
	AllowClaim              bool                                      `json:"allow_claim"`
	Notes                   string                                    `json:"notes"`
	UniqueIdentifier        string                                    `json:"unique_identifier"`
	GalleryImageReference   []AzurermDevTestWindowsVirtualMachineSpec `json:"gallery_image_reference"`
	Fqdn                    string                                    `json:"fqdn"`
	LabName                 string                                    `json:"lab_name"`
	ResourceGroupName       string                                    `json:"resource_group_name"`
	Username                string                                    `json:"username"`
	LabSubnetName           string                                    `json:"lab_subnet_name"`
	DisallowPublicIpAddress bool                                      `json:"disallow_public_ip_address"`
	LabVirtualNetworkId     string                                    `json:"lab_virtual_network_id"`
	Name                    string                                    `json:"name"`
	Size                    string                                    `json:"size"`
	Password                string                                    `json:"password"`
	InboundNatRule          []AzurermDevTestWindowsVirtualMachineSpec `json:"inbound_nat_rule"`
	Tags                    map[string]string                         `json:"tags"`
}

type AzurermDevTestWindowsVirtualMachineStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermDevTestWindowsVirtualMachineList is a list of AzurermDevTestWindowsVirtualMachines
type AzurermDevTestWindowsVirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDevTestWindowsVirtualMachine CRD objects
	Items []AzurermDevTestWindowsVirtualMachine `json:"items,omitempty"`
}
