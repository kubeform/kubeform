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

type ServicebusSubscriptionRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServicebusSubscriptionRuleSpec   `json:"spec,omitempty"`
	Status            ServicebusSubscriptionRuleStatus `json:"status,omitempty"`
}

type ServicebusSubscriptionRuleSpecCorrelationFilter struct {
	// +optional
	ContentType string `json:"content_type,omitempty"`
	// +optional
	CorrelationId string `json:"correlation_id,omitempty"`
	// +optional
	Label string `json:"label,omitempty"`
	// +optional
	MessageId string `json:"message_id,omitempty"`
	// +optional
	ReplyTo string `json:"reply_to,omitempty"`
	// +optional
	ReplyToSessionId string `json:"reply_to_session_id,omitempty"`
	// +optional
	SessionId string `json:"session_id,omitempty"`
	// +optional
	To string `json:"to,omitempty"`
}

type ServicebusSubscriptionRuleSpec struct {
	// +optional
	Action string `json:"action,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CorrelationFilter *[]ServicebusSubscriptionRuleSpec `json:"correlation_filter,omitempty"`
	FilterType        string                            `json:"filter_type"`
	Name              string                            `json:"name"`
	NamespaceName     string                            `json:"namespace_name"`
	ResourceGroupName string                            `json:"resource_group_name"`
	// +optional
	SqlFilter        string `json:"sql_filter,omitempty"`
	SubscriptionName string `json:"subscription_name"`
	TopicName        string `json:"topic_name"`
}

type ServicebusSubscriptionRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
