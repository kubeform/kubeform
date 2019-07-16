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

type CosmosdbAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CosmosdbAccountSpec   `json:"spec,omitempty"`
	Status            CosmosdbAccountStatus `json:"status,omitempty"`
}

type CosmosdbAccountSpecCapabilities struct {
	Name string `json:"name"`
}

type CosmosdbAccountSpecConsistencyPolicy struct {
	ConsistencyLevel string `json:"consistency_level"`
	// +optional
	MaxIntervalInSeconds int `json:"max_interval_in_seconds,omitempty"`
	// +optional
	MaxStalenessPrefix int `json:"max_staleness_prefix,omitempty"`
}

type CosmosdbAccountSpecFailoverPolicy struct {
	Location string `json:"location"`
	Priority int    `json:"priority"`
}

type CosmosdbAccountSpecVirtualNetworkRule struct {
	Id string `json:"id"`
}

type CosmosdbAccountSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Capabilities *[]CosmosdbAccountSpec `json:"capabilities,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	ConsistencyPolicy []CosmosdbAccountSpec `json:"consistency_policy"`
	// +optional
	EnableAutomaticFailover bool `json:"enable_automatic_failover,omitempty"`
	// +optional
	EnableMultipleWriteLocations bool `json:"enable_multiple_write_locations,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	// Deprecated
	FailoverPolicy *[]CosmosdbAccountSpec `json:"failover_policy,omitempty"`
	// +optional
	IpRangeFilter string `json:"ip_range_filter,omitempty"`
	// +optional
	IsVirtualNetworkFilterEnabled bool `json:"is_virtual_network_filter_enabled,omitempty"`
	// +optional
	Kind              string `json:"kind,omitempty"`
	Location          string `json:"location"`
	Name              string `json:"name"`
	OfferType         string `json:"offer_type"`
	ResourceGroupName string `json:"resource_group_name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	VirtualNetworkRule *[]CosmosdbAccountSpec `json:"virtual_network_rule,omitempty"`
}

type CosmosdbAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CosmosdbAccountList is a list of CosmosdbAccounts
type CosmosdbAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CosmosdbAccount CRD objects
	Items []CosmosdbAccount `json:"items,omitempty"`
}
