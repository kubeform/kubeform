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

type AzurermDevTestVirtualNetwork struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDevTestVirtualNetworkSpec   `json:"spec,omitempty"`
	Status            AzurermDevTestVirtualNetworkStatus `json:"status,omitempty"`
}

type AzurermDevTestVirtualNetworkSpecSubnet struct {
	UseInVirtualMachineCreation string `json:"use_in_virtual_machine_creation"`
	UsePublicIpAddress          string `json:"use_public_ip_address"`
	Name                        string `json:"name"`
}

type AzurermDevTestVirtualNetworkSpec struct {
	Tags              map[string]string                  `json:"tags"`
	UniqueIdentifier  string                             `json:"unique_identifier"`
	Name              string                             `json:"name"`
	LabName           string                             `json:"lab_name"`
	ResourceGroupName string                             `json:"resource_group_name"`
	Description       string                             `json:"description"`
	Subnet            []AzurermDevTestVirtualNetworkSpec `json:"subnet"`
}

type AzurermDevTestVirtualNetworkStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermDevTestVirtualNetworkList is a list of AzurermDevTestVirtualNetworks
type AzurermDevTestVirtualNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDevTestVirtualNetwork CRD objects
	Items []AzurermDevTestVirtualNetwork `json:"items,omitempty"`
}
