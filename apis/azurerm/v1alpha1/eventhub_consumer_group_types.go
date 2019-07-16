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

type EventhubConsumerGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventhubConsumerGroupSpec   `json:"spec,omitempty"`
	Status            EventhubConsumerGroupStatus `json:"status,omitempty"`
}

type EventhubConsumerGroupSpec struct {
	EventhubName string `json:"eventhub_name"`
	// +optional
	// Deprecated
	Location          string `json:"location,omitempty"`
	Name              string `json:"name"`
	NamespaceName     string `json:"namespace_name"`
	ResourceGroupName string `json:"resource_group_name"`
	// +optional
	UserMetadata string `json:"user_metadata,omitempty"`
}

type EventhubConsumerGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
