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

type AzurermIothub struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermIothubSpec   `json:"spec,omitempty"`
	Status            AzurermIothubStatus `json:"status,omitempty"`
}

type AzurermIothubSpecRoute struct {
	Source        string   `json:"source"`
	Condition     string   `json:"condition"`
	EndpointNames []string `json:"endpoint_names"`
	Enabled       bool     `json:"enabled"`
	Name          string   `json:"name"`
}

type AzurermIothubSpecIpFilterRule struct {
	Name   string `json:"name"`
	IpMask string `json:"ip_mask"`
	Action string `json:"action"`
}

type AzurermIothubSpecEndpoint struct {
	ContainerName           string `json:"container_name"`
	Encoding                string `json:"encoding"`
	FileNameFormat          string `json:"file_name_format"`
	Type                    string `json:"type"`
	ConnectionString        string `json:"connection_string"`
	Name                    string `json:"name"`
	BatchFrequencyInSeconds int    `json:"batch_frequency_in_seconds"`
	MaxChunkSizeInBytes     int    `json:"max_chunk_size_in_bytes"`
}

type AzurermIothubSpecSku struct {
	Tier     string `json:"tier"`
	Capacity int    `json:"capacity"`
	Name     string `json:"name"`
}

type AzurermIothubSpecSharedAccessPolicy struct {
	PrimaryKey   string `json:"primary_key"`
	SecondaryKey string `json:"secondary_key"`
	Permissions  string `json:"permissions"`
	KeyName      string `json:"key_name"`
}

type AzurermIothubSpecFallbackRoute struct {
	Source        string   `json:"source"`
	Condition     string   `json:"condition"`
	EndpointNames []string `json:"endpoint_names"`
	Enabled       bool     `json:"enabled"`
}

type AzurermIothubSpec struct {
	EventHubEventsEndpoint     string              `json:"event_hub_events_endpoint"`
	Route                      []AzurermIothubSpec `json:"route"`
	IpFilterRule               []AzurermIothubSpec `json:"ip_filter_rule"`
	ResourceGroupName          string              `json:"resource_group_name"`
	Type                       string              `json:"type"`
	EventHubOperationsEndpoint string              `json:"event_hub_operations_endpoint"`
	EventHubEventsPath         string              `json:"event_hub_events_path"`
	Endpoint                   []AzurermIothubSpec `json:"endpoint"`
	Name                       string              `json:"name"`
	Sku                        []AzurermIothubSpec `json:"sku"`
	Tags                       map[string]string   `json:"tags"`
	Location                   string              `json:"location"`
	Hostname                   string              `json:"hostname"`
	EventHubOperationsPath     string              `json:"event_hub_operations_path"`
	SharedAccessPolicy         []AzurermIothubSpec `json:"shared_access_policy"`
	FallbackRoute              []AzurermIothubSpec `json:"fallback_route"`
}

type AzurermIothubStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermIothubList is a list of AzurermIothubs
type AzurermIothubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermIothub CRD objects
	Items []AzurermIothub `json:"items,omitempty"`
}
