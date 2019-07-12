package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermEventhubNamespace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermEventhubNamespaceSpec   `json:"spec,omitempty"`
	Status            AzurermEventhubNamespaceStatus `json:"status,omitempty"`
}

type AzurermEventhubNamespaceSpec struct {
	Location                         string            `json:"location"`
	ResourceGroupName                string            `json:"resource_group_name"`
	KafkaEnabled                     bool              `json:"kafka_enabled"`
	MaximumThroughputUnits           int               `json:"maximum_throughput_units"`
	DefaultPrimaryConnectionString   string            `json:"default_primary_connection_string"`
	Tags                             map[string]string `json:"tags"`
	DefaultSecondaryKey              string            `json:"default_secondary_key"`
	Name                             string            `json:"name"`
	Sku                              string            `json:"sku"`
	Capacity                         int               `json:"capacity"`
	AutoInflateEnabled               bool              `json:"auto_inflate_enabled"`
	DefaultSecondaryConnectionString string            `json:"default_secondary_connection_string"`
	DefaultPrimaryKey                string            `json:"default_primary_key"`
}

type AzurermEventhubNamespaceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermEventhubNamespaceList is a list of AzurermEventhubNamespaces
type AzurermEventhubNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermEventhubNamespace CRD objects
	Items []AzurermEventhubNamespace `json:"items,omitempty"`
}
