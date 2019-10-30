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

type LambdaAlias struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LambdaAliasSpec   `json:"spec,omitempty"`
	Status            LambdaAliasStatus `json:"status,omitempty"`
}

type LambdaAliasSpecRoutingConfig struct {
	// +optional
	AdditionalVersionWeights map[string]json.Number `json:"additionalVersionWeights,omitempty" tf:"additional_version_weights,omitempty"`
}

type LambdaAliasSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	Description     string `json:"description,omitempty" tf:"description,omitempty"`
	FunctionName    string `json:"functionName" tf:"function_name"`
	FunctionVersion string `json:"functionVersion" tf:"function_version"`
	// +optional
	InvokeArn string `json:"invokeArn,omitempty" tf:"invoke_arn,omitempty"`
	Name      string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	RoutingConfig []LambdaAliasSpecRoutingConfig `json:"routingConfig,omitempty" tf:"routing_config,omitempty"`
}

type LambdaAliasStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LambdaAliasSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LambdaAliasList is a list of LambdaAliass
type LambdaAliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LambdaAlias CRD objects
	Items []LambdaAlias `json:"items,omitempty"`
}
