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

type ComputeTargetHttpsProxy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeTargetHttpsProxySpec   `json:"spec,omitempty"`
	Status            ComputeTargetHttpsProxyStatus `json:"status,omitempty"`
}

type ComputeTargetHttpsProxySpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	// +optional
	QuicOverride    string   `json:"quic_override,omitempty"`
	SslCertificates []string `json:"ssl_certificates"`
	// +optional
	SslPolicy string `json:"ssl_policy,omitempty"`
	UrlMap    string `json:"url_map"`
}

type ComputeTargetHttpsProxyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeTargetHttpsProxyList is a list of ComputeTargetHttpsProxys
type ComputeTargetHttpsProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeTargetHttpsProxy CRD objects
	Items []ComputeTargetHttpsProxy `json:"items,omitempty"`
}
