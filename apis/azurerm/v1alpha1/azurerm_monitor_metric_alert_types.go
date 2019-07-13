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

type AzurermMonitorMetricAlert struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermMonitorMetricAlertSpec   `json:"spec,omitempty"`
	Status            AzurermMonitorMetricAlertStatus `json:"status,omitempty"`
}

type AzurermMonitorMetricAlertSpecAction struct {
	ActionGroupId     string            `json:"action_group_id"`
	WebhookProperties map[string]string `json:"webhook_properties"`
}

type AzurermMonitorMetricAlertSpecCriteriaDimension struct {
	Name     string   `json:"name"`
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

type AzurermMonitorMetricAlertSpecCriteria struct {
	Threshold       float64                                 `json:"threshold"`
	Dimension       []AzurermMonitorMetricAlertSpecCriteria `json:"dimension"`
	MetricNamespace string                                  `json:"metric_namespace"`
	MetricName      string                                  `json:"metric_name"`
	Aggregation     string                                  `json:"aggregation"`
	Operator        string                                  `json:"operator"`
}

type AzurermMonitorMetricAlertSpec struct {
	Severity          int                             `json:"severity"`
	Name              string                          `json:"name"`
	Scopes            []string                        `json:"scopes"`
	Action            []AzurermMonitorMetricAlertSpec `json:"action"`
	AutoMitigate      bool                            `json:"auto_mitigate"`
	Description       string                          `json:"description"`
	Tags              map[string]string               `json:"tags"`
	ResourceGroupName string                          `json:"resource_group_name"`
	Criteria          []AzurermMonitorMetricAlertSpec `json:"criteria"`
	Enabled           bool                            `json:"enabled"`
	Frequency         string                          `json:"frequency"`
	WindowSize        string                          `json:"window_size"`
}



type AzurermMonitorMetricAlertStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermMonitorMetricAlertList is a list of AzurermMonitorMetricAlerts
type AzurermMonitorMetricAlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermMonitorMetricAlert CRD objects
	Items []AzurermMonitorMetricAlert `json:"items,omitempty"`
}