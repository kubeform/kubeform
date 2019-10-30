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

type SagemakerModel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SagemakerModelSpec   `json:"spec,omitempty"`
	Status            SagemakerModelStatus `json:"status,omitempty"`
}

type SagemakerModelSpecContainer struct {
	// +optional
	ContainerHostname string `json:"containerHostname,omitempty" tf:"container_hostname,omitempty"`
	// +optional
	Environment map[string]string `json:"environment,omitempty" tf:"environment,omitempty"`
	Image       string            `json:"image" tf:"image"`
	// +optional
	ModelDataURL string `json:"modelDataURL,omitempty" tf:"model_data_url,omitempty"`
}

type SagemakerModelSpecPrimaryContainer struct {
	// +optional
	ContainerHostname string `json:"containerHostname,omitempty" tf:"container_hostname,omitempty"`
	// +optional
	Environment map[string]string `json:"environment,omitempty" tf:"environment,omitempty"`
	Image       string            `json:"image" tf:"image"`
	// +optional
	ModelDataURL string `json:"modelDataURL,omitempty" tf:"model_data_url,omitempty"`
}

type SagemakerModelSpecVpcConfig struct {
	SecurityGroupIDS []string `json:"securityGroupIDS" tf:"security_group_ids"`
	Subnets          []string `json:"subnets" tf:"subnets"`
}

type SagemakerModelSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	Container []SagemakerModelSpecContainer `json:"container,omitempty" tf:"container,omitempty"`
	// +optional
	EnableNetworkIsolation bool   `json:"enableNetworkIsolation,omitempty" tf:"enable_network_isolation,omitempty"`
	ExecutionRoleArn       string `json:"executionRoleArn" tf:"execution_role_arn"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	PrimaryContainer []SagemakerModelSpecPrimaryContainer `json:"primaryContainer,omitempty" tf:"primary_container,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	VpcConfig []SagemakerModelSpecVpcConfig `json:"vpcConfig,omitempty" tf:"vpc_config,omitempty"`
}

type SagemakerModelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SagemakerModelSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SagemakerModelList is a list of SagemakerModels
type SagemakerModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SagemakerModel CRD objects
	Items []SagemakerModel `json:"items,omitempty"`
}
