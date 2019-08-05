package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
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
	KeyID string `json:"keyID" tf:"key_id"`
}

type S3BucketInventorySpecDestinationBucketEncryptionSseS3 struct{}

type S3BucketInventorySpecDestinationBucketEncryption struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SseKms []S3BucketInventorySpecDestinationBucketEncryptionSseKms `json:"sseKms,omitempty" tf:"sse_kms,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	SseS3 []S3BucketInventorySpecDestinationBucketEncryptionSseS3 `json:"sseS3,omitempty" tf:"sse_s3,omitempty"`
}

type S3BucketInventorySpecDestinationBucket struct {
	// +optional
	AccountID string `json:"accountID,omitempty" tf:"account_id,omitempty"`
	BucketArn string `json:"bucketArn" tf:"bucket_arn"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Encryption []S3BucketInventorySpecDestinationBucketEncryption `json:"encryption,omitempty" tf:"encryption,omitempty"`
	Format     string                                             `json:"format" tf:"format"`
	// +optional
	Prefix string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type S3BucketInventorySpecDestination struct {
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Bucket []S3BucketInventorySpecDestinationBucket `json:"bucket" tf:"bucket"`
}

type S3BucketInventorySpecFilter struct {
	// +optional
	Prefix string `json:"prefix,omitempty" tf:"prefix,omitempty"`
}

type S3BucketInventorySpecSchedule struct {
	Frequency string `json:"frequency" tf:"frequency"`
}

type S3BucketInventorySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Bucket string `json:"bucket" tf:"bucket"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Destination []S3BucketInventorySpecDestination `json:"destination" tf:"destination"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Filter                 []S3BucketInventorySpecFilter `json:"filter,omitempty" tf:"filter,omitempty"`
	IncludedObjectVersions string                        `json:"includedObjectVersions" tf:"included_object_versions"`
	Name                   string                        `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	OptionalFields []string `json:"optionalFields,omitempty" tf:"optional_fields,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:MinItems=1
	Schedule []S3BucketInventorySpecSchedule `json:"schedule" tf:"schedule"`
}

type S3BucketInventoryStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *S3BucketInventorySpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
