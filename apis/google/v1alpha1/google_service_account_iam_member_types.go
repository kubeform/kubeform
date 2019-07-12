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

type GoogleServiceAccountIamMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleServiceAccountIamMemberSpec   `json:"spec,omitempty"`
	Status            GoogleServiceAccountIamMemberStatus `json:"status,omitempty"`
}

type GoogleServiceAccountIamMemberSpec struct {
	ServiceAccountId string `json:"service_account_id"`
	Role             string `json:"role"`
	Member           string `json:"member"`
	Etag             string `json:"etag"`
}

type GoogleServiceAccountIamMemberStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleServiceAccountIamMemberList is a list of GoogleServiceAccountIamMembers
type GoogleServiceAccountIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleServiceAccountIamMember CRD objects
	Items []GoogleServiceAccountIamMember `json:"items,omitempty"`
}
