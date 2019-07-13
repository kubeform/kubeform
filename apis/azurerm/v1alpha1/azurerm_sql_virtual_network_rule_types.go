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

type AzurermSqlVirtualNetworkRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermSqlVirtualNetworkRuleSpec   `json:"spec,omitempty"`
	Status            AzurermSqlVirtualNetworkRuleStatus `json:"status,omitempty"`
}

type AzurermSqlVirtualNetworkRuleSpec struct {
	Name                             string `json:"name"`
	ResourceGroupName                string `json:"resource_group_name"`
	ServerName                       string `json:"server_name"`
	SubnetId                         string `json:"subnet_id"`
	IgnoreMissingVnetServiceEndpoint bool   `json:"ignore_missing_vnet_service_endpoint"`
}



type AzurermSqlVirtualNetworkRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermSqlVirtualNetworkRuleList is a list of AzurermSqlVirtualNetworkRules
type AzurermSqlVirtualNetworkRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermSqlVirtualNetworkRule CRD objects
	Items []AzurermSqlVirtualNetworkRule `json:"items,omitempty"`
}