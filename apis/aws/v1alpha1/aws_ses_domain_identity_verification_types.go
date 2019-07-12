package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesDomainIdentityVerification struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSesDomainIdentityVerificationSpec   `json:"spec,omitempty"`
	Status            AwsSesDomainIdentityVerificationStatus `json:"status,omitempty"`
}

type AwsSesDomainIdentityVerificationSpec struct {
	Arn    string `json:"arn"`
	Domain string `json:"domain"`
}

type AwsSesDomainIdentityVerificationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesDomainIdentityVerificationList is a list of AwsSesDomainIdentityVerifications
type AwsSesDomainIdentityVerificationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSesDomainIdentityVerification CRD objects
	Items []AwsSesDomainIdentityVerification `json:"items,omitempty"`
}
