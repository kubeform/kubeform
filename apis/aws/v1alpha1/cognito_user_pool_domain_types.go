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

type CognitoUserPoolDomain struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CognitoUserPoolDomainSpec   `json:"spec,omitempty"`
	Status            CognitoUserPoolDomainStatus `json:"status,omitempty"`
}

type CognitoUserPoolDomainSpec struct {
	// +optional
	CertificateArn string `json:"certificate_arn,omitempty"`
	Domain         string `json:"domain"`
	UserPoolId     string `json:"user_pool_id"`
}

type CognitoUserPoolDomainStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CognitoUserPoolDomainList is a list of CognitoUserPoolDomains
type CognitoUserPoolDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CognitoUserPoolDomain CRD objects
	Items []CognitoUserPoolDomain `json:"items,omitempty"`
}