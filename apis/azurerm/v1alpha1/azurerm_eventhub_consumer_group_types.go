package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermEventhubConsumerGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermEventhubConsumerGroupSpec   `json:"spec,omitempty"`
	Status            AzurermEventhubConsumerGroupStatus `json:"status,omitempty"`
}

type AzurermEventhubConsumerGroupSpec struct {
	Location          string `json:"location"`
	UserMetadata      string `json:"user_metadata"`
	Name              string `json:"name"`
	NamespaceName     string `json:"namespace_name"`
	EventhubName      string `json:"eventhub_name"`
	ResourceGroupName string `json:"resource_group_name"`
}



type AzurermEventhubConsumerGroupStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermEventhubConsumerGroupList is a list of AzurermEventhubConsumerGroups
type AzurermEventhubConsumerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermEventhubConsumerGroup CRD objects
	Items []AzurermEventhubConsumerGroup `json:"items,omitempty"`
}