package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GooglePubsubTopicIamPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GooglePubsubTopicIamPolicySpec   `json:"spec,omitempty"`
	Status            GooglePubsubTopicIamPolicyStatus `json:"status,omitempty"`
}

type GooglePubsubTopicIamPolicySpec struct {
	PolicyData string `json:"policy_data"`
	Etag       string `json:"etag"`
	Topic      string `json:"topic"`
	Project    string `json:"project"`
}



type GooglePubsubTopicIamPolicyStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GooglePubsubTopicIamPolicyList is a list of GooglePubsubTopicIamPolicys
type GooglePubsubTopicIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GooglePubsubTopicIamPolicy CRD objects
	Items []GooglePubsubTopicIamPolicy `json:"items,omitempty"`
}