package v1alpha1

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
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
	ApiId string `json:"api_id"`
	Stage string `json:"stage"`
}

type ApiGatewayUsagePlanSpecQuotaSettings struct {
	Limit int `json:"limit"`
	// +optional
	Offset int    `json:"offset,omitempty"`
	Period string `json:"period"`
}

type ApiGatewayUsagePlanSpecThrottleSettings struct {
	// +optional
	BurstLimit int `json:"burst_limit,omitempty"`
	// +optional
	RateLimit json.Number `json:"rate_limit,omitempty"`
}

type ApiGatewayUsagePlanSpec struct {
	// +optional
	ApiStages *[]ApiGatewayUsagePlanSpec `json:"api_stages,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	// +optional
	ProductCode string `json:"product_code,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	QuotaSettings *[]ApiGatewayUsagePlanSpec `json:"quota_settings,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	ThrottleSettings *[]ApiGatewayUsagePlanSpec `json:"throttle_settings,omitempty"`
}

type ApiGatewayUsagePlanStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
