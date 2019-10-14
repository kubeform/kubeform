package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type ApiGatewayMethodSettings struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayMethodSettingsSpec   `json:"spec,omitempty"`
	Status            ApiGatewayMethodSettingsStatus `json:"status,omitempty"`
}

type ApiGatewayMethodSettingsSpecSettings struct {
	// +optional
	CacheDataEncrypted bool `json:"cacheDataEncrypted,omitempty" tf:"cache_data_encrypted,omitempty"`
	// +optional
	CacheTtlInSeconds int `json:"cacheTtlInSeconds,omitempty" tf:"cache_ttl_in_seconds,omitempty"`
	// +optional
	CachingEnabled bool `json:"cachingEnabled,omitempty" tf:"caching_enabled,omitempty"`
	// +optional
	DataTraceEnabled bool `json:"dataTraceEnabled,omitempty" tf:"data_trace_enabled,omitempty"`
	// +optional
	LoggingLevel string `json:"loggingLevel,omitempty" tf:"logging_level,omitempty"`
	// +optional
	MetricsEnabled bool `json:"metricsEnabled,omitempty" tf:"metrics_enabled,omitempty"`
	// +optional
	RequireAuthorizationForCacheControl bool `json:"requireAuthorizationForCacheControl,omitempty" tf:"require_authorization_for_cache_control,omitempty"`
	// +optional
	ThrottlingBurstLimit int `json:"throttlingBurstLimit,omitempty" tf:"throttling_burst_limit,omitempty"`
	// +optional
	ThrottlingRateLimit json.Number `json:"throttlingRateLimit,omitempty" tf:"throttling_rate_limit,omitempty"`
	// +optional
	UnauthorizedCacheControlHeaderStrategy string `json:"unauthorizedCacheControlHeaderStrategy,omitempty" tf:"unauthorized_cache_control_header_strategy,omitempty"`
}

type ApiGatewayMethodSettingsSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	MethodPath string `json:"methodPath" tf:"method_path"`
	RestAPIID  string `json:"restAPIID" tf:"rest_api_id"`
	// +kubebuilder:validation:MaxItems=1
	Settings  []ApiGatewayMethodSettingsSpecSettings `json:"settings" tf:"settings"`
	StageName string                                 `json:"stageName" tf:"stage_name"`
}

type ApiGatewayMethodSettingsStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ApiGatewayMethodSettingsSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiGatewayMethodSettingsList is a list of ApiGatewayMethodSettingss
type ApiGatewayMethodSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiGatewayMethodSettings CRD objects
	Items []ApiGatewayMethodSettings `json:"items,omitempty"`
}
