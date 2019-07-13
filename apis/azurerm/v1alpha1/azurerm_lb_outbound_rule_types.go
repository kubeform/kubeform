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
	ResourceGroupName       string                      `json:"resource_group_name"`
	EnableTcpReset          bool                        `json:"enable_tcp_reset"`
	Name                    string                      `json:"name"`
	LoadbalancerId          string                      `json:"loadbalancer_id"`
	FrontendIpConfiguration []AzurermLbOutboundRuleSpec `json:"frontend_ip_configuration"`
	BackendAddressPoolId    string                      `json:"backend_address_pool_id"`
	Protocol                string                      `json:"protocol"`
	AllocatedOutboundPorts  int                         `json:"allocated_outbound_ports"`
	IdleTimeoutInMinutes    int                         `json:"idle_timeout_in_minutes"`
}



type AzurermLbOutboundRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermLbOutboundRuleList is a list of AzurermLbOutboundRules
type AzurermLbOutboundRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermLbOutboundRule CRD objects
	Items []AzurermLbOutboundRule `json:"items,omitempty"`
}