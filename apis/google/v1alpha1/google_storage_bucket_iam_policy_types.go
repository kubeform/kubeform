package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleStorageBucketIamPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleStorageBucketIamPolicySpec   `json:"spec,omitempty"`
	Status            GoogleStorageBucketIamPolicyStatus `json:"status,omitempty"`
}

type GoogleStorageBucketIamPolicySpec struct {
	Bucket     string `json:"bucket"`
	PolicyData string `json:"policy_data"`
	Etag       string `json:"etag"`
}



type GoogleStorageBucketIamPolicyStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleStorageBucketIamPolicyList is a list of GoogleStorageBucketIamPolicys
type GoogleStorageBucketIamPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleStorageBucketIamPolicy CRD objects
	Items []GoogleStorageBucketIamPolicy `json:"items,omitempty"`
}