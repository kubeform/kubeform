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

type ComputeTargetHTTPProxy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeTargetHTTPProxySpec   `json:"spec,omitempty"`
	Status            ComputeTargetHTTPProxyStatus `json:"status,omitempty"`
}

type ComputeTargetHTTPProxySpec struct {
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +optional
	Project     string                    `json:"project,omitempty" tf:"project,omitempty"`
	UrlMap      string                    `json:"urlMap" tf:"url_map"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type ComputeTargetHTTPProxyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeTargetHTTPProxyList is a list of ComputeTargetHTTPProxys
type ComputeTargetHTTPProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeTargetHTTPProxy CRD objects
	Items []ComputeTargetHTTPProxy `json:"items,omitempty"`
}
