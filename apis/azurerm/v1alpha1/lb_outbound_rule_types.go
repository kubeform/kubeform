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

type LbOutboundRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbOutboundRuleSpec   `json:"spec,omitempty"`
	Status            LbOutboundRuleStatus `json:"status,omitempty"`
}

type LbOutboundRuleSpecFrontendIpConfiguration struct {
	Name string `json:"name"`
}

type LbOutboundRuleSpec struct {
	// +optional
	AllocatedOutboundPorts int    `json:"allocated_outbound_ports,omitempty"`
	BackendAddressPoolId   string `json:"backend_address_pool_id"`
	// +optional
	EnableTcpReset bool `json:"enable_tcp_reset,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	FrontendIpConfiguration *[]LbOutboundRuleSpec `json:"frontend_ip_configuration,omitempty"`
	// +optional
	IdleTimeoutInMinutes int    `json:"idle_timeout_in_minutes,omitempty"`
	LoadbalancerId       string `json:"loadbalancer_id"`
	Name                 string `json:"name"`
	Protocol             string `json:"protocol"`
	ResourceGroupName    string `json:"resource_group_name"`
}

type LbOutboundRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LbOutboundRuleList is a list of LbOutboundRules
type LbOutboundRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LbOutboundRule CRD objects
	Items []LbOutboundRule `json:"items,omitempty"`
}
