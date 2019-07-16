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

type ApiGatewayMethodSettings struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayMethodSettingsSpec   `json:"spec,omitempty"`
	Status            ApiGatewayMethodSettingsStatus `json:"status,omitempty"`
}

type ApiGatewayMethodSettingsSpecSettings struct {
	// +optional
	CacheDataEncrypted bool `json:"cache_data_encrypted,omitempty"`
	// +optional
	CacheTtlInSeconds int `json:"cache_ttl_in_seconds,omitempty"`
	// +optional
	CachingEnabled bool `json:"caching_enabled,omitempty"`
	// +optional
	DataTraceEnabled bool `json:"data_trace_enabled,omitempty"`
	// +optional
	LoggingLevel string `json:"logging_level,omitempty"`
	// +optional
	MetricsEnabled bool `json:"metrics_enabled,omitempty"`
	// +optional
	RequireAuthorizationForCacheControl bool `json:"require_authorization_for_cache_control,omitempty"`
	// +optional
	ThrottlingBurstLimit int `json:"throttling_burst_limit,omitempty"`
	// +optional
	ThrottlingRateLimit json.Number `json:"throttling_rate_limit,omitempty"`
	// +optional
	UnauthorizedCacheControlHeaderStrategy string `json:"unauthorized_cache_control_header_strategy,omitempty"`
}

type ApiGatewayMethodSettingsSpec struct {
	MethodPath string `json:"method_path"`
	RestApiId  string `json:"rest_api_id"`
	// +kubebuilder:validation:MaxItems=1
	Settings  []ApiGatewayMethodSettingsSpec `json:"settings"`
	StageName string                         `json:"stage_name"`
}

type ApiGatewayMethodSettingsStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
