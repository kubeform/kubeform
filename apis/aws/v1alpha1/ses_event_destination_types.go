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

type SesEventDestination struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SesEventDestinationSpec   `json:"spec,omitempty"`
	Status            SesEventDestinationStatus `json:"status,omitempty"`
}

type SesEventDestinationSpecCloudwatchDestination struct {
	DefaultValue  string `json:"defaultValue" tf:"default_value"`
	DimensionName string `json:"dimensionName" tf:"dimension_name"`
	ValueSource   string `json:"valueSource" tf:"value_source"`
}

type SesEventDestinationSpecKinesisDestination struct {
	RoleArn   string `json:"roleArn" tf:"role_arn"`
	StreamArn string `json:"streamArn" tf:"stream_arn"`
}

type SesEventDestinationSpecSnsDestination struct {
	TopicArn string `json:"topicArn" tf:"topic_arn"`
}

type SesEventDestinationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CloudwatchDestination []SesEventDestinationSpecCloudwatchDestination `json:"cloudwatchDestination,omitempty" tf:"cloudwatch_destination,omitempty"`
	ConfigurationSetName  string                                         `json:"configurationSetName" tf:"configuration_set_name"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	KinesisDestination []SesEventDestinationSpecKinesisDestination `json:"kinesisDestination,omitempty" tf:"kinesis_destination,omitempty"`
	MatchingTypes      []string                                    `json:"matchingTypes" tf:"matching_types"`
	Name               string                                      `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SnsDestination []SesEventDestinationSpecSnsDestination `json:"snsDestination,omitempty" tf:"sns_destination,omitempty"`
}

type SesEventDestinationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SesEventDestinationSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SesEventDestinationList is a list of SesEventDestinations
type SesEventDestinationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SesEventDestination CRD objects
	Items []SesEventDestination `json:"items,omitempty"`
}
