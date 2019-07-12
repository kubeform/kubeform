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

type GooglePubsubTopicIamMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GooglePubsubTopicIamMemberSpec   `json:"spec,omitempty"`
	Status            GooglePubsubTopicIamMemberStatus `json:"status,omitempty"`
}

type GooglePubsubTopicIamMemberSpec struct {
	Member  string `json:"member"`
	Topic   string `json:"topic"`
	Project string `json:"project"`
	Etag    string `json:"etag"`
	Role    string `json:"role"`
}

type GooglePubsubTopicIamMemberStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GooglePubsubTopicIamMemberList is a list of GooglePubsubTopicIamMembers
type GooglePubsubTopicIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GooglePubsubTopicIamMember CRD objects
	Items []GooglePubsubTopicIamMember `json:"items,omitempty"`
}
