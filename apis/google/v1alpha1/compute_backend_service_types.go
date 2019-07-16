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

type ComputeBackendService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeBackendServiceSpec   `json:"spec,omitempty"`
	Status            ComputeBackendServiceStatus `json:"status,omitempty"`
}

type ComputeBackendServiceSpecBackend struct {
	// +optional
	BalancingMode string `json:"balancing_mode,omitempty"`
	// +optional
	CapacityScaler json.Number `json:"capacity_scaler,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Group string `json:"group,omitempty"`
	// +optional
	MaxConnections int `json:"max_connections,omitempty"`
	// +optional
	MaxConnectionsPerInstance int `json:"max_connections_per_instance,omitempty"`
	// +optional
	MaxRate int `json:"max_rate,omitempty"`
	// +optional
	MaxRatePerInstance json.Number `json:"max_rate_per_instance,omitempty"`
	// +optional
	MaxUtilization json.Number `json:"max_utilization,omitempty"`
}

type ComputeBackendServiceSpecIap struct {
	Oauth2ClientId     string `json:"oauth2_client_id"`
	Oauth2ClientSecret string `json:"oauth2_client_secret"`
}

type ComputeBackendServiceSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Backend *[]ComputeBackendServiceSpec `json:"backend,omitempty"`
	// +optional
	ConnectionDrainingTimeoutSec int `json:"connection_draining_timeout_sec,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	// Deprecated
	CustomRequestHeaders []string `json:"custom_request_headers,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	EnableCdn bool `json:"enable_cdn,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	HealthChecks []string `json:"health_checks"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Iap  *[]ComputeBackendServiceSpec `json:"iap,omitempty"`
	Name string                       `json:"name"`
	// +optional
	SecurityPolicy string `json:"security_policy,omitempty"`
}

type ComputeBackendServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
