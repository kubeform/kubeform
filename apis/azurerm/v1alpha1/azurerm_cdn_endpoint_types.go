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

type AzurermCdnEndpointSpecGeoFilter struct {
	CountryCodes []string `json:"country_codes"`
	RelativePath string   `json:"relative_path"`
	Action       string   `json:"action"`
}

type AzurermCdnEndpointSpecOrigin struct {
	HttpsPort int    `json:"https_port"`
	Name      string `json:"name"`
	HostName  string `json:"host_name"`
	HttpPort  int    `json:"http_port"`
}

type AzurermCdnEndpointSpec struct {
	QuerystringCachingBehaviour string                   `json:"querystring_caching_behaviour"`
	GeoFilter                   []AzurermCdnEndpointSpec `json:"geo_filter"`
	IsHttpAllowed               bool                     `json:"is_http_allowed"`
	Origin                      []AzurermCdnEndpointSpec `json:"origin"`
	OriginPath                  string                   `json:"origin_path"`
	ContentTypesToCompress      []string                 `json:"content_types_to_compress"`
	IsCompressionEnabled        bool                     `json:"is_compression_enabled"`
	Tags                        map[string]string        `json:"tags"`
	ResourceGroupName           string                   `json:"resource_group_name"`
	OriginHostHeader            string                   `json:"origin_host_header"`
	ProbePath                   string                   `json:"probe_path"`
	OptimizationType            string                   `json:"optimization_type"`
	HostName                    string                   `json:"host_name"`
	Name                        string                   `json:"name"`
	ProfileName                 string                   `json:"profile_name"`
	Location                    string                   `json:"location"`
	IsHttpsAllowed              bool                     `json:"is_https_allowed"`
}

type AzurermCdnEndpointStatus struct {
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
