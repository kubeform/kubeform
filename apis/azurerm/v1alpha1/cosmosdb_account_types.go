package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
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
	Name string `json:"name" tf:"name"`
}

type CosmosdbAccountSpecConsistencyPolicy struct {
	ConsistencyLevel string `json:"consistencyLevel" tf:"consistency_level"`
	// +optional
	MaxIntervalInSeconds int `json:"maxIntervalInSeconds,omitempty" tf:"max_interval_in_seconds,omitempty"`
	// +optional
	MaxStalenessPrefix int `json:"maxStalenessPrefix,omitempty" tf:"max_staleness_prefix,omitempty"`
}

type CosmosdbAccountSpecFailoverPolicy struct {
	// +optional
	ID       string `json:"ID,omitempty" tf:"id,omitempty"`
	Location string `json:"location" tf:"location"`
	Priority int    `json:"priority" tf:"priority"`
}

type CosmosdbAccountSpecGeoLocation struct {
	FailoverPriority int `json:"failoverPriority" tf:"failover_priority"`
	// +optional
	ID       string `json:"ID,omitempty" tf:"id,omitempty"`
	Location string `json:"location" tf:"location"`
	// +optional
	Prefix string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type CosmosdbAccountSpecVirtualNetworkRule struct {
	ID string `json:"ID" tf:"id"`
}

type CosmosdbAccountSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	Capabilities []CosmosdbAccountSpecCapabilities `json:"capabilities,omitempty" tf:"capabilities,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ConsistencyPolicy []CosmosdbAccountSpecConsistencyPolicy `json:"consistencyPolicy" tf:"consistency_policy"`
	// +optional
	EnableAutomaticFailover bool `json:"enableAutomaticFailover,omitempty" tf:"enable_automatic_failover,omitempty"`
	// +optional
	EnableMultipleWriteLocations bool `json:"enableMultipleWriteLocations,omitempty" tf:"enable_multiple_write_locations,omitempty"`
	// +optional
	Endpoint string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`
	// +optional
	// Deprecated
	FailoverPolicy []CosmosdbAccountSpecFailoverPolicy `json:"failoverPolicy,omitempty" tf:"failover_policy,omitempty"`
	// +optional
	GeoLocation []CosmosdbAccountSpecGeoLocation `json:"geoLocation,omitempty" tf:"geo_location,omitempty"`
	// +optional
	IpRangeFilter string `json:"ipRangeFilter,omitempty" tf:"ip_range_filter,omitempty"`
	// +optional
	IsVirtualNetworkFilterEnabled bool `json:"isVirtualNetworkFilterEnabled,omitempty" tf:"is_virtual_network_filter_enabled,omitempty"`
	// +optional
	Kind      string `json:"kind,omitempty" tf:"kind,omitempty"`
	Location  string `json:"location" tf:"location"`
	Name      string `json:"name" tf:"name"`
	OfferType string `json:"offerType" tf:"offer_type"`
	// +optional
	PrimaryMasterKey string `json:"-" sensitive:"true" tf:"primary_master_key,omitempty"`
	// +optional
	PrimaryReadonlyMasterKey string `json:"-" sensitive:"true" tf:"primary_readonly_master_key,omitempty"`
	// +optional
	ReadEndpoints     []string `json:"readEndpoints,omitempty" tf:"read_endpoints,omitempty"`
	ResourceGroupName string   `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecondaryMasterKey string `json:"-" sensitive:"true" tf:"secondary_master_key,omitempty"`
	// +optional
	SecondaryReadonlyMasterKey string `json:"-" sensitive:"true" tf:"secondary_readonly_master_key,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	VirtualNetworkRule []CosmosdbAccountSpecVirtualNetworkRule `json:"virtualNetworkRule,omitempty" tf:"virtual_network_rule,omitempty"`
	// +optional
	WriteEndpoints []string `json:"writeEndpoints,omitempty" tf:"write_endpoints,omitempty"`
}

type CosmosdbAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CosmosdbAccountSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
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
