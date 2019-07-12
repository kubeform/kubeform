package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GooglePubsubTopicIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GooglePubsubTopicIamBindingSpec   `json:"spec,omitempty"`
	Status            GooglePubsubTopicIamBindingStatus `json:"status,omitempty"`
}

type GooglePubsubTopicIamBindingSpec struct {
	Role    string   `json:"role"`
	Members []string `json:"members"`
	Etag    string   `json:"etag"`
	Topic   string   `json:"topic"`
	Project string   `json:"project"`
}



type GooglePubsubTopicIamBindingStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GooglePubsubTopicIamBindingList is a list of GooglePubsubTopicIamBindings
type GooglePubsubTopicIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GooglePubsubTopicIamBinding CRD objects
	Items []GooglePubsubTopicIamBinding `json:"items,omitempty"`
}