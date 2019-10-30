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

type ConfigConfigurationRecorder struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigConfigurationRecorderSpec   `json:"spec,omitempty"`
	Status            ConfigConfigurationRecorderStatus `json:"status,omitempty"`
}

type ConfigConfigurationRecorderSpecRecordingGroup struct {
	// +optional
	AllSupported bool `json:"allSupported,omitempty" tf:"all_supported,omitempty"`
	// +optional
	IncludeGlobalResourceTypes bool `json:"includeGlobalResourceTypes,omitempty" tf:"include_global_resource_types,omitempty"`
	// +optional
	ResourceTypes []string `json:"resourceTypes,omitempty" tf:"resource_types,omitempty"`
}

type ConfigConfigurationRecorderSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RecordingGroup []ConfigConfigurationRecorderSpecRecordingGroup `json:"recordingGroup,omitempty" tf:"recording_group,omitempty"`
	RoleArn        string                                          `json:"roleArn" tf:"role_arn"`
}

type ConfigConfigurationRecorderStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ConfigConfigurationRecorderSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ConfigConfigurationRecorderList is a list of ConfigConfigurationRecorders
type ConfigConfigurationRecorderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ConfigConfigurationRecorder CRD objects
	Items []ConfigConfigurationRecorder `json:"items,omitempty"`
}
