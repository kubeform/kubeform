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

type ComputeTargetSslProxy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeTargetSslProxySpec   `json:"spec,omitempty"`
	Status            ComputeTargetSslProxyStatus `json:"status,omitempty"`
}

type ComputeTargetSslProxySpec struct {
	BackendService string `json:"backend_service"`
	// +optional
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
	// +optional
	ProxyHeader string `json:"proxy_header,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	SslCertificates []string `json:"ssl_certificates"`
	// +optional
	SslPolicy string `json:"ssl_policy,omitempty"`
}

type ComputeTargetSslProxyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeTargetSslProxyList is a list of ComputeTargetSslProxys
type ComputeTargetSslProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeTargetSslProxy CRD objects
	Items []ComputeTargetSslProxy `json:"items,omitempty"`
}
