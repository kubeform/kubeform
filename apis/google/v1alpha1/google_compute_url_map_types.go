package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeUrlMap struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeUrlMapSpec   `json:"spec,omitempty"`
	Status            GoogleComputeUrlMapStatus `json:"status,omitempty"`
}

type GoogleComputeUrlMapSpecTest struct {
	Host        string `json:"host"`
	Path        string `json:"path"`
	Service     string `json:"service"`
	Description string `json:"description"`
}

type GoogleComputeUrlMapSpecHostRule struct {
	Description string   `json:"description"`
	Hosts       []string `json:"hosts"`
	PathMatcher string   `json:"path_matcher"`
}

type GoogleComputeUrlMapSpecPathMatcherPathRule struct {
	Paths   []string `json:"paths"`
	Service string   `json:"service"`
}

type GoogleComputeUrlMapSpecPathMatcher struct {
	DefaultService string                               `json:"default_service"`
	Description    string                               `json:"description"`
	Name           string                               `json:"name"`
	PathRule       []GoogleComputeUrlMapSpecPathMatcher `json:"path_rule"`
}

type GoogleComputeUrlMapSpec struct {
	Test           []GoogleComputeUrlMapSpec `json:"test"`
	Name           string                    `json:"name"`
	Fingerprint    string                    `json:"fingerprint"`
	HostRule       []GoogleComputeUrlMapSpec `json:"host_rule"`
	MapId          string                    `json:"map_id"`
	PathMatcher    []GoogleComputeUrlMapSpec `json:"path_matcher"`
	Project        string                    `json:"project"`
	DefaultService string                    `json:"default_service"`
	Description    string                    `json:"description"`
	SelfLink       string                    `json:"self_link"`
}



type GoogleComputeUrlMapStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeUrlMapList is a list of GoogleComputeUrlMaps
type GoogleComputeUrlMapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeUrlMap CRD objects
	Items []GoogleComputeUrlMap `json:"items,omitempty"`
}