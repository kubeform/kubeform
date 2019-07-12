package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSimpledbDomain struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSimpledbDomainSpec   `json:"spec,omitempty"`
	Status            AwsSimpledbDomainStatus `json:"status,omitempty"`
}

type AwsSimpledbDomainSpec struct {
	Name string `json:"name"`
}



type AwsSimpledbDomainStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSimpledbDomainList is a list of AwsSimpledbDomains
type AwsSimpledbDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSimpledbDomain CRD objects
	Items []AwsSimpledbDomain `json:"items,omitempty"`
}