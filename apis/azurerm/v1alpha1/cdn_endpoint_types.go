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

type CdnEndpoint struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CdnEndpointSpec   `json:"spec,omitempty"`
	Status            CdnEndpointStatus `json:"status,omitempty"`
}

type CdnEndpointSpecGeoFilter struct {
	Action       string   `json:"action" tf:"action"`
	CountryCodes []string `json:"countryCodes" tf:"country_codes"`
	RelativePath string   `json:"relativePath" tf:"relative_path"`
}

type CdnEndpointSpecOrigin struct {
	HostName string `json:"hostName" tf:"host_name"`
	// +optional
	HttpPort int `json:"httpPort,omitempty" tf:"http_port,omitempty"`
	// +optional
	HttpsPort int    `json:"httpsPort,omitempty" tf:"https_port,omitempty"`
	Name      string `json:"name" tf:"name"`
}

type CdnEndpointSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true
	ContentTypesToCompress []string `json:"contentTypesToCompress,omitempty" tf:"content_types_to_compress,omitempty"`
	// +optional
	GeoFilter []CdnEndpointSpecGeoFilter `json:"geoFilter,omitempty" tf:"geo_filter,omitempty"`
	HostName  string                     `json:"hostName" tf:"host_name"`
	// +optional
	IsCompressionEnabled bool `json:"isCompressionEnabled,omitempty" tf:"is_compression_enabled,omitempty"`
	// +optional
	IsHTTPAllowed bool `json:"isHTTPAllowed,omitempty" tf:"is_http_allowed,omitempty"`
	// +optional
	IsHTTPSAllowed bool   `json:"isHTTPSAllowed,omitempty" tf:"is_https_allowed,omitempty"`
	Location       string `json:"location" tf:"location"`
	Name           string `json:"name" tf:"name"`
	// +optional
	OptimizationType string `json:"optimizationType,omitempty" tf:"optimization_type,omitempty"`
	// +kubebuilder:validation:UniqueItems=true
	Origin []CdnEndpointSpecOrigin `json:"origin" tf:"origin"`
	// +optional
	OriginHostHeader string `json:"originHostHeader,omitempty" tf:"origin_host_header,omitempty"`
	// +optional
	OriginPath string `json:"originPath,omitempty" tf:"origin_path,omitempty"`
	// +optional
	ProbePath   string `json:"probePath,omitempty" tf:"probe_path,omitempty"`
	ProfileName string `json:"profileName" tf:"profile_name"`
	// +optional
	QuerystringCachingBehaviour string `json:"querystringCachingBehaviour,omitempty" tf:"querystring_caching_behaviour,omitempty"`
	ResourceGroupName           string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type CdnEndpointStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CdnEndpointSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
