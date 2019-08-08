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

type Iothub struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IothubSpec   `json:"spec,omitempty"`
	Status            IothubStatus `json:"status,omitempty"`
}

type IothubSpecEndpoint struct {
	// +optional
	BatchFrequencyInSeconds int    `json:"batchFrequencyInSeconds,omitempty" tf:"batch_frequency_in_seconds,omitempty"`
	ConnectionString        string `json:"-" sensitive:"true" tf:"connection_string"`
	// +optional
	ContainerName string `json:"containerName,omitempty" tf:"container_name,omitempty"`
	// +optional
	Encoding string `json:"encoding,omitempty" tf:"encoding,omitempty"`
	// +optional
	FileNameFormat string `json:"fileNameFormat,omitempty" tf:"file_name_format,omitempty"`
	// +optional
	MaxChunkSizeInBytes int    `json:"maxChunkSizeInBytes,omitempty" tf:"max_chunk_size_in_bytes,omitempty"`
	Name                string `json:"name" tf:"name"`
	Type                string `json:"type" tf:"type"`
}

type IothubSpecFallbackRoute struct {
	// +optional
	Condition string `json:"condition,omitempty" tf:"condition,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	EndpointNames []string `json:"endpointNames,omitempty" tf:"endpoint_names,omitempty"`
	// +optional
	Source string `json:"source,omitempty" tf:"source,omitempty"`
}

type IothubSpecIpFilterRule struct {
	Action string `json:"action" tf:"action"`
	IpMask string `json:"ipMask" tf:"ip_mask"`
	Name   string `json:"name" tf:"name"`
}

type IothubSpecRoute struct {
	// +optional
	Condition     string   `json:"condition,omitempty" tf:"condition,omitempty"`
	Enabled       bool     `json:"enabled" tf:"enabled"`
	EndpointNames []string `json:"endpointNames" tf:"endpoint_names"`
	Name          string   `json:"name" tf:"name"`
	Source        string   `json:"source" tf:"source"`
}

type IothubSpecSharedAccessPolicy struct {
	// +optional
	KeyName string `json:"keyName,omitempty" tf:"key_name,omitempty"`
	// +optional
	Permissions string `json:"permissions,omitempty" tf:"permissions,omitempty"`
	// +optional
	PrimaryKey string `json:"-" sensitive:"true" tf:"primary_key,omitempty"`
	// +optional
	SecondaryKey string `json:"-" sensitive:"true" tf:"secondary_key,omitempty"`
}

type IothubSpecSku struct {
	Capacity int    `json:"capacity" tf:"capacity"`
	Name     string `json:"name" tf:"name"`
	Tier     string `json:"tier" tf:"tier"`
}

type IothubSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	Endpoint []IothubSpecEndpoint `json:"endpoint,omitempty" tf:"endpoint,omitempty"`
	// +optional
	EventHubEventsEndpoint string `json:"eventHubEventsEndpoint,omitempty" tf:"event_hub_events_endpoint,omitempty"`
	// +optional
	EventHubEventsPath string `json:"eventHubEventsPath,omitempty" tf:"event_hub_events_path,omitempty"`
	// +optional
	EventHubOperationsEndpoint string `json:"eventHubOperationsEndpoint,omitempty" tf:"event_hub_operations_endpoint,omitempty"`
	// +optional
	EventHubOperationsPath string `json:"eventHubOperationsPath,omitempty" tf:"event_hub_operations_path,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	FallbackRoute []IothubSpecFallbackRoute `json:"fallbackRoute,omitempty" tf:"fallback_route,omitempty"`
	// +optional
	Hostname string `json:"hostname,omitempty" tf:"hostname,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	IpFilterRule      []IothubSpecIpFilterRule `json:"ipFilterRule,omitempty" tf:"ip_filter_rule,omitempty"`
	Location          string                   `json:"location" tf:"location"`
	Name              string                   `json:"name" tf:"name"`
	ResourceGroupName string                   `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Route []IothubSpecRoute `json:"route,omitempty" tf:"route,omitempty"`
	// +optional
	SharedAccessPolicy []IothubSpecSharedAccessPolicy `json:"sharedAccessPolicy,omitempty" tf:"shared_access_policy,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Sku []IothubSpecSku `json:"sku" tf:"sku"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type IothubStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *IothubSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IothubList is a list of Iothubs
type IothubList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Iothub CRD objects
	Items []Iothub `json:"items,omitempty"`
}
