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

type ApiGatewayUsagePlan struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayUsagePlanSpec   `json:"spec,omitempty"`
	Status            ApiGatewayUsagePlanStatus `json:"status,omitempty"`
}

type ApiGatewayUsagePlanSpecApiStages struct {
	ApiID string `json:"apiID" tf:"api_id"`
	Stage string `json:"stage" tf:"stage"`
}

type ApiGatewayUsagePlanSpecQuotaSettings struct {
	Limit int `json:"limit" tf:"limit"`
	// +optional
	Offset int    `json:"offset,omitempty" tf:"offset,omitempty"`
	Period string `json:"period" tf:"period"`
}

type ApiGatewayUsagePlanSpecThrottleSettings struct {
	// +optional
	BurstLimit int `json:"burstLimit,omitempty" tf:"burst_limit,omitempty"`
	// +optional
	RateLimit json.Number `json:"rateLimit,omitempty" tf:"rate_limit,omitempty"`
}

type ApiGatewayUsagePlanSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	ApiStages []ApiGatewayUsagePlanSpecApiStages `json:"apiStages,omitempty" tf:"api_stages,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +optional
	ProductCode string `json:"productCode,omitempty" tf:"product_code,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	QuotaSettings []ApiGatewayUsagePlanSpecQuotaSettings `json:"quotaSettings,omitempty" tf:"quota_settings,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ThrottleSettings []ApiGatewayUsagePlanSpecThrottleSettings `json:"throttleSettings,omitempty" tf:"throttle_settings,omitempty"`
}

type ApiGatewayUsagePlanStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ApiGatewayUsagePlanSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiGatewayUsagePlanList is a list of ApiGatewayUsagePlans
type ApiGatewayUsagePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiGatewayUsagePlan CRD objects
	Items []ApiGatewayUsagePlan `json:"items,omitempty"`
}
