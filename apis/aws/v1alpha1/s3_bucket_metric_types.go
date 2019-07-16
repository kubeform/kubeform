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

type S3BucketMetric struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              S3BucketMetricSpec   `json:"spec,omitempty"`
	Status            S3BucketMetricStatus `json:"status,omitempty"`
}

type S3BucketMetricSpecFilter struct {
	// +optional
	Prefix string `json:"prefix,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type S3BucketMetricSpec struct {
	Bucket string `json:"bucket"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Filter *[]S3BucketMetricSpec `json:"filter,omitempty"`
	Name   string                `json:"name"`
}

type S3BucketMetricStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// S3BucketMetricList is a list of S3BucketMetrics
type S3BucketMetricList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of S3BucketMetric CRD objects
	Items []S3BucketMetric `json:"items,omitempty"`
}
