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

type AwsSesDomainIdentity struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSesDomainIdentitySpec   `json:"spec,omitempty"`
	Status            AwsSesDomainIdentityStatus `json:"status,omitempty"`
}

type AwsSesDomainIdentitySpec struct {
	Arn               string `json:"arn"`
	Domain            string `json:"domain"`
	VerificationToken string `json:"verification_token"`
}



type AwsSesDomainIdentityStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSesDomainIdentityList is a list of AwsSesDomainIdentitys
type AwsSesDomainIdentityList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSesDomainIdentity CRD objects
	Items []AwsSesDomainIdentity `json:"items,omitempty"`
}