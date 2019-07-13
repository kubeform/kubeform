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

type AwsS3BucketInventory struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsS3BucketInventorySpec   `json:"spec,omitempty"`
	Status            AwsS3BucketInventoryStatus `json:"status,omitempty"`
}

type AwsS3BucketInventorySpecFilter struct {
	Prefix string `json:"prefix"`
}

type AwsS3BucketInventorySpecDestinationBucketEncryptionSseS3 struct{}

type AwsS3BucketInventorySpecDestinationBucketEncryptionSseKms struct {
	KeyId string `json:"key_id"`
}

type AwsS3BucketInventorySpecDestinationBucketEncryption struct {
	SseS3  []AwsS3BucketInventorySpecDestinationBucketEncryption `json:"sse_s3"`
	SseKms []AwsS3BucketInventorySpecDestinationBucketEncryption `json:"sse_kms"`
}

type AwsS3BucketInventorySpecDestinationBucket struct {
	Encryption []AwsS3BucketInventorySpecDestinationBucket `json:"encryption"`
	Format     string                                      `json:"format"`
	BucketArn  string                                      `json:"bucket_arn"`
	AccountId  string                                      `json:"account_id"`
	Prefix     string                                      `json:"prefix"`
}

type AwsS3BucketInventorySpecDestination struct {
	Bucket []AwsS3BucketInventorySpecDestination `json:"bucket"`
}

type AwsS3BucketInventorySpecSchedule struct {
	Frequency string `json:"frequency"`
}

type AwsS3BucketInventorySpec struct {
	Filter                 []AwsS3BucketInventorySpec `json:"filter"`
	Destination            []AwsS3BucketInventorySpec `json:"destination"`
	Schedule               []AwsS3BucketInventorySpec `json:"schedule"`
	IncludedObjectVersions string                     `json:"included_object_versions"`
	OptionalFields         []string                   `json:"optional_fields"`
	Bucket                 string                     `json:"bucket"`
	Name                   string                     `json:"name"`
	Enabled                bool                       `json:"enabled"`
}



type AwsS3BucketInventoryStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsS3BucketInventoryList is a list of AwsS3BucketInventorys
type AwsS3BucketInventoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsS3BucketInventory CRD objects
	Items []AwsS3BucketInventory `json:"items,omitempty"`
}