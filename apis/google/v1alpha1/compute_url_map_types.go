package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ComputeURLMap struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeURLMapSpec   `json:"spec,omitempty"`
	Status            ComputeURLMapStatus `json:"status,omitempty"`
}

type ComputeURLMapSpecHostRule struct {
	// +optional
	Description string   `json:"description,omitempty" tf:"description,omitempty"`
	Hosts       []string `json:"hosts" tf:"hosts"`
	PathMatcher string   `json:"pathMatcher" tf:"path_matcher"`
}

type ComputeURLMapSpecPathMatcherPathRule struct {
	Paths   []string `json:"paths" tf:"paths"`
	Service string   `json:"service" tf:"service"`
}

type ComputeURLMapSpecPathMatcher struct {
	DefaultService string `json:"defaultService" tf:"default_service"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +optional
	PathRule []ComputeURLMapSpecPathMatcherPathRule `json:"pathRule,omitempty" tf:"path_rule,omitempty"`
}

type ComputeURLMapSpecTest struct {
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Host        string `json:"host" tf:"host"`
	Path        string `json:"path" tf:"path"`
	Service     string `json:"service" tf:"service"`
}

type ComputeURLMapSpec struct {
	DefaultService string `json:"defaultService" tf:"default_service"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	HostRule []ComputeURLMapSpecHostRule `json:"hostRule,omitempty" tf:"host_rule,omitempty"`
	Name     string                      `json:"name" tf:"name"`
	// +optional
	PathMatcher []ComputeURLMapSpecPathMatcher `json:"pathMatcher,omitempty" tf:"path_matcher,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Test        []ComputeURLMapSpecTest   `json:"test,omitempty" tf:"test,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ComputeURLMapStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeURLMapList is a list of ComputeURLMaps
type ComputeURLMapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeURLMap CRD objects
	Items []ComputeURLMap `json:"items,omitempty"`
}
