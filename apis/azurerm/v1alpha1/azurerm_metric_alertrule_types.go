package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermMetricAlertrule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermMetricAlertruleSpec   `json:"spec,omitempty"`
	Status            AzurermMetricAlertruleStatus `json:"status,omitempty"`
}

type AzurermMetricAlertruleSpecEmailAction struct {
	SendToServiceOwners bool     `json:"send_to_service_owners"`
	CustomEmails        []string `json:"custom_emails"`
}

type AzurermMetricAlertruleSpecWebhookAction struct {
	ServiceUri string            `json:"service_uri"`
	Properties map[string]string `json:"properties"`
}

type AzurermMetricAlertruleSpec struct {
	MetricName        string                       `json:"metric_name"`
	Name              string                       `json:"name"`
	ResourceGroupName string                       `json:"resource_group_name"`
	Description       string                       `json:"description"`
	Period            string                       `json:"period"`
	ResourceId        string                       `json:"resource_id"`
	Operator          string                       `json:"operator"`
	Threshold         float64                      `json:"threshold"`
	Aggregation       string                       `json:"aggregation"`
	EmailAction       []AzurermMetricAlertruleSpec `json:"email_action"`
	Tags              map[string]string            `json:"tags"`
	Location          string                       `json:"location"`
	Enabled           bool                         `json:"enabled"`
	WebhookAction     []AzurermMetricAlertruleSpec `json:"webhook_action"`
}

type AzurermMetricAlertruleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermMetricAlertruleList is a list of AzurermMetricAlertrules
type AzurermMetricAlertruleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermMetricAlertrule CRD objects
	Items []AzurermMetricAlertrule `json:"items,omitempty"`
}
