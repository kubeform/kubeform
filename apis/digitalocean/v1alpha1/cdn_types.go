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

type Cdn struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CdnSpec   `json:"spec,omitempty"`
	Status            CdnStatus `json:"status,omitempty"`
}

type CdnSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	CertificateID string `json:"certificateID,omitempty" tf:"certificate_id,omitempty"`
	CreatedAt     string `json:"createdAt" tf:"created_at"`
	// +optional
	CustomDomain string `json:"customDomain,omitempty" tf:"custom_domain,omitempty"`
	Endpoint     string `json:"endpoint" tf:"endpoint"`
	Origin       string `json:"origin" tf:"origin"`
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
	State *apis.State `json:"state,omitempty"`
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
