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

type AzurermFirewall struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermFirewallSpec   `json:"spec,omitempty"`
	Status            AzurermFirewallStatus `json:"status,omitempty"`
}

type AzurermFirewallSpecIpConfiguration struct {
	Name                      string `json:"name"`
	SubnetId                  string `json:"subnet_id"`
	InternalPublicIpAddressId string `json:"internal_public_ip_address_id"`
	PublicIpAddressId         string `json:"public_ip_address_id"`
	PrivateIpAddress          string `json:"private_ip_address"`
}

type AzurermFirewallSpec struct {
	ResourceGroupName string                `json:"resource_group_name"`
	IpConfiguration   []AzurermFirewallSpec `json:"ip_configuration"`
	Tags              map[string]string     `json:"tags"`
	Name              string                `json:"name"`
	Location          string                `json:"location"`
}



type AzurermFirewallStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermFirewallList is a list of AzurermFirewalls
type AzurermFirewallList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermFirewall CRD objects
	Items []AzurermFirewall `json:"items,omitempty"`
}