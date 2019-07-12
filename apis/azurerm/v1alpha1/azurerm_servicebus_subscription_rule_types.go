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
	MessageId        string `json:"message_id"`
	To               string `json:"to"`
	ReplyTo          string `json:"reply_to"`
	Label            string `json:"label"`
	SessionId        string `json:"session_id"`
	ReplyToSessionId string `json:"reply_to_session_id"`
	ContentType      string `json:"content_type"`
	CorrelationId    string `json:"correlation_id"`
}

type AzurermServicebusSubscriptionRuleSpec struct {
	ResourceGroupName string                                  `json:"resource_group_name"`
	Action            string                                  `json:"action"`
	SqlFilter         string                                  `json:"sql_filter"`
	CorrelationFilter []AzurermServicebusSubscriptionRuleSpec `json:"correlation_filter"`
	Name              string                                  `json:"name"`
	NamespaceName     string                                  `json:"namespace_name"`
	TopicName         string                                  `json:"topic_name"`
	SubscriptionName  string                                  `json:"subscription_name"`
	FilterType        string                                  `json:"filter_type"`
}

type AzurermServicebusSubscriptionRuleStatus struct {
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
