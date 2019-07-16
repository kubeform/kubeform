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

type ComputeTargetHttpProxy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeTargetHttpProxySpec   `json:"spec,omitempty"`
	Status            ComputeTargetHttpProxyStatus `json:"status,omitempty"`
}

type ComputeTargetHttpProxySpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	UrlMap      string `json:"url_map"`
}

type ComputeTargetHttpProxyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeTargetHttpProxyList is a list of ComputeTargetHttpProxys
type ComputeTargetHttpProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeTargetHttpProxy CRD objects
	Items []ComputeTargetHttpProxy `json:"items,omitempty"`
}
