package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsS3BucketInventory struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsS3BucketInventorySpec   `json:"spec,omitempty"`
	Status            AwsS3BucketInventoryStatus `json:"status,omitempty"`
}

type AwsS3BucketInventorySpecFilter struct {
	Prefix string `json:"prefix"`
}

type AwsS3BucketInventorySpecDestinationBucketEncryptionSseKms struct {
	KeyId string `json:"key_id"`
}

type AwsS3BucketInventorySpecDestinationBucketEncryptionSseS3 struct{}

type AwsS3BucketInventorySpecDestinationBucketEncryption struct {
	SseKms []AwsS3BucketInventorySpecDestinationBucketEncryption `json:"sse_kms"`
	SseS3  []AwsS3BucketInventorySpecDestinationBucketEncryption `json:"sse_s3"`
}

type AwsS3BucketInventorySpecDestinationBucket struct {
	Format     string                                      `json:"format"`
	BucketArn  string                                      `json:"bucket_arn"`
	AccountId  string                                      `json:"account_id"`
	Prefix     string                                      `json:"prefix"`
	Encryption []AwsS3BucketInventorySpecDestinationBucket `json:"encryption"`
}

type AwsS3BucketInventorySpecDestination struct {
	Bucket []AwsS3BucketInventorySpecDestination `json:"bucket"`
}

type AwsS3BucketInventorySpecSchedule struct {
	Frequency string `json:"frequency"`
}

type AwsS3BucketInventorySpec struct {
	Enabled                bool                       `json:"enabled"`
	Filter                 []AwsS3BucketInventorySpec `json:"filter"`
	Destination            []AwsS3BucketInventorySpec `json:"destination"`
	Schedule               []AwsS3BucketInventorySpec `json:"schedule"`
	IncludedObjectVersions string                     `json:"included_object_versions"`
	OptionalFields         []string                   `json:"optional_fields"`
	Bucket                 string                     `json:"bucket"`
	Name                   string                     `json:"name"`
}

type AwsS3BucketInventoryStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsS3BucketInventoryList is a list of AwsS3BucketInventorys
type AwsS3BucketInventoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsS3BucketInventory CRD objects
	Items []AwsS3BucketInventory `json:"items,omitempty"`
}
