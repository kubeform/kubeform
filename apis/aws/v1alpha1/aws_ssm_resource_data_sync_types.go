package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmResourceDataSync struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSsmResourceDataSyncSpec   `json:"spec,omitempty"`
	Status            AwsSsmResourceDataSyncStatus `json:"status,omitempty"`
}

type AwsSsmResourceDataSyncSpecS3Destination struct {
	KmsKeyArn  string `json:"kms_key_arn"`
	BucketName string `json:"bucket_name"`
	Prefix     string `json:"prefix"`
	Region     string `json:"region"`
	SyncFormat string `json:"sync_format"`
}

type AwsSsmResourceDataSyncSpec struct {
	Name          string                       `json:"name"`
	S3Destination []AwsSsmResourceDataSyncSpec `json:"s3_destination"`
}

type AwsSsmResourceDataSyncStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmResourceDataSyncList is a list of AwsSsmResourceDataSyncs
type AwsSsmResourceDataSyncList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSsmResourceDataSync CRD objects
	Items []AwsSsmResourceDataSync `json:"items,omitempty"`
}
