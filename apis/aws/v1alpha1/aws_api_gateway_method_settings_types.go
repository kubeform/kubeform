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

type AwsApiGatewayMethodSettings struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsApiGatewayMethodSettingsSpec   `json:"spec,omitempty"`
	Status            AwsApiGatewayMethodSettingsStatus `json:"status,omitempty"`
}

type AwsApiGatewayMethodSettingsSpecSettings struct {
	CacheTtlInSeconds                      int     `json:"cache_ttl_in_seconds"`
	MetricsEnabled                         bool    `json:"metrics_enabled"`
	LoggingLevel                           string  `json:"logging_level"`
	DataTraceEnabled                       bool    `json:"data_trace_enabled"`
	ThrottlingBurstLimit                   int     `json:"throttling_burst_limit"`
	CachingEnabled                         bool    `json:"caching_enabled"`
	ThrottlingRateLimit                    float64 `json:"throttling_rate_limit"`
	CacheDataEncrypted                     bool    `json:"cache_data_encrypted"`
	RequireAuthorizationForCacheControl    bool    `json:"require_authorization_for_cache_control"`
	UnauthorizedCacheControlHeaderStrategy string  `json:"unauthorized_cache_control_header_strategy"`
}

type AwsApiGatewayMethodSettingsSpec struct {
	Settings   []AwsApiGatewayMethodSettingsSpec `json:"settings"`
	RestApiId  string                            `json:"rest_api_id"`
	StageName  string                            `json:"stage_name"`
	MethodPath string                            `json:"method_path"`
}

type AwsApiGatewayMethodSettingsStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsApiGatewayMethodSettingsList is a list of AwsApiGatewayMethodSettingss
type AwsApiGatewayMethodSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsApiGatewayMethodSettings CRD objects
	Items []AwsApiGatewayMethodSettings `json:"items,omitempty"`
}
