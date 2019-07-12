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

type GoogleProjectUsageExportBucket struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleProjectUsageExportBucketSpec   `json:"spec,omitempty"`
	Status            GoogleProjectUsageExportBucketStatus `json:"status,omitempty"`
}

type GoogleProjectUsageExportBucketSpec struct {
	BucketName string `json:"bucket_name"`
	Prefix     string `json:"prefix"`
	Project    string `json:"project"`
}

type GoogleProjectUsageExportBucketStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleProjectUsageExportBucketList is a list of GoogleProjectUsageExportBuckets
type GoogleProjectUsageExportBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleProjectUsageExportBucket CRD objects
	Items []GoogleProjectUsageExportBucket `json:"items,omitempty"`
}
