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
	Period            string                              `json:"period"`
	MetricName        string                              `json:"metric_name"`
	Operator          string                              `json:"operator"`
	Threshold         float64                             `json:"threshold"`
	Aggregation       string                              `json:"aggregation"`
	WebhookAction     []AzurermMonitorMetricAlertruleSpec `json:"webhook_action"`
	ResourceGroupName string                              `json:"resource_group_name"`
	Enabled           bool                                `json:"enabled"`
	EmailAction       []AzurermMonitorMetricAlertruleSpec `json:"email_action"`
	Tags              map[string]string                   `json:"tags"`
	Name              string                              `json:"name"`
	Location          string                              `json:"location"`
	ResourceId        string                              `json:"resource_id"`
}



type AzurermMonitorMetricAlertruleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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