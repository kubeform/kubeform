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

type GoogleComputeTargetHttpProxy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeTargetHttpProxySpec   `json:"spec,omitempty"`
	Status            GoogleComputeTargetHttpProxyStatus `json:"status,omitempty"`
}

type GoogleComputeTargetHttpProxySpec struct {
	Description       string `json:"description"`
	CreationTimestamp string `json:"creation_timestamp"`
	ProxyId           int    `json:"proxy_id"`
	Project           string `json:"project"`
	SelfLink          string `json:"self_link"`
	Name              string `json:"name"`
	UrlMap            string `json:"url_map"`
}



type GoogleComputeTargetHttpProxyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeTargetHttpProxyList is a list of GoogleComputeTargetHttpProxys
type GoogleComputeTargetHttpProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeTargetHttpProxy CRD objects
	Items []GoogleComputeTargetHttpProxy `json:"items,omitempty"`
}