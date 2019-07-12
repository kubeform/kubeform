package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermMonitorMetricAlertrule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermMonitorMetricAlertruleSpec   `json:"spec,omitempty"`
	Status            AzurermMonitorMetricAlertruleStatus `json:"status,omitempty"`
}

type AzurermMonitorMetricAlertruleSpecEmailAction struct {
	CustomEmails        []string `json:"custom_emails"`
	SendToServiceOwners bool     `json:"send_to_service_owners"`
}

type AzurermMonitorMetricAlertruleSpecWebhookAction struct {
	ServiceUri string            `json:"service_uri"`
	Properties map[string]string `json:"properties"`
}

type AzurermMonitorMetricAlertruleSpec struct {
	Name              string                              `json:"name"`
	ResourceId        string                              `json:"resource_id"`
	Operator          string                              `json:"operator"`
	Location          string                              `json:"location"`
	MetricName        string                              `json:"metric_name"`
	EmailAction       []AzurermMonitorMetricAlertruleSpec `json:"email_action"`
	Tags              map[string]string                   `json:"tags"`
	ResourceGroupName string                              `json:"resource_group_name"`
	Enabled           bool                                `json:"enabled"`
	Period            string                              `json:"period"`
	Description       string                              `json:"description"`
	Threshold         float64                             `json:"threshold"`
	Aggregation       string                              `json:"aggregation"`
	WebhookAction     []AzurermMonitorMetricAlertruleSpec `json:"webhook_action"`
}

type AzurermMonitorMetricAlertruleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermMonitorMetricAlertruleList is a list of AzurermMonitorMetricAlertrules
type AzurermMonitorMetricAlertruleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermMonitorMetricAlertrule CRD objects
	Items []AzurermMonitorMetricAlertrule `json:"items,omitempty"`
}
