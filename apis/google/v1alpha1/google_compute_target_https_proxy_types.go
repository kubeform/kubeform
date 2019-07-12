package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeTargetHttpsProxy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeTargetHttpsProxySpec   `json:"spec,omitempty"`
	Status            GoogleComputeTargetHttpsProxyStatus `json:"status,omitempty"`
}

type GoogleComputeTargetHttpsProxySpec struct {
	Project           string   `json:"project"`
	Description       string   `json:"description"`
	QuicOverride      string   `json:"quic_override"`
	SslPolicy         string   `json:"ssl_policy"`
	CreationTimestamp string   `json:"creation_timestamp"`
	ProxyId           int      `json:"proxy_id"`
	Name              string   `json:"name"`
	SslCertificates   []string `json:"ssl_certificates"`
	UrlMap            string   `json:"url_map"`
	SelfLink          string   `json:"self_link"`
}



type GoogleComputeTargetHttpsProxyStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeTargetHttpsProxyList is a list of GoogleComputeTargetHttpsProxys
type GoogleComputeTargetHttpsProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeTargetHttpsProxy CRD objects
	Items []GoogleComputeTargetHttpsProxy `json:"items,omitempty"`
}