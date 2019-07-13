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

type GoogleComputeTargetSslProxy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeTargetSslProxySpec   `json:"spec,omitempty"`
	Status            GoogleComputeTargetSslProxyStatus `json:"status,omitempty"`
}

type GoogleComputeTargetSslProxySpec struct {
	Description       string   `json:"description"`
	SslPolicy         string   `json:"ssl_policy"`
	CreationTimestamp string   `json:"creation_timestamp"`
	Name              string   `json:"name"`
	SslCertificates   []string `json:"ssl_certificates"`
	ProxyHeader       string   `json:"proxy_header"`
	ProxyId           int      `json:"proxy_id"`
	Project           string   `json:"project"`
	SelfLink          string   `json:"self_link"`
	BackendService    string   `json:"backend_service"`
}



type GoogleComputeTargetSslProxyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleComputeTargetSslProxyList is a list of GoogleComputeTargetSslProxys
type GoogleComputeTargetSslProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeTargetSslProxy CRD objects
	Items []GoogleComputeTargetSslProxy `json:"items,omitempty"`
}