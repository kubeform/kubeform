package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type LinodeDomainRecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinodeDomainRecordSpec   `json:"spec,omitempty"`
	Status            LinodeDomainRecordStatus `json:"status,omitempty"`
}

type LinodeDomainRecordSpec struct {
	DomainId   int    `json:"domain_id"`
	TtlSec     int    `json:"ttl_sec"`
	Priority   int    `json:"priority"`
	Weight     int    `json:"weight"`
	Port       int    `json:"port"`
	Name       string `json:"name"`
	RecordType string `json:"record_type"`
	Target     string `json:"target"`
	Protocol   string `json:"protocol"`
	Service    string `json:"service"`
	Tag        string `json:"tag"`
}



type LinodeDomainRecordStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LinodeDomainRecordList is a list of LinodeDomainRecords
type LinodeDomainRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LinodeDomainRecord CRD objects
	Items []LinodeDomainRecord `json:"items,omitempty"`
}