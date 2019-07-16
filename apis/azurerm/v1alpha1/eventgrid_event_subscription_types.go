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

type EventgridEventSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventgridEventSubscriptionSpec   `json:"spec,omitempty"`
	Status            EventgridEventSubscriptionStatus `json:"status,omitempty"`
}

type EventgridEventSubscriptionSpecEventhubEndpoint struct {
	EventhubId string `json:"eventhub_id"`
}

type EventgridEventSubscriptionSpecHybridConnectionEndpoint struct {
	HybridConnectionId string `json:"hybrid_connection_id"`
}

type EventgridEventSubscriptionSpecStorageBlobDeadLetterDestination struct {
	StorageAccountId         string `json:"storage_account_id"`
	StorageBlobContainerName string `json:"storage_blob_container_name"`
}

type EventgridEventSubscriptionSpecStorageQueueEndpoint struct {
	QueueName        string `json:"queue_name"`
	StorageAccountId string `json:"storage_account_id"`
}

type EventgridEventSubscriptionSpecSubjectFilter struct {
	// +optional
	CaseSensitive bool `json:"case_sensitive,omitempty"`
	// +optional
	SubjectBeginsWith string `json:"subject_begins_with,omitempty"`
	// +optional
	SubjectEndsWith string `json:"subject_ends_with,omitempty"`
}

type EventgridEventSubscriptionSpecWebhookEndpoint struct {
	Url string `json:"url"`
}

type EventgridEventSubscriptionSpec struct {
	// +optional
	EventDeliverySchema string `json:"event_delivery_schema,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EventhubEndpoint *[]EventgridEventSubscriptionSpec `json:"eventhub_endpoint,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HybridConnectionEndpoint *[]EventgridEventSubscriptionSpec `json:"hybrid_connection_endpoint,omitempty"`
	// +optional
	Labels []string `json:"labels,omitempty"`
	Name   string   `json:"name"`
	Scope  string   `json:"scope"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	StorageBlobDeadLetterDestination *[]EventgridEventSubscriptionSpec `json:"storage_blob_dead_letter_destination,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	StorageQueueEndpoint *[]EventgridEventSubscriptionSpec `json:"storage_queue_endpoint,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SubjectFilter *[]EventgridEventSubscriptionSpec `json:"subject_filter,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	WebhookEndpoint *[]EventgridEventSubscriptionSpec `json:"webhook_endpoint,omitempty"`
}

type EventgridEventSubscriptionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EventgridEventSubscriptionList is a list of EventgridEventSubscriptions
type EventgridEventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EventgridEventSubscription CRD objects
	Items []EventgridEventSubscription `json:"items,omitempty"`
}
