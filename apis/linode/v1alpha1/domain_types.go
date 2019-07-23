package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type Domain struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainSpec   `json:"spec,omitempty"`
	Status            DomainStatus `json:"status,omitempty"`
}

type DomainSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AxfrIPS []string `json:"axfrIPS,omitempty" tf:"axfr_ips,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Domain      string `json:"domain" tf:"domain"`
	// +optional
	ExpireSec int `json:"expireSec,omitempty" tf:"expire_sec,omitempty"`
	// +optional
	Group string `json:"group,omitempty" tf:"group,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	MasterIPS []string `json:"masterIPS,omitempty" tf:"master_ips,omitempty"`
	// +optional
	RefreshSec int `json:"refreshSec,omitempty" tf:"refresh_sec,omitempty"`
	// +optional
	RetrySec int `json:"retrySec,omitempty" tf:"retry_sec,omitempty"`
	// +optional
	SoaEmail string `json:"soaEmail,omitempty" tf:"soa_email,omitempty"`
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Tags []string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	TtlSec int    `json:"ttlSec,omitempty" tf:"ttl_sec,omitempty"`
	Type   string `json:"type" tf:"type"`
}

type DomainStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DomainList is a list of Domains
type DomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Domain CRD objects
	Items []Domain `json:"items,omitempty"`
}
