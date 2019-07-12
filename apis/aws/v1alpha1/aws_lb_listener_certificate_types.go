package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsLbListenerCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsLbListenerCertificateSpec   `json:"spec,omitempty"`
	Status            AwsLbListenerCertificateStatus `json:"status,omitempty"`
}

type AwsLbListenerCertificateSpec struct {
	ListenerArn    string `json:"listener_arn"`
	CertificateArn string `json:"certificate_arn"`
}

type AwsLbListenerCertificateStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbListenerCertificateList is a list of AwsLbListenerCertificates
type AwsLbListenerCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsLbListenerCertificate CRD objects
	Items []AwsLbListenerCertificate `json:"items,omitempty"`
}
