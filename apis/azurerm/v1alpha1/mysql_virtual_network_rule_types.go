package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type MysqlVirtualNetworkRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MysqlVirtualNetworkRuleSpec   `json:"spec,omitempty"`
	Status            MysqlVirtualNetworkRuleStatus `json:"status,omitempty"`
}

type MysqlVirtualNetworkRuleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	ServerName        string `json:"serverName" tf:"server_name"`
	SubnetID          string `json:"subnetID" tf:"subnet_id"`
}

type MysqlVirtualNetworkRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *MysqlVirtualNetworkRuleSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MysqlVirtualNetworkRuleList is a list of MysqlVirtualNetworkRules
type MysqlVirtualNetworkRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MysqlVirtualNetworkRule CRD objects
	Items []MysqlVirtualNetworkRule `json:"items,omitempty"`
}
