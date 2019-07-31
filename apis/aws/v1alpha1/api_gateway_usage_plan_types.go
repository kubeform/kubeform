package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

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

	// +optional
	ApiStages []ApiGatewayUsagePlanSpecApiStages `json:"apiStages,omitempty" tf:"api_stages,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +optional
	ProductCode string `json:"productCode,omitempty" tf:"product_code,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	QuotaSettings []ApiGatewayUsagePlanSpecQuotaSettings `json:"quotaSettings,omitempty" tf:"quota_settings,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	ThrottleSettings []ApiGatewayUsagePlanSpecThrottleSettings `json:"throttleSettings,omitempty" tf:"throttle_settings,omitempty"`
}

type ApiGatewayUsagePlanStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
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
