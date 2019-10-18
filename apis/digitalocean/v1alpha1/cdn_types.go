package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Cdn struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CdnSpec   `json:"spec,omitempty"`
	Status            CdnStatus `json:"status,omitempty"`
}

type CdnSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// ID of a DigitalOcean managed TLS certificate for use with custom domains
	// +optional
	CertificateID string `json:"certificateID,omitempty" tf:"certificate_id,omitempty"`
	// The date and time (ISO8601) of when the CDN endpoint was created.
	// +optional
	CreatedAt string `json:"createdAt,omitempty" tf:"created_at,omitempty"`
	// fully qualified domain name (FQDN) for custom subdomain, (requires certificate_id)
	// +optional
	CustomDomain string `json:"customDomain,omitempty" tf:"custom_domain,omitempty"`
	// fully qualified domain name (FQDN) to serve the CDN content
	// +optional
	Endpoint string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`
	// fully qualified domain name (FQDN) for the origin server
	Origin string `json:"origin" tf:"origin"`
	// The amount of time the content is cached in the CDN
	// +optional
	Ttl int `json:"ttl,omitempty" tf:"ttl,omitempty"`
}

type CdnStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *CdnSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CdnList is a list of Cdns
type CdnList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Cdn CRD objects
	Items []Cdn `json:"items,omitempty"`
}
