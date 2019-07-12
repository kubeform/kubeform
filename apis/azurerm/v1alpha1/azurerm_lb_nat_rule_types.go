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

type AzurermLbNatRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermLbNatRuleSpec   `json:"spec,omitempty"`
	Status            AzurermLbNatRuleStatus `json:"status,omitempty"`
}

type AzurermLbNatRuleSpec struct {
	EnableFloatingIp            bool   `json:"enable_floating_ip"`
	BackendPort                 int    `json:"backend_port"`
	FrontendIpConfigurationId   string `json:"frontend_ip_configuration_id"`
	Name                        string `json:"name"`
	Location                    string `json:"location"`
	Protocol                    string `json:"protocol"`
	FrontendPort                int    `json:"frontend_port"`
	FrontendIpConfigurationName string `json:"frontend_ip_configuration_name"`
	BackendIpConfigurationId    string `json:"backend_ip_configuration_id"`
	ResourceGroupName           string `json:"resource_group_name"`
	LoadbalancerId              string `json:"loadbalancer_id"`
}

type AzurermLbNatRuleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermLbNatRuleList is a list of AzurermLbNatRules
type AzurermLbNatRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermLbNatRule CRD objects
	Items []AzurermLbNatRule `json:"items,omitempty"`
}
