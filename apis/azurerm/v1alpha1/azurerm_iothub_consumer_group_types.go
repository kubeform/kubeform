package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermIothubConsumerGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermIothubConsumerGroupSpec   `json:"spec,omitempty"`
	Status            AzurermIothubConsumerGroupStatus `json:"status,omitempty"`
}

type AzurermIothubConsumerGroupSpec struct {
	Name                 string `json:"name"`
	IothubName           string `json:"iothub_name"`
	EventhubEndpointName string `json:"eventhub_endpoint_name"`
	ResourceGroupName    string `json:"resource_group_name"`
}



type AzurermIothubConsumerGroupStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermIothubConsumerGroupList is a list of AzurermIothubConsumerGroups
type AzurermIothubConsumerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermIothubConsumerGroup CRD objects
	Items []AzurermIothubConsumerGroup `json:"items,omitempty"`
}