package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	BackendService string `json:"backendService" tf:"backend_service"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	ProxyHeader string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	SslCertificates []string `json:"sslCertificates" tf:"ssl_certificates"`
	// +optional
	SslPolicy string `json:"sslPolicy,omitempty" tf:"ssl_policy,omitempty"`
}

type ComputeTargetSslProxyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State *ComputeTargetSslProxySpec `json:"state,omitempty"`
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
