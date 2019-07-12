package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSecurityhubAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSecurityhubAccountSpec   `json:"spec,omitempty"`
	Status            AwsSecurityhubAccountStatus `json:"status,omitempty"`
}

type AwsSecurityhubAccountSpec struct{}



type AwsSecurityhubAccountStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSecurityhubAccountList is a list of AwsSecurityhubAccounts
type AwsSecurityhubAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSecurityhubAccount CRD objects
	Items []AwsSecurityhubAccount `json:"items,omitempty"`
}