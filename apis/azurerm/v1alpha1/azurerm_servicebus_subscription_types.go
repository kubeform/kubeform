package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermServicebusSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermServicebusSubscriptionSpec   `json:"spec,omitempty"`
	Status            AzurermServicebusSubscriptionStatus `json:"status,omitempty"`
}

type AzurermServicebusSubscriptionSpec struct {
	TopicName                                 string `json:"topic_name"`
	Location                                  string `json:"location"`
	ResourceGroupName                         string `json:"resource_group_name"`
	AutoDeleteOnIdle                          string `json:"auto_delete_on_idle"`
	DefaultMessageTtl                         string `json:"default_message_ttl"`
	MaxDeliveryCount                          int    `json:"max_delivery_count"`
	Name                                      string `json:"name"`
	NamespaceName                             string `json:"namespace_name"`
	LockDuration                              string `json:"lock_duration"`
	EnableBatchedOperations                   bool   `json:"enable_batched_operations"`
	ForwardTo                                 string `json:"forward_to"`
	DeadLetteringOnMessageExpiration          bool   `json:"dead_lettering_on_message_expiration"`
	RequiresSession                           bool   `json:"requires_session"`
	DeadLetteringOnFilterEvaluationExceptions bool   `json:"dead_lettering_on_filter_evaluation_exceptions"`
}



type AzurermServicebusSubscriptionStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermServicebusSubscriptionList is a list of AzurermServicebusSubscriptions
type AzurermServicebusSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermServicebusSubscription CRD objects
	Items []AzurermServicebusSubscription `json:"items,omitempty"`
}