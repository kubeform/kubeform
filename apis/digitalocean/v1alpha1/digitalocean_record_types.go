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

type DigitaloceanRecord struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DigitaloceanRecordSpec   `json:"spec,omitempty"`
	Status            DigitaloceanRecordStatus `json:"status,omitempty"`
}

type DigitaloceanRecordSpec struct {
	Ttl      int    `json:"ttl"`
	Fqdn     string `json:"fqdn"`
	Flags    int    `json:"flags"`
	Type     string `json:"type"`
	Domain   string `json:"domain"`
	Name     string `json:"name"`
	Port     int    `json:"port"`
	Weight   int    `json:"weight"`
	Tag      string `json:"tag"`
	Priority int    `json:"priority"`
	Value    string `json:"value"`
}



type DigitaloceanRecordStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DigitaloceanRecordList is a list of DigitaloceanRecords
type DigitaloceanRecordList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DigitaloceanRecord CRD objects
	Items []DigitaloceanRecord `json:"items,omitempty"`
}