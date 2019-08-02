package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type EventhubConsumerGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventhubConsumerGroupSpec   `json:"spec,omitempty"`
	Status            EventhubConsumerGroupStatus `json:"status,omitempty"`
}

type EventhubConsumerGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	EventhubName string `json:"eventhubName" tf:"eventhub_name"`
	// +optional
	// Deprecated
	Location          string `json:"location,omitempty" tf:"location,omitempty"`
	Name              string `json:"name" tf:"name"`
	NamespaceName     string `json:"namespaceName" tf:"namespace_name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	UserMetadata string `json:"userMetadata,omitempty" tf:"user_metadata,omitempty"`
}

type EventhubConsumerGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *EventhubConsumerGroupSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EventhubConsumerGroupList is a list of EventhubConsumerGroups
type EventhubConsumerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EventhubConsumerGroup CRD objects
	Items []EventhubConsumerGroup `json:"items,omitempty"`
}
