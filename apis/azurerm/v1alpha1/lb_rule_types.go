package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type LbRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbRuleSpec   `json:"spec,omitempty"`
	Status            LbRuleStatus `json:"status,omitempty"`
}

type LbRuleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	BackendAddressPoolID string `json:"backendAddressPoolID,omitempty" tf:"backend_address_pool_id,omitempty"`
	BackendPort          int    `json:"backendPort" tf:"backend_port"`
	// +optional
	DisableOutboundSnat bool `json:"disableOutboundSnat,omitempty" tf:"disable_outbound_snat,omitempty"`
	// +optional
	EnableFloatingIP bool `json:"enableFloatingIP,omitempty" tf:"enable_floating_ip,omitempty"`
	// +optional
	FrontendIPConfigurationID   string `json:"frontendIPConfigurationID,omitempty" tf:"frontend_ip_configuration_id,omitempty"`
	FrontendIPConfigurationName string `json:"frontendIPConfigurationName" tf:"frontend_ip_configuration_name"`
	FrontendPort                int    `json:"frontendPort" tf:"frontend_port"`
	// +optional
	IdleTimeoutInMinutes int `json:"idleTimeoutInMinutes,omitempty" tf:"idle_timeout_in_minutes,omitempty"`
	// +optional
	LoadDistribution string `json:"loadDistribution,omitempty" tf:"load_distribution,omitempty"`
	LoadbalancerID   string `json:"loadbalancerID" tf:"loadbalancer_id"`
	// +optional
	// Deprecated
	Location string `json:"location,omitempty" tf:"location,omitempty"`
	Name     string `json:"name" tf:"name"`
	// +optional
	ProbeID           string `json:"probeID,omitempty" tf:"probe_id,omitempty"`
	Protocol          string `json:"protocol" tf:"protocol"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
}

type LbRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LbRuleSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LbRuleList is a list of LbRules
type LbRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LbRule CRD objects
	Items []LbRule `json:"items,omitempty"`
}
