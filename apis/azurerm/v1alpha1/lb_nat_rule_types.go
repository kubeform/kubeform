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

type LbNatRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbNatRuleSpec   `json:"spec,omitempty"`
	Status            LbNatRuleStatus `json:"status,omitempty"`
}

type LbNatRuleSpec struct {
	BackendPort                 int    `json:"backend_port"`
	FrontendIpConfigurationName string `json:"frontend_ip_configuration_name"`
	FrontendPort                int    `json:"frontend_port"`
	LoadbalancerId              string `json:"loadbalancer_id"`
	// +optional
	// Deprecated
	Location          string `json:"location,omitempty"`
	Name              string `json:"name"`
	Protocol          string `json:"protocol"`
	ResourceGroupName string `json:"resource_group_name"`
}

type LbNatRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LbNatRuleList is a list of LbNatRules
type LbNatRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LbNatRule CRD objects
	Items []LbNatRule `json:"items,omitempty"`
}
