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

type AwsServiceDiscoveryService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsServiceDiscoveryServiceSpec   `json:"spec,omitempty"`
	Status            AwsServiceDiscoveryServiceStatus `json:"status,omitempty"`
}

type AwsServiceDiscoveryServiceSpecHealthCheckCustomConfig struct {
	FailureThreshold int `json:"failure_threshold"`
}

type AwsServiceDiscoveryServiceSpecDnsConfigDnsRecords struct {
	Ttl  int    `json:"ttl"`
	Type string `json:"type"`
}

type AwsServiceDiscoveryServiceSpecDnsConfig struct {
	NamespaceId   string                                    `json:"namespace_id"`
	DnsRecords    []AwsServiceDiscoveryServiceSpecDnsConfig `json:"dns_records"`
	RoutingPolicy string                                    `json:"routing_policy"`
}

type AwsServiceDiscoveryServiceSpecHealthCheckConfig struct {
	FailureThreshold int    `json:"failure_threshold"`
	ResourcePath     string `json:"resource_path"`
	Type             string `json:"type"`
}

type AwsServiceDiscoveryServiceSpec struct {
	HealthCheckCustomConfig []AwsServiceDiscoveryServiceSpec `json:"health_check_custom_config"`
	Arn                     string                           `json:"arn"`
	Name                    string                           `json:"name"`
	Description             string                           `json:"description"`
	NamespaceId             string                           `json:"namespace_id"`
	DnsConfig               []AwsServiceDiscoveryServiceSpec `json:"dns_config"`
	HealthCheckConfig       []AwsServiceDiscoveryServiceSpec `json:"health_check_config"`
}



type AwsServiceDiscoveryServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsServiceDiscoveryServiceList is a list of AwsServiceDiscoveryServices
type AwsServiceDiscoveryServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsServiceDiscoveryService CRD objects
	Items []AwsServiceDiscoveryService `json:"items,omitempty"`
}