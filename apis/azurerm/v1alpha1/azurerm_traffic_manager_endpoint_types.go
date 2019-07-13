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

type AzurermTrafficManagerEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermTrafficManagerEndpointSpec   `json:"spec,omitempty"`
	Status            AzurermTrafficManagerEndpointStatus `json:"status,omitempty"`
}

type AzurermTrafficManagerEndpointSpec struct {
	ProfileName           string   `json:"profile_name"`
	Weight                int      `json:"weight"`
	EndpointLocation      string   `json:"endpoint_location"`
	MinChildEndpoints     int      `json:"min_child_endpoints"`
	Name                  string   `json:"name"`
	ResourceGroupName     string   `json:"resource_group_name"`
	Type                  string   `json:"type"`
	Target                string   `json:"target"`
	TargetResourceId      string   `json:"target_resource_id"`
	EndpointStatus        string   `json:"endpoint_status"`
	Priority              int      `json:"priority"`
	GeoMappings           []string `json:"geo_mappings"`
	EndpointMonitorStatus string   `json:"endpoint_monitor_status"`
}



type AzurermTrafficManagerEndpointStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermTrafficManagerEndpointList is a list of AzurermTrafficManagerEndpoints
type AzurermTrafficManagerEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermTrafficManagerEndpoint CRD objects
	Items []AzurermTrafficManagerEndpoint `json:"items,omitempty"`
}