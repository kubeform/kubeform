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

type SsmResourceDataSync struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SsmResourceDataSyncSpec   `json:"spec,omitempty"`
	Status            SsmResourceDataSyncStatus `json:"status,omitempty"`
}

type SsmResourceDataSyncSpecS3Destination struct {
	BucketName string `json:"bucketName" tf:"bucket_name"`
	// +optional
	KmsKeyArn string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`
	// +optional
	Prefix string `json:"prefix,omitempty" tf:"prefix,omitempty"`
	Region string `json:"region" tf:"region"`
	// +optional
	SyncFormat string `json:"syncFormat,omitempty" tf:"sync_format,omitempty"`
}

type SsmResourceDataSyncSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Name string `json:"name" tf:"name"`
	// +kubebuilder:validation:MaxItems=1
	S3Destination []SsmResourceDataSyncSpecS3Destination `json:"s3Destination" tf:"s3_destination"`
}

type SsmResourceDataSyncStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SsmResourceDataSyncSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
