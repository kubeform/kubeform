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

type AwsCloudwatchLogMetricFilter struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsCloudwatchLogMetricFilterSpec   `json:"spec,omitempty"`
	Status            AwsCloudwatchLogMetricFilterStatus `json:"status,omitempty"`
}

type AwsCloudwatchLogMetricFilterSpecMetricTransformation struct {
	Name         string `json:"name"`
	Namespace    string `json:"namespace"`
	Value        string `json:"value"`
	DefaultValue string `json:"default_value"`
}

type AwsCloudwatchLogMetricFilterSpec struct {
	Name                 string                             `json:"name"`
	Pattern              string                             `json:"pattern"`
	LogGroupName         string                             `json:"log_group_name"`
	MetricTransformation []AwsCloudwatchLogMetricFilterSpec `json:"metric_transformation"`
}

type AwsCloudwatchLogMetricFilterStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsCloudwatchLogMetricFilterList is a list of AwsCloudwatchLogMetricFilters
type AwsCloudwatchLogMetricFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsCloudwatchLogMetricFilter CRD objects
	Items []AwsCloudwatchLogMetricFilter `json:"items,omitempty"`
}
