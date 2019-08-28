package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ProjectUsageExportBucket struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectUsageExportBucketSpec   `json:"spec,omitempty"`
	Status            ProjectUsageExportBucketStatus `json:"status,omitempty"`
}

type ProjectUsageExportBucketSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	BucketName string `json:"bucketName" tf:"bucket_name"`
	// +optional
	Prefix string `json:"prefix,omitempty" tf:"prefix,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
}



type ProjectUsageExportBucketStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *ProjectUsageExportBucketSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ProjectUsageExportBucketList is a list of ProjectUsageExportBuckets
type ProjectUsageExportBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ProjectUsageExportBucket CRD objects
	Items []ProjectUsageExportBucket `json:"items,omitempty"`
}