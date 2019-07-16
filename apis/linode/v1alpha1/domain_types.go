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

type Domain struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainSpec   `json:"spec,omitempty"`
	Status            DomainStatus `json:"status,omitempty"`
}

type DomainSpec struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	AxfrIps []string `json:"axfr_ips,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	Domain      string `json:"domain"`
	// +optional
	ExpireSec int `json:"expire_sec,omitempty"`
	// +optional
	Group string `json:"group,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	MasterIps []string `json:"master_ips,omitempty"`
	// +optional
	RefreshSec int `json:"refresh_sec,omitempty"`
	// +optional
	RetrySec int `json:"retry_sec,omitempty"`
	// +optional
	SoaEmail string `json:"soa_email,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Tags []string `json:"tags,omitempty"`
	// +optional
	TtlSec int    `json:"ttl_sec,omitempty"`
	Type   string `json:"type"`
}

type DomainStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
