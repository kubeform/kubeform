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

type ServicebusSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicebusSubscriptionSpec   `json:"spec,omitempty"`
	Status            ServicebusSubscriptionStatus `json:"status,omitempty"`
}

type ServicebusSubscriptionSpec struct {
	// +optional
	// Deprecated
	DeadLetteringOnFilterEvaluationExceptions bool `json:"dead_lettering_on_filter_evaluation_exceptions,omitempty"`
	// +optional
	DeadLetteringOnMessageExpiration bool `json:"dead_lettering_on_message_expiration,omitempty"`
	// +optional
	EnableBatchedOperations bool `json:"enable_batched_operations,omitempty"`
	// +optional
	ForwardTo string `json:"forward_to,omitempty"`
	// +optional
	// Deprecated
	Location         string `json:"location,omitempty"`
	MaxDeliveryCount int    `json:"max_delivery_count"`
	Name             string `json:"name"`
	NamespaceName    string `json:"namespace_name"`
	// +optional
	RequiresSession   bool   `json:"requires_session,omitempty"`
	ResourceGroupName string `json:"resource_group_name"`
	TopicName         string `json:"topic_name"`
}

type ServicebusSubscriptionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServicebusSubscriptionList is a list of ServicebusSubscriptions
type ServicebusSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServicebusSubscription CRD objects
	Items []ServicebusSubscription `json:"items,omitempty"`
}
