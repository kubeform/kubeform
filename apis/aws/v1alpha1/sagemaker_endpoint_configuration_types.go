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
	"encoding/json"

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

type SagemakerEndpointConfiguration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SagemakerEndpointConfigurationSpec   `json:"spec,omitempty"`
	Status            SagemakerEndpointConfigurationStatus `json:"status,omitempty"`
}

type SagemakerEndpointConfigurationSpecProductionVariants struct {
	// +optional
	AcceleratorType      string `json:"acceleratorType,omitempty" tf:"accelerator_type,omitempty"`
	InitialInstanceCount int    `json:"initialInstanceCount" tf:"initial_instance_count"`
	// +optional
	InitialVariantWeight json.Number `json:"initialVariantWeight,omitempty" tf:"initial_variant_weight,omitempty"`
	InstanceType         string      `json:"instanceType" tf:"instance_type"`
	ModelName            string      `json:"modelName" tf:"model_name"`
	// +optional
	VariantName string `json:"variantName,omitempty" tf:"variant_name,omitempty"`
}

type SagemakerEndpointConfigurationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	KmsKeyArn string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
	// +optional
	Name               string                                                 `json:"name,omitempty" tf:"name,omitempty"`
	ProductionVariants []SagemakerEndpointConfigurationSpecProductionVariants `json:"productionVariants" tf:"production_variants"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SagemakerEndpointConfigurationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SagemakerEndpointConfigurationSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SagemakerEndpointConfigurationList is a list of SagemakerEndpointConfigurations
type SagemakerEndpointConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SagemakerEndpointConfiguration CRD objects
	Items []SagemakerEndpointConfiguration `json:"items,omitempty"`
}
