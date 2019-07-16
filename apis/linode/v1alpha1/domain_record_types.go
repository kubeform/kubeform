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

type DomainRecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainRecordSpec   `json:"spec,omitempty"`
	Status            DomainRecordStatus `json:"status,omitempty"`
}

type DomainRecordSpec struct {
	DomainId int    `json:"domain_id"`
	Name     string `json:"name"`
	// +optional
	Port int `json:"port,omitempty"`
	// +optional
	Priority int `json:"priority,omitempty"`
	// +optional
	Protocol   string `json:"protocol,omitempty"`
	RecordType string `json:"record_type"`
	// +optional
	Service string `json:"service,omitempty"`
	// +optional
	Tag    string `json:"tag,omitempty"`
	Target string `json:"target"`
	// +optional
	TtlSec int `json:"ttl_sec,omitempty"`
	// +optional
	Weight int `json:"weight,omitempty"`
}

type DomainRecordStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
