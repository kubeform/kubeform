package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GooglePubsubSubscriptionIamMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GooglePubsubSubscriptionIamMemberSpec   `json:"spec,omitempty"`
	Status            GooglePubsubSubscriptionIamMemberStatus `json:"status,omitempty"`
}

type GooglePubsubSubscriptionIamMemberSpec struct {
	Project      string `json:"project"`
	Role         string `json:"role"`
	Member       string `json:"member"`
	Etag         string `json:"etag"`
	Subscription string `json:"subscription"`
}

type GooglePubsubSubscriptionIamMemberStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GooglePubsubSubscriptionIamMemberList is a list of GooglePubsubSubscriptionIamMembers
type GooglePubsubSubscriptionIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GooglePubsubSubscriptionIamMember CRD objects
	Items []GooglePubsubSubscriptionIamMember `json:"items,omitempty"`
}
