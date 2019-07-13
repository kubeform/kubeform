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
	FrontendPort int    `json:"frontend_port"`
	Protocol     string `json:"protocol"`
	BackendPort  int    `json:"backend_port"`
}

type AzurermDevTestLinuxVirtualMachineSpecGalleryImageReference struct {
	Publisher string `json:"publisher"`
	Sku       string `json:"sku"`
	Version   string `json:"version"`
	Offer     string `json:"offer"`
}

type AzurermDevTestLinuxVirtualMachineSpec struct {
	Name                    string                                  `json:"name"`
	StorageType             string                                  `json:"storage_type"`
	SshKey                  string                                  `json:"ssh_key"`
	InboundNatRule          []AzurermDevTestLinuxVirtualMachineSpec `json:"inbound_nat_rule"`
	UniqueIdentifier        string                                  `json:"unique_identifier"`
	LabName                 string                                  `json:"lab_name"`
	Location                string                                  `json:"location"`
	LabSubnetName           string                                  `json:"lab_subnet_name"`
	LabVirtualNetworkId     string                                  `json:"lab_virtual_network_id"`
	AllowClaim              bool                                    `json:"allow_claim"`
	Username                string                                  `json:"username"`
	Tags                    map[string]string                       `json:"tags"`
	ResourceGroupName       string                                  `json:"resource_group_name"`
	Size                    string                                  `json:"size"`
	DisallowPublicIpAddress bool                                    `json:"disallow_public_ip_address"`
	Password                string                                  `json:"password"`
	GalleryImageReference   []AzurermDevTestLinuxVirtualMachineSpec `json:"gallery_image_reference"`
	Notes                   string                                  `json:"notes"`
	Fqdn                    string                                  `json:"fqdn"`
}



type AzurermDevTestLinuxVirtualMachineStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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