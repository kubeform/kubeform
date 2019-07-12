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

type AwsS3BucketPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsS3BucketPolicySpec   `json:"spec,omitempty"`
	Status            AwsS3BucketPolicyStatus `json:"status,omitempty"`
}

type AwsS3BucketPolicySpec struct {
	Bucket string `json:"bucket"`
	Policy string `json:"policy"`
}

type AwsS3BucketPolicyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsS3BucketPolicyList is a list of AwsS3BucketPolicys
type AwsS3BucketPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsS3BucketPolicy CRD objects
	Items []AwsS3BucketPolicy `json:"items,omitempty"`
}
