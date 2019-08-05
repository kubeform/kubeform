package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ComputeTargetHTTPSProxy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeTargetHTTPSProxySpec   `json:"spec,omitempty"`
	Status            ComputeTargetHTTPSProxyStatus `json:"status,omitempty"`
}

type ComputeTargetHTTPSProxySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	CreationTimestamp string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	ProxyID int `json:"proxyID,omitempty" tf:"proxy_id,omitempty"`
	// +optional
	QuicOverride string `json:"quicOverride,omitempty" tf:"quic_override,omitempty"`
	// +optional
	SelfLink        string   `json:"selfLink,omitempty" tf:"self_link,omitempty"`
	SslCertificates []string `json:"sslCertificates" tf:"ssl_certificates"`
	// +optional
	SslPolicy string `json:"sslPolicy,omitempty" tf:"ssl_policy,omitempty"`
	UrlMap    string `json:"urlMap" tf:"url_map"`
}

type ComputeTargetHTTPSProxyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeTargetHTTPSProxySpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeTargetHTTPSProxyList is a list of ComputeTargetHTTPSProxys
type ComputeTargetHTTPSProxyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeTargetHTTPSProxy CRD objects
	Items []ComputeTargetHTTPSProxy `json:"items,omitempty"`
}
