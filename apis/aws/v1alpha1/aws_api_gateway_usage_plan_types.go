package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AwsApiGatewayUsagePlan struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsApiGatewayUsagePlanSpec   `json:"spec,omitempty"`
	Status            AwsApiGatewayUsagePlanStatus `json:"status,omitempty"`
}

type AwsApiGatewayUsagePlanSpecApiStages struct {
	Stage string `json:"stage"`
	ApiId string `json:"api_id"`
}

type AwsApiGatewayUsagePlanSpecQuotaSettings struct {
	Period string `json:"period"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}

type AwsApiGatewayUsagePlanSpecThrottleSettings struct {
	BurstLimit int     `json:"burst_limit"`
	RateLimit  float64 `json:"rate_limit"`
}

type AwsApiGatewayUsagePlanSpec struct {
	Description      string                       `json:"description"`
	ApiStages        []AwsApiGatewayUsagePlanSpec `json:"api_stages"`
	QuotaSettings    []AwsApiGatewayUsagePlanSpec `json:"quota_settings"`
	ThrottleSettings []AwsApiGatewayUsagePlanSpec `json:"throttle_settings"`
	ProductCode      string                       `json:"product_code"`
	Name             string                       `json:"name"`
}



type AwsApiGatewayUsagePlanStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsApiGatewayUsagePlanList is a list of AwsApiGatewayUsagePlans
type AwsApiGatewayUsagePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsApiGatewayUsagePlan CRD objects
	Items []AwsApiGatewayUsagePlan `json:"items,omitempty"`
}