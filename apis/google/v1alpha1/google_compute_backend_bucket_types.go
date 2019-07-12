package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleComputeBackendBucket struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleComputeBackendBucketSpec   `json:"spec,omitempty"`
	Status            GoogleComputeBackendBucketStatus `json:"status,omitempty"`
}

type GoogleComputeBackendBucketSpec struct {
	Description       string `json:"description"`
	EnableCdn         bool   `json:"enable_cdn"`
	CreationTimestamp string `json:"creation_timestamp"`
	Project           string `json:"project"`
	SelfLink          string `json:"self_link"`
	BucketName        string `json:"bucket_name"`
	Name              string `json:"name"`
}



type GoogleComputeBackendBucketStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GoogleComputeBackendBucketList is a list of GoogleComputeBackendBuckets
type GoogleComputeBackendBucketList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleComputeBackendBucket CRD objects
	Items []GoogleComputeBackendBucket `json:"items,omitempty"`
}