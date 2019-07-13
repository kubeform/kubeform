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

type AwsSsmResourceDataSync struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSsmResourceDataSyncSpec   `json:"spec,omitempty"`
	Status            AwsSsmResourceDataSyncStatus `json:"status,omitempty"`
}

type AwsSsmResourceDataSyncSpecS3Destination struct {
	SyncFormat string `json:"sync_format"`
	KmsKeyArn  string `json:"kms_key_arn"`
	BucketName string `json:"bucket_name"`
	Prefix     string `json:"prefix"`
	Region     string `json:"region"`
}

type AwsSsmResourceDataSyncSpec struct {
	Name          string                       `json:"name"`
	S3Destination []AwsSsmResourceDataSyncSpec `json:"s3_destination"`
}



type AwsSsmResourceDataSyncStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSsmResourceDataSyncList is a list of AwsSsmResourceDataSyncs
type AwsSsmResourceDataSyncList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSsmResourceDataSync CRD objects
	Items []AwsSsmResourceDataSync `json:"items,omitempty"`
}