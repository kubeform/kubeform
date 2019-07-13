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

type AzurermCdnEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermCdnEndpointSpec   `json:"spec,omitempty"`
	Status            AzurermCdnEndpointStatus `json:"status,omitempty"`
}

type AzurermCdnEndpointSpecOrigin struct {
	Name      string `json:"name"`
	HostName  string `json:"host_name"`
	HttpPort  int    `json:"http_port"`
	HttpsPort int    `json:"https_port"`
}

type AzurermCdnEndpointSpecGeoFilter struct {
	RelativePath string   `json:"relative_path"`
	Action       string   `json:"action"`
	CountryCodes []string `json:"country_codes"`
}

type AzurermCdnEndpointSpec struct {
	OptimizationType            string                   `json:"optimization_type"`
	Location                    string                   `json:"location"`
	OriginHostHeader            string                   `json:"origin_host_header"`
	Origin                      []AzurermCdnEndpointSpec `json:"origin"`
	ContentTypesToCompress      []string                 `json:"content_types_to_compress"`
	IsCompressionEnabled        bool                     `json:"is_compression_enabled"`
	ProbePath                   string                   `json:"probe_path"`
	GeoFilter                   []AzurermCdnEndpointSpec `json:"geo_filter"`
	Name                        string                   `json:"name"`
	QuerystringCachingBehaviour string                   `json:"querystring_caching_behaviour"`
	HostName                    string                   `json:"host_name"`
	Tags                        map[string]string        `json:"tags"`
	ResourceGroupName           string                   `json:"resource_group_name"`
	ProfileName                 string                   `json:"profile_name"`
	IsHttpAllowed               bool                     `json:"is_http_allowed"`
	IsHttpsAllowed              bool                     `json:"is_https_allowed"`
	OriginPath                  string                   `json:"origin_path"`
}



type AzurermCdnEndpointStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermCdnEndpointList is a list of AzurermCdnEndpoints
type AzurermCdnEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermCdnEndpoint CRD objects
	Items []AzurermCdnEndpoint `json:"items,omitempty"`
}