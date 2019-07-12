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

type AzurermDevTestLinuxVirtualMachine struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDevTestLinuxVirtualMachineSpec   `json:"spec,omitempty"`
	Status            AzurermDevTestLinuxVirtualMachineStatus `json:"status,omitempty"`
}

type AzurermDevTestLinuxVirtualMachineSpecInboundNatRule struct {
	Protocol     string `json:"protocol"`
	BackendPort  int    `json:"backend_port"`
	FrontendPort int    `json:"frontend_port"`
}

type AzurermDevTestLinuxVirtualMachineSpecGalleryImageReference struct {
	Sku       string `json:"sku"`
	Version   string `json:"version"`
	Offer     string `json:"offer"`
	Publisher string `json:"publisher"`
}

type AzurermDevTestLinuxVirtualMachineSpec struct {
	ResourceGroupName       string                                  `json:"resource_group_name"`
	LabVirtualNetworkId     string                                  `json:"lab_virtual_network_id"`
	SshKey                  string                                  `json:"ssh_key"`
	InboundNatRule          []AzurermDevTestLinuxVirtualMachineSpec `json:"inbound_nat_rule"`
	Fqdn                    string                                  `json:"fqdn"`
	Name                    string                                  `json:"name"`
	Username                string                                  `json:"username"`
	DisallowPublicIpAddress bool                                    `json:"disallow_public_ip_address"`
	Password                string                                  `json:"password"`
	Location                string                                  `json:"location"`
	Size                    string                                  `json:"size"`
	StorageType             string                                  `json:"storage_type"`
	GalleryImageReference   []AzurermDevTestLinuxVirtualMachineSpec `json:"gallery_image_reference"`
	Notes                   string                                  `json:"notes"`
	UniqueIdentifier        string                                  `json:"unique_identifier"`
	LabName                 string                                  `json:"lab_name"`
	LabSubnetName           string                                  `json:"lab_subnet_name"`
	AllowClaim              bool                                    `json:"allow_claim"`
	Tags                    map[string]string                       `json:"tags"`
}

type AzurermDevTestLinuxVirtualMachineStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermDevTestLinuxVirtualMachineList is a list of AzurermDevTestLinuxVirtualMachines
type AzurermDevTestLinuxVirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDevTestLinuxVirtualMachine CRD objects
	Items []AzurermDevTestLinuxVirtualMachine `json:"items,omitempty"`
}
