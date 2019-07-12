package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermNetworkProfile struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermNetworkProfileSpec   `json:"spec,omitempty"`
	Status            AzurermNetworkProfileStatus `json:"status,omitempty"`
}

type AzurermNetworkProfileSpecContainerNetworkInterfaceIpConfiguration struct {
	Name     string `json:"name"`
	SubnetId string `json:"subnet_id"`
}

type AzurermNetworkProfileSpecContainerNetworkInterface struct {
	Name            string                                               `json:"name"`
	IpConfiguration []AzurermNetworkProfileSpecContainerNetworkInterface `json:"ip_configuration"`
}

type AzurermNetworkProfileSpec struct {
	ContainerNetworkInterface    []AzurermNetworkProfileSpec `json:"container_network_interface"`
	ContainerNetworkInterfaceIds []string                    `json:"container_network_interface_ids"`
	Tags                         map[string]string           `json:"tags"`
	Name                         string                      `json:"name"`
	Location                     string                      `json:"location"`
	ResourceGroupName            string                      `json:"resource_group_name"`
}



type AzurermNetworkProfileStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermNetworkProfileList is a list of AzurermNetworkProfiles
type AzurermNetworkProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermNetworkProfile CRD objects
	Items []AzurermNetworkProfile `json:"items,omitempty"`
}