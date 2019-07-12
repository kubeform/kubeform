package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsS3BucketPublicAccessBlock struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsS3BucketPublicAccessBlockSpec   `json:"spec,omitempty"`
	Status            AwsS3BucketPublicAccessBlockStatus `json:"status,omitempty"`
}

type AwsS3BucketPublicAccessBlockSpec struct {
	Bucket                string `json:"bucket"`
	BlockPublicAcls       bool   `json:"block_public_acls"`
	BlockPublicPolicy     bool   `json:"block_public_policy"`
	IgnorePublicAcls      bool   `json:"ignore_public_acls"`
	RestrictPublicBuckets bool   `json:"restrict_public_buckets"`
}



type AwsS3BucketPublicAccessBlockStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsS3BucketPublicAccessBlockList is a list of AwsS3BucketPublicAccessBlocks
type AwsS3BucketPublicAccessBlockList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsS3BucketPublicAccessBlock CRD objects
	Items []AwsS3BucketPublicAccessBlock `json:"items,omitempty"`
}