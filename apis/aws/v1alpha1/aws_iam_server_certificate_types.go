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

type AwsIamServerCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIamServerCertificateSpec   `json:"spec,omitempty"`
	Status            AwsIamServerCertificateStatus `json:"status,omitempty"`
}

type AwsIamServerCertificateSpec struct {
	Name             string `json:"name"`
	NamePrefix       string `json:"name_prefix"`
	Arn              string `json:"arn"`
	CertificateBody  string `json:"certificate_body"`
	CertificateChain string `json:"certificate_chain"`
	Path             string `json:"path"`
	PrivateKey       string `json:"private_key"`
}



type AwsIamServerCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsIamServerCertificateList is a list of AwsIamServerCertificates
type AwsIamServerCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIamServerCertificate CRD objects
	Items []AwsIamServerCertificate `json:"items,omitempty"`
}