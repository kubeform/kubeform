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
	IpConfiguration []AzurermNetworkProfileSpecContainerNetworkInterface `json:"ip_configuration"`
	Name            string                                               `json:"name"`
}

type AzurermNetworkProfileSpec struct {
	Tags                         map[string]string           `json:"tags"`
	Name                         string                      `json:"name"`
	Location                     string                      `json:"location"`
	ResourceGroupName            string                      `json:"resource_group_name"`
	ContainerNetworkInterface    []AzurermNetworkProfileSpec `json:"container_network_interface"`
	ContainerNetworkInterfaceIds []string                    `json:"container_network_interface_ids"`
}



type AzurermNetworkProfileStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermNetworkProfileList is a list of AzurermNetworkProfiles
type AzurermNetworkProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermNetworkProfile CRD objects
	Items []AzurermNetworkProfile `json:"items,omitempty"`
}