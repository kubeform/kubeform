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

type AzurermMonitorMetricAlertrule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermMonitorMetricAlertruleSpec   `json:"spec,omitempty"`
	Status            AzurermMonitorMetricAlertruleStatus `json:"status,omitempty"`
}

type AzurermMonitorMetricAlertruleSpecWebhookAction struct {
	ServiceUri string            `json:"service_uri"`
	Properties map[string]string `json:"properties"`
}

type AzurermMonitorMetricAlertruleSpecEmailAction struct {
	SendToServiceOwners bool     `json:"send_to_service_owners"`
	CustomEmails        []string `json:"custom_emails"`
}

type AzurermMonitorMetricAlertruleSpec struct {
	Description       string                              `json:"description"`
	Enabled           bool                                `json:"enabled"`
	MetricName        string                              `json:"metric_name"`
	Threshold         float64                             `json:"threshold"`
	Period            string                              `json:"period"`
	Name              string                              `json:"name"`
	Location          string                              `json:"location"`
	Aggregation       string                              `json:"aggregation"`
	ResourceGroupName string                              `json:"resource_group_name"`
	WebhookAction     []AzurermMonitorMetricAlertruleSpec `json:"webhook_action"`
	Tags              map[string]string                   `json:"tags"`
	ResourceId        string                              `json:"resource_id"`
	Operator          string                              `json:"operator"`
	EmailAction       []AzurermMonitorMetricAlertruleSpec `json:"email_action"`
}

type AzurermMonitorMetricAlertruleStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermMonitorMetricAlertruleList is a list of AzurermMonitorMetricAlertrules
type AzurermMonitorMetricAlertruleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermMonitorMetricAlertrule CRD objects
	Items []AzurermMonitorMetricAlertrule `json:"items,omitempty"`
}
