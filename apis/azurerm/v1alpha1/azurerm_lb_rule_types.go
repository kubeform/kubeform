package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermLbRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermLbRuleSpec   `json:"spec,omitempty"`
	Status            AzurermLbRuleStatus `json:"status,omitempty"`
}

type AzurermLbRuleSpec struct {
	LoadbalancerId              string `json:"loadbalancer_id"`
	Protocol                    string `json:"protocol"`
	EnableFloatingIp            bool   `json:"enable_floating_ip"`
	DisableOutboundSnat         bool   `json:"disable_outbound_snat"`
	ResourceGroupName           string `json:"resource_group_name"`
	FrontendIpConfigurationId   string `json:"frontend_ip_configuration_id"`
	BackendPort                 int    `json:"backend_port"`
	BackendAddressPoolId        string `json:"backend_address_pool_id"`
	ProbeId                     string `json:"probe_id"`
	LoadDistribution            string `json:"load_distribution"`
	Name                        string `json:"name"`
	Location                    string `json:"location"`
	FrontendIpConfigurationName string `json:"frontend_ip_configuration_name"`
	FrontendPort                int    `json:"frontend_port"`
	IdleTimeoutInMinutes        int    `json:"idle_timeout_in_minutes"`
}

type AzurermLbRuleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermLbRuleList is a list of AzurermLbRules
type AzurermLbRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermLbRule CRD objects
	Items []AzurermLbRule `json:"items,omitempty"`
}
