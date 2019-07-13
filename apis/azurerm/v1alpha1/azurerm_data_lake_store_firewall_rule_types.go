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

type AzurermDataLakeStoreFirewallRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDataLakeStoreFirewallRuleSpec   `json:"spec,omitempty"`
	Status            AzurermDataLakeStoreFirewallRuleStatus `json:"status,omitempty"`
}

type AzurermDataLakeStoreFirewallRuleSpec struct {
	ResourceGroupName string `json:"resource_group_name"`
	StartIpAddress    string `json:"start_ip_address"`
	EndIpAddress      string `json:"end_ip_address"`
	Name              string `json:"name"`
	AccountName       string `json:"account_name"`
}



type AzurermDataLakeStoreFirewallRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermDataLakeStoreFirewallRuleList is a list of AzurermDataLakeStoreFirewallRules
type AzurermDataLakeStoreFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDataLakeStoreFirewallRule CRD objects
	Items []AzurermDataLakeStoreFirewallRule `json:"items,omitempty"`
}