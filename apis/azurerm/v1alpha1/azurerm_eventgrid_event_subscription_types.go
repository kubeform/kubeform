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

type AzurermEventgridEventSubscriptionSpecWebhookEndpoint struct {
	Url string `json:"url"`
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

type AzurermEventgridEventSubscriptionSpecStorageQueueEndpoint struct {
	StorageAccountId string `json:"storage_account_id"`
	QueueName        string `json:"queue_name"`
}

type AzurermEventgridEventSubscriptionSpecStorageBlobDeadLetterDestination struct {
	StorageBlobContainerName string `json:"storage_blob_container_name"`
	StorageAccountId         string `json:"storage_account_id"`
}

type AzurermEventgridEventSubscriptionSpecEventhubEndpoint struct {
	EventhubId string `json:"eventhub_id"`
}

type AzurermEventgridEventSubscriptionSpecHybridConnectionEndpoint struct {
	HybridConnectionId string `json:"hybrid_connection_id"`
}

type AzurermEventgridEventSubscriptionSpec struct {
	WebhookEndpoint                  []AzurermEventgridEventSubscriptionSpec `json:"webhook_endpoint"`
	SubjectFilter                    []AzurermEventgridEventSubscriptionSpec `json:"subject_filter"`
	RetryPolicy                      []AzurermEventgridEventSubscriptionSpec `json:"retry_policy"`
	Name                             string                                  `json:"name"`
	EventDeliverySchema              string                                  `json:"event_delivery_schema"`
	TopicName                        string                                  `json:"topic_name"`
	StorageQueueEndpoint             []AzurermEventgridEventSubscriptionSpec `json:"storage_queue_endpoint"`
	StorageBlobDeadLetterDestination []AzurermEventgridEventSubscriptionSpec `json:"storage_blob_dead_letter_destination"`
	Labels                           []string                                `json:"labels"`
	Scope                            string                                  `json:"scope"`
	EventhubEndpoint                 []AzurermEventgridEventSubscriptionSpec `json:"eventhub_endpoint"`
	HybridConnectionEndpoint         []AzurermEventgridEventSubscriptionSpec `json:"hybrid_connection_endpoint"`
	IncludedEventTypes               []string                                `json:"included_event_types"`
}



type AzurermEventgridEventSubscriptionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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