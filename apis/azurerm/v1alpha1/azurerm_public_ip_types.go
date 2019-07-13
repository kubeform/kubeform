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

type AzurermPublicIp struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermPublicIpSpec   `json:"spec,omitempty"`
	Status            AzurermPublicIpStatus `json:"status,omitempty"`
}

type AzurermPublicIpSpec struct {
	IpVersion                 string            `json:"ip_version"`
	Sku                       string            `json:"sku"`
	Zones                     []string          `json:"zones"`
	Tags                      map[string]string `json:"tags"`
	Name                      string            `json:"name"`
	AllocationMethod          string            `json:"allocation_method"`
	PublicIpAddressAllocation string            `json:"public_ip_address_allocation"`
	Location                  string            `json:"location"`
	PublicIpPrefixId          string            `json:"public_ip_prefix_id"`
	IpAddress                 string            `json:"ip_address"`
	IdleTimeoutInMinutes      int               `json:"idle_timeout_in_minutes"`
	DomainNameLabel           string            `json:"domain_name_label"`
	Fqdn                      string            `json:"fqdn"`
	ResourceGroupName         string            `json:"resource_group_name"`
	ReverseFqdn               string            `json:"reverse_fqdn"`
}



type AzurermPublicIpStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermPublicIpList is a list of AzurermPublicIps
type AzurermPublicIpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermPublicIp CRD objects
	Items []AzurermPublicIp `json:"items,omitempty"`
}