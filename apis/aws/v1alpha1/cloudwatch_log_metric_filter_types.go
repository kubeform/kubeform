package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	DefaultValue string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`
	Name         string `json:"name" tf:"name"`
	Namespace    string `json:"namespace" tf:"namespace"`
	Value        string `json:"value" tf:"value"`
}

type CloudwatchLogMetricFilterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	LogGroupName string `json:"logGroupName" tf:"log_group_name"`
	// +kubebuilder:validation:MaxItems=1
	MetricTransformation []CloudwatchLogMetricFilterSpecMetricTransformation `json:"metricTransformation" tf:"metric_transformation"`
	Name                 string                                              `json:"name" tf:"name"`
	Pattern              string                                              `json:"pattern" tf:"pattern"`
}

type CloudwatchLogMetricFilterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
