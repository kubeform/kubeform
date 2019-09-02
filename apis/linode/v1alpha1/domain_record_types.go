package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
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

	// The ID of the Domain to access.
	DomainID int `json:"domainID" tf:"domain_id"`
	// The name of this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the subdomain being associated with an IP address.
	Name string `json:"name" tf:"name"`
	// The port this Record points to.
	// +optional
	Port int `json:"port,omitempty" tf:"port,omitempty"`
	// The priority of the target host. Lower values are preferred.
	// +optional
	Priority int `json:"priority,omitempty" tf:"priority,omitempty"`
	// The protocol this Record's service communicates with. Only valid for SRV records.
	// +optional
	Protocol string `json:"protocol,omitempty" tf:"protocol,omitempty"`
	// The type of Record this is in the DNS system. For example, A records associate a domain name with an IPv4 address, and AAAA records associate a domain name with an IPv6 address.
	RecordType string `json:"recordType" tf:"record_type"`
	// The service this Record identified. Only valid for SRV records.
	// +optional
	Service string `json:"service,omitempty" tf:"service,omitempty"`
	// The tag portion of a CAA record. It is invalid to set this on other record types.
	// +optional
	Tag string `json:"tag,omitempty" tf:"tag,omitempty"`
	// The target for this Record. This field's actual usage depends on the type of record this represents. For A and AAAA records, this is the address the named Domain should resolve to.
	Target string `json:"target" tf:"target"`
	// 'Time to Live' - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 0, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	// +optional
	TtlSec int `json:"ttlSec,omitempty" tf:"ttl_sec,omitempty"`
	// The relative weight of this Record. Higher values are preferred.
	// +optional
	Weight int `json:"weight,omitempty" tf:"weight,omitempty"`
}



type DomainRecordStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
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