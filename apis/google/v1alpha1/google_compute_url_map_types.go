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

type GoogleComputeUrlMap struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeUrlMapSpec   `json:"spec,omitempty"`
	Status            GoogleComputeUrlMapStatus `json:"status,omitempty"`
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

type GoogleComputeUrlMapSpecTest struct {
	Host        string `json:"host"`
	Path        string `json:"path"`
	Service     string `json:"service"`
	Description string `json:"description"`
}

type GoogleComputeUrlMapSpec struct {
	DefaultService string                    `json:"default_service"`
	Fingerprint    string                    `json:"fingerprint"`
	HostRule       []GoogleComputeUrlMapSpec `json:"host_rule"`
	PathMatcher    []GoogleComputeUrlMapSpec `json:"path_matcher"`
	SelfLink       string                    `json:"self_link"`
	Test           []GoogleComputeUrlMapSpec `json:"test"`
	Name           string                    `json:"name"`
	Description    string                    `json:"description"`
	MapId          string                    `json:"map_id"`
	Project        string                    `json:"project"`
}



type GoogleComputeUrlMapStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeUrlMapList is a list of GoogleComputeUrlMaps
type GoogleComputeUrlMapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeUrlMap CRD objects
	Items []GoogleComputeUrlMap `json:"items,omitempty"`
}