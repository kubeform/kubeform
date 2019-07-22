package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type LbOutboundRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbOutboundRuleSpec   `json:"spec,omitempty"`
	Status            LbOutboundRuleStatus `json:"status,omitempty"`
}

type LbOutboundRuleSpecFrontendIPConfiguration struct {
	Name string `json:"name" tf:"name"`
}

type LbOutboundRuleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	// +optional
	AllocatedOutboundPorts int    `json:"allocatedOutboundPorts,omitempty" tf:"allocated_outbound_ports,omitempty"`
	BackendAddressPoolID   string `json:"backendAddressPoolID" tf:"backend_address_pool_id"`
	// +optional
	EnableTcpReset bool `json:"enableTcpReset,omitempty" tf:"enable_tcp_reset,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	FrontendIPConfiguration []LbOutboundRuleSpecFrontendIPConfiguration `json:"frontendIPConfiguration,omitempty" tf:"frontend_ip_configuration,omitempty"`
	// +optional
	IdleTimeoutInMinutes int    `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`
	LoadbalancerID       string `json:"loadbalancerID" tf:"loadbalancer_id"`
	Name                 string `json:"name" tf:"name"`
	Protocol             string `json:"protocol" tf:"protocol"`
	ResourceGroupName    string `json:"resourceGroupName" tf:"resource_group_name"`
}

type LbOutboundRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
