package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type TrafficManagerProfile struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TrafficManagerProfileSpec   `json:"spec,omitempty"`
	Status            TrafficManagerProfileStatus `json:"status,omitempty"`
}

type TrafficManagerProfileSpecDnsConfig struct {
	RelativeName string `json:"relativeName" tf:"relative_name"`
	Ttl          int    `json:"ttl" tf:"ttl"`
}

type TrafficManagerProfileSpecMonitorConfig struct {
	// +optional
	Path     string `json:"path,omitempty" tf:"path,omitempty"`
	Port     int    `json:"port" tf:"port"`
	Protocol string `json:"protocol" tf:"protocol"`
}

type TrafficManagerProfileSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	DnsConfig []TrafficManagerProfileSpecDnsConfig `json:"dnsConfig" tf:"dns_config"`
	// +optional
	Fqdn          string                                   `json:"fqdn,omitempty" tf:"fqdn,omitempty"`
	MonitorConfig []TrafficManagerProfileSpecMonitorConfig `json:"monitorConfig" tf:"monitor_config"`
	Name          string                                   `json:"name" tf:"name"`
	// +optional
	ProfileStatus     string `json:"profileStatus,omitempty" tf:"profile_status,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags                 map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	TrafficRoutingMethod string            `json:"trafficRoutingMethod" tf:"traffic_routing_method"`
}

type TrafficManagerProfileStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *TrafficManagerProfileSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// TrafficManagerProfileList is a list of TrafficManagerProfiles
type TrafficManagerProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of TrafficManagerProfile CRD objects
	Items []TrafficManagerProfile `json:"items,omitempty"`
}
