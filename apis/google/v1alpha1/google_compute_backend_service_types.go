package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeBackendService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeBackendServiceSpec   `json:"spec,omitempty"`
	Status            GoogleComputeBackendServiceStatus `json:"status,omitempty"`
}

type GoogleComputeBackendServiceSpecCdnPolicyCacheKeyPolicy struct {
	QueryStringBlacklist []string `json:"query_string_blacklist"`
	QueryStringWhitelist []string `json:"query_string_whitelist"`
	IncludeHost          bool     `json:"include_host"`
	IncludeProtocol      bool     `json:"include_protocol"`
	IncludeQueryString   bool     `json:"include_query_string"`
}

type GoogleComputeBackendServiceSpecCdnPolicy struct {
	CacheKeyPolicy []GoogleComputeBackendServiceSpecCdnPolicy `json:"cache_key_policy"`
}

type GoogleComputeBackendServiceSpecIap struct {
	Oauth2ClientId     string `json:"oauth2_client_id"`
	Oauth2ClientSecret string `json:"oauth2_client_secret"`
}

type GoogleComputeBackendServiceSpecBackend struct {
	Description               string  `json:"description"`
	MaxConnectionsPerInstance int     `json:"max_connections_per_instance"`
	Group                     string  `json:"group"`
	BalancingMode             string  `json:"balancing_mode"`
	CapacityScaler            float64 `json:"capacity_scaler"`
	MaxUtilization            float64 `json:"max_utilization"`
	MaxRate                   int     `json:"max_rate"`
	MaxRatePerInstance        float64 `json:"max_rate_per_instance"`
	MaxConnections            int     `json:"max_connections"`
}

type GoogleComputeBackendServiceSpec struct {
	CdnPolicy                    []GoogleComputeBackendServiceSpec `json:"cdn_policy"`
	CustomRequestHeaders         []string                          `json:"custom_request_headers"`
	SecurityPolicy               string                            `json:"security_policy"`
	Iap                          []GoogleComputeBackendServiceSpec `json:"iap"`
	Backend                      []GoogleComputeBackendServiceSpec `json:"backend"`
	Description                  string                            `json:"description"`
	Fingerprint                  string                            `json:"fingerprint"`
	TimeoutSec                   int                               `json:"timeout_sec"`
	EnableCdn                    bool                              `json:"enable_cdn"`
	Protocol                     string                            `json:"protocol"`
	Region                       string                            `json:"region"`
	SessionAffinity              string                            `json:"session_affinity"`
	ConnectionDrainingTimeoutSec int                               `json:"connection_draining_timeout_sec"`
	Name                         string                            `json:"name"`
	HealthChecks                 []string                          `json:"health_checks"`
	PortName                     string                            `json:"port_name"`
	Project                      string                            `json:"project"`
	SelfLink                     string                            `json:"self_link"`
}



type GoogleComputeBackendServiceStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeBackendServiceList is a list of GoogleComputeBackendServices
type GoogleComputeBackendServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeBackendService CRD objects
	Items []GoogleComputeBackendService `json:"items,omitempty"`
}