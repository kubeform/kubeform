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

type AzurermDataLakeAnalyticsFirewallRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDataLakeAnalyticsFirewallRuleSpec   `json:"spec,omitempty"`
	Status            AzurermDataLakeAnalyticsFirewallRuleStatus `json:"status,omitempty"`
}

type AzurermDataLakeAnalyticsFirewallRuleSpec struct {
	ResourceGroupName string `json:"resource_group_name"`
	StartIpAddress    string `json:"start_ip_address"`
	EndIpAddress      string `json:"end_ip_address"`
	Name              string `json:"name"`
	AccountName       string `json:"account_name"`
}



type AzurermDataLakeAnalyticsFirewallRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermDataLakeAnalyticsFirewallRuleList is a list of AzurermDataLakeAnalyticsFirewallRules
type AzurermDataLakeAnalyticsFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDataLakeAnalyticsFirewallRule CRD objects
	Items []AzurermDataLakeAnalyticsFirewallRule `json:"items,omitempty"`
}