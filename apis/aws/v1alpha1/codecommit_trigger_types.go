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

type CodecommitTrigger struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodecommitTriggerSpec   `json:"spec,omitempty"`
	Status            CodecommitTriggerStatus `json:"status,omitempty"`
}

type CodecommitTriggerSpecTrigger struct {
	// +optional
	Branches []string `json:"branches,omitempty" tf:"branches,omitempty"`
	// +optional
	CustomData     string   `json:"customData,omitempty" tf:"custom_data,omitempty"`
	DestinationArn string   `json:"destinationArn" tf:"destination_arn"`
	Events         []string `json:"events" tf:"events"`
	Name           string   `json:"name" tf:"name"`
}

type CodecommitTriggerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ConfigurationID string `json:"configurationID,omitempty" tf:"configuration_id,omitempty"`
	RepositoryName  string `json:"repositoryName" tf:"repository_name"`
	// +kubebuilder:validation:MaxItems=10
	Trigger []CodecommitTriggerSpecTrigger `json:"trigger" tf:"trigger"`
}

type CodecommitTriggerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CodecommitTriggerSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CodecommitTriggerList is a list of CodecommitTriggers
type CodecommitTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CodecommitTrigger CRD objects
	Items []CodecommitTrigger `json:"items,omitempty"`
}
