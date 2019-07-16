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

type SsmResourceDataSync struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SsmResourceDataSyncSpec   `json:"spec,omitempty"`
	Status            SsmResourceDataSyncStatus `json:"status,omitempty"`
}

type SsmResourceDataSyncSpecS3Destination struct {
	BucketName string `json:"bucket_name"`
	// +optional
	KmsKeyArn string `json:"kms_key_arn,omitempty"`
	// +optional
	Prefix string `json:"prefix,omitempty"`
	Region string `json:"region"`
	// +optional
	SyncFormat string `json:"sync_format,omitempty"`
}

type SsmResourceDataSyncSpec struct {
	Name string `json:"name"`
	// +kubebuilder:validation:MaxItems=1
	S3Destination []SsmResourceDataSyncSpec `json:"s3_destination"`
}

type SsmResourceDataSyncStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SsmResourceDataSyncList is a list of SsmResourceDataSyncs
type SsmResourceDataSyncList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SsmResourceDataSync CRD objects
	Items []SsmResourceDataSync `json:"items,omitempty"`
}
