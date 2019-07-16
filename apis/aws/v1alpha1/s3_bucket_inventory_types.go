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

type S3BucketInventory struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              S3BucketInventorySpec   `json:"spec,omitempty"`
	Status            S3BucketInventoryStatus `json:"status,omitempty"`
}

type S3BucketInventorySpecDestinationBucketEncryptionSseKms struct {
	KeyId string `json:"key_id"`
}

type S3BucketInventorySpecDestinationBucketEncryptionSseS3 struct{}

type S3BucketInventorySpecDestinationBucketEncryption struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SseKms *[]S3BucketInventorySpecDestinationBucketEncryption `json:"sse_kms,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SseS3 *[]S3BucketInventorySpecDestinationBucketEncryption `json:"sse_s3,omitempty"`
}

type S3BucketInventorySpecDestinationBucket struct {
	// +optional
	AccountId string `json:"account_id,omitempty"`
	BucketArn string `json:"bucket_arn"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Encryption *[]S3BucketInventorySpecDestinationBucket `json:"encryption,omitempty"`
	Format     string                                    `json:"format"`
	// +optional
	Prefix string `json:"prefix,omitempty"`
}

type S3BucketInventorySpecDestination struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Bucket []S3BucketInventorySpecDestination `json:"bucket"`
}

type S3BucketInventorySpecFilter struct {
	// +optional
	Prefix string `json:"prefix,omitempty"`
}

type S3BucketInventorySpecSchedule struct {
	Frequency string `json:"frequency"`
}

type S3BucketInventorySpec struct {
	Bucket string `json:"bucket"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Destination []S3BucketInventorySpec `json:"destination"`
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Filter                 *[]S3BucketInventorySpec `json:"filter,omitempty"`
	IncludedObjectVersions string                   `json:"included_object_versions"`
	Name                   string                   `json:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	OptionalFields []string `json:"optional_fields,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Schedule []S3BucketInventorySpec `json:"schedule"`
}

type S3BucketInventoryStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// S3BucketInventoryList is a list of S3BucketInventorys
type S3BucketInventoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of S3BucketInventory CRD objects
	Items []S3BucketInventory `json:"items,omitempty"`
}
