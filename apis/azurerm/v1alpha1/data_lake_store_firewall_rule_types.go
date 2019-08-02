package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type DataLakeStoreFirewallRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataLakeStoreFirewallRuleSpec   `json:"spec,omitempty"`
	Status            DataLakeStoreFirewallRuleStatus `json:"status,omitempty"`
}

type DataLakeStoreFirewallRuleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	AccountName       string `json:"accountName" tf:"account_name"`
	EndIPAddress      string `json:"endIPAddress" tf:"end_ip_address"`
	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	StartIPAddress    string `json:"startIPAddress" tf:"start_ip_address"`
}

type DataLakeStoreFirewallRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *DataLakeStoreFirewallRuleSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DataLakeStoreFirewallRuleList is a list of DataLakeStoreFirewallRules
type DataLakeStoreFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DataLakeStoreFirewallRule CRD objects
	Items []DataLakeStoreFirewallRule `json:"items,omitempty"`
}
