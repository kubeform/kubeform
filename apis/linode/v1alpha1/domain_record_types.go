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

type DomainRecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainRecordSpec   `json:"spec,omitempty"`
	Status            DomainRecordStatus `json:"status,omitempty"`
}

type DomainRecordSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	DomainID int    `json:"domainID" tf:"domain_id"`
	Name     string `json:"name" tf:"name"`
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	Priority int `json:"priority,omitempty" tf:"priority,omitempty"`
	// +optional
	Protocol   string `json:"protocol,omitempty" tf:"protocol,omitempty"`
	RecordType string `json:"recordType" tf:"record_type"`
	// +optional
	Service string `json:"service,omitempty" tf:"service,omitempty"`
	// +optional
	Tag    string `json:"tag,omitempty" tf:"tag,omitempty"`
	Target string `json:"target" tf:"target"`
	// +optional
	TtlSec int `json:"ttlSec,omitempty" tf:"ttl_sec,omitempty"`
	// +optional
	Weight int `json:"weight,omitempty" tf:"weight,omitempty"`
}

type DomainRecordStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DomainRecordSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DomainRecordList is a list of DomainRecords
type DomainRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DomainRecord CRD objects
	Items []DomainRecord `json:"items,omitempty"`
}
