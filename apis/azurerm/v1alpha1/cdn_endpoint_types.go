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

type CdnEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CdnEndpointSpec   `json:"spec,omitempty"`
	Status            CdnEndpointStatus `json:"status,omitempty"`
}

type CdnEndpointSpecGeoFilter struct {
	Action       string   `json:"action"`
	CountryCodes []string `json:"country_codes"`
	RelativePath string   `json:"relative_path"`
}

type CdnEndpointSpecOrigin struct {
	HostName string `json:"host_name"`
	// +optional
	HttpPort int `json:"http_port,omitempty"`
	// +optional
	HttpsPort int    `json:"https_port,omitempty"`
	Name      string `json:"name"`
}

type CdnEndpointSpec struct {
	// +optional
	GeoFilter *[]CdnEndpointSpec `json:"geo_filter,omitempty"`
	// +optional
	IsCompressionEnabled bool `json:"is_compression_enabled,omitempty"`
	// +optional
	IsHttpAllowed bool `json:"is_http_allowed,omitempty"`
	// +optional
	IsHttpsAllowed bool   `json:"is_https_allowed,omitempty"`
	Location       string `json:"location"`
	Name           string `json:"name"`
	// +optional
	OptimizationType string `json:"optimization_type,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	Origin      []CdnEndpointSpec `json:"origin"`
	ProfileName string            `json:"profile_name"`
	// +optional
	QuerystringCachingBehaviour string `json:"querystring_caching_behaviour,omitempty"`
	ResourceGroupName           string `json:"resource_group_name"`
}

type CdnEndpointStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CdnEndpointList is a list of CdnEndpoints
type CdnEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CdnEndpoint CRD objects
	Items []CdnEndpoint `json:"items,omitempty"`
}
