package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ServiceDiscoveryService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceDiscoveryServiceSpec   `json:"spec,omitempty"`
	Status            ServiceDiscoveryServiceStatus `json:"status,omitempty"`
}

type ServiceDiscoveryServiceSpecDnsConfigDnsRecords struct {
	Ttl  int    `json:"ttl" tf:"ttl"`
	Type string `json:"type" tf:"type"`
}

type ServiceDiscoveryServiceSpecDnsConfig struct {
	DnsRecords  []ServiceDiscoveryServiceSpecDnsConfigDnsRecords `json:"dnsRecords" tf:"dns_records"`
	NamespaceID string                                           `json:"namespaceID" tf:"namespace_id"`
	// +optional
	RoutingPolicy string `json:"routingPolicy,omitempty" tf:"routing_policy,omitempty"`
}

type ServiceDiscoveryServiceSpecHealthCheckConfig struct {
	// +optional
	FailureThreshold int `json:"failureThreshold,omitempty" tf:"failure_threshold,omitempty"`
	// +optional
	ResourcePath string `json:"resourcePath,omitempty" tf:"resource_path,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type ServiceDiscoveryServiceSpecHealthCheckCustomConfig struct {
	// +optional
	FailureThreshold int `json:"failureThreshold,omitempty" tf:"failure_threshold,omitempty"`
}

type ServiceDiscoveryServiceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DnsConfig []ServiceDiscoveryServiceSpecDnsConfig `json:"dnsConfig,omitempty" tf:"dns_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HealthCheckConfig []ServiceDiscoveryServiceSpecHealthCheckConfig `json:"healthCheckConfig,omitempty" tf:"health_check_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HealthCheckCustomConfig []ServiceDiscoveryServiceSpecHealthCheckCustomConfig `json:"healthCheckCustomConfig,omitempty" tf:"health_check_custom_config,omitempty"`
	Name                    string                                               `json:"name" tf:"name"`
	// +optional
	NamespaceID string `json:"namespaceID,omitempty" tf:"namespace_id,omitempty"`
}

type ServiceDiscoveryServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServiceDiscoveryServiceList is a list of ServiceDiscoveryServices
type ServiceDiscoveryServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServiceDiscoveryService CRD objects
	Items []ServiceDiscoveryService `json:"items,omitempty"`
}
