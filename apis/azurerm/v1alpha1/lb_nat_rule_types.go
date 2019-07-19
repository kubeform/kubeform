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

type LbNATRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbNATRuleSpec   `json:"spec,omitempty"`
	Status            LbNATRuleStatus `json:"status,omitempty"`
}

type LbNATRuleSpec struct {
	BackendPort int `json:"backendPort" tf:"backend_port"`
	// +optional
	EnableFloatingIP            bool   `json:"enableFloatingIP,omitempty" tf:"enable_floating_ip,omitempty"`
	FrontendIPConfigurationName string `json:"frontendIPConfigurationName" tf:"frontend_ip_configuration_name"`
	FrontendPort                int    `json:"frontendPort" tf:"frontend_port"`
	LoadbalancerID              string `json:"loadbalancerID" tf:"loadbalancer_id"`
	// +optional
	// Deprecated
	Location          string                    `json:"location,omitempty" tf:"location,omitempty"`
	Name              string                    `json:"name" tf:"name"`
	Protocol          string                    `json:"protocol" tf:"protocol"`
	ResourceGroupName string                    `json:"resourceGroupName" tf:"resource_group_name"`
	ProviderRef       core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type LbNATRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LbNATRuleList is a list of LbNATRules
type LbNATRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LbNATRule CRD objects
	Items []LbNATRule `json:"items,omitempty"`
}
