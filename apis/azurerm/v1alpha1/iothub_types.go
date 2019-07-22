package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
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
	BatchFrequencyInSeconds int `json:"batchFrequencyInSeconds,omitempty" tf:"batch_frequency_in_seconds,omitempty"`
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

type IothubSpecSku struct {
	Capacity int    `json:"capacity" tf:"capacity"`
	Name     string `json:"name" tf:"name"`
	Tier     string `json:"tier" tf:"tier"`
}

type IothubSpec struct {
	Secret      *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`
	ProviderRef core.LocalObjectReference  `json:"providerRef" tf:"-"`
	// +optional
	Endpoint []IothubSpecEndpoint `json:"endpoint,omitempty" tf:"endpoint,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	FallbackRoute []IothubSpecFallbackRoute `json:"fallbackRoute,omitempty" tf:"fallback_route,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	IpFilterRule      []IothubSpecIpFilterRule `json:"ipFilterRule,omitempty" tf:"ip_filter_rule,omitempty"`
	Location          string                   `json:"location" tf:"location"`
	Name              string                   `json:"name" tf:"name"`
	ResourceGroupName string                   `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Route []IothubSpecRoute `json:"route,omitempty" tf:"route,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Sku []IothubSpecSku `json:"sku" tf:"sku"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type IothubStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
