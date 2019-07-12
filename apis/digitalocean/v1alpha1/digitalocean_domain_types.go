package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DigitaloceanDomain struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DigitaloceanDomainSpec   `json:"spec,omitempty"`
	Status            DigitaloceanDomainStatus `json:"status,omitempty"`
}

type DigitaloceanDomainSpec struct {
	Urn       string `json:"urn"`
	Name      string `json:"name"`
	IpAddress string `json:"ip_address"`
}



type DigitaloceanDomainStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DigitaloceanDomainList is a list of DigitaloceanDomains
type DigitaloceanDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DigitaloceanDomain CRD objects
	Items []DigitaloceanDomain `json:"items,omitempty"`
}