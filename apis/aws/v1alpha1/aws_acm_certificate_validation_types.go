package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAcmCertificateValidation struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAcmCertificateValidationSpec   `json:"spec,omitempty"`
	Status            AwsAcmCertificateValidationStatus `json:"status,omitempty"`
}

type AwsAcmCertificateValidationSpec struct {
	CertificateArn        string   `json:"certificate_arn"`
	ValidationRecordFqdns []string `json:"validation_record_fqdns"`
}

type AwsAcmCertificateValidationStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAcmCertificateValidationList is a list of AwsAcmCertificateValidations
type AwsAcmCertificateValidationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAcmCertificateValidation CRD objects
	Items []AwsAcmCertificateValidation `json:"items,omitempty"`
}
