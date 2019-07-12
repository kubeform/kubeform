package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleKmsCryptoKeyIamMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleKmsCryptoKeyIamMemberSpec   `json:"spec,omitempty"`
	Status            GoogleKmsCryptoKeyIamMemberStatus `json:"status,omitempty"`
}

type GoogleKmsCryptoKeyIamMemberSpec struct {
	Member      string `json:"member"`
	Etag        string `json:"etag"`
	CryptoKeyId string `json:"crypto_key_id"`
	Role        string `json:"role"`
}



type GoogleKmsCryptoKeyIamMemberStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleKmsCryptoKeyIamMemberList is a list of GoogleKmsCryptoKeyIamMembers
type GoogleKmsCryptoKeyIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleKmsCryptoKeyIamMember CRD objects
	Items []GoogleKmsCryptoKeyIamMember `json:"items,omitempty"`
}