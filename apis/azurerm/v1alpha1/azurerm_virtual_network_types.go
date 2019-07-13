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

type AzurermVirtualNetwork struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermVirtualNetworkSpec   `json:"spec,omitempty"`
	Status            AzurermVirtualNetworkStatus `json:"status,omitempty"`
}

type AzurermVirtualNetworkSpecSubnet struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	AddressPrefix string `json:"address_prefix"`
	SecurityGroup string `json:"security_group"`
}

type AzurermVirtualNetworkSpecDdosProtectionPlan struct {
	Id     string `json:"id"`
	Enable bool   `json:"enable"`
}

type AzurermVirtualNetworkSpec struct {
	DnsServers         []string                    `json:"dns_servers"`
	Subnet             []AzurermVirtualNetworkSpec `json:"subnet"`
	Tags               map[string]string           `json:"tags"`
	Name               string                      `json:"name"`
	ResourceGroupName  string                      `json:"resource_group_name"`
	Location           string                      `json:"location"`
	AddressSpace       []string                    `json:"address_space"`
	DdosProtectionPlan []AzurermVirtualNetworkSpec `json:"ddos_protection_plan"`
}



type AzurermVirtualNetworkStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermVirtualNetworkList is a list of AzurermVirtualNetworks
type AzurermVirtualNetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermVirtualNetwork CRD objects
	Items []AzurermVirtualNetwork `json:"items,omitempty"`
}