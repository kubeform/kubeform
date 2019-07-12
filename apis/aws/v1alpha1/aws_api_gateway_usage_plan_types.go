package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayUsagePlan struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsApiGatewayUsagePlanSpec   `json:"spec,omitempty"`
	Status            AwsApiGatewayUsagePlanStatus `json:"status,omitempty"`
}

type AwsApiGatewayUsagePlanSpecApiStages struct {
	ApiId string `json:"api_id"`
	Stage string `json:"stage"`
}

type AwsApiGatewayUsagePlanSpecQuotaSettings struct {
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
	Period string `json:"period"`
}

type AwsApiGatewayUsagePlanSpecThrottleSettings struct {
	RateLimit  float64 `json:"rate_limit"`
	BurstLimit int     `json:"burst_limit"`
}

type AwsApiGatewayUsagePlanSpec struct {
	ProductCode      string                       `json:"product_code"`
	Name             string                       `json:"name"`
	Description      string                       `json:"description"`
	ApiStages        []AwsApiGatewayUsagePlanSpec `json:"api_stages"`
	QuotaSettings    []AwsApiGatewayUsagePlanSpec `json:"quota_settings"`
	ThrottleSettings []AwsApiGatewayUsagePlanSpec `json:"throttle_settings"`
}



type AwsApiGatewayUsagePlanStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayUsagePlanList is a list of AwsApiGatewayUsagePlans
type AwsApiGatewayUsagePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsApiGatewayUsagePlan CRD objects
	Items []AwsApiGatewayUsagePlan `json:"items,omitempty"`
}