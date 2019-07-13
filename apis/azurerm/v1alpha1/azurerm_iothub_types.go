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

type AzurermIothubSpecFallbackRoute struct {
	EndpointNames []string `json:"endpoint_names"`
	Enabled       bool     `json:"enabled"`
	Source        string   `json:"source"`
	Condition     string   `json:"condition"`
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
	Name     string `json:"name"`
	Tier     string `json:"tier"`
	Capacity int    `json:"capacity"`
}

type AzurermIothubSpecSharedAccessPolicy struct {
	KeyName      string `json:"key_name"`
	PrimaryKey   string `json:"primary_key"`
	SecondaryKey string `json:"secondary_key"`
	Permissions  string `json:"permissions"`
}

type AzurermIothubSpecRoute struct {
	Enabled       bool     `json:"enabled"`
	Name          string   `json:"name"`
	Source        string   `json:"source"`
	Condition     string   `json:"condition"`
	EndpointNames []string `json:"endpoint_names"`
}

type AzurermIothubSpec struct {
	Name                       string              `json:"name"`
	Location                   string              `json:"location"`
	ResourceGroupName          string              `json:"resource_group_name"`
	EventHubOperationsPath     string              `json:"event_hub_operations_path"`
	FallbackRoute              []AzurermIothubSpec `json:"fallback_route"`
	IpFilterRule               []AzurermIothubSpec `json:"ip_filter_rule"`
	EventHubEventsEndpoint     string              `json:"event_hub_events_endpoint"`
	EventHubOperationsEndpoint string              `json:"event_hub_operations_endpoint"`
	Endpoint                   []AzurermIothubSpec `json:"endpoint"`
	Tags                       map[string]string   `json:"tags"`
	Sku                        []AzurermIothubSpec `json:"sku"`
	Type                       string              `json:"type"`
	Hostname                   string              `json:"hostname"`
	EventHubEventsPath         string              `json:"event_hub_events_path"`
	SharedAccessPolicy         []AzurermIothubSpec `json:"shared_access_policy"`
	Route                      []AzurermIothubSpec `json:"route"`
}



type AzurermIothubStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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