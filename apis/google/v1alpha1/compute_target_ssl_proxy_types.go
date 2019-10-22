package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type ComputeTargetSSLProxy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeTargetSSLProxySpec   `json:"spec,omitempty"`
	Status            ComputeTargetSSLProxyStatus `json:"status,omitempty"`
}

type ComputeTargetSSLProxySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	BackendService string `json:"backendService" tf:"backend_service"`
	// +optional
	CreationTimestamp string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	ProxyHeader string `json:"proxyHeader,omitempty" tf:"proxy_header,omitempty"`
	// +optional
	ProxyID int `json:"proxyID,omitempty" tf:"proxy_id,omitempty"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	SslCertificates []string `json:"sslCertificates" tf:"ssl_certificates"`
	// +optional
	SslPolicy string `json:"sslPolicy,omitempty" tf:"ssl_policy,omitempty"`
}

type ComputeTargetSSLProxyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeTargetSSLProxySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeTargetSSLProxyList is a list of ComputeTargetSSLProxys
type ComputeTargetSSLProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeTargetSSLProxy CRD objects
	Items []ComputeTargetSSLProxy `json:"items,omitempty"`
}
