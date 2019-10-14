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
// +kubebuilder:subresource:status

type EventhubNamespace_ struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventhubNamespace_Spec   `json:"spec,omitempty"`
	Status            EventhubNamespace_Status `json:"status,omitempty"`
}

type EventhubNamespace_Spec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	AutoInflateEnabled bool `json:"autoInflateEnabled,omitempty" tf:"auto_inflate_enabled,omitempty"`
	// +optional
	Capacity int `json:"capacity,omitempty" tf:"capacity,omitempty"`
	// +optional
	DefaultPrimaryConnectionString string `json:"-" sensitive:"true" tf:"default_primary_connection_string,omitempty"`
	// +optional
	DefaultPrimaryKey string `json:"-" sensitive:"true" tf:"default_primary_key,omitempty"`
	// +optional
	DefaultSecondaryConnectionString string `json:"-" sensitive:"true" tf:"default_secondary_connection_string,omitempty"`
	// +optional
	DefaultSecondaryKey string `json:"-" sensitive:"true" tf:"default_secondary_key,omitempty"`
	// +optional
	KafkaEnabled bool   `json:"kafkaEnabled,omitempty" tf:"kafka_enabled,omitempty"`
	Location     string `json:"location" tf:"location"`
	// +optional
	MaximumThroughputUnits int    `json:"maximumThroughputUnits,omitempty" tf:"maximum_throughput_units,omitempty"`
	Name                   string `json:"name" tf:"name"`
	ResourceGroupName      string `json:"resourceGroupName" tf:"resource_group_name"`
	Sku                    string `json:"sku" tf:"sku"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type EventhubNamespace_Status struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *EventhubNamespace_Spec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EventhubNamespace_List is a list of EventhubNamespace_s
type EventhubNamespace_List struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EventhubNamespace_ CRD objects
	Items []EventhubNamespace_ `json:"items,omitempty"`
}
