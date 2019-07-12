package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermTrafficManagerProfile struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermTrafficManagerProfileSpec   `json:"spec,omitempty"`
	Status            AzurermTrafficManagerProfileStatus `json:"status,omitempty"`
}

type AzurermTrafficManagerProfileSpecDnsConfig struct {
	Ttl          int    `json:"ttl"`
	RelativeName string `json:"relative_name"`
}

type AzurermTrafficManagerProfileSpecMonitorConfig struct {
	Port     int    `json:"port"`
	Path     string `json:"path"`
	Protocol string `json:"protocol"`
}

type AzurermTrafficManagerProfileSpec struct {
	Tags                 map[string]string                  `json:"tags"`
	Name                 string                             `json:"name"`
	ResourceGroupName    string                             `json:"resource_group_name"`
	ProfileStatus        string                             `json:"profile_status"`
	TrafficRoutingMethod string                             `json:"traffic_routing_method"`
	DnsConfig            []AzurermTrafficManagerProfileSpec `json:"dns_config"`
	Fqdn                 string                             `json:"fqdn"`
	MonitorConfig        []AzurermTrafficManagerProfileSpec `json:"monitor_config"`
}



type AzurermTrafficManagerProfileStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermTrafficManagerProfileList is a list of AzurermTrafficManagerProfiles
type AzurermTrafficManagerProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermTrafficManagerProfile CRD objects
	Items []AzurermTrafficManagerProfile `json:"items,omitempty"`
}