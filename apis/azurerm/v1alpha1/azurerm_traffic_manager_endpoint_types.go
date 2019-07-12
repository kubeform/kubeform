package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermTrafficManagerEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermTrafficManagerEndpointSpec   `json:"spec,omitempty"`
	Status            AzurermTrafficManagerEndpointStatus `json:"status,omitempty"`
}

type AzurermTrafficManagerEndpointSpec struct {
	ResourceGroupName     string   `json:"resource_group_name"`
	Type                  string   `json:"type"`
	TargetResourceId      string   `json:"target_resource_id"`
	Weight                int      `json:"weight"`
	MinChildEndpoints     int      `json:"min_child_endpoints"`
	GeoMappings           []string `json:"geo_mappings"`
	ProfileName           string   `json:"profile_name"`
	Target                string   `json:"target"`
	EndpointStatus        string   `json:"endpoint_status"`
	Priority              int      `json:"priority"`
	EndpointLocation      string   `json:"endpoint_location"`
	EndpointMonitorStatus string   `json:"endpoint_monitor_status"`
	Name                  string   `json:"name"`
}



type AzurermTrafficManagerEndpointStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermTrafficManagerEndpointList is a list of AzurermTrafficManagerEndpoints
type AzurermTrafficManagerEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermTrafficManagerEndpoint CRD objects
	Items []AzurermTrafficManagerEndpoint `json:"items,omitempty"`
}