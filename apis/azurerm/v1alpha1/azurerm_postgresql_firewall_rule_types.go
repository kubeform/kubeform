package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermPostgresqlFirewallRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermPostgresqlFirewallRuleSpec   `json:"spec,omitempty"`
	Status            AzurermPostgresqlFirewallRuleStatus `json:"status,omitempty"`
}

type AzurermPostgresqlFirewallRuleSpec struct {
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	ServerName        string `json:"server_name"`
	StartIpAddress    string `json:"start_ip_address"`
	EndIpAddress      string `json:"end_ip_address"`
}



type AzurermPostgresqlFirewallRuleStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermPostgresqlFirewallRuleList is a list of AzurermPostgresqlFirewallRules
type AzurermPostgresqlFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermPostgresqlFirewallRule CRD objects
	Items []AzurermPostgresqlFirewallRule `json:"items,omitempty"`
}