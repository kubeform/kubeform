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

type PostgresqlVirtualNetworkRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PostgresqlVirtualNetworkRuleSpec   `json:"spec,omitempty"`
	Status            PostgresqlVirtualNetworkRuleStatus `json:"status,omitempty"`
}

type PostgresqlVirtualNetworkRuleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	IgnoreMissingVnetServiceEndpoint bool   `json:"ignoreMissingVnetServiceEndpoint,omitempty" tf:"ignore_missing_vnet_service_endpoint,omitempty"`
	Name                             string `json:"name" tf:"name"`
	ResourceGroupName                string `json:"resourceGroupName" tf:"resource_group_name"`
	ServerName                       string `json:"serverName" tf:"server_name"`
	SubnetID                         string `json:"subnetID" tf:"subnet_id"`
}

type PostgresqlVirtualNetworkRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *PostgresqlVirtualNetworkRuleSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PostgresqlVirtualNetworkRuleList is a list of PostgresqlVirtualNetworkRules
type PostgresqlVirtualNetworkRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PostgresqlVirtualNetworkRule CRD objects
	Items []PostgresqlVirtualNetworkRule `json:"items,omitempty"`
}
