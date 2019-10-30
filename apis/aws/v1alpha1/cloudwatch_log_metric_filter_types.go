/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

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

	ID string `json:"id,omitempty" tf:"id,omitempty"`

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
	// +optional
	Output *CloudwatchLogMetricFilterSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
