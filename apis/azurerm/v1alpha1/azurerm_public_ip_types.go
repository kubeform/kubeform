package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermPublicIp struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermPublicIpSpec   `json:"spec,omitempty"`
	Status            AzurermPublicIpStatus `json:"status,omitempty"`
}

type AzurermPublicIpSpec struct {
	IdleTimeoutInMinutes      int               `json:"idle_timeout_in_minutes"`
	PublicIpAddressAllocation string            `json:"public_ip_address_allocation"`
	Sku                       string            `json:"sku"`
	DomainNameLabel           string            `json:"domain_name_label"`
	Name                      string            `json:"name"`
	ResourceGroupName         string            `json:"resource_group_name"`
	AllocationMethod          string            `json:"allocation_method"`
	Tags                      map[string]string `json:"tags"`
	Location                  string            `json:"location"`
	IpAddress                 string            `json:"ip_address"`
	PublicIpPrefixId          string            `json:"public_ip_prefix_id"`
	Zones                     []string          `json:"zones"`
	IpVersion                 string            `json:"ip_version"`
	Fqdn                      string            `json:"fqdn"`
	ReverseFqdn               string            `json:"reverse_fqdn"`
}

type AzurermPublicIpStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermPublicIpList is a list of AzurermPublicIps
type AzurermPublicIpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermPublicIp CRD objects
	Items []AzurermPublicIp `json:"items,omitempty"`
}
