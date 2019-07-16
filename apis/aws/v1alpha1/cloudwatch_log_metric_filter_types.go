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

type CloudwatchLogMetricFilter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudwatchLogMetricFilterSpec   `json:"spec,omitempty"`
	Status            CloudwatchLogMetricFilterStatus `json:"status,omitempty"`
}

type CloudwatchLogMetricFilterSpecMetricTransformation struct {
	// +optional
	DefaultValue string `json:"default_value,omitempty"`
	Name         string `json:"name"`
	Namespace    string `json:"namespace"`
	Value        string `json:"value"`
}

type CloudwatchLogMetricFilterSpec struct {
	LogGroupName string `json:"log_group_name"`
	// +kubebuilder:validation:MaxItems=1
	MetricTransformation []CloudwatchLogMetricFilterSpec `json:"metric_transformation"`
	Name                 string                          `json:"name"`
	Pattern              string                          `json:"pattern"`
}

type CloudwatchLogMetricFilterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudwatchLogMetricFilterList is a list of CloudwatchLogMetricFilters
type CloudwatchLogMetricFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudwatchLogMetricFilter CRD objects
	Items []CloudwatchLogMetricFilter `json:"items,omitempty"`
}
