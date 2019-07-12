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

type AzurermMetricAlertrule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermMetricAlertruleSpec   `json:"spec,omitempty"`
	Status            AzurermMetricAlertruleStatus `json:"status,omitempty"`
}

type AzurermMetricAlertruleSpecWebhookAction struct {
	ServiceUri string            `json:"service_uri"`
	Properties map[string]string `json:"properties"`
}

type AzurermMetricAlertruleSpecEmailAction struct {
	SendToServiceOwners bool     `json:"send_to_service_owners"`
	CustomEmails        []string `json:"custom_emails"`
}

type AzurermMetricAlertruleSpec struct {
	WebhookAction     []AzurermMetricAlertruleSpec `json:"webhook_action"`
	Operator          string                       `json:"operator"`
	EmailAction       []AzurermMetricAlertruleSpec `json:"email_action"`
	MetricName        string                       `json:"metric_name"`
	Aggregation       string                       `json:"aggregation"`
	Description       string                       `json:"description"`
	ResourceId        string                       `json:"resource_id"`
	Threshold         float64                      `json:"threshold"`
	ResourceGroupName string                       `json:"resource_group_name"`
	Enabled           bool                         `json:"enabled"`
	Period            string                       `json:"period"`
	Tags              map[string]string            `json:"tags"`
	Name              string                       `json:"name"`
	Location          string                       `json:"location"`
}

type AzurermMetricAlertruleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermMetricAlertruleList is a list of AzurermMetricAlertrules
type AzurermMetricAlertruleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermMetricAlertrule CRD objects
	Items []AzurermMetricAlertrule `json:"items,omitempty"`
}
