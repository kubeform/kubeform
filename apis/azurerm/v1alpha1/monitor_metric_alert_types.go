package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type MonitorMetricAlert struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorMetricAlertSpec   `json:"spec,omitempty"`
	Status            MonitorMetricAlertStatus `json:"status,omitempty"`
}

type MonitorMetricAlertSpecAction struct {
	ActionGroupID string `json:"actionGroupID" tf:"action_group_id"`
	// +optional
	WebhookProperties map[string]string `json:"webhookProperties,omitempty" tf:"webhook_properties,omitempty"`
}

type MonitorMetricAlertSpecCriteriaDimension struct {
	Name     string `json:"name" tf:"name"`
	Operator string `json:"operator" tf:"operator"`
	// +kubebuilder:validation:MinItems=1
	Values []string `json:"values" tf:"values"`
}

type MonitorMetricAlertSpecCriteria struct {
	Aggregation string `json:"aggregation" tf:"aggregation"`
	// +optional
	Dimension       []MonitorMetricAlertSpecCriteriaDimension `json:"dimension,omitempty" tf:"dimension,omitempty"`
	MetricName      string                                    `json:"metricName" tf:"metric_name"`
	MetricNamespace string                                    `json:"metricNamespace" tf:"metric_namespace"`
	Operator        string                                    `json:"operator" tf:"operator"`
	Threshold       json.Number                               `json:"threshold" tf:"threshold"`
}

type MonitorMetricAlertSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Action []MonitorMetricAlertSpecAction `json:"action,omitempty" tf:"action,omitempty"`
	// +optional
	AutoMitigate bool `json:"autoMitigate,omitempty" tf:"auto_mitigate,omitempty"`
	// +kubebuilder:validation:MinItems=1
	Criteria []MonitorMetricAlertSpecCriteria `json:"criteria" tf:"criteria"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	Frequency         string `json:"frequency,omitempty" tf:"frequency,omitempty"`
	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Scopes []string `json:"scopes" tf:"scopes"`
	// +optional
	Severity int `json:"severity,omitempty" tf:"severity,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	WindowSize string `json:"windowSize,omitempty" tf:"window_size,omitempty"`
}

type MonitorMetricAlertStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *MonitorMetricAlertSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
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
