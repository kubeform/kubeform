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

type ComputeUrlMap struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeUrlMapSpec   `json:"spec,omitempty"`
	Status            ComputeUrlMapStatus `json:"status,omitempty"`
}

type ComputeUrlMapSpecHostRule struct {
	// +optional
	Description string   `json:"description,omitempty"`
	Hosts       []string `json:"hosts"`
	PathMatcher string   `json:"path_matcher"`
}

type ComputeUrlMapSpecPathMatcherPathRule struct {
	Paths   []string `json:"paths"`
	Service string   `json:"service"`
}

type ComputeUrlMapSpecPathMatcher struct {
	DefaultService string `json:"default_service"`
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	// +optional
	PathRule *[]ComputeUrlMapSpecPathMatcher `json:"path_rule,omitempty"`
}

type ComputeUrlMapSpecTest struct {
	// +optional
	Description string `json:"description,omitempty"`
	Host        string `json:"host"`
	Path        string `json:"path"`
	Service     string `json:"service"`
}

type ComputeUrlMapSpec struct {
	DefaultService string `json:"default_service"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	HostRule *[]ComputeUrlMapSpec `json:"host_rule,omitempty"`
	Name     string               `json:"name"`
	// +optional
	PathMatcher *[]ComputeUrlMapSpec `json:"path_matcher,omitempty"`
	// +optional
	Test *[]ComputeUrlMapSpec `json:"test,omitempty"`
}

type ComputeUrlMapStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeUrlMapList is a list of ComputeUrlMaps
type ComputeUrlMapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeUrlMap CRD objects
	Items []ComputeUrlMap `json:"items,omitempty"`
}
