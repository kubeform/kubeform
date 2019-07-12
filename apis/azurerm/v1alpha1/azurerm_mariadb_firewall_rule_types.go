package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermMariadbFirewallRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermMariadbFirewallRuleSpec   `json:"spec,omitempty"`
	Status            AzurermMariadbFirewallRuleStatus `json:"status,omitempty"`
}

type AzurermMariadbFirewallRuleSpec struct {
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	ServerName        string `json:"server_name"`
	StartIpAddress    string `json:"start_ip_address"`
	EndIpAddress      string `json:"end_ip_address"`
}



type AzurermMariadbFirewallRuleStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermMariadbFirewallRuleList is a list of AzurermMariadbFirewallRules
type AzurermMariadbFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermMariadbFirewallRule CRD objects
	Items []AzurermMariadbFirewallRule `json:"items,omitempty"`
}