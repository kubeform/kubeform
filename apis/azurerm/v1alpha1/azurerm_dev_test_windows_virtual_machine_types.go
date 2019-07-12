package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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
	Protocol     string `json:"protocol"`
	BackendPort  int    `json:"backend_port"`
	FrontendPort int    `json:"frontend_port"`
}

type AzurermDevTestWindowsVirtualMachineSpec struct {
	Fqdn                    string                                    `json:"fqdn"`
	Location                string                                    `json:"location"`
	LabSubnetName           string                                    `json:"lab_subnet_name"`
	AllowClaim              bool                                      `json:"allow_claim"`
	GalleryImageReference   []AzurermDevTestWindowsVirtualMachineSpec `json:"gallery_image_reference"`
	InboundNatRule          []AzurermDevTestWindowsVirtualMachineSpec `json:"inbound_nat_rule"`
	Notes                   string                                    `json:"notes"`
	Username                string                                    `json:"username"`
	Password                string                                    `json:"password"`
	Name                    string                                    `json:"name"`
	LabName                 string                                    `json:"lab_name"`
	Size                    string                                    `json:"size"`
	StorageType             string                                    `json:"storage_type"`
	DisallowPublicIpAddress bool                                      `json:"disallow_public_ip_address"`
	ResourceGroupName       string                                    `json:"resource_group_name"`
	LabVirtualNetworkId     string                                    `json:"lab_virtual_network_id"`
	Tags                    map[string]string                         `json:"tags"`
	UniqueIdentifier        string                                    `json:"unique_identifier"`
}

type AzurermDevTestWindowsVirtualMachineStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermDevTestWindowsVirtualMachineList is a list of AzurermDevTestWindowsVirtualMachines
type AzurermDevTestWindowsVirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDevTestWindowsVirtualMachine CRD objects
	Items []AzurermDevTestWindowsVirtualMachine `json:"items,omitempty"`
}
