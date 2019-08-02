package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ServicebusSubscriptionRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicebusSubscriptionRuleSpec   `json:"spec,omitempty"`
	Status            ServicebusSubscriptionRuleStatus `json:"status,omitempty"`
}

type ServicebusSubscriptionRuleSpecCorrelationFilter struct {
	// +optional
	ContentType string `json:"contentType,omitempty" tf:"content_type,omitempty"`
	// +optional
	CorrelationID string `json:"correlationID,omitempty" tf:"correlation_id,omitempty"`
	// +optional
	Label string `json:"label,omitempty" tf:"label,omitempty"`
	// +optional
	MessageID string `json:"messageID,omitempty" tf:"message_id,omitempty"`
	// +optional
	ReplyTo string `json:"replyTo,omitempty" tf:"reply_to,omitempty"`
	// +optional
	ReplyToSessionID string `json:"replyToSessionID,omitempty" tf:"reply_to_session_id,omitempty"`
	// +optional
	SessionID string `json:"sessionID,omitempty" tf:"session_id,omitempty"`
	// +optional
	To string `json:"to,omitempty" tf:"to,omitempty"`
}

type ServicebusSubscriptionRuleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Action string `json:"action,omitempty" tf:"action,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CorrelationFilter []ServicebusSubscriptionRuleSpecCorrelationFilter `json:"correlationFilter,omitempty" tf:"correlation_filter,omitempty"`
	FilterType        string                                            `json:"filterType" tf:"filter_type"`
	Name              string                                            `json:"name" tf:"name"`
	NamespaceName     string                                            `json:"namespaceName" tf:"namespace_name"`
	ResourceGroupName string                                            `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SqlFilter        string `json:"sqlFilter,omitempty" tf:"sql_filter,omitempty"`
	SubscriptionName string `json:"subscriptionName" tf:"subscription_name"`
	TopicName        string `json:"topicName" tf:"topic_name"`
}

type ServicebusSubscriptionRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *ServicebusSubscriptionRuleSpec `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServicebusSubscriptionRuleList is a list of ServicebusSubscriptionRules
type ServicebusSubscriptionRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServicebusSubscriptionRule CRD objects
	Items []ServicebusSubscriptionRule `json:"items,omitempty"`
}
