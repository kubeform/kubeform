package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DigitaloceanRecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DigitaloceanRecordSpec   `json:"spec,omitempty"`
	Status            DigitaloceanRecordStatus `json:"status,omitempty"`
}

type DigitaloceanRecordSpec struct {
	Type     string `json:"type"`
	Name     string `json:"name"`
	Weight   int    `json:"weight"`
	Ttl      int    `json:"ttl"`
	Tag      string `json:"tag"`
	Domain   string `json:"domain"`
	Port     int    `json:"port"`
	Priority int    `json:"priority"`
	Value    string `json:"value"`
	Fqdn     string `json:"fqdn"`
	Flags    int    `json:"flags"`
}

type DigitaloceanRecordStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DigitaloceanRecordList is a list of DigitaloceanRecords
type DigitaloceanRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DigitaloceanRecord CRD objects
	Items []DigitaloceanRecord `json:"items,omitempty"`
}
