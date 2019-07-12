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
	Fqdn     string `json:"fqdn"`
	Tag      string `json:"tag"`
	Port     int    `json:"port"`
	Priority int    `json:"priority"`
	Weight   int    `json:"weight"`
	Ttl      int    `json:"ttl"`
	Value    string `json:"value"`
	Flags    int    `json:"flags"`
	Type     string `json:"type"`
	Domain   string `json:"domain"`
	Name     string `json:"name"`
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