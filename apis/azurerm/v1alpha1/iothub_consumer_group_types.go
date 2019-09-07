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

type IothubConsumerGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IothubConsumerGroupSpec   `json:"spec,omitempty"`
	Status            IothubConsumerGroupStatus `json:"status,omitempty"`
}

type IothubConsumerGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	EventhubEndpointName string `json:"eventhubEndpointName" tf:"eventhub_endpoint_name"`
	IothubName           string `json:"iothubName" tf:"iothub_name"`
	Name                 string `json:"name" tf:"name"`
	ResourceGroupName    string `json:"resourceGroupName" tf:"resource_group_name"`
}

type IothubConsumerGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *IothubConsumerGroupSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IothubConsumerGroupList is a list of IothubConsumerGroups
type IothubConsumerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IothubConsumerGroup CRD objects
	Items []IothubConsumerGroup `json:"items,omitempty"`
}