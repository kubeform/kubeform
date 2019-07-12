package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermLbOutboundRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermLbOutboundRuleSpec   `json:"spec,omitempty"`
	Status            AzurermLbOutboundRuleStatus `json:"status,omitempty"`
}

type AzurermLbOutboundRuleSpecFrontendIpConfiguration struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type AzurermLbOutboundRuleSpec struct {
	IdleTimeoutInMinutes    int                         `json:"idle_timeout_in_minutes"`
	FrontendIpConfiguration []AzurermLbOutboundRuleSpec `json:"frontend_ip_configuration"`
	BackendAddressPoolId    string                      `json:"backend_address_pool_id"`
	EnableTcpReset          bool                        `json:"enable_tcp_reset"`
	AllocatedOutboundPorts  int                         `json:"allocated_outbound_ports"`
	Name                    string                      `json:"name"`
	ResourceGroupName       string                      `json:"resource_group_name"`
	LoadbalancerId          string                      `json:"loadbalancer_id"`
	Protocol                string                      `json:"protocol"`
}



type AzurermLbOutboundRuleStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermLbOutboundRuleList is a list of AzurermLbOutboundRules
type AzurermLbOutboundRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermLbOutboundRule CRD objects
	Items []AzurermLbOutboundRule `json:"items,omitempty"`
}