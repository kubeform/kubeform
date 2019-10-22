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

type Domain struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainSpec   `json:"spec,omitempty"`
	Status            DomainStatus `json:"status,omitempty"`
}

type DomainSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// The list of IPs that may perform a zone transfer for this Domain. This is potentially dangerous, and should be set to an empty list unless you intend to use it.
	// +optional
	AxfrIPS []string `json:"axfrIPS,omitempty" tf:"axfr_ips,omitempty"`
	// A description for this Domain. This is for display purposes only.
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// The domain this Domain represents. These must be unique in our system; you cannot have two Domains representing the same domain.
	Domain string `json:"domain" tf:"domain"`
	// The amount of time in seconds that may pass before this Domain is no longer authoritative. Valid values are 0, 00, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	// +optional
	ExpireSec int `json:"expireSec,omitempty" tf:"expire_sec,omitempty"`
	// The group this Domain belongs to. This is for display purposes only.
	// +optional
	Group string `json:"group,omitempty" tf:"group,omitempty"`
	// The IP addresses representing the master DNS for this Domain.
	// +optional
	MasterIPS []string `json:"masterIPS,omitempty" tf:"master_ips,omitempty"`
	// The amount of time in seconds before this Domain should be refreshed. Valid values are 0, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	// +optional
	RefreshSec int `json:"refreshSec,omitempty" tf:"refresh_sec,omitempty"`
	// The interval, in seconds, at which a failed refresh should be retried. Valid values are 0, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	// +optional
	RetrySec int `json:"retrySec,omitempty" tf:"retry_sec,omitempty"`
	// Start of Authority email address. This is required for master Domains.
	// +optional
	SoaEmail string `json:"soaEmail,omitempty" tf:"soa_email,omitempty"`
	// Used to control whether this Domain is currently being rendered.
	// +optional
	Status string `json:"status,omitempty" tf:"status,omitempty"`
	// An array of tags applied to this object. Tags are for organizational purposes only.
	// +optional
	Tags []string `json:"tags,omitempty" tf:"tags,omitempty"`
	// 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 0, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	// +optional
	TtlSec int `json:"ttlSec,omitempty" tf:"ttl_sec,omitempty"`
	// If this Domain represents the authoritative source of information for the domain it describes, or if it is a read-only copy of a master (also called a slave).
	Type string `json:"type" tf:"type"`
}

type DomainStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DomainSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
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
