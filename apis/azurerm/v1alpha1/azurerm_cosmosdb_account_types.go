package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermCosmosdbAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermCosmosdbAccountSpec   `json:"spec,omitempty"`
	Status            AzurermCosmosdbAccountStatus `json:"status,omitempty"`
}

type AzurermCosmosdbAccountSpecGeoLocation struct {
	Prefix           string `json:"prefix"`
	Id               string `json:"id"`
	Location         string `json:"location"`
	FailoverPriority int    `json:"failover_priority"`
}

type AzurermCosmosdbAccountSpecFailoverPolicy struct {
	Id       string `json:"id"`
	Location string `json:"location"`
	Priority int    `json:"priority"`
}

type AzurermCosmosdbAccountSpecCapabilities struct {
	Name string `json:"name"`
}

type AzurermCosmosdbAccountSpecVirtualNetworkRule struct {
	Id string `json:"id"`
}

type AzurermCosmosdbAccountSpecConsistencyPolicy struct {
	ConsistencyLevel     string `json:"consistency_level"`
	MaxIntervalInSeconds int    `json:"max_interval_in_seconds"`
	MaxStalenessPrefix   int    `json:"max_staleness_prefix"`
}

type AzurermCosmosdbAccountSpec struct {
	Location                      string                       `json:"location"`
	Kind                          string                       `json:"kind"`
	GeoLocation                   []AzurermCosmosdbAccountSpec `json:"geo_location"`
	Endpoint                      string                       `json:"endpoint"`
	SecondaryMasterKey            string                       `json:"secondary_master_key"`
	Name                          string                       `json:"name"`
	ResourceGroupName             string                       `json:"resource_group_name"`
	FailoverPolicy                []AzurermCosmosdbAccountSpec `json:"failover_policy"`
	PrimaryMasterKey              string                       `json:"primary_master_key"`
	PrimaryReadonlyMasterKey      string                       `json:"primary_readonly_master_key"`
	ConnectionStrings             []string                     `json:"connection_strings"`
	EnableAutomaticFailover       bool                         `json:"enable_automatic_failover"`
	Capabilities                  []AzurermCosmosdbAccountSpec `json:"capabilities"`
	IsVirtualNetworkFilterEnabled bool                         `json:"is_virtual_network_filter_enabled"`
	VirtualNetworkRule            []AzurermCosmosdbAccountSpec `json:"virtual_network_rule"`
	EnableMultipleWriteLocations  bool                         `json:"enable_multiple_write_locations"`
	SecondaryReadonlyMasterKey    string                       `json:"secondary_readonly_master_key"`
	Tags                          map[string]string            `json:"tags"`
	OfferType                     string                       `json:"offer_type"`
	IpRangeFilter                 string                       `json:"ip_range_filter"`
	ConsistencyPolicy             []AzurermCosmosdbAccountSpec `json:"consistency_policy"`
	ReadEndpoints                 []string                     `json:"read_endpoints"`
	WriteEndpoints                []string                     `json:"write_endpoints"`
}



type AzurermCosmosdbAccountStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermCosmosdbAccountList is a list of AzurermCosmosdbAccounts
type AzurermCosmosdbAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermCosmosdbAccount CRD objects
	Items []AzurermCosmosdbAccount `json:"items,omitempty"`
}