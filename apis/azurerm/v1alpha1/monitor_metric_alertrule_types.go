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

type MonitorMetricAlertrule struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorMetricAlertruleSpec   `json:"spec,omitempty"`
	Status            MonitorMetricAlertruleStatus `json:"status,omitempty"`
}

type MonitorMetricAlertruleSpec struct {
	Aggregation string `json:"aggregation"`
	// +optional
	Enabled           bool        `json:"enabled,omitempty"`
	Location          string      `json:"location"`
	MetricName        string      `json:"metric_name"`
	Name              string      `json:"name"`
	Operator          string      `json:"operator"`
	Period            string      `json:"period"`
	ResourceGroupName string      `json:"resource_group_name"`
	ResourceId        string      `json:"resource_id"`
	Threshold         json.Number `json:"threshold"`
}

type MonitorMetricAlertruleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MonitorMetricAlertruleList is a list of MonitorMetricAlertrules
type MonitorMetricAlertruleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MonitorMetricAlertrule CRD objects
	Items []MonitorMetricAlertrule `json:"items,omitempty"`
}