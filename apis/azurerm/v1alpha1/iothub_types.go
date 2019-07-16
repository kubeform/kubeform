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

type Iothub struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IothubSpec   `json:"spec,omitempty"`
	Status            IothubStatus `json:"status,omitempty"`
}

type IothubSpecEndpoint struct {
	// +optional
	BatchFrequencyInSeconds int    `json:"batch_frequency_in_seconds,omitempty"`
	ConnectionString        string `json:"connection_string"`
	// +optional
	ContainerName string `json:"container_name,omitempty"`
	// +optional
	Encoding string `json:"encoding,omitempty"`
	// +optional
	FileNameFormat string `json:"file_name_format,omitempty"`
	// +optional
	MaxChunkSizeInBytes int    `json:"max_chunk_size_in_bytes,omitempty"`
	Name                string `json:"name"`
	Type                string `json:"type"`
}

type IothubSpecFileUpload struct {
	ConnectionString string `json:"connection_string"`
	ContainerName    string `json:"container_name"`
	// +optional
	MaxDeliveryCount int `json:"max_delivery_count,omitempty"`
	// +optional
	Notifications bool `json:"notifications,omitempty"`
}

type IothubSpecIpFilterRule struct {
	Action string `json:"action"`
	IpMask string `json:"ip_mask"`
	Name   string `json:"name"`
}

type IothubSpecRoute struct {
	// +optional
	Condition     string   `json:"condition,omitempty"`
	Enabled       bool     `json:"enabled"`
	EndpointNames []string `json:"endpoint_names"`
	Name          string   `json:"name"`
	Source        string   `json:"source"`
}

type IothubSpecSku struct {
	Capacity int    `json:"capacity"`
	Name     string `json:"name"`
	Tier     string `json:"tier"`
}

type IothubSpec struct {
	// +optional
	Endpoint *[]IothubSpec `json:"endpoint,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	FileUpload *[]IothubSpec `json:"file_upload,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	IpFilterRule      *[]IothubSpec `json:"ip_filter_rule,omitempty"`
	Location          string        `json:"location"`
	Name              string        `json:"name"`
	ResourceGroupName string        `json:"resource_group_name"`
	// +optional
	Route *[]IothubSpec `json:"route,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	Sku []IothubSpec `json:"sku"`
}

type IothubStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
