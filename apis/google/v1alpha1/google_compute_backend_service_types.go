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

type GoogleComputeBackendService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeBackendServiceSpec   `json:"spec,omitempty"`
	Status            GoogleComputeBackendServiceStatus `json:"status,omitempty"`
}

type GoogleComputeBackendServiceSpecCdnPolicyCacheKeyPolicy struct {
	IncludeHost          bool     `json:"include_host"`
	IncludeProtocol      bool     `json:"include_protocol"`
	IncludeQueryString   bool     `json:"include_query_string"`
	QueryStringBlacklist []string `json:"query_string_blacklist"`
	QueryStringWhitelist []string `json:"query_string_whitelist"`
}

type GoogleComputeBackendServiceSpecCdnPolicy struct {
	CacheKeyPolicy []GoogleComputeBackendServiceSpecCdnPolicy `json:"cache_key_policy"`
}

type GoogleComputeBackendServiceSpecIap struct {
	Oauth2ClientId     string `json:"oauth2_client_id"`
	Oauth2ClientSecret string `json:"oauth2_client_secret"`
}

type GoogleComputeBackendServiceSpecBackend struct {
	MaxUtilization            float64 `json:"max_utilization"`
	Group                     string  `json:"group"`
	CapacityScaler            float64 `json:"capacity_scaler"`
	Description               string  `json:"description"`
	MaxRate                   int     `json:"max_rate"`
	MaxConnections            int     `json:"max_connections"`
	BalancingMode             string  `json:"balancing_mode"`
	MaxRatePerInstance        float64 `json:"max_rate_per_instance"`
	MaxConnectionsPerInstance int     `json:"max_connections_per_instance"`
}

type GoogleComputeBackendServiceSpec struct {
	HealthChecks                 []string                          `json:"health_checks"`
	CdnPolicy                    []GoogleComputeBackendServiceSpec `json:"cdn_policy"`
	SelfLink                     string                            `json:"self_link"`
	TimeoutSec                   int                               `json:"timeout_sec"`
	Project                      string                            `json:"project"`
	Name                         string                            `json:"name"`
	Iap                          []GoogleComputeBackendServiceSpec `json:"iap"`
	Backend                      []GoogleComputeBackendServiceSpec `json:"backend"`
	PortName                     string                            `json:"port_name"`
	CustomRequestHeaders         []string                          `json:"custom_request_headers"`
	Description                  string                            `json:"description"`
	EnableCdn                    bool                              `json:"enable_cdn"`
	SessionAffinity              string                            `json:"session_affinity"`
	ConnectionDrainingTimeoutSec int                               `json:"connection_draining_timeout_sec"`
	Fingerprint                  string                            `json:"fingerprint"`
	Protocol                     string                            `json:"protocol"`
	Region                       string                            `json:"region"`
	SecurityPolicy               string                            `json:"security_policy"`
}



type GoogleComputeBackendServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeBackendServiceList is a list of GoogleComputeBackendServices
type GoogleComputeBackendServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeBackendService CRD objects
	Items []GoogleComputeBackendService `json:"items,omitempty"`
}