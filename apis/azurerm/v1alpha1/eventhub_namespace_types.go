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

type EventhubNamespace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventhubNamespaceSpec   `json:"spec,omitempty"`
	Status            EventhubNamespaceStatus `json:"status,omitempty"`
}

type EventhubNamespaceSpec struct {
	// +optional
	AutoInflateEnabled bool `json:"autoInflateEnabled,omitempty" tf:"auto_inflate_enabled,omitempty"`
	// +optional
	Capacity int `json:"capacity,omitempty" tf:"capacity,omitempty"`
	// +optional
	KafkaEnabled bool   `json:"kafkaEnabled,omitempty" tf:"kafka_enabled,omitempty"`
	Location     string `json:"location" tf:"location"`
	// +optional
	MaximumThroughputUnits int    `json:"maximumThroughputUnits,omitempty" tf:"maximum_throughput_units,omitempty"`
	Name                   string `json:"name" tf:"name"`
	ResourceGroupName      string `json:"resourceGroupName" tf:"resource_group_name"`
	Sku                    string `json:"sku" tf:"sku"`
	// +optional
	Tags        map[string]string         `json:"tags,omitempty" tf:"tags,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type EventhubNamespaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EventhubNamespaceList is a list of EventhubNamespaces
type EventhubNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EventhubNamespace CRD objects
	Items []EventhubNamespace `json:"items,omitempty"`
}
