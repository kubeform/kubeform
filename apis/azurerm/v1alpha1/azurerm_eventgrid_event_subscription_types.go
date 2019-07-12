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

type AzurermEventgridEventSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermEventgridEventSubscriptionSpec   `json:"spec,omitempty"`
	Status            AzurermEventgridEventSubscriptionStatus `json:"status,omitempty"`
}

type AzurermEventgridEventSubscriptionSpecStorageBlobDeadLetterDestination struct {
	StorageAccountId         string `json:"storage_account_id"`
	StorageBlobContainerName string `json:"storage_blob_container_name"`
}

type AzurermEventgridEventSubscriptionSpecEventhubEndpoint struct {
	EventhubId string `json:"eventhub_id"`
}

type AzurermEventgridEventSubscriptionSpecHybridConnectionEndpoint struct {
	HybridConnectionId string `json:"hybrid_connection_id"`
}

type AzurermEventgridEventSubscriptionSpecWebhookEndpoint struct {
	Url string `json:"url"`
}

type AzurermEventgridEventSubscriptionSpecStorageQueueEndpoint struct {
	StorageAccountId string `json:"storage_account_id"`
	QueueName        string `json:"queue_name"`
}

type AzurermEventgridEventSubscriptionSpecSubjectFilter struct {
	SubjectBeginsWith string `json:"subject_begins_with"`
	SubjectEndsWith   string `json:"subject_ends_with"`
	CaseSensitive     bool   `json:"case_sensitive"`
}

type AzurermEventgridEventSubscriptionSpecRetryPolicy struct {
	MaxDeliveryAttempts int `json:"max_delivery_attempts"`
	EventTimeToLive     int `json:"event_time_to_live"`
}

type AzurermEventgridEventSubscriptionSpec struct {
	IncludedEventTypes               []string                                `json:"included_event_types"`
	StorageBlobDeadLetterDestination []AzurermEventgridEventSubscriptionSpec `json:"storage_blob_dead_letter_destination"`
	Name                             string                                  `json:"name"`
	EventDeliverySchema              string                                  `json:"event_delivery_schema"`
	TopicName                        string                                  `json:"topic_name"`
	EventhubEndpoint                 []AzurermEventgridEventSubscriptionSpec `json:"eventhub_endpoint"`
	HybridConnectionEndpoint         []AzurermEventgridEventSubscriptionSpec `json:"hybrid_connection_endpoint"`
	WebhookEndpoint                  []AzurermEventgridEventSubscriptionSpec `json:"webhook_endpoint"`
	Labels                           []string                                `json:"labels"`
	Scope                            string                                  `json:"scope"`
	StorageQueueEndpoint             []AzurermEventgridEventSubscriptionSpec `json:"storage_queue_endpoint"`
	SubjectFilter                    []AzurermEventgridEventSubscriptionSpec `json:"subject_filter"`
	RetryPolicy                      []AzurermEventgridEventSubscriptionSpec `json:"retry_policy"`
}

type AzurermEventgridEventSubscriptionStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermEventgridEventSubscriptionList is a list of AzurermEventgridEventSubscriptions
type AzurermEventgridEventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermEventgridEventSubscription CRD objects
	Items []AzurermEventgridEventSubscription `json:"items,omitempty"`
}
