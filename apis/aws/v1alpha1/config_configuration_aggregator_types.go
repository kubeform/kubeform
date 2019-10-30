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

type ConfigConfigurationAggregator struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConfigConfigurationAggregatorSpec   `json:"spec,omitempty"`
	Status            ConfigConfigurationAggregatorStatus `json:"status,omitempty"`
}

type ConfigConfigurationAggregatorSpecAccountAggregationSource struct {
	// +kubebuilder:validation:MinItems=1
	AccountIDS []string `json:"accountIDS" tf:"account_ids"`
	// +optional
	AllRegions bool `json:"allRegions,omitempty" tf:"all_regions,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	Regions []string `json:"regions,omitempty" tf:"regions,omitempty"`
}

type ConfigConfigurationAggregatorSpecOrganizationAggregationSource struct {
	// +optional
	AllRegions bool `json:"allRegions,omitempty" tf:"all_regions,omitempty"`
	// +optional
	// +kubebuilder:validation:MinItems=1
	Regions []string `json:"regions,omitempty" tf:"regions,omitempty"`
	RoleArn string   `json:"roleArn" tf:"role_arn"`
}

type ConfigConfigurationAggregatorSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	AccountAggregationSource []ConfigConfigurationAggregatorSpecAccountAggregationSource `json:"accountAggregationSource,omitempty" tf:"account_aggregation_source,omitempty"`
	// +optional
	Arn  string `json:"arn,omitempty" tf:"arn,omitempty"`
	Name string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	OrganizationAggregationSource []ConfigConfigurationAggregatorSpecOrganizationAggregationSource `json:"organizationAggregationSource,omitempty" tf:"organization_aggregation_source,omitempty"`
}

type ConfigConfigurationAggregatorStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ConfigConfigurationAggregatorSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ConfigConfigurationAggregatorList is a list of ConfigConfigurationAggregators
type ConfigConfigurationAggregatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ConfigConfigurationAggregator CRD objects
	Items []ConfigConfigurationAggregator `json:"items,omitempty"`
}
