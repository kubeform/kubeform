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

type GoogleBillingAccountIamMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleBillingAccountIamMemberSpec   `json:"spec,omitempty"`
	Status            GoogleBillingAccountIamMemberStatus `json:"status,omitempty"`
}

type GoogleBillingAccountIamMemberSpec struct {
	Etag             string `json:"etag"`
	Role             string `json:"role"`
	Member           string `json:"member"`
	BillingAccountId string `json:"billing_account_id"`
}



type GoogleBillingAccountIamMemberStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleBillingAccountIamMemberList is a list of GoogleBillingAccountIamMembers
type GoogleBillingAccountIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleBillingAccountIamMember CRD objects
	Items []GoogleBillingAccountIamMember `json:"items,omitempty"`
}