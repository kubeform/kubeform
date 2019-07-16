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

type ServiceDiscoveryService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceDiscoveryServiceSpec   `json:"spec,omitempty"`
	Status            ServiceDiscoveryServiceStatus `json:"status,omitempty"`
}

type ServiceDiscoveryServiceSpecDnsConfigDnsRecords struct {
	Ttl  int    `json:"ttl"`
	Type string `json:"type"`
}

type ServiceDiscoveryServiceSpecDnsConfig struct {
	DnsRecords  []ServiceDiscoveryServiceSpecDnsConfig `json:"dns_records"`
	NamespaceId string                                 `json:"namespace_id"`
	// +optional
	RoutingPolicy string `json:"routing_policy,omitempty"`
}

type ServiceDiscoveryServiceSpecHealthCheckConfig struct {
	// +optional
	FailureThreshold int `json:"failure_threshold,omitempty"`
	// +optional
	ResourcePath string `json:"resource_path,omitempty"`
	// +optional
	Type string `json:"type,omitempty"`
}

type ServiceDiscoveryServiceSpecHealthCheckCustomConfig struct {
	// +optional
	FailureThreshold int `json:"failure_threshold,omitempty"`
}

type ServiceDiscoveryServiceSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	DnsConfig *[]ServiceDiscoveryServiceSpec `json:"dns_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HealthCheckConfig *[]ServiceDiscoveryServiceSpec `json:"health_check_config,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	HealthCheckCustomConfig *[]ServiceDiscoveryServiceSpec `json:"health_check_custom_config,omitempty"`
	Name                    string                         `json:"name"`
}

type ServiceDiscoveryServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
