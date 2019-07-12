package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermDevTestLinuxVirtualMachine struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDevTestLinuxVirtualMachineSpec   `json:"spec,omitempty"`
	Status            AzurermDevTestLinuxVirtualMachineStatus `json:"status,omitempty"`
}

type AzurermDevTestLinuxVirtualMachineSpecGalleryImageReference struct {
	Offer     string `json:"offer"`
	Publisher string `json:"publisher"`
	Sku       string `json:"sku"`
	Version   string `json:"version"`
}

type AzurermDevTestLinuxVirtualMachineSpecInboundNatRule struct {
	Protocol     string `json:"protocol"`
	BackendPort  int    `json:"backend_port"`
	FrontendPort int    `json:"frontend_port"`
}

type AzurermDevTestLinuxVirtualMachineSpec struct {
	ResourceGroupName       string                                  `json:"resource_group_name"`
	StorageType             string                                  `json:"storage_type"`
	LabVirtualNetworkId     string                                  `json:"lab_virtual_network_id"`
	Fqdn                    string                                  `json:"fqdn"`
	UniqueIdentifier        string                                  `json:"unique_identifier"`
	LabName                 string                                  `json:"lab_name"`
	AllowClaim              bool                                    `json:"allow_claim"`
	DisallowPublicIpAddress bool                                    `json:"disallow_public_ip_address"`
	GalleryImageReference   []AzurermDevTestLinuxVirtualMachineSpec `json:"gallery_image_reference"`
	Notes                   string                                  `json:"notes"`
	Name                    string                                  `json:"name"`
	Location                string                                  `json:"location"`
	SshKey                  string                                  `json:"ssh_key"`
	InboundNatRule          []AzurermDevTestLinuxVirtualMachineSpec `json:"inbound_nat_rule"`
	Size                    string                                  `json:"size"`
	Username                string                                  `json:"username"`
	LabSubnetName           string                                  `json:"lab_subnet_name"`
	Password                string                                  `json:"password"`
	Tags                    map[string]string                       `json:"tags"`
}



type AzurermDevTestLinuxVirtualMachineStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermDevTestLinuxVirtualMachineList is a list of AzurermDevTestLinuxVirtualMachines
type AzurermDevTestLinuxVirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDevTestLinuxVirtualMachine CRD objects
	Items []AzurermDevTestLinuxVirtualMachine `json:"items,omitempty"`
}