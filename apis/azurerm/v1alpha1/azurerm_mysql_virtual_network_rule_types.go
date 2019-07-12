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

type AzurermMysqlVirtualNetworkRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermMysqlVirtualNetworkRuleSpec   `json:"spec,omitempty"`
	Status            AzurermMysqlVirtualNetworkRuleStatus `json:"status,omitempty"`
}

type AzurermMysqlVirtualNetworkRuleSpec struct {
	ResourceGroupName string `json:"resource_group_name"`
	ServerName        string `json:"server_name"`
	SubnetId          string `json:"subnet_id"`
	Name              string `json:"name"`
}

type AzurermMysqlVirtualNetworkRuleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermMysqlVirtualNetworkRuleList is a list of AzurermMysqlVirtualNetworkRules
type AzurermMysqlVirtualNetworkRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermMysqlVirtualNetworkRule CRD objects
	Items []AzurermMysqlVirtualNetworkRule `json:"items,omitempty"`
}
