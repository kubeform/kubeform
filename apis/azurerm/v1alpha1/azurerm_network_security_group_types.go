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

type AzurermNetworkSecurityGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermNetworkSecurityGroupSpec   `json:"spec,omitempty"`
	Status            AzurermNetworkSecurityGroupStatus `json:"status,omitempty"`
}

type AzurermNetworkSecurityGroupSpecSecurityRule struct {
	SourceApplicationSecurityGroupIds      []string `json:"source_application_security_group_ids"`
	Name                                   string   `json:"name"`
	SourcePortRange                        string   `json:"source_port_range"`
	DestinationPortRanges                  []string `json:"destination_port_ranges"`
	SourceAddressPrefixes                  []string `json:"source_address_prefixes"`
	Protocol                               string   `json:"protocol"`
	DestinationPortRange                   string   `json:"destination_port_range"`
	Priority                               int      `json:"priority"`
	Direction                              string   `json:"direction"`
	Access                                 string   `json:"access"`
	Description                            string   `json:"description"`
	SourcePortRanges                       []string `json:"source_port_ranges"`
	DestinationAddressPrefix               string   `json:"destination_address_prefix"`
	DestinationApplicationSecurityGroupIds []string `json:"destination_application_security_group_ids"`
	SourceAddressPrefix                    string   `json:"source_address_prefix"`
	DestinationAddressPrefixes             []string `json:"destination_address_prefixes"`
}

type AzurermNetworkSecurityGroupSpec struct {
	SecurityRule      []AzurermNetworkSecurityGroupSpec `json:"security_rule"`
	Tags              map[string]string                 `json:"tags"`
	Name              string                            `json:"name"`
	Location          string                            `json:"location"`
	ResourceGroupName string                            `json:"resource_group_name"`
}

type AzurermNetworkSecurityGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermNetworkSecurityGroupList is a list of AzurermNetworkSecurityGroups
type AzurermNetworkSecurityGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermNetworkSecurityGroup CRD objects
	Items []AzurermNetworkSecurityGroup `json:"items,omitempty"`
}
