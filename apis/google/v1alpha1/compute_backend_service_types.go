package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ComputeBackendService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeBackendServiceSpec   `json:"spec,omitempty"`
	Status            ComputeBackendServiceStatus `json:"status,omitempty"`
}

type ComputeBackendServiceSpecBackend struct {
	// +optional
	BalancingMode string `json:"balancingMode,omitempty" tf:"balancing_mode,omitempty"`
	// +optional
	CapacityScaler json.Number `json:"capacityScaler,omitempty" tf:"capacity_scaler,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Group string `json:"group,omitempty" tf:"group,omitempty"`
	// +optional
	MaxConnections int `json:"maxConnections,omitempty" tf:"max_connections,omitempty"`
	// +optional
	MaxConnectionsPerInstance int `json:"maxConnectionsPerInstance,omitempty" tf:"max_connections_per_instance,omitempty"`
	// +optional
	MaxRate int `json:"maxRate,omitempty" tf:"max_rate,omitempty"`
	// +optional
	MaxRatePerInstance json.Number `json:"maxRatePerInstance,omitempty" tf:"max_rate_per_instance,omitempty"`
	// +optional
	MaxUtilization json.Number `json:"maxUtilization,omitempty" tf:"max_utilization,omitempty"`
}

type ComputeBackendServiceSpecCdnPolicyCacheKeyPolicy struct {
	// +optional
	IncludeHost bool `json:"includeHost,omitempty" tf:"include_host,omitempty"`
	// +optional
	IncludeProtocol bool `json:"includeProtocol,omitempty" tf:"include_protocol,omitempty"`
	// +optional
	IncludeQueryString bool `json:"includeQueryString,omitempty" tf:"include_query_string,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	QueryStringBlacklist []string `json:"queryStringBlacklist,omitempty" tf:"query_string_blacklist,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	QueryStringWhitelist []string `json:"queryStringWhitelist,omitempty" tf:"query_string_whitelist,omitempty"`
}

type ComputeBackendServiceSpecCdnPolicy struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CacheKeyPolicy []ComputeBackendServiceSpecCdnPolicyCacheKeyPolicy `json:"cacheKeyPolicy,omitempty" tf:"cache_key_policy,omitempty"`
}

type ComputeBackendServiceSpecIap struct {
	Oauth2ClientID string `json:"oauth2ClientID" tf:"oauth2_client_id"`
}

type ComputeBackendServiceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Backend []ComputeBackendServiceSpecBackend `json:"backend,omitempty" tf:"backend,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CdnPolicy []ComputeBackendServiceSpecCdnPolicy `json:"cdnPolicy,omitempty" tf:"cdn_policy,omitempty"`
	// +optional
	ConnectionDrainingTimeoutSec int `json:"connectionDrainingTimeoutSec,omitempty" tf:"connection_draining_timeout_sec,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	// Deprecated
	CustomRequestHeaders []string `json:"customRequestHeaders,omitempty" tf:"custom_request_headers,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	EnableCdn bool `json:"enableCdn,omitempty" tf:"enable_cdn,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	HealthChecks []string `json:"healthChecks" tf:"health_checks"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Iap  []ComputeBackendServiceSpecIap `json:"iap,omitempty" tf:"iap,omitempty"`
	Name string                         `json:"name" tf:"name"`
	// +optional
	PortName string `json:"portName,omitempty" tf:"port_name,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Protocol string `json:"protocol,omitempty" tf:"protocol,omitempty"`
	// +optional
	SecurityPolicy string `json:"securityPolicy,omitempty" tf:"security_policy,omitempty"`
	// +optional
	SessionAffinity string `json:"sessionAffinity,omitempty" tf:"session_affinity,omitempty"`
	// +optional
	TimeoutSec int `json:"timeoutSec,omitempty" tf:"timeout_sec,omitempty"`
}

type ComputeBackendServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeBackendServiceList is a list of ComputeBackendServices
type ComputeBackendServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeBackendService CRD objects
	Items []ComputeBackendService `json:"items,omitempty"`
}
