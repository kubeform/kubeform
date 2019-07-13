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

type GoogleComputeTargetHttpsProxy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeTargetHttpsProxySpec   `json:"spec,omitempty"`
	Status            GoogleComputeTargetHttpsProxyStatus `json:"status,omitempty"`
}

type GoogleComputeTargetHttpsProxySpec struct {
	UrlMap            string   `json:"url_map"`
	Description       string   `json:"description"`
	CreationTimestamp string   `json:"creation_timestamp"`
	ProxyId           int      `json:"proxy_id"`
	Project           string   `json:"project"`
	SelfLink          string   `json:"self_link"`
	Name              string   `json:"name"`
	SslCertificates   []string `json:"ssl_certificates"`
	QuicOverride      string   `json:"quic_override"`
	SslPolicy         string   `json:"ssl_policy"`
}



type GoogleComputeTargetHttpsProxyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeTargetHttpsProxyList is a list of GoogleComputeTargetHttpsProxys
type GoogleComputeTargetHttpsProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeTargetHttpsProxy CRD objects
	Items []GoogleComputeTargetHttpsProxy `json:"items,omitempty"`
}