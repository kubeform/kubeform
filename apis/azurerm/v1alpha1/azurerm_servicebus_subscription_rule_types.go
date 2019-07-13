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

type AzurermServicebusSubscriptionRule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermServicebusSubscriptionRuleSpec   `json:"spec,omitempty"`
	Status            AzurermServicebusSubscriptionRuleStatus `json:"status,omitempty"`
}

type AzurermServicebusSubscriptionRuleSpecCorrelationFilter struct {
	Label            string `json:"label"`
	SessionId        string `json:"session_id"`
	ReplyToSessionId string `json:"reply_to_session_id"`
	ContentType      string `json:"content_type"`
	CorrelationId    string `json:"correlation_id"`
	MessageId        string `json:"message_id"`
	To               string `json:"to"`
	ReplyTo          string `json:"reply_to"`
}

type AzurermServicebusSubscriptionRuleSpec struct {
	Name              string                                  `json:"name"`
	SubscriptionName  string                                  `json:"subscription_name"`
	CorrelationFilter []AzurermServicebusSubscriptionRuleSpec `json:"correlation_filter"`
	ResourceGroupName string                                  `json:"resource_group_name"`
	NamespaceName     string                                  `json:"namespace_name"`
	TopicName         string                                  `json:"topic_name"`
	FilterType        string                                  `json:"filter_type"`
	Action            string                                  `json:"action"`
	SqlFilter         string                                  `json:"sql_filter"`
}



type AzurermServicebusSubscriptionRuleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermServicebusSubscriptionRuleList is a list of AzurermServicebusSubscriptionRules
type AzurermServicebusSubscriptionRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermServicebusSubscriptionRule CRD objects
	Items []AzurermServicebusSubscriptionRule `json:"items,omitempty"`
}