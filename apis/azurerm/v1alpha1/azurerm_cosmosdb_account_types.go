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

type AzurermCosmosdbAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermCosmosdbAccountSpec   `json:"spec,omitempty"`
	Status            AzurermCosmosdbAccountStatus `json:"status,omitempty"`
}

type AzurermCosmosdbAccountSpecVirtualNetworkRule struct {
	Id string `json:"id"`
}

type AzurermCosmosdbAccountSpecFailoverPolicy struct {
	Id       string `json:"id"`
	Location string `json:"location"`
	Priority int    `json:"priority"`
}

type AzurermCosmosdbAccountSpecConsistencyPolicy struct {
	MaxIntervalInSeconds int    `json:"max_interval_in_seconds"`
	MaxStalenessPrefix   int    `json:"max_staleness_prefix"`
	ConsistencyLevel     string `json:"consistency_level"`
}

type AzurermCosmosdbAccountSpecGeoLocation struct {
	Prefix           string `json:"prefix"`
	Id               string `json:"id"`
	Location         string `json:"location"`
	FailoverPriority int    `json:"failover_priority"`
}

type AzurermCosmosdbAccountSpecCapabilities struct {
	Name string `json:"name"`
}

type AzurermCosmosdbAccountSpec struct {
	VirtualNetworkRule            []AzurermCosmosdbAccountSpec `json:"virtual_network_rule"`
	SecondaryMasterKey            string                       `json:"secondary_master_key"`
	SecondaryReadonlyMasterKey    string                       `json:"secondary_readonly_master_key"`
	IpRangeFilter                 string                       `json:"ip_range_filter"`
	FailoverPolicy                []AzurermCosmosdbAccountSpec `json:"failover_policy"`
	Kind                          string                       `json:"kind"`
	WriteEndpoints                []string                     `json:"write_endpoints"`
	PrimaryMasterKey              string                       `json:"primary_master_key"`
	Name                          string                       `json:"name"`
	Location                      string                       `json:"location"`
	EnableAutomaticFailover       bool                         `json:"enable_automatic_failover"`
	ConsistencyPolicy             []AzurermCosmosdbAccountSpec `json:"consistency_policy"`
	GeoLocation                   []AzurermCosmosdbAccountSpec `json:"geo_location"`
	EnableMultipleWriteLocations  bool                         `json:"enable_multiple_write_locations"`
	ResourceGroupName             string                       `json:"resource_group_name"`
	Tags                          map[string]string            `json:"tags"`
	IsVirtualNetworkFilterEnabled bool                         `json:"is_virtual_network_filter_enabled"`
	Endpoint                      string                       `json:"endpoint"`
	ReadEndpoints                 []string                     `json:"read_endpoints"`
	PrimaryReadonlyMasterKey      string                       `json:"primary_readonly_master_key"`
	ConnectionStrings             []string                     `json:"connection_strings"`
	OfferType                     string                       `json:"offer_type"`
	Capabilities                  []AzurermCosmosdbAccountSpec `json:"capabilities"`
}



type AzurermCosmosdbAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermCosmosdbAccountList is a list of AzurermCosmosdbAccounts
type AzurermCosmosdbAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermCosmosdbAccount CRD objects
	Items []AzurermCosmosdbAccount `json:"items,omitempty"`
}