package v1alpha1

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type MonitorMetricAlert struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorMetricAlertSpec   `json:"spec,omitempty"`
	Status            MonitorMetricAlertStatus `json:"status,omitempty"`
}

type MonitorMetricAlertSpecAction struct {
	ActionGroupId string `json:"action_group_id"`
	// +optional
	WebhookProperties map[string]string `json:"webhook_properties,omitempty"`
}

type MonitorMetricAlertSpecCriteriaDimension struct {
	Name     string `json:"name"`
	Operator string `json:"operator"`
	// +kubebuilder:validation:MinItems=1
	Values []string `json:"values"`
}

type MonitorMetricAlertSpecCriteria struct {
	Aggregation string `json:"aggregation"`
	// +optional
	Dimension       *[]MonitorMetricAlertSpecCriteria `json:"dimension,omitempty"`
	MetricName      string                            `json:"metric_name"`
	MetricNamespace string                            `json:"metric_namespace"`
	Operator        string                            `json:"operator"`
	Threshold       json.Number                       `json:"threshold"`
}

type MonitorMetricAlertSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Action *[]MonitorMetricAlertSpec `json:"action,omitempty"`
	// +optional
	AutoMitigate bool `json:"auto_mitigate,omitempty"`
	// +kubebuilder:validation:MinItems=1
	Criteria []MonitorMetricAlertSpec `json:"criteria"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +optional
	Frequency         string `json:"frequency,omitempty"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	Scopes []string `json:"scopes"`
	// +optional
	Severity int `json:"severity,omitempty"`
	// +optional
	WindowSize string `json:"window_size,omitempty"`
}

type MonitorMetricAlertStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MonitorMetricAlertList is a list of MonitorMetricAlerts
type MonitorMetricAlertList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MonitorMetricAlert CRD objects
	Items []MonitorMetricAlert `json:"items,omitempty"`
}
